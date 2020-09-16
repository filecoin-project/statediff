package statediff

import (
	"context"
	"sync"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/lib/blockstore"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
)

type proxyingBlockstore struct {
	ctx context.Context
	api api.FullNode

	bsLock sync.RWMutex
	blockstore.Blockstore
}

func (pb *proxyingBlockstore) Get(cid cid.Cid) (blocks.Block, error) {
	pb.bsLock.RLock()
	if block, err := pb.Blockstore.Get(cid); err == nil {
		pb.bsLock.RUnlock()
		return block, err
	}
	pb.bsLock.RUnlock()

	// fmt.Printf("fetching cid via rpc: %v\n", cid)
	item, err := pb.api.ChainReadObj(pb.ctx, cid)
	if err != nil {
		return nil, err
	}
	block, err := blocks.NewBlockWithCid(item, cid)
	if err != nil {
		return nil, err
	}

	pb.bsLock.Lock()
	defer pb.bsLock.Unlock()
	err = pb.Blockstore.Put(block)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func StoreFor(ctx context.Context, client api.FullNode) blockstore.Blockstore {
	ds := ds.NewMapDatastore()

	bs := &proxyingBlockstore{
		ctx:        ctx,
		api:        client,
		Blockstore: blockstore.NewBlockstore(ds),
	}
	return bs
}
