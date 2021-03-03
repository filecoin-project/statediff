package lib

import (
	"context"
	"fmt"
	"sync"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	abitypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/types"
	"github.com/ipfs/go-cid"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/urfave/cli/v2"
)

// Datasource is the interface for loading data to share with stateexplorer when needed.
type Datasource interface {
	Store() blockstore.Blockstore
	Head(ctx context.Context) cid.Cid
	CidAtHeight(ctx context.Context, h int64) (cid.Cid, error)
	Reload() error
}

type lazyDs struct {
	ctx *cli.Context

	client api.FullNode
	head   statediff.StateRootFunc
	store  blockstore.Blockstore

	// when loaded from a car snapshot there's no gettipsetbyheight metadata index.
	// we build it into tipsetMap in those cases.
	tipsetMap   map[int64]cid.Cid
	tipsetMutex sync.RWMutex
}

var lazyInst *lazyDs

func (l *lazyDs) Store() blockstore.Blockstore {
	return l.store
}

func (l *lazyDs) Head(ctx context.Context) cid.Cid {
	h := l.head(ctx)
	if len(h) > 0 {
		return h[0]
	}
	return cid.Undef
}

func (l *lazyDs) Reload() error {
	client, head, store, err := GetBlockstore(l.ctx)
	if err != nil {
		return err
	}
	l.client = client
	l.head = head
	l.store = store
	return nil
}

func (l *lazyDs) index(ctx context.Context) error {
	fmt.Printf("running jank tipset index\n")
	n := l.head(ctx)
	if len(n) == 0 {
		return fmt.Errorf("no head func for indexing")
	}
	l.tipsetMutex.Lock()
	defer l.tipsetMutex.Unlock()
	if l.tipsetMap == nil {
		l.tipsetMap = make(map[int64]cid.Cid)
	}
	c := n[0]
	cooldown := -1
	for true {
		p := types.Type.LotusBlockHeader__Repr.NewBuilder()
		if err := statediff.Load(ctx, c, l.store, p); err != nil {
			return fmt.Errorf("no header @ %s: %v", c, err)
		}
		n := p.Build()
		h, err := n.LookupByString("Height")
		if err != nil {
			return fmt.Errorf("no Height on header: %v", err)
		}
		hn, err := h.AsInt()
		if err != nil {
			return fmt.Errorf("height could not be interpreted as int: %v", err)
		}
		// re-do recent index when we're adding new things to handle recent reorgs
		if _, ok := l.tipsetMap[hn]; ok {
			if cooldown < 0 {
				cooldown = 10
			}
			cooldown--
			if cooldown == 0 {
				return nil
			}
		}
		l.tipsetMap[hn] = c
		if hn == 1 {
			return nil
		}

		parents, err := n.LookupByString("Parents")
		if err != nil {
			return fmt.Errorf("no parents on %s: %v", c, err)
		}
		_, parentCid, err := parents.ListIterator().Next()
		if err != nil {
			return fmt.Errorf("parents on %s empty: %v", c, err)
		}
		pcc, err := parentCid.AsLink()
		if err != nil {
			return fmt.Errorf("parents on %s not link: %v", c, err)
		}
		c = pcc.(cidlink.Link).Cid
	}
	return nil
}

func (l *lazyDs) CidAtHeight(ctx context.Context, h int64) (cid.Cid, error) {
	l.tipsetMutex.RLock()
	if l.tipsetMap == nil && l.client != nil {
		l.tipsetMutex.RUnlock()
		ts, err := l.client.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(h), abitypes.TipSetKey{})
		if err != nil {
			if _, err := l.client.ChainHead(ctx); err != nil {
				return cid.Undef, l.Reload()
			}
			return cid.Undef, err
		}
		return ts.Cids()[0], nil
	} else if l.tipsetMap != nil {
		c, ok := l.tipsetMap[h]
		l.tipsetMutex.RUnlock()
		if !ok {
			if err := l.index(ctx); err != nil {
				return cid.Undef, err
			}
			l.tipsetMutex.RLock()
			defer l.tipsetMutex.RUnlock()
			return l.tipsetMap[h], nil
		}
		return c, nil
	}
	l.tipsetMutex.RUnlock()
	return cid.Undef, nil
}

func Lazy(c *cli.Context) (Datasource, error) {
	if lazyInst != nil {
		return lazyInst, nil
	}
	var err error
	client, head, store, err := GetBlockstore(c)
	if err != nil {
		return nil, err
	}

	if dsc, ok := store.(Datasource); ok {
		return dsc, nil
	}

	lazyInst = &lazyDs{
		ctx:    c,
		client: client,
		head:   head,
		store:  store,
	}

	if client == nil {
		if err := lazyInst.index(c.Context); err != nil {
			return nil, err
		}
	}
	return lazyInst, nil
}
