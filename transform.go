package statediff

import (
	"bytes"
	"context"
	"fmt"
	"regexp"

	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	abi "github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	hamt "github.com/ipfs/go-hamt-ipld"
	"github.com/ipfs/go-ipfs-blockstore"
	cbor "github.com/ipfs/go-ipld-cbor"
	cbg "github.com/whyrusleeping/cbor-gen"

	lotusTypes "github.com/filecoin-project/lotus/chain/types"

	accountActor "github.com/filecoin-project/specs-actors/actors/builtin/account"
	cronActor "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	initActor "github.com/filecoin-project/specs-actors/actors/builtin/init"
	marketActor "github.com/filecoin-project/specs-actors/actors/builtin/market"
	storageMinerActor "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	multisigActor "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	paychActor "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	storagePowerActor "github.com/filecoin-project/specs-actors/actors/builtin/power"
	rewardActor "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	verifiedRegistryActor "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
)

type LotusType string

const (
	LotusTypeTipset                            LotusType = "tipset"
	LotusTypeStateroot                         LotusType = "stateRoot"
	AccountActorState                          LotusType = "accountActor"
	CronActorState                             LotusType = "cronActor"
	InitActorState                             LotusType = "initActor"
	InitActorAddresses                         LotusType = "initActorAddresses"
	MarketActorState                           LotusType = "storageMarketActor"
	MarketActorProposals                       LotusType = "storageMarketActor.Proposals"
	MarketActorStates                          LotusType = "storageMarketActor.State"
	MarketActorPendingProposals                LotusType = "storageMarketActor.PendingProposals"
	MarketActorEscrowTable                     LotusType = "storageMarketActor.EscrowTable"
	MarketActorLockedTable                     LotusType = "storageMarketActor.LockedTable"
	MarketActorDealOpsByEpoch                  LotusType = "storageMarketActor.DealOpsByEpoch"
	MultisigActorState                         LotusType = "multisigActor"
	MultisigActorPending                       LotusType = "multisigActor.PendingTxns"
	StorageMinerActorState                     LotusType = "storageMinerActor"
	StorageMinerActorInfo                      LotusType = "storageMinerActor.Info"
	StorageMinerActorVestingFunds              LotusType = "storageMinerActor.VestingFunds"
	StorageMinerActorPreCommittedSectors       LotusType = "storageMinerActor.PreCommittedSectors"
	StorageMinerActorPreCommittedSectorsExpiry LotusType = "storageMinerActor.PreCommittedSectorsExpiry"
	StorageMinerActorAllocatedSectors          LotusType = "storageMinerActor.AllocatedSectors"
	StorageMinerActorSectors                   LotusType = "storageMinerActor.Sectors"
	StorageMinerActorDeadlines                 LotusType = "storageMinerActor.Deadlines"
	StorageMinerActorDeadline                  LotusType = "storageMinerActor.Deadlines.Due"
	StorageMinerActorDeadlinePartitions        LotusType = "storageMinerActor.Deadlines.Due.Partitions"
	StorageMinerActorDeadlineExpiry            LotusType = "storageMinerActor.Deadlines.Due.ExpirationsEpochs"
	StoragePowerActorState                     LotusType = "storagePowerActor"
	StoragePowerActorCronEventQueue            LotusType = "storagePowerCronEventQueue"
	StoragePowerActorClaims                    LotusType = "storagePowerClaims"
	RewardActorState                           LotusType = "rewardActor"
	VerifiedRegistryActorState                 LotusType = "verifiedRegistryActor"
	VerifiedRegistryActorVerifiers             LotusType = "verifiedRegistryActor.Verifiers"
	VerifiedRegistryActorVerifiedClients       LotusType = "verifiedRegistryActor.VerifiedClients"
	PaymentChannelActorState                   LotusType = "paymentChannelActor"
	PaymentChannelActorLaneStates              LotusType = "paymentChannelActor.LaneStates"
)

var simplifyingRe = regexp.MustCompile(`\[\d+\]`)

// Transform will unmarshal cbor data based on a provided type hint.
func Transform(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as string) (interface{}, error) {
	as = string(simplifyingRe.ReplaceAll([]byte(as), []byte("")))

	// First select types which do their own store loading.
	switch LotusType(as) {
	case LotusTypeStateroot:
		return transformStateRoot(ctx, c, store)
	case InitActorAddresses:
		return transformInitActor(ctx, c, store)
	case StorageMinerActorPreCommittedSectors:
		return transformMinerActorPreCommittedSectors(ctx, c, store)
	case StorageMinerActorPreCommittedSectorsExpiry:
		return transformMinerActorPreCommittedSectorsExpiry(ctx, c, store)
	case StorageMinerActorSectors:
		return transformMinerActorSectors(ctx, c, store)
	case StorageMinerActorDeadlinePartitions:
		return transformMinerActorDeadlinePartitions(ctx, c, store)
	case StorageMinerActorDeadlineExpiry:
		return transformMinerActorDeadlineExpiry(ctx, c, store)
	case StoragePowerActorCronEventQueue:
		return transformPowerActorEventQueue(ctx, c, store)
	case StoragePowerActorClaims:
		return transformPowerActorClaims(ctx, c, store)
	case MarketActorProposals:
		return transformMarketProposals(ctx, c, store)
	case MarketActorStates:
		return transformMarketStates(ctx, c, store)
	case MarketActorPendingProposals:
		return transformMarketPendingProposals(ctx, c, store)
	case MarketActorEscrowTable:
		fallthrough
	case MarketActorLockedTable:
		return transformMarketBalanceTable(ctx, c, store)
	case MultisigActorPending:
		return transformMultisigPending(ctx, c, store)
	// TODO: MarketActorDealOpsByEpoch (multimap)
	case VerifiedRegistryActorVerifiers:
		fallthrough
	case VerifiedRegistryActorVerifiedClients:
		return transformVerifiedRegistryDataCaps(ctx, c, store)
	case PaymentChannelActorLaneStates:
		return transformPaymentChannelLaneStates(ctx, c, store)
	default:
	}

	block, err := store.Get(c)
	if err != nil {
		return nil, err
	}
	data := block.RawData()

	// Then select types which use block data.
	switch LotusType(as) {
	case LotusTypeTipset:
		dest := lotusTypes.BlockHeader{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case AccountActorState:
		dest := accountActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case CronActorState:
		dest := cronActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case InitActorState:
		dest := initActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case MarketActorState:
		dest := marketActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case MultisigActorState:
		dest := multisigActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case StorageMinerActorState:
		dest := storageMinerActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case StorageMinerActorInfo:
		dest := storageMinerActor.MinerInfo{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case StorageMinerActorVestingFunds:
		dest := storageMinerActor.VestingFunds{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case StorageMinerActorAllocatedSectors:
		dest := bitfield.BitField{}
		err := cbor.DecodeInto(data, &dest)
		return JSONBitField{dest}, err
	case StorageMinerActorDeadlines:
		dest := storageMinerActor.Deadlines{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case StorageMinerActorDeadline:
		dest := storageMinerActor.Deadline{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case StoragePowerActorState:
		dest := storagePowerActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case RewardActorState:
		dest := rewardActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case VerifiedRegistryActorState:
		dest := verifiedRegistryActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	case PaymentChannelActorState:
		dest := paychActor.State{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	default:
		var dest interface{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	}
}

func transformStateRoot(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return nil, err
	}
	m := make(map[string]*lotusTypes.Actor)
	node.ForEach(ctx, func(k string, val interface{}) error {
		actor := lotusTypes.Actor{}
		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}
		err := cbor.DecodeInto(asDef.Raw, &actor)
		if err != nil {
			return err
		}
		a, _ := addr.NewFromBytes([]byte(k))
		m[a.String()] = &actor
		return nil
	})
	return m, nil
}

func transformInitActor(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return nil, err
	}
	m := make(map[string]uint64)
	var actorID cbg.CborInt
	node.ForEach(ctx, func(k string, val interface{}) error {
		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}
		err := cbor.DecodeInto(asDef.Raw, &actorID)
		if err != nil {
			return err
		}
		a, _ := addr.NewFromBytes([]byte(k))
		m[a.String()] = uint64(actorID)
		return nil
	})
	return m, nil
}

func transformMinerActorPreCommittedSectors(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	table, err := adt.AsMap(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]storageMinerActor.SectorPreCommitOnChainInfo)
	var value storageMinerActor.SectorPreCommitOnChainInfo
	var key cbg.CborInt
	if err := table.ForEach(&value, func(k string) error {
		(&key).UnmarshalCBOR(bytes.NewBuffer([]byte(k)))
		m[int64(key)] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMinerActorPreCommittedSectorsExpiry(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]JSONBitField)
	value := bitfield.BitField{}
	if err := list.ForEach(&value, func(k int64) error {
		m[k] = JSONBitField{value}
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMinerActorSectors(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]storageMinerActor.SectorOnChainInfo)
	value := storageMinerActor.SectorOnChainInfo{}
	if err := list.ForEach(&value, func(k int64) error {
		m[k] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMinerActorDeadlinePartitions(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]storageMinerActor.Partition)
	value := storageMinerActor.Partition{}
	if err := list.ForEach(&value, func(k int64) error {
		m[k] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMinerActorDeadlineExpiry(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]JSONBitField)
	value := bitfield.BitField{}
	if err := list.ForEach(&value, func(k int64) error {
		m[k] = JSONBitField{value}
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformPowerActorEventQueue(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	node, err := adt.AsMultimap(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}
	m := make(map[uint64]map[int64]storagePowerActor.CronEvent)
	var key cbg.CborInt
	if err := node.ForAll(func(k string, val *adt.Array) error {
		eval := storagePowerActor.CronEvent{}
		items := make(map[int64]storagePowerActor.CronEvent)
		if err := val.ForEach(&eval, func(i int64) error {
			items[i] = eval
			return nil
		}); err != nil {
			return err
		}
		(&key).UnmarshalCBOR(bytes.NewBuffer([]byte(k)))
		m[uint64(key)] = items
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformPowerActorClaims(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return nil, err
	}
	m := make(map[string]storagePowerActor.Claim)
	var claim storagePowerActor.Claim
	node.ForEach(ctx, func(k string, val interface{}) error {
		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}
		err := cbor.DecodeInto(asDef.Raw, &claim)
		if err != nil {
			return err
		}
		a, _ := addr.NewFromBytes([]byte(k))
		m[a.String()] = claim
		return nil
	})
	return m, nil
}

func transformVerifiedRegistryDataCaps(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	node, err := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
	if err != nil {
		return nil, err
	}
	m := make(map[string]verifiedRegistryActor.DataCap)
	var dataCap verifiedRegistryActor.DataCap
	node.ForEach(ctx, func(k string, val interface{}) error {
		asDef, ok := val.(*cbg.Deferred)
		if !ok {
			return fmt.Errorf("unexpected non-cbg.Deferred")
		}
		err := cbor.DecodeInto(asDef.Raw, &dataCap)
		if err != nil {
			return err
		}
		a, _ := addr.NewFromBytes([]byte(k))
		m[a.String()] = dataCap
		return nil
	})
	return m, nil
}

func transformMarketPendingProposals(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	mapper, err := adt.AsMap(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[string]marketActor.DealProposal)
	cidr := cid.Undef
	value := marketActor.DealProposal{}
	if err := mapper.ForEach(&value, func(c string) error {
		cidr.UnmarshalBinary([]byte(c))
		m[cidr.String()] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMarketProposals(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]marketActor.DealProposal)
	value := marketActor.DealProposal{}
	if err := list.ForEach(&value, func(k int64) error {
		m[k] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMarketStates(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]marketActor.DealState)
	value := marketActor.DealState{}
	if err := list.ForEach(&value, func(k int64) error {
		m[k] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMarketBalanceTable(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	table, err := adt.AsMap(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[string]abi.TokenAmount)
	var value abi.TokenAmount
	if err := table.ForEach(&value, func(k string) error {
		a, _ := addr.NewFromBytes([]byte(k))
		m[a.String()] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformMultisigPending(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	table, err := adt.AsMap(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]multisigActor.Transaction)
	var value multisigActor.Transaction
	var key cbg.CborInt
	if err := table.ForEach(&value, func(k string) error {
		(&key).UnmarshalCBOR(bytes.NewBuffer([]byte(k)))
		m[int64(key)] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}

func transformPaymentChannelLaneStates(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (interface{}, error) {
	cborStore := cbor.NewCborStore(store)
	list, err := adt.AsArray(adt.WrapStore(ctx, cborStore), c)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]paychActor.LaneState)
	value := paychActor.LaneState{}
	if err := list.ForEach(&value, func(k int64) error {
		m[k] = value
		return nil
	}); err != nil {
		return nil, err
	}
	return m, nil
}
