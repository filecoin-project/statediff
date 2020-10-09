package statediff

import (
	"bytes"

	"github.com/ipfs/go-cid"
	hamt "github.com/ipfs/go-hamt-ipld"
	cbornode "github.com/ipfs/go-ipld-cbor"
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
