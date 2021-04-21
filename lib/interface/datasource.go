package iface

import (
	"context"

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/ipfs/go-cid"
)

// Datasource is the interface for loading data to share with stateexplorer when needed.
type Datasource interface {
	Store() blockstore.Blockstore
	Head(ctx context.Context) cid.Cid
	CidAtHeight(ctx context.Context, h int64) (cid.Cid, error)
	Reload() error
}
