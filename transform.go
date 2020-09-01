package statediff

import (
	"context"
	"fmt"

	addr "github.com/filecoin-project/go-address"
	cbor "github.com/ipfs/go-ipld-cbor"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/ipfs/go-cid"
	hamt "github.com/ipfs/go-hamt-ipld"
	"github.com/ipfs/go-ipfs-blockstore"

	lotusTypes "github.com/filecoin-project/lotus/chain/types"

	accountActor "github.com/filecoin-project/specs-actors/actors/builtin/account"
	cronActor "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	initActor "github.com/filecoin-project/specs-actors/actors/builtin/init"
	marketActor "github.com/filecoin-project/specs-actors/actors/builtin/market"
	storageMinerActor "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	storagePowerActor "github.com/filecoin-project/specs-actors/actors/builtin/power"
	rewardActor "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	verifiedRegistryActor "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
)

type LotusType string
const (
	LotusTypeTipset LotusType = "tipset"
	LotusTypeStateroot LotusType = "stateRoot"
	AccountActorState LotusType = "accountActor"
	CronActorState LotusType = "cronActor"
	InitActorState LotusType = "initActor"
	InitActorAddresses LotusType = "initActorAddresses"
	MarketActorState LotusType = "marketActor"
	StorageMinerActorState LotusType = "storageMinerActor"
	StoragePowerActorState LotusType = "storagePowerActor"
	RewardActorState LotusType = "rewardActor"
	VerifiedRegistryActorState LotusType = "verifiedRegistryActor"
)

// Transform will unmarshal cbor data based on a provided type hint.
func Transform(ctx context.Context, c cid.Cid, store blockstore.Blockstore, as string) (interface{}, error) {
	// First select types which do their own store loading.
	switch LotusType(as) {
	case LotusTypeStateroot:
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
			m[fmt.Sprintf("%x",k)] = &actor
			return nil
		})
		return m, nil
	case InitActorAddresses:
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
	case StorageMinerActorState:
		dest := storageMinerActor.State{}
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
	default:
		var dest interface{}
		err := cbor.DecodeInto(data, &dest)
		return dest, err
	}
}
