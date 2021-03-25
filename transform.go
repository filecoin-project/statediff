package statediff

//go:generate go run ./types/gen ./types ./cmd/stateql/lib

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strings"

	hamt "github.com/filecoin-project/go-hamt-ipld/v2"
	hamtv3 "github.com/filecoin-project/go-hamt-ipld/v3"
	abi "github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/multiformats/go-multihash"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/statediff/types"

	"github.com/filecoin-project/lotus/blockstore"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	adtv2 "github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	adtv3 "github.com/filecoin-project/specs-actors/v3/actors/util/adt"
)

// LotusType represents known types
type LotusType string

// LotusType enum
const (
	LotusTypeUnknown                           LotusType = "unknown"
	LotusTypeTipset                            LotusType = "tipset"
	LotusTypeMsgMeta                           LotusType = "msgMeta"
	LotusTypeMsgList                           LotusType = "msgMeta.BlsMessages"
	LotusTypeMessage                           LotusType = "msgMeta.BlsMessages[]"
	LotusTypeSignedMessage                     LotusType = "msgMeta.SecpkMessages[]"
	LotusTypeStateroot                         LotusType = "stateRoot"
	LotusVersionedStateroot                    LotusType = "versionedStateRoot"
	AccountActorState                          LotusType = "accountActor"
	CronActorState                             LotusType = "cronActor"
	InitActorState                             LotusType = "initActor"
	InitActorV3State                           LotusType = "initActorV3"
	InitActorAddresses                         LotusType = "initActorAddresses"
	InitActorV3Addresses                       LotusType = "initActorV3Addresses"
	MarketActorState                           LotusType = "storageMarketActor"
	MarketActorV2State                         LotusType = "storageMarketActorV2"
	MarketActorV3State                         LotusType = "storageMarketActorV3"
	MarketActorProposals                       LotusType = "storageMarketActor.Proposals"
	MarketActorV2Proposals                     LotusType = "storageMarketActorV2.Proposals"
	MarketActorV3Proposals                     LotusType = "storageMarketActorV3.Proposals"
	MarketActorStates                          LotusType = "storageMarketActor.States"
	MarketActorV3States                        LotusType = "storageMarketActorV3.States"
	MarketActorPendingProposals                LotusType = "storageMarketActor.PendingProposals"
	MarketActorV2PendingProposals              LotusType = "storageMarketActorV2.PendingProposals"
	MarketActorV3PendingProposals              LotusType = "storageMarketActorV3.PendingProposals"
	MarketActorEscrowTable                     LotusType = "storageMarketActor.EscrowTable"
	MarketActorLockedTable                     LotusType = "storageMarketActor.LockedTable"
	MarketActorV3EscrowTable                   LotusType = "storageMarketActorV3.EscrowTable"
	MarketActorV3LockedTable                   LotusType = "storageMarketActorV3.LockedTable"
	MarketActorDealOpsByEpoch                  LotusType = "storageMarketActor.DealOpsByEpoch"
	MarketActorV3DealOpsByEpoch                LotusType = "storageMarketActorV3.DealOpsByEpoch"
	MultisigActorState                         LotusType = "multisigActor"
	MultisigActorV3State                       LotusType = "multisigActorV3"
	MultisigActorPending                       LotusType = "multisigActor.PendingTxns"
	MultisigActorV3Pending                     LotusType = "multisigActorV3.PendingTxns"
	StorageMinerActorState                     LotusType = "storageMinerActor"
	StorageMinerActorInfo                      LotusType = "storageMinerActor.Info"
	StorageMinerActorVestingFunds              LotusType = "storageMinerActor.VestingFunds"
	StorageMinerActorPreCommittedSectors       LotusType = "storageMinerActor.PreCommittedSectors"
	StorageMinerActorV3PreCommittedSectors     LotusType = "storageMinerActorV3.PreCommittedSectors"
	StorageMinerActorPreCommittedSectorsExpiry LotusType = "storageMinerActor.PreCommittedSectorsExpiry"
	StorageMinerActorAllocatedSectors          LotusType = "storageMinerActor.AllocatedSectors"
	StorageMinerActorSectors                   LotusType = "storageMinerActor.Sectors"
	StorageMinerActorV2Sectors                 LotusType = "storageMinerActorV2.Sectors"
	StorageMinerActorV3Sectors                 LotusType = "storageMinerActorV3.Sectors"
	StorageMinerActorDeadlines                 LotusType = "storageMinerActor.Deadlines"
	StorageMinerActorDeadline                  LotusType = "storageMinerActor.Deadlines.Due[]"
	StorageMinerActorDeadlinePartitions        LotusType = "storageMinerActor.Deadlines.Due.Partitions"
	StorageMinerActorDeadlinePartitionExpiry   LotusType = "storageMinerActor.Deadlines.Due.Partitions.ExpirationsEpochs"
	StorageMinerActorDeadlinePartitionEarly    LotusType = "storageMinerActor.Deadlines.Due.Partitions.EarlyTerminated"
	StorageMinerActorDeadlineExpiry            LotusType = "storageMinerActor.Deadlines.Due.ExpirationsEpochs"
	StorageMinerActorV2State                   LotusType = "storageMinerActorV2"
	StorageMinerActorV2Info                    LotusType = "storageMinerActorV2.Info"
	StorageMinerActorV2Deadlines               LotusType = "storageMinerActorV2.Deadlines"
	StorageMinerActorV2Deadline                LotusType = "storageMinerActorV2.Deadlines.Due[]"
	StorageMinerActorV2DeadlinePartitions      LotusType = "storageMinerActorV2.Deadlines.Due.Partitions"
	StorageMinerActorV3State                   LotusType = "storageMinerActorV3"
	StorageMinerActorV3Deadlines               LotusType = "storageMinerActorV3.Deadlines"
	StorageMinerActorV3Deadline                LotusType = "storageMinerActorV3.Deadlines.Due[]"
	StorageMinerActorV3DeadlineExpiry          LotusType = "storageMinerActorV3.Deadlines.Due.ExpirationsEpochs"
	StorageMinerActorV3DeadlinePartitions      LotusType = "storageMinerActorV3.Deadlines.Due.Partitions"
	StorageMinerActorV3DeadlinePartitionExpiry LotusType = "storageMinerActorV3.Deadlines.Due.Partitions.ExpirationsEpochs"
	StorageMinerActorV3DeadlinePartitionEarly  LotusType = "storageMinerActorV3.Deadlines.Due.Partitions.EarlyTerminated"
	StoragePowerActorState                     LotusType = "storagePowerActor"
	StoragePowerActorV2State                   LotusType = "storagePowerActorV2"
	StoragePowerActorV3State                   LotusType = "storagePowerActorV3"
	StoragePowerActorCronEventQueue            LotusType = "storagePowerCronEventQueue"
	StoragePowerActorV3CronEventQueue          LotusType = "storagePowerV3CronEventQueue"
	StoragePowerActorClaims                    LotusType = "storagePowerClaims"
	StoragePowerActorV2Claims                  LotusType = "storagePowerClaimsV2"
	StoragePowerActorV3Claims                  LotusType = "storagePowerClaimsV3"
	RewardActorState                           LotusType = "rewardActor"
	RewardActorV2State                         LotusType = "rewardActorV2"
	VerifiedRegistryActorState                 LotusType = "verifiedRegistryActor"
	VerifiedRegistryActorVerifiers             LotusType = "verifiedRegistryActor.Verifiers"
	VerifiedRegistryActorVerifiedClients       LotusType = "verifiedRegistryActor.VerifiedClients"
	VerifiedRegistryActorV3State               LotusType = "verifiedRegistryActorV3"
	VerifiedRegistryActorV3Verifiers           LotusType = "verifiedRegistryActorV3.Verifiers"
	VerifiedRegistryActorV3VerifiedClients     LotusType = "verifiedRegistryActorV3.VerifiedClients"
	PaymentChannelActorState                   LotusType = "paymentChannelActor"
	PaymentChannelActorLaneStates              LotusType = "paymentChannelActor.LaneStates"
	PaymentChannelActorV3State                 LotusType = "paymentChannelActorV3"
	PaymentChannelActorV3LaneStates            LotusType = "paymentChannelActorV3.LaneStates"
)

// LotusTypeAliases lists non-direct mapped aliases
var LotusTypeAliases = map[string]LotusType{
	"tipset.ParentStateRoot":                                         LotusTypeStateroot,
	"initActor.AddressMap":                                           InitActorAddresses,
	"initActorV3.AddressMap":                                         InitActorV3Addresses,
	"storagePowerActor.CronEventQueue":                               StoragePowerActorCronEventQueue,
	"storagePowerActorV2.CronEventQueue":                             StoragePowerActorCronEventQueue,
	"storagePowerActorV3.CronEventQueue":                             StoragePowerActorV3CronEventQueue,
	"storagePowerActor.Claims":                                       StoragePowerActorClaims,
	"storagePowerActorV2.Claims":                                     StoragePowerActorV2Claims,
	"storagePowerActorV3.Claims":                                     StoragePowerActorV3Claims,
	"storageMarketActorV2.States":                                    MarketActorStates,
	"storageMarketActorV2.EscrowTable":                               MarketActorEscrowTable,
	"storageMarketActorV2.LockedTable":                               MarketActorLockedTable,
	"storageMarketActorV2.DealOpsByEpoch":                            MarketActorDealOpsByEpoch,
	"storageMinerActor.Deadlines.Due.ExpirationEpochs":               StorageMinerActorDeadlineExpiry,
	"storageMinerActor.Deadlines.Due.Partitions.ExpirationEpochs":    StorageMinerActorDeadlinePartitionExpiry,
	"storageMinerActorV2.PreCommittedSectors":                        StorageMinerActorPreCommittedSectors,
	"storageMinerActorV2.PreCommittedSectorsExpiry":                  StorageMinerActorPreCommittedSectorsExpiry,
	"storageMinerActorV2.AllocatedSectors":                           StorageMinerActorAllocatedSectors,
	"storageMinerActorV2.VestingFunds":                               StorageMinerActorVestingFunds,
	"storageMinerActorV2.Deadlines.Due.ExpirationEpochs":             StorageMinerActorDeadlineExpiry,
	"storageMinerActorV2.Deadlines.Due.Partitions.ExpirationsEpochs": StorageMinerActorDeadlinePartitionExpiry,
	"storageMinerActorV2.Deadlines.Due.Partitions.EarlyTerminated":   StorageMinerActorDeadlinePartitionEarly,
	"storageMinerActorV3.Info":                                       StorageMinerActorV2Info,
	"storageMinerActorV3.VestingFunds":                               StorageMinerActorVestingFunds,
	"storageMinerActorV3.AllocatedSectors":                           StorageMinerActorAllocatedSectors,
	"msgMeta.SecpkMessages":                                          LotusTypeMsgList,
}

func cidOf(name string) string {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, _ := builder.Sum([]byte(name))
	return c.String()
}

// LotusActorCodes for actor states
var LotusActorCodes = map[string]LotusType{
	"bafkqaddgnfwc6mjpon4xg5dfnu":                 LotusType("systemActor"),
	"bafkqactgnfwc6mjpnfxgs5a":                    InitActorState,
	"bafkqaddgnfwc6mjpojsxoylsmq":                 RewardActorState,
	"bafkqactgnfwc6mjpmnzg63q":                    CronActorState,
	"bafkqaetgnfwc6mjpon2g64tbm5sxa33xmvza":       StoragePowerActorState,
	"bafkqae3gnfwc6mjpon2g64tbm5sw2ylsnnsxi":      MarketActorState,
	"bafkqaftgnfwc6mjpozsxe2lgnfswi4tfm5uxg5dspe": VerifiedRegistryActorState,
	"bafkqadlgnfwc6mjpmfrwg33vnz2a":               AccountActorState,
	"bafkqadtgnfwc6mjpnv2wy5djonuwo":              MultisigActorState,
	"bafkqafdgnfwc6mjpobqxs3lfnz2gg2dbnzxgk3a":    PaymentChannelActorState,
	"bafkqaetgnfwc6mjpon2g64tbm5sw22lomvza":       StorageMinerActorState,
	// v2
	"bafkqaddgnfwc6mrpon4xg5dfnu":                 LotusType("systemActor"),
	"bafkqactgnfwc6mrpnfxgs5a":                    InitActorState,
	"bafkqaddgnfwc6mrpojsxoylsmq":                 RewardActorV2State,
	"bafkqactgnfwc6mrpmnzg63q":                    CronActorState,
	"bafkqaetgnfwc6mrpon2g64tbm5sxa33xmvza":       StoragePowerActorV2State,
	"bafkqae3gnfwc6mrpon2g64tbm5sw2ylsnnsxi":      MarketActorV2State,
	"bafkqaftgnfwc6mrpozsxe2lgnfswi4tfm5uxg5dspe": VerifiedRegistryActorState,
	"bafkqadlgnfwc6mrpmfrwg33vnz2a":               AccountActorState,
	"bafkqadtgnfwc6mrpnv2wy5djonuwo":              MultisigActorState,
	"bafkqafdgnfwc6mrpobqxs3lfnz2gg2dbnzxgk3a":    PaymentChannelActorState,
	"bafkqaetgnfwc6mrpon2g64tbm5sw22lomvza":       StorageMinerActorV2State,
	// v3
	cidOf("fil/3/system"):           LotusType("systemActor"),
	cidOf("fil/3/init"):             InitActorV3State,
	cidOf("fil/3/cron"):             CronActorState,
	cidOf("fil/3/storagepower"):     StoragePowerActorV3State,
	cidOf("fil/3/storageminer"):     StorageMinerActorV3State,
	cidOf("fil/3/storagemarket"):    MarketActorV3State,
	cidOf("fil/3/paymentchannel"):   PaymentChannelActorV3State,
	cidOf("fil/3/reward"):           RewardActorV2State,
	cidOf("fil/3/verifiedregistry"): VerifiedRegistryActorV3State,
	cidOf("fil/3/account"):          AccountActorState,
	cidOf("fil/3/multisig"):         MultisigActorV3State,
}

var lotusMessageMap = basicnode.Prototype__Map{}

// LotusPrototypes provide expected node types for each state type.
var LotusPrototypes = map[LotusType]ipld.NodePrototype{
	LotusTypeUnknown:                  types.Type.Any__Repr,
	LotusTypeTipset:                   types.Type.LotusBlockHeader__Repr,
	LotusVersionedStateroot:           types.Type.LotusStateRoot__Repr,
	LotusTypeMsgMeta:                  types.Type.LotusMsgMeta__Repr,
	LotusTypeMessage:                  lotusMessageMap,
	LotusTypeSignedMessage:            types.Type.LotusSignedMessage__Repr,
	AccountActorState:                 types.Type.AccountV0State__Repr,
	CronActorState:                    types.Type.CronV0State__Repr,
	InitActorState:                    types.Type.InitV0State__Repr,
	InitActorV3State:                  types.Type.InitV3State__Repr,
	MarketActorState:                  types.Type.MarketV0State__Repr,
	MarketActorV2State:                types.Type.MarketV2State__Repr,
	MarketActorV3State:                types.Type.MarketV3State__Repr,
	MultisigActorState:                types.Type.MultisigV0State__Repr,
	MultisigActorV3State:              types.Type.MultisigV3State__Repr,
	StorageMinerActorState:            types.Type.MinerV0State__Repr,
	StorageMinerActorInfo:             types.Type.MinerV0Info__Repr,
	StorageMinerActorVestingFunds:     types.Type.MinerV0VestingFunds__Repr,
	StorageMinerActorAllocatedSectors: types.Type.BitField__Repr,
	StorageMinerActorDeadlines:        types.Type.MinerV0Deadlines__Repr,
	StorageMinerActorDeadline:         types.Type.MinerV0Deadline__Repr,
	StorageMinerActorV2State:          types.Type.MinerV2State__Repr,
	StorageMinerActorV2Info:           types.Type.MinerV2Info__Repr,
	StorageMinerActorV2Deadlines:      types.Type.MinerV2Deadlines__Repr,
	StorageMinerActorV2Deadline:       types.Type.MinerV2Deadline__Repr,
	StorageMinerActorV3State:          types.Type.MinerV3State__Repr,
	StorageMinerActorV3Deadlines:      types.Type.MinerV3Deadlines__Repr,
	StorageMinerActorV3Deadline:       types.Type.MinerV3Deadline__Repr,
	StoragePowerActorState:            types.Type.PowerV0State__Repr,
	StoragePowerActorV2State:          types.Type.PowerV2State__Repr,
	StoragePowerActorV3State:          types.Type.PowerV3State__Repr,
	RewardActorState:                  types.Type.RewardV0State__Repr,
	RewardActorV2State:                types.Type.RewardV2State__Repr,
	VerifiedRegistryActorState:        types.Type.VerifregV0State__Repr,
	VerifiedRegistryActorV3State:      types.Type.VerifregV3State__Repr,
	PaymentChannelActorState:          types.Type.PaychV0State__Repr,
	PaymentChannelActorV3State:        types.Type.PaychV3State__Repr,
	// Complex types
	LotusTypeMsgList:                           types.Type.List__LinkLotusMessage__Repr,
	LotusTypeStateroot:                         hamt.NewTypedHamt(types.Type.RawAddress__Repr, types.Type.LotusActors__Repr),
	InitActorAddresses:                         types.Type.Map__ActorID__Repr,
	InitActorV3Addresses:                       types.Type.Map__V3ActorID__Repr,
	StorageMinerActorPreCommittedSectors:       types.Type.Map__SectorPreCommitOnChainInfo__Repr,
	StorageMinerActorV3PreCommittedSectors:     types.Type.Map__V3SectorPreCommitOnChainInfo__Repr,
	StorageMinerActorDeadlinePartitionEarly:    types.Type.Map__BitField__Repr,
	StorageMinerActorV3DeadlinePartitionEarly:  types.Type.MapV3__BitField__Repr,
	StorageMinerActorPreCommittedSectorsExpiry: types.Type.Map__BitField__Repr,
	StorageMinerActorSectors:                   types.Type.Map__SectorOnChainInfo__Repr,
	StorageMinerActorV2Sectors:                 types.Type.Map__SectorV2OnChainInfo__Repr,
	StorageMinerActorV3Sectors:                 types.Type.Map__SectorV3OnChainInfo__Repr,
	StorageMinerActorDeadlinePartitions:        types.Type.Map__MinerV0Partition__Repr,
	StorageMinerActorV2DeadlinePartitions:      types.Type.Map__MinerV2Partition__Repr,
	StorageMinerActorV3DeadlinePartitions:      types.Type.Map__MinerV3Partition__Repr,
	StorageMinerActorDeadlinePartitionExpiry:   types.Type.Map__MinerV0ExpirationSet__Repr,
	StorageMinerActorV3DeadlinePartitionExpiry: types.Type.Map__MinerV3ExpirationSet__Repr,
	StorageMinerActorDeadlineExpiry:            types.Type.Map__BitField__Repr,
	StorageMinerActorV3DeadlineExpiry:          types.Type.MapV3__BitField__Repr,
	StoragePowerActorCronEventQueue:            types.Type.Multimap__PowerV0CronEvent__Repr,

	StoragePowerActorV3CronEventQueue:      types.Type.Multimap__PowerV3CronEvent__Repr,
	StoragePowerActorClaims:                types.Type.Map__PowerV0Claim__Repr,
	StoragePowerActorV2Claims:              types.Type.Map__PowerV2Claim__Repr,
	StoragePowerActorV3Claims:              types.Type.Map__PowerV3Claim__Repr,
	VerifiedRegistryActorVerifiers:         types.Type.Map__DataCap__Repr,
	VerifiedRegistryActorVerifiedClients:   types.Type.Map__DataCap__Repr,
	VerifiedRegistryActorV3Verifiers:       types.Type.Map__V3DataCap__Repr,
	VerifiedRegistryActorV3VerifiedClients: types.Type.Map__V3DataCap__Repr,
	MarketActorPendingProposals:            types.Type.Map__MarketV0DealProposal__Repr,
	MarketActorV2PendingProposals:          types.Type.Map__MarketV2DealProposal__Repr,
	MarketActorV3PendingProposals:          types.Type.Map__MarketV3DealProposal__Repr,
	MarketActorProposals:                   types.Type.Map__MarketV0RawDealProposal__Repr,
	MarketActorV2Proposals:                 types.Type.Map__MarketV2RawDealProposal__Repr,
	MarketActorV3Proposals:                 types.Type.Map__MarketV3RawDealProposal__Repr,
	MarketActorStates:                      types.Type.Map__MarketV0DealState__Repr,
	MarketActorV3States:                    types.Type.Map__MarketV3DealState__Repr,
	MarketActorEscrowTable:                 types.Type.Map__BalanceTable__Repr,
	MarketActorV3EscrowTable:               types.Type.Map__V3BalanceTable__Repr,
	MarketActorLockedTable:                 types.Type.Map__BalanceTable__Repr,
	MarketActorV3LockedTable:               types.Type.Map__V3BalanceTable__Repr,
	MarketActorDealOpsByEpoch:              types.Type.Map__List__DealID__Repr,
	MarketActorV3DealOpsByEpoch:            types.Type.MapV3__List__DealID__Repr,
	MultisigActorPending:                   types.Type.Map__MultisigV0Transaction__Repr,
	MultisigActorV3Pending:                 types.Type.Map__MultisigV3Transaction__Repr,
	PaymentChannelActorLaneStates:          types.Type.Map__PaychV0LaneState__Repr,
	PaymentChannelActorV3LaneStates:        types.Type.Map__PaychV3LaneState__Repr,
}

// Loader is the conformer for our wrapper around ADLs
type Loader func(context.Context, cid.Cid, blockstore.Blockstore, ipld.NodeAssembler) error

var complexLoaders = map[ipld.NodePrototype]Loader{
	types.Type.LotusStateRoot__Repr:                    maybeLoadUnversionedStateRoot,
	LotusPrototypes[LotusTypeStateroot]:                loadHamt,
	LotusPrototypes[InitActorAddresses]:                loadHamt,
	types.Type.Map__LotusActors__Repr:                  maybeLoadV3Actors,
	types.Type.Mapv3__LotusActors__Repr:                loadV3Map,
	types.Type.List__LinkLotusMessage__Repr:            loadList,
	types.Type.Map__ActorID__Repr:                      loadMap,
	types.Type.Map__V3ActorID__Repr:                    loadV3Map,
	types.Type.Map__SectorPreCommitOnChainInfo__Repr:   transformMinerActorPreCommittedSectors,
	types.Type.Map__V3SectorPreCommitOnChainInfo__Repr: transformMinerActorV3PreCommittedSectors,
	types.Type.Map__BitField__Repr:                     loadListAsMap,
	types.Type.MapV3__BitField__Repr:                   loadV3ListAsMap(5),
	types.Type.Map__SectorOnChainInfo__Repr:            loadListAsMap,
	types.Type.Map__SectorV2OnChainInfo__Repr:          loadListAsMap,
	types.Type.Map__SectorV3OnChainInfo__Repr:          loadV3ListAsMap(5),
	types.Type.Map__MinerV0Partition__Repr:             loadListAsMap,
	types.Type.Map__MinerV2Partition__Repr:             loadListAsMap,
	types.Type.Map__MinerV3Partition__Repr:             loadV3ListAsMap(5),
	types.Type.Map__MinerV0ExpirationSet__Repr:         loadListAsMap,
	types.Type.Map__MinerV3ExpirationSet__Repr:         loadV3ListAsMap(6),
	types.Type.Multimap__PowerV0CronEvent__Repr:        transformPowerActorEventQueue,
	types.Type.Multimap__PowerV3CronEvent__Repr:        transformPowerActorV3EventQueue,
	types.Type.Map__PowerV0Claim__Repr:                 loadMap,
	types.Type.Map__PowerV2Claim__Repr:                 loadMap,
	types.Type.Map__PowerV3Claim__Repr:                 loadV3Map,
	types.Type.Map__DataCap__Repr:                      transformVerifiedRegistryDataCaps,
	types.Type.Map__V3DataCap__Repr:                    transformVerifiedRegistryV3DataCaps,
	types.Type.Map__MarketV0DealProposal__Repr:         loadListAsMap,
	types.Type.Map__MarketV2DealProposal__Repr:         loadMap,
	types.Type.Map__MarketV3DealProposal__Repr:         loadV3Map,
	types.Type.Map__MarketV0RawDealProposal__Repr:      loadMap,
	types.Type.Map__MarketV2RawDealProposal__Repr:      transformMarketV2Proposals,
	types.Type.Map__MarketV3RawDealProposal__Repr:      loadV3ListAsMap(5),
	types.Type.Map__MarketV0DealState__Repr:            loadListAsMap,
	types.Type.Map__MarketV3DealState__Repr:            loadV3ListAsMap(6),
	types.Type.Map__BalanceTable__Repr:                 transformMarketBalanceTable,
	types.Type.Map__V3BalanceTable__Repr:               transformMarketV3BalanceTable,
	types.Type.Map__List__DealID__Repr:                 transformMarketDealOpsByEpoch,
	types.Type.MapV3__List__DealID__Repr:               transformMarketDealOpsByEpochV3,
	types.Type.Map__MultisigV0Transaction__Repr:        transformMultisigPending,
	types.Type.Map__MultisigV3Transaction__Repr:        transformMultisigPendingV3,
	types.Type.Map__PaychV0LaneState__Repr:             transformPaymentChannelLaneStates,
	types.Type.Map__PaychV3LaneState__Repr:             transformPaymentChannelV3LaneStates,
}

func init() {
	complexLoaders[lotusMessageMap] = loadMessage
}

var typedLinks = map[ipld.NodePrototype]ipld.NodePrototype{
	types.Type.Link__BalanceTable:            types.Type.Map__BalanceTable__Repr,
	types.Type.Link__V3BalanceTable:          types.Type.Map__V3BalanceTable__Repr,
	types.Type.Link__BitField:                types.Type.BitField__Repr,
	types.Type.Link__DataCap:                 types.Type.Map__DataCap__Repr,
	types.Type.Link__V3DataCap:               types.Type.Map__V3DataCap__Repr,
	types.Type.Link__LotusStateRoot:          types.Type.LotusStateRoot__Repr,
	types.Type.Link__MapActorID:              types.Type.Map__ActorID__Repr,
	types.Type.Link__MarketV0DealProposal:    types.Type.Map__MarketV0DealProposal__Repr,
	types.Type.Link__MarketV2DealProposal:    types.Type.Map__MarketV2DealProposal__Repr,
	types.Type.Link__MarketV3DealProposal:    types.Type.Map__MarketV3DealProposal__Repr,
	types.Type.Link__MarketV0DealState:       types.Type.Map__MarketV0DealState__Repr,
	types.Type.Link__MarketV3DealState:       types.Type.Map__MarketV3DealState__Repr,
	types.Type.Link__MarketV0RawDealProposal: types.Type.Map__MarketV0RawDealProposal__Repr,
	types.Type.Link__MarketV2RawDealProposal: types.Type.Map__MarketV2RawDealProposal__Repr,
	types.Type.Link__MarketV3RawDealProposal: types.Type.Map__MarketV3RawDealProposal__Repr,
	types.Type.Link__MultimapDealID:          types.Type.Map__List__DealID__Repr,
	types.Type.Link__MarketV3MultimapDealID:  types.Type.MapV3__List__DealID__Repr,
	types.Type.Link__MinerV0Deadlines:        types.Type.List__MinerV0DeadlineLink__Repr,
	types.Type.Link__MinerV0Deadline:         types.Type.MinerV0Deadline__Repr,
	types.Type.Link__MinerV0ExpirationSet:    types.Type.Map__MinerV0ExpirationSet__Repr,
	types.Type.Link__MinerV0Info:             types.Type.MinerV0Info__Repr,
	types.Type.Link__MinerV0Partition:        types.Type.Map__MinerV0Partition__Repr,
	types.Type.Link__MinerV0SectorInfo:       types.Type.Map__SectorOnChainInfo__Repr,
	types.Type.Link__MinerV2SectorInfo:       types.Type.Map__SectorV2OnChainInfo__Repr,
	types.Type.Link__MinerV0SectorPreCommits: types.Type.Map__SectorPreCommitOnChainInfo__Repr,
	types.Type.Link__MinerV0VestingFunds:     types.Type.MinerV0VestingFunds__Repr,
	types.Type.Link__MinerV2Deadlines:        types.Type.List__MinerV2DeadlineLink__Repr,
	types.Type.Link__MinerV2Deadline:         types.Type.MinerV2Deadline__Repr,
	types.Type.Link__MinerV2Info:             types.Type.MinerV2Info__Repr,
	types.Type.Link__MinerV2Partition:        types.Type.MinerV2Partition__Repr,
	types.Type.Link__MinerV3SectorPreCommits: types.Type.Map__V3SectorPreCommitOnChainInfo__Repr,
	types.Type.Link__MinerV3Deadlines:        types.Type.List__MinerV3DeadlineLink__Repr,
	types.Type.Link__MinerV3Deadline:         types.Type.MinerV3Deadline__Repr,
	types.Type.Link__MinerV3Partition:        types.Type.MinerV3Partition__Repr,
	types.Type.Link__MinerV3ExpirationSet:    types.Type.Map__MinerV3ExpirationSet__Repr,
	types.Type.Link__MinerV3SectorInfo:       types.Type.Map__SectorV3OnChainInfo__Repr,
	types.Type.Link__MultisigV0Transaction:   types.Type.Map__MultisigV0Transaction__Repr,
	types.Type.Link__MultisigV3Transaction:   types.Type.Map__MultisigV3Transaction__Repr,
	types.Type.Link__PaychV0LaneState:        types.Type.Map__PaychV0LaneState__Repr,
	types.Type.Link__PaychV3LaneState:        types.Type.Map__PaychV3LaneState__Repr,
	types.Type.Link__PowerV0Claim:            types.Type.Map__PowerV0Claim__Repr,
	types.Type.Link__PowerV2Claim:            types.Type.Map__PowerV2Claim__Repr,
	types.Type.Link__PowerV3Claim:            types.Type.Map__PowerV3Claim__Repr,
	types.Type.Link__PowerV0CronEvent:        types.Type.Multimap__PowerV0CronEvent__Repr,
	types.Type.Link__PowerV3CronEvent:        types.Type.Multimap__PowerV3CronEvent__Repr,
	types.Type.Link__AccountV0State:          types.Type.AccountV0State__Repr,
	types.Type.Link__CronV0State:             types.Type.CronV0State__Repr,
	types.Type.Link__InitV0State:             types.Type.InitV0State__Repr,
	types.Type.Link__InitV3State:             types.Type.InitV3State__Repr,
	types.Type.Link__MarketV0State:           types.Type.MarketV0State__Repr,
	types.Type.Link__MarketV2State:           types.Type.MarketV2State__Repr,
	types.Type.Link__MarketV3State:           types.Type.MarketV3State__Repr,
	types.Type.Link__MinerV0State:            types.Type.MinerV0State__Repr,
	types.Type.Link__MinerV2State:            types.Type.MinerV2State__Repr,
	types.Type.Link__MinerV3State:            types.Type.MinerV3State__Repr,
	types.Type.Link__MultisigV0State:         types.Type.MultisigV0State__Repr,
	types.Type.Link__MultisigV3State:         types.Type.MultisigV3State__Repr,
	types.Type.Link__PaychV0State:            types.Type.PaychV0State__Repr,
	types.Type.Link__PaychV3State:            types.Type.PaychV3State__Repr,
	types.Type.Link__PowerV0State:            types.Type.PowerV0State__Repr,
	types.Type.Link__PowerV2State:            types.Type.PowerV2State__Repr,
	types.Type.Link__PowerV3State:            types.Type.PowerV3State__Repr,
	types.Type.Link__RewardV0State:           types.Type.RewardV0State__Repr,
	types.Type.Link__RewardV2State:           types.Type.RewardV2State__Repr,
	types.Type.Link__VerifregV0State:         types.Type.VerifregV0State__Repr,
	types.Type.Link__VerifregV3State:         types.Type.VerifregV3State__Repr,
	types.Type.Link__LotusMsgMeta:            types.Type.LotusMsgMeta__Repr,
	types.Type.Link__ListLotusMessage:        types.Type.List__LinkLotusMessage__Repr,
	types.Type.Link__LotusMessage:            types.Type.LotusMessage__Repr,
	types.Type.Link__LotusActors:             types.Type.Map__LotusActors__Repr,
}

// LinkDest fills in a gap in current schema: what type does a `LinkReference` point to
func LinkDest(n ipld.Node) ipld.NodePrototype {
	proto, ok := typedLinks[n.Prototype()]
	if ok {
		return proto
	}
	return nil
}

var (
	simplifyingRe  = regexp.MustCompile(`\[\d+\]\.`)
	simplifyingRe2 = regexp.MustCompile(`\.\d+\.`)
	finalRe        = regexp.MustCompile(`\[\d+\]`)
)

// ResolveType maps incoming type strings to enum known types
func ResolveType(as string) LotusType {
	as = strings.ReplaceAll(as, string('/'), string('.'))
	as = string(simplifyingRe2.ReplaceAll(simplifyingRe.ReplaceAll([]byte(as), []byte(".")), []byte(".")))
	as = string(finalRe.ReplaceAll([]byte(as), []byte("[]")))
	if alias, ok := LotusTypeAliases[as]; ok {
		as = string(alias)
	}
	return LotusType(as)
}

// Load will load c into a given assembler.
func Load(ctx context.Context, c cid.Cid, store blockstore.Blockstore, into ipld.NodeAssembler) error {
	prototype := into.Prototype()
	if complexLoader, ok := complexLoaders[prototype]; ok {
		return complexLoader(ctx, c, store, into)
	}

	block, err := store.Get(c)
	if err != nil {
		return err
	}
	data := block.RawData()

	if err := dagcbor.Decoder(into, bytes.NewBuffer(data)); err != nil {
		return err
	}
	return nil
}

// Transform will unmarshal cbor data based on a provided type hint.
func Transform(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as string) (ipld.Node, error) {
	ca := GetGlobalNodeCache()
	k := string(c.Bytes()) + as
	if n := ca.Get(k); n != nil {
		return n, nil
	}

	proto, ok := LotusPrototypes[ResolveType(as)]
	if !ok {
		return nil, fmt.Errorf("unknown type: %s (parsed to %s)", as, ResolveType(as))
	}
	assembler := proto.NewBuilder()
	if err := Load(ctx, c, store, assembler); err != nil {
		return nil, err
	}
	n := assembler.Build()
	if n != nil {
		ca.Add(k, n)
	}
	return n, nil
}

// TypeOfActor returns the type a given actor will be based on its binary address and stateroot.
func TypeOfActor(stateroot ipld.Node, actor string) (string, error) {
	actr, err := stateroot.LookupByString(actor)
	if err != nil {
		return "", err
	}
	return actorTypeOf(actr)
}

func actorTypeOf(n ipld.Node) (string, error) {
	ref, err := n.LookupByString("Code")
	if err != nil {
		return "", err
	}
	l, err := ref.AsLink()
	if err != nil {
		return "", err
	}
	asCid, ok := l.(cidlink.Link)
	if !ok {
		return "", fmt.Errorf("%s is not a cid", l.String())
	}
	code, ok := LotusActorCodes[asCid.Cid.String()]
	if !ok {
		return "", fmt.Errorf("%s was not a known code", asCid.Cid.String())
	}
	return string(code), nil
}

// TypeActorHead takes a LotusActor node, and returns a `Link__<actortype>State`
// link reference corresponding to the untyped `Head` Link when the type can be
// inferred from the `Code` of the actor.
func TypeActorHead(actor ipld.Node) (ipld.Node, error) {
	t, err := actorTypeOf(actor)
	if err != nil {
		return nil, err
	}
	proto, ok := LotusPrototypes[LotusType(t)]
	if !ok {
		return nil, fmt.Errorf("Unknown prototype for %s", t)
	}
	linkNode, err := actor.LookupByString("Head")
	if err != nil {
		return nil, err
	}
	link, err := linkNode.AsLink()
	if err != nil {
		return nil, err
	}
	for k, v := range typedLinks {
		if v == proto {
			typedBuilder := k.NewBuilder()
			typedBuilder.AssignLink(link)
			return typedBuilder.Build(), nil
		}
	}
	return nil, fmt.Errorf("unknown type of actor: %s", t)
}

// Try directly loading a 'ParentStateRoot' as versioned state root. If that fails, fall back to a v0 sythetic one.
func maybeLoadUnversionedStateRoot(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as ipld.NodeAssembler) error {
	block, err := store.Get(c)
	if err != nil {
		return err
	}
	data := block.RawData()

	tmpAs := as.Prototype().NewBuilder()
	if err := dagcbor.Decoder(tmpAs, bytes.NewBuffer(data)); err != nil {
		list, err := as.BeginList(3)
		if err != nil {
			return err
		}
		if err := list.AssembleValue().AssignInt(0); err != nil {
			return err
		}
		if err := list.AssembleValue().AssignLink(cidlink.Link{Cid: c}); err != nil {
			return err
		}
		if err := list.AssembleValue().AssignLink(cidlink.Link{cid.NewCidV1(cid.Raw, multihash.Multihash([]byte("Synthetic")))}); err != nil {
			return err
		}
		return list.Finish()
	}
	return as.AssignNode(tmpAs.Build())
}

// Try directly loading a 'v3 hamt' as actors map. if that fails fall back to v2
func maybeLoadV3Actors(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as ipld.NodeAssembler) error {
	tmpAs := as.Prototype().NewBuilder()
	if err := loadV3Map(ctx, c, store, tmpAs); err != nil {
		return loadMap(ctx, c, store, as)
	}
	return as.AssignNode(tmpAs.Build())
}

// AllowDegradedADLNodes uses adl walks that ignore store.Get fails
var AllowDegradedADLNodes = false

func loadV3Map(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamtv3.LoadNode(ctx, cborStore, c, hamtv3.UseTreeBitWidth(5))
	if err != nil {
		return fmt.Errorf("hamt load node errored: %v", err)
	}
	mapper, err := as.BeginMap(0)
	if err != nil {
		return err
	}

	mapcb := func(k string, val *cbg.Deferred) error {
		v, err := mapper.AssembleEntry(k)
		if err != nil {
			return err
		}

		value := mapper.ValuePrototype(k).NewBuilder()
		if err := dagcbor.Decoder(value, bytes.NewBuffer(val.Raw)); err != nil {
			return err
		}
		return v.AssignNode(value.Build())
	}

	if AllowDegradedADLNodes {
		if err := degradedForEachV3(node, mapcb, cborStore); err != nil {
			return fmt.Errorf("Hamt browse failed: %v", err)
		}
	} else {
		if err := node.ForEach(ctx, mapcb); err != nil {
			return fmt.Errorf("Hamt browse failed: %v", err)
		}
	}
	if err := mapper.Finish(); err != nil {
		return err
	}
	return nil
}

func loadMap(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return fmt.Errorf("hamt load node errored: %v", err)
	}
	mapper, err := as.BeginMap(0)
	if err != nil {
		return err
	}

	mapcb := func(k string, val interface{}) error {
		v, err := mapper.AssembleEntry(k)
		if err != nil {
			return err
		}

		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}

		value := mapper.ValuePrototype(k).NewBuilder()
		if err := dagcbor.Decoder(value, bytes.NewBuffer(asDef.Raw)); err != nil {
			return err
		}
		return v.AssignNode(value.Build())
	}

	if AllowDegradedADLNodes {
		if err := degradedForEach(node, mapcb, cborStore); err != nil {
			return fmt.Errorf("Hamt browse failed: %v", err)
		}
	} else {
		if err := node.ForEach(ctx, mapcb); err != nil {
			return fmt.Errorf("Hamt browse failed: %v", err)
		}
	}
	if err := mapper.Finish(); err != nil {
		return err
	}
	return nil
}

func loadHamt(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return fmt.Errorf("hamt load node errored: %v", err)
	}
	return as.AssignNode(node)
}

func loadList(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return err
	}

	lister, err := assembler.BeginList(0)
	if err != nil {
		return err
	}

	value := cbg.Deferred{}
	if err := list.ForEach(&value, func(k int64) error {
		vb := lister.AssembleValue()
		return dagcbor.Decoder(vb, bytes.NewBuffer(value.Raw))
	}); err != nil {
		return err
	}
	return lister.Finish()
}

func loadListAsMap(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	value := cbg.Deferred{}
	if err := list.ForEach(&value, func(k int64) error {
		v, err := mapper.AssembleEntry(fmt.Sprintf("%d", k))
		if err != nil {
			return err
		}

		return dagcbor.Decoder(v, bytes.NewBuffer(value.Raw))
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func loadV3ListAsMap(bitwidth int) func(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	return func(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
		cborStore := cbor.NewCborStore(store)
		list, err := adtv3.AsArray(adt.WrapStore(ctx, cborStore), c, bitwidth)
		if err != nil {
			return err
		}

		mapper, err := assembler.BeginMap(0)
		if err != nil {
			return err
		}

		value := cbg.Deferred{}
		if err := list.ForEach(&value, func(k int64) error {
			v, err := mapper.AssembleEntry(fmt.Sprintf("%d", k))
			if err != nil {
				return err
			}

			return dagcbor.Decoder(v, bytes.NewBuffer(value.Raw))
		}); err != nil {
			return err
		}
		return mapper.Finish()
	}
}

type StateRootFunc func(context.Context) []cid.Cid

var StateRootKey = "stateroot"

func loadMessage(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	mapper, err := assembler.BeginMap(3)
	if err != nil {
		return err
	}

	messager, err := mapper.AssembleEntry("Message")
	if err != nil {
		return err
	}
	block, err := store.Get(c)
	if err != nil {
		return err
	}
	data := block.RawData()

	mproto := types.Type.LotusMessage__Repr.NewBuilder()
	if err := dagcbor.Decoder(mproto, bytes.NewBuffer(data)); err != nil {
		return err
	}
	msgNode := mproto.Build()
	if err := messager.AssignNode(msgNode); err != nil {
		return err
	}

	params, err := mapper.AssembleEntry("Params")
	if err != nil {
		return err
	}

	// try to learn stateroot from context.
	actorType := "miner"
	sr := ctx.Value(StateRootKey)
	if sr != nil {
		srf, ok := sr.(StateRootFunc)
		if ok {
			blkCid := srf(ctx)[0]
			sr, err := GetStateRoot(ctx, blkCid, store)
			if err != nil {
				return err
			}
			toAddr, err := msgNode.LookupByString("To")
			if err != nil {
				return err
			}
			addrBytes, err := toAddr.AsBytes()
			if err != nil {
				return err
			}
			actorType, err = TypeOfActor(sr, string(addrBytes))
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid context")
		}
	} else if at, ok := ctx.Value("actorType").(string); ok {
		actorType = at
	} else {
		return fmt.Errorf("unspecified target actor type")
	}

	pn, method, err := ParamFor(LotusType(actorType), msgNode)
	if err != nil {
		return err
	}
	if err := params.AssignNode(pn); err != nil {
		return err
	}

	dat, err := mapper.AssembleEntry("MethodName")
	if err != nil {
		return err
	}
	dat.AssignString(method)

	dat, err = mapper.AssembleEntry("DestinationActorType")
	if err != nil {
		return err
	}
	dat.AssignString(actorType)

	return mapper.Finish()
}

func transformMinerActorPreCommittedSectors(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val interface{}) error {
		i := big.NewInt(0)
		i.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(i.String())
		if err != nil {
			return err
		}

		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}

		actor := types.Type.MinerV0SectorPreCommitOnChainInfo__Repr.NewBuilder()
		if err := dagcbor.Decoder(actor, bytes.NewBuffer(asDef.Raw)); err != nil {
			return err
		}
		return v.AssignNode(actor.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformMinerActorV3PreCommittedSectors(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamtv3.LoadNode(ctx, cborStore, c, hamtv3.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val *cbg.Deferred) error {
		i := big.NewInt(0)
		i.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(i.String())
		if err != nil {
			return err
		}

		actor := types.Type.MinerV0SectorPreCommitOnChainInfo__Repr.NewBuilder()
		if err := dagcbor.Decoder(actor, bytes.NewBuffer(val.Raw)); err != nil {
			return err
		}
		return v.AssignNode(actor.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformPowerActorEventQueue(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := adt.AsMultimap(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return err
	}
	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForAll(func(k string, val *adt.Array) error {
		bi := big.NewInt(0)
		bi.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(bi.String())
		if err != nil {
			return err
		}

		amt := types.Type.Map__PowerV0CronEvent__Repr.NewBuilder()
		amtM, err := amt.BeginMap(0)
		if err != nil {
			return err
		}

		var eval cbg.Deferred
		if err := val.ForEach(&eval, func(i int64) error {
			subv, err := amtM.AssembleEntry(fmt.Sprintf("%d", i))
			if err != nil {
				return err
			}

			actor := types.Type.PowerV0CronEvent__Repr.NewBuilder()
			if err := dagcbor.Decoder(actor, bytes.NewBuffer(eval.Raw)); err != nil {
				return fmt.Errorf("decoding of event leaf failed: %w", err)
			}
			if err := subv.AssignNode(actor.Build()); err != nil {
				return fmt.Errorf("Aassign of thign sad:%w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := amtM.Finish(); err != nil {
			return err
		}
		return v.AssignNode(amt.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformPowerActorV3EventQueue(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := adtv3.AsMultimap(adtv3.WrapStore(ctx, cborStore), c, 6, 6)
	if err != nil {
		return err
	}
	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForAll(func(k string, val *adtv3.Array) error {
		bi := big.NewInt(0)
		bi.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(bi.String())
		if err != nil {
			return err
		}

		amt := types.Type.Map__PowerV0CronEvent__Repr.NewBuilder()
		amtM, err := amt.BeginMap(0)
		if err != nil {
			return err
		}

		var eval cbg.Deferred
		if err := val.ForEach(&eval, func(i int64) error {
			subv, err := amtM.AssembleEntry(fmt.Sprintf("%d", i))
			if err != nil {
				return err
			}

			actor := types.Type.PowerV0CronEvent__Repr.NewBuilder()
			if err := dagcbor.Decoder(actor, bytes.NewBuffer(eval.Raw)); err != nil {
				return fmt.Errorf("decoding of event leaf failed: %w", err)
			}
			if err := subv.AssignNode(actor.Build()); err != nil {
				return fmt.Errorf("Aassign of thign sad:%w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := amtM.Finish(); err != nil {
			return err
		}
		return v.AssignNode(amt.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformVerifiedRegistryDataCaps(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val interface{}) error {
		v, err := mapper.AssembleEntry(k)
		if err != nil {
			return err
		}

		// Deferred parsing of big.Int
		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}

		return v.AssignBytes(asDef.Raw)
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformVerifiedRegistryV3DataCaps(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamtv3.LoadNode(ctx, cborStore, c, hamtv3.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val *cbg.Deferred) error {
		v, err := mapper.AssembleEntry(k)
		if err != nil {
			return err
		}

		return v.AssignBytes(val.Raw)
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformMarketV2Proposals(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	list, err := adtv2.AsArray(adtv2.WrapStore(ctx, cborStore), c)
	if err != nil {
		return fmt.Errorf("couldn't load amt:%w", err)
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	value := cbg.Deferred{}
	if err := list.ForEach(&value, func(k int64) error {
		v, err := mapper.AssembleEntry(fmt.Sprintf("%d", k))
		if err != nil {
			return err
		}

		actor := types.Type.MarketV2DealProposal__Repr.NewBuilder()
		if err := dagcbor.Decoder(actor, bytes.NewBuffer(value.Raw)); err != nil {
			return fmt.Errorf("unmarshal of individual proposal failed: %w", err)
		}
		return v.AssignNode(actor.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformMarketBalanceTable(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val interface{}) error {
		v, err := mapper.AssembleEntry(k)
		if err != nil {
			return err
		}

		// Deferred parsing of big.Int
		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}

		return v.AssignBytes(asDef.Raw)
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformMarketV3BalanceTable(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamtv3.LoadNode(ctx, cborStore, c, hamtv3.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val *cbg.Deferred) error {
		v, err := mapper.AssembleEntry(k)
		if err != nil {
			return err
		}

		return v.AssignBytes(val.Raw)
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformMarketDealOpsByEpoch(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	adtStore := adt.WrapStore(ctx, cbor.NewCborStore(store))
	table, err := adt.AsMap(adtStore, c)
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	var value cbg.CborCid
	if err := table.ForEach(&value, func(k string) error {
		set, err := adt.AsSet(adtStore, cid.Cid(value))
		if err != nil {
			return err
		}

		b := big.NewInt(0)
		b.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(b.String())
		if err != nil {
			return err
		}

		amt := types.Type.List__DealID__Repr.NewBuilder()
		amtL, err := amt.BeginList(0)
		if err != nil {
			return err
		}

		set.ForEach(func(d string) error {
			key, err := abi.ParseUIntKey(d)
			if err != nil {
				return err
			}
			return amtL.AssembleValue().AssignInt(int64(key))
		})

		if err := amtL.Finish(); err != nil {
			return err
		}

		return v.AssignNode(amt.Build())
	}); err != nil {
		return err
	}

	return mapper.Finish()
}

func transformMarketDealOpsByEpochV3(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	adtStore := adt.WrapStore(ctx, cbor.NewCborStore(store))
	table, err := adtv3.AsMap(adtStore, c, 5)
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	var value cbg.CborCid
	if err := table.ForEach(&value, func(k string) error {
		set, err := adtv3.AsSet(adtStore, cid.Cid(value), 5)
		if err != nil {
			return err
		}

		b := big.NewInt(0)
		b.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(b.String())
		if err != nil {
			return err
		}

		amt := types.Type.List__DealID__Repr.NewBuilder()
		amtL, err := amt.BeginList(0)
		if err != nil {
			return err
		}

		set.ForEach(func(d string) error {
			key, err := abi.ParseUIntKey(d)
			if err != nil {
				return err
			}
			return amtL.AssembleValue().AssignInt(int64(key))
		})

		if err := amtL.Finish(); err != nil {
			return err
		}

		return v.AssignNode(amt.Build())
	}); err != nil {
		return err
	}

	return mapper.Finish()
}

func transformMultisigPending(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val interface{}) error {
		i := big.NewInt(0)
		i.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(i.String())
		if err != nil {
			return err
		}

		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}

		actor := types.Type.MultisigV0Transaction__Repr.NewBuilder()
		if err := dagcbor.Decoder(actor, bytes.NewBuffer(asDef.Raw)); err != nil {
			return err
		}
		return v.AssignNode(actor.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformMultisigPendingV3(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	node, err := hamtv3.LoadNode(ctx, cborStore, c, hamtv3.UseTreeBitWidth(5))
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	if err := node.ForEach(ctx, func(k string, val *cbg.Deferred) error {
		i := big.NewInt(0)
		i.SetBytes([]byte(k))
		v, err := mapper.AssembleEntry(i.String())
		if err != nil {
			return err
		}

		actor := types.Type.MultisigV0Transaction__Repr.NewBuilder()
		if err := dagcbor.Decoder(actor, bytes.NewBuffer(val.Raw)); err != nil {
			return err
		}
		return v.AssignNode(actor.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformPaymentChannelLaneStates(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	value := cbg.Deferred{}
	if err := list.ForEach(&value, func(k int64) error {
		v, err := mapper.AssembleEntry(fmt.Sprintf("%d", k))
		if err != nil {
			return err
		}

		actor := types.Type.PaychV0LaneState__Repr.NewBuilder()
		if err := dagcbor.Decoder(actor, bytes.NewBuffer(value.Raw)); err != nil {
			return err
		}
		return v.AssignNode(actor.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}

func transformPaymentChannelV3LaneStates(ctx context.Context, c cid.Cid, store blockstore.Blockstore, assembler ipld.NodeAssembler) error {
	cborStore := cbor.NewCborStore(store)
	list, err := adtv3.AsArray(adt.WrapStore(ctx, cborStore), c, 3)
	if err != nil {
		return err
	}

	mapper, err := assembler.BeginMap(0)
	if err != nil {
		return err
	}

	value := cbg.Deferred{}
	if err := list.ForEach(&value, func(k int64) error {
		v, err := mapper.AssembleEntry(fmt.Sprintf("%d", k))
		if err != nil {
			return err
		}

		actor := types.Type.PaychV0LaneState__Repr.NewBuilder()
		if err := dagcbor.Decoder(actor, bytes.NewBuffer(value.Raw)); err != nil {
			return err
		}
		return v.AssignNode(actor.Build())
	}); err != nil {
		return err
	}
	return mapper.Finish()
}
