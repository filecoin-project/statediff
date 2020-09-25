package statediff

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strings"

	addr "github.com/filecoin-project/go-address"
	bitfield "github.com/filecoin-project/go-bitfield"
	cboriface "github.com/filecoin-project/go-state-types/cbor"
	chainState "github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	accountActor "github.com/filecoin-project/specs-actors/actors/builtin/account"
	cronActor "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	initActor "github.com/filecoin-project/specs-actors/actors/builtin/init"
	marketActor "github.com/filecoin-project/specs-actors/actors/builtin/market"
	storageMinerActor "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	storagePowerActor "github.com/filecoin-project/specs-actors/actors/builtin/power"
	rewardActor "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	verifiedRegistryActor "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	hamt "github.com/ipfs/go-hamt-ipld"
	cbor "github.com/ipfs/go-ipld-cbor"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/willscott/go-cmp/cmp"
)

type config struct {
	ExpandActors   bool
	ActorCidFilter []cid.Cid
}

type Option func(c *config)

// ExpandActors indicates that exploration should recurse into actor states when
// they differ, rather than simply indicating which state roots have diverged.
func ExpandActors(c *config) {
	c.ExpandActors = true
}

// ExpandActorByCid expands only specific actor states that diverge. Cids that
// diverge can bes specified by the `CodeID` to expand a specific type of actor
// e.g. `builtin.AccountActorCodeID`. or a specific instance of an actor based
// on the specific `Actor.Head`.
func ExpandActorByCid(cids []cid.Cid) Option {
	return func(c *config) {
		c.ExpandActors = true
		c.ActorCidFilter = cids
	}
}

// Parse a user entered fuzzy definition for actor expansion.
func WithActorExpansionFromUser(arg string) (Option, error) {
	if arg == "all" {
		return ExpandActors, nil
	}
	parts := strings.Split(arg, ",")

	cids := make([]cid.Cid, 0, len(parts))
	for _, part := range parts {
		c, err := cid.Parse(part)
		if err != nil {
			for k, v := range builtinCodeClasses {
				if part == k {
					cids = append(cids, v)
					break
				}
			}
			return nil, fmt.Errorf("Could not parse %s: %w", part, err)
		} else {
			cids = append(cids, c)
		}
	}
	return ExpandActorByCid(cids), nil
}

var builtinCodeClasses = map[string]cid.Cid{
	"init":             builtin.InitActorCodeID,
	"cron":             builtin.CronActorCodeID,
	"account":          builtin.AccountActorCodeID,
	"storagePower":     builtin.StoragePowerActorCodeID,
	"storageMiner":     builtin.StorageMinerActorCodeID,
	"storageMarket":    builtin.StorageMarketActorCodeID,
	"paymentChannel":   builtin.PaymentChannelActorCodeID,
	"multisig":         builtin.MultisigActorCodeID,
	"reward":           builtin.RewardActorCodeID,
	"verifiedRegistry": builtin.VerifiedRegistryActorCodeID,
}

func Diff(ctx context.Context, store blockstore.Blockstore, a, b cid.Cid, opts ...Option) string {
	conf := config{}
	for _, o := range opts {
		o(&conf)
	}

	cborStore := cbor.NewCborStore(store)
	adtStore := adt.WrapStore(ctx, cborStore)

	initActorTransformer := func(act initActor.State) *initActorState {
		am, _ := adt.AsMap(adtStore, act.AddressMap)
		var val cbg.CborInt
		m := make(map[string]uint64)
		am.ForEach(&val, func(k string) error {
			address, _ := addr.NewFromBytes([]byte(k))
			m[address.String()] = uint64(val)
			return nil
		})
		return &initActorState{
			NextID:      act.NextID.String(),
			NetworkName: act.NetworkName,
			ADTRoot:     act.AddressMap.String(),
			ADT:         m,
		}
	}

	stateTreeNamer := getInitFor(ctx, cborStore, a, initActorTransformer)

	hamtActorExpander := func(n *hamtNode) map[string]*types.Actor {
		m := make(map[string]*types.Actor)
		kv := hamt.KV{}
		n.Node.ForEach(ctx, func(k string, val interface{}) error {
			kv.Key = []byte(k)
			if v, ok := val.(*cbg.Deferred); !ok {
				return fmt.Errorf("unexpected hamt node %v", val)
			} else {
				kv.Value = v
			}

			var template types.Actor
			cbor.DecodeInto(kv.Value.Raw, &template)

			if nk, ok := stateTreeNamer[k]; ok {
				k = nk
			}
			m[k] = &template

			return nil
		})
		return m
	}

	initHamtTransformer := func(n *hamtNode) map[string]string {
		m := make(map[string]string)
		n.Node.ForEach(ctx, func(k string, val interface{}) error {
			m[k] = fmt.Sprintf("%v", val)
			return nil
		})
		return m
	}

	actorTransformer := func(act *types.Actor) *statefulActor {
		var state interface{}
		block, _ := store.Get(act.Head)

		state = block

		found := true
		if len(conf.ActorCidFilter) > 0 {
			found = false
			for _, cid := range conf.ActorCidFilter {
				if cid.Equals(act.Code) || cid.Equals(act.Head) {
					found = true
					break
				}
			}
		}

		if found {
			switch act.Code {
			case builtin.InitActorCodeID:
				var initState initActor.State
				cbor.DecodeInto(block.RawData(), &initState)
				state = initState
			case builtin.RewardActorCodeID:
				var rewardState rewardActor.State
				cbor.DecodeInto(block.RawData(), &rewardState)
				state = rewardState
			case builtin.StoragePowerActorCodeID:
				var storagePowerState storagePowerActor.State
				cbor.DecodeInto(block.RawData(), &storagePowerState)
				state = storagePowerState
			case builtin.StorageMinerActorCodeID:
				var storageMinerState storageMinerActor.State
				cbor.DecodeInto(block.RawData(), &storageMinerState)
				state = storageMinerState
			case builtin.CronActorCodeID:
				var cronState cronActor.State
				cbor.DecodeInto(block.RawData(), &cronState)
				state = cronState
			case builtin.StorageMarketActorCodeID:
				var marketState marketActor.State
				cbor.DecodeInto(block.RawData(), &marketState)
				state = marketState
			case builtin.VerifiedRegistryActorCodeID:
				var registryState verifiedRegistryActor.State
				cbor.DecodeInto(block.RawData(), &registryState)
				state = registryState
			case builtin.AccountActorCodeID:
				var accountState accountActor.State
				cbor.DecodeInto(block.RawData(), &accountState)
				state = accountState
			default:
				state = block
			}
		}
		return &statefulActor{
			Type:    builtin.ActorNameByCode(act.Code),
			State:   state,
			Nonce:   act.Nonce,
			Balance: act.Balance.String(),
		}
	}

	cidMap := make(map[string]reflect.Type)
	cidMap[`.*BasicBlock\)\.cid`] = reflect.TypeOf("")        // don't recurse on uninterpreted data.
	cidMap[`^{cid.Cid}$`] = reflect.TypeOf((*hamt.Node)(nil)) // handle state root expansion as a top-level hamt

	// storageMinerActor mappings
	cidMap[`\(miner\.State\)\.Info$`] = reflect.TypeOf((*storageMinerActor.MinerInfo)(nil))
	cidMap[`\(miner\.State\)\.VestingFunds$`] = reflect.TypeOf("") // TODO: use VestingFund after next spec-actors release.
	cidMap[`\(miner\.State\)\.PreCommittedSectors$`] = reflect.TypeOf(make(map[string]*storageMinerActor.SectorPreCommitOnChainInfo))
	cidMap[`\(miner\\.State\)\.PreCommittedSectors.*SealedCID$`] = reflect.TypeOf("")
	cidMap[`\(miner\.State\)\.PreCommittedSectorsExpiry$`] = reflect.TypeOf("")
	cidMap[`\(miner\.State\)\.AllocatedSectors$`] = reflect.TypeOf((*bitfield.BitField)(nil))
	cidMap[`\(miner\.State\)\.Sectors$`] = reflect.TypeOf((*adt.Array)(nil))
	cidMap[`\(miner\.State\)\.Sectors.*SealedCID$`] = reflect.TypeOf("")
	cidMap[`\(miner\.State\)\.Deadlines$`] = reflect.TypeOf((*storageMinerActor.Deadlines)(nil))
	cidMap[`\(miner\.State\)\.Deadlines.*Due([^\.]*)$`] = reflect.TypeOf((*storageMinerActor.Deadline)(nil))
	cidMap[`\(miner\.State\).*Partitions$`] = reflect.TypeOf(make([]*storageMinerActor.Partition, 0))
	cidMap[`miner\.Deadline\)\.ExpirationsEpochs$`] = reflect.TypeOf(make([]*bitfield.BitField, 0))
	cidMap[`\(miner\.State\).*Partitions.*ExpirationsEpochs$`] = reflect.TypeOf(make([]*storageMinerActor.ExpirationSet, 0))
	cidMap[`\(miner\.State\).*EarlyTerminated$`] = reflect.TypeOf(make([]*bitfield.BitField, 0))

	// storagePowerActor mappings
	cidMap[`\(power\.State\)\.CronEventQueue$`] = reflect.TypeOf("") // TODO: support for 'multimap'
	cidMap[`\(power\.State\)\.Claims$`] = reflect.TypeOf(make(map[string]*storagePowerActor.Claim))
	cidMap[`\(power\.State\)\.ProofValidationBatch$`] = reflect.TypeOf("") // TODO: used?

	// marketActor mappings
	cidMap[`\(market\.State\)\.Proposals$`] = reflect.TypeOf(make([]*marketActor.DealProposal, 0))
	cidMap[`\(market\.State\).*PieceCID$`] = reflect.TypeOf("")
	cidMap[`\(market\.State\)\.States$`] = reflect.TypeOf(make([]*marketActor.DealState, 0))
	cidMap[`\(market\.State\)\.PendingProposals$`] = reflect.TypeOf(make(map[string]*marketActor.DealProposal))
	cidMap[`\(market\.State\)\.EscrowTable$`] = reflect.TypeOf(make(map[string]*big.Int)) // adt.BalanceTable
	cidMap[`\(market\.State\)\.LockedTable$`] = reflect.TypeOf(make(map[string]*big.Int)) // adt.BalanceTable
	cidMap[`\(market\.State\)\.DealOpsByEpoch$`] = reflect.TypeOf("")                     // TODO: support for 'multimap'

	// verifiedRegistryActor mappings
	cidMap[`\(verifreg\.State\)\.Verifiers$`] = reflect.TypeOf(make(map[string]*verifiedRegistryActor.DataCap))
	cidMap[`\(verifreg\.State\)\.VerifiedClients$`] = reflect.TypeOf(make(map[string]*verifiedRegistryActor.DataCap))

	cmpOpts := []cmp.Option{
		cmp.Comparer(bigIntComparer),
		cmp.AllowUnexported(blocks.BasicBlock{}),
		cmp.Transformer("bitfield.Bitfield", bitfieldTransformer),
		cmp.Transformer("bitfield.InlineBitfield", directBitfieldTransformer),
		cmp.Transformer("address.Address", addressTransformer),
		cmp.Transformer("initActor.State", initActorTransformer),
		cmp.FilterPath(filterIn("initActor"), cmp.Transformer("init.State", initHamtTransformer)),
		cmp.FilterPath(topFilter, cmp.Transformer("state.StateTree", hamtActorExpander)),
	}
	if conf.ExpandActors {
		cmpOpts = append(cmpOpts, cmp.Transformer("types.Actor", actorTransformer))
	} else {
		cidMap[`\.Code$`] = reflect.TypeOf("")
		cidMap[`\.Head$`] = reflect.TypeOf("")
	}
	cmpOpts = append(cmpOpts, cidTransformer(ctx, store, cborStore, cidMap)...)
	coreDiff := cmp.Diff(a, b, cmpOpts...)

	header := fmt.Sprintf("--- %s\n+++ %s\n@@ -1,1 +1,1 @@\n", a, b)
	return header + coreDiff
}

func cidTransformer(ctx context.Context, store blockstore.Blockstore, cborStore cbor.IpldStore, atlas map[string]reflect.Type) []cmp.Option {
	var options []cmp.Option
	pathFilter := func(matcher string) func(p cmp.Path) bool {
		return func(p cmp.Path) bool {
			ok, _ := regexp.MatchString(matcher, p.GoString())
			return ok
		}
	}

	for r, t := range atlas {
		name := strings.Trim(t.String(), "*")
		if t.Kind() == reflect.Map {
			name = "amtmap." + strings.Trim(t.Elem().String(), "*")
		} else if t.Kind() == reflect.Slice {
			name = "amtarray." + strings.Trim(t.Elem().String(), "*")
		}
		boundFunc := (func(name string, t reflect.Type) func(c cid.Cid) interface{} {
			return func(c cid.Cid) interface{} {
				if c == cid.Undef {
					return "<Undefined>"
				}
				if name == "hamt.Node" {
					// special case for maps / other complex data types that require multiple cid walks.
					node, _ := hamt.LoadNode(ctx, cborStore, c, hamt.UseTreeBitWidth(5))
					return &hamtNode{Cid: c, Node: node}
				} else if name == "string" {
					// special case for not expanding.
					return c.String()
				} else if strings.HasPrefix(name, "amtmap.") {
					adtStore := adt.WrapStore(ctx, cborStore)
					am, err := adt.AsMap(adtStore, c)
					if err != nil {
						panic(fmt.Sprintf("loading %s failed: %v", name, err))
					}
					val := reflect.New(t.Elem().Elem())
					asUnmarshaller, ok := val.Interface().(cboriface.Unmarshaler)
					if !ok {
						if t.Elem().String() == "*big.Int" {
							var um cbg.CborInt
							asUnmarshaller = &um
						} else {
							panic(fmt.Sprintf("%s must implement CBORUnmarshaler", t.Elem().String()))
						}
					}
					m := reflect.MakeMap(t)
					am.ForEach(asUnmarshaller, func(k string) error {
						m.SetMapIndex(reflect.ValueOf(k), val)
						return nil
					})
					return m.Interface()
				} else if strings.HasPrefix(name, "amtarray.") {
					// Note: this implementation throws away the idx's of the array.
					// it lets you see the more compact array represenation, but may lead to
					// incorrect ordering (and makes it hard to trace the real index)
					adtStore := adt.WrapStore(ctx, cborStore)
					arr, err := adt.AsArray(adtStore, c)
					if err != nil {
						panic(fmt.Sprintf("loading %s failed: %v", name, err))
					}
					val := reflect.New(t.Elem().Elem())
					asUnmarshaller, ok := val.Interface().(cboriface.Unmarshaler)
					if !ok {
						panic(fmt.Sprintf("%s must implement CBORUnmarshaler", t.Elem().String()))
					}
					arrLen := int(arr.Length())
					m := reflect.MakeSlice(t, arrLen, arrLen)
					i := 0
					arr.ForEach(asUnmarshaller, func(idx int64) error {
						m.Index(i).Set(val)
						i += 1
						return nil
					})
					return m.Interface()
				}

				val := reflect.New(t.Elem())
				asUnmarshaller, ok := val.Interface().(cboriface.Unmarshaler)
				block, _ := store.Get(c)
				if !ok {
					return block.RawData()
				}
				if err := asUnmarshaller.UnmarshalCBOR(bytes.NewBuffer(block.RawData())); err != nil {
					panic(fmt.Sprintf("Unable to interpret %s as a %v", c, t.Elem().String()))
				}
				return val.Interface()
			}
		})(name, t)
		options = append(options, cmp.FilterPath(pathFilter(r), cmp.CompareExpander(name, cidEqual, cidAsString, boundFunc)))
	}
	return options
}

func cidEqual(a, b cid.Cid) bool {
	return a.Equals(b)
}

func cidAsString(a cid.Cid) string {
	return a.String()
}

func getInitFor(ctx context.Context, store cbor.IpldStore, root cid.Cid, helper func(act initActor.State) *initActorState) map[string]string {
	inverseMap := make(map[string]string)
	inverseMap[string(builtin.InitActorAddr.Bytes())] = "<InitActor>"
	inverseMap[string(builtin.RewardActorAddr.Bytes())] = "<RewardActor>"
	inverseMap[string(builtin.CronActorAddr.Bytes())] = "<CronActor>"
	inverseMap[string(builtin.StoragePowerActorAddr.Bytes())] = "<StoragePowerActor>"
	inverseMap[string(builtin.StorageMarketActorAddr.Bytes())] = "<StorageMarketActor>"
	inverseMap[string(builtin.VerifiedRegistryActorAddr.Bytes())] = "<VerifiedRegistryActor>"
	inverseMap[string(builtin.BurntFundsActorAddr.Bytes())] = "<BurntFundsActor>"
	tree, err := chainState.LoadStateTree(store, root)
	if err != nil {
		fmt.Printf("failed to load root State Tree for account mapping\n")
		return inverseMap
	}
	initAct, err := tree.GetActor(builtin.InitActorAddr)

	var initState initActor.State
	if err := store.Get(ctx, initAct.Head, &initState); err != nil {
		fmt.Printf("failed to load Init acct @%v: %v\n", initAct.Head, err)
		return inverseMap
	}
	forward := helper(initState)
	for k, v := range forward.ADT {
		address, _ := addr.NewIDAddress(v)
		inverseMap[string(address.Bytes())] = k
	}
	return inverseMap
}

type hamtNode struct {
	Cid  cid.Cid
	Node *hamt.Node
}

func bigIntComparer(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func addressTransformer(a addr.Address) string {
	return a.String()
}

func topFilter(p cmp.Path) bool {
	return !strings.Contains(p.GoString(), "Actor")
}

func filterIn(substr string) func(cmp.Path) bool {
	return func(p cmp.Path) bool {
		return strings.Contains(p.GoString(), substr)
	}
}

func bitfieldTransformer(b *bitfield.BitField) string {
	data, _ := b.MarshalJSON()
	return string(data)
}

func directBitfieldTransformer(b bitfield.BitField) string {
	data, _ := b.MarshalJSON()
	return string(data)
}

type statefulActor struct {
	Type    string
	State   interface{}
	Nonce   uint64
	Balance string
}

type initActorState struct {
	ADT         map[string]uint64
	ADTRoot     string
	NextID      string
	NetworkName string
}
