package statediff

import (
	"bytes"

	hamt "github.com/filecoin-project/go-hamt-ipld/v2"
	hamtv3 "github.com/filecoin-project/go-hamt-ipld/v3"
	"github.com/ipfs/go-cid"
	cbornode "github.com/ipfs/go-ipld-cbor"
	cbg "github.com/whyrusleeping/cbor-gen"
)

// this is like `hamt.ForEach` but ignores blocks within the hamt that fail to load.
func degradedForEach(n *hamt.Node, cb func(k string, val interface{}) error, store *cbornode.BasicIpldStore) error {
	for _, p := range n.Pointers {
		for _, e := range p.KVs {
			if err := cb(string(e.Key), e.Value); err != nil {
				return err
			}
		}
		if p.Link != cid.Undef {
			if nxt, err := store.Blocks.Get(p.Link); err == nil {
				subN := hamt.Node{}
				subN.UnmarshalCBOR(bytes.NewBuffer(nxt.RawData()))
				if err := degradedForEach(&subN, cb, store); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func degradedForEachV3(n *hamtv3.Node, cb func(k string, val *cbg.Deferred) error, store *cbornode.BasicIpldStore) error {
	for _, p := range n.Pointers {
		for _, e := range p.KVs {
			if err := cb(string(e.Key), e.Value); err != nil {
				return err
			}
		}
		if p.Link != cid.Undef {
			if nxt, err := store.Blocks.Get(p.Link); err == nil {
				subN := hamtv3.Node{}
				subN.UnmarshalCBOR(bytes.NewBuffer(nxt.RawData()))
				if err := degradedForEachV3(&subN, cb, store); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
