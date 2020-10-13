package statediff

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
)

// GetStateRoot gets a stateroot node from a cid
// The cid can be the block of a tipset, the stateroot itself, or the
// versioned stateroot
func GetStateRoot(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (ipld.Node, error) {
	blk, err := Transform(ctx, c, store, "tipset")
	if err == nil {
		srLink, err := blk.LookupByString("ParentStateRoot")
		if err != nil {
			return nil, err
		}
		srAsLink, err := srLink.AsLink()
		if err != nil {
			return nil, err
		}
		srAsCidLink, ok := srAsLink.(cidlink.Link)
		if !ok {
			return nil, fmt.Errorf("not a cid link: %v", srAsLink)
		}
		return getStateRootMaybeVersioned(ctx, srAsCidLink.Cid, store)
	}
	return getStateRootMaybeVersioned(ctx, c, store)
}

func getStateRootMaybeVersioned(ctx context.Context, c cid.Cid, store blockstore.Blockstore) (ipld.Node, error) {
	sr, err := Transform(ctx, c, store, "stateRoot")
	if err == nil {
		return sr, nil
	}
	vsr, err2 := Transform(ctx, c, store, "versionedStateRoot")
	if err2 != nil {
		return nil, fmt.Errorf("cid was not a stateroot: %v / %v", err, err2)
	}
	srLink, err := vsr.LookupByString("Actors")
	if err != nil {
		return nil, err
	}
	srAsLink, err := srLink.AsLink()
	if err != nil {
		return nil, err
	}
	srAsCidLink, ok := srAsLink.(cidlink.Link)
	if !ok {
		return nil, fmt.Errorf("not a cid link: %v", srAsLink)
	}
	return getStateRootMaybeVersioned(ctx, srAsCidLink.Cid, store)
}
