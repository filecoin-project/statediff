package statediff

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/statediff/types"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/willscott/go-cmp/cmp"
)

type state struct {
	ctx       context.Context
	maxExpand int
	expanded  map[ipld.NodePrototype]int
	store     blockstore.Blockstore
}

func (s *state) matchingNode(a, b ipld.Node) bool {
	if a.Prototype() != b.Prototype() {
		return false
	}
	switch a.Kind() {
	case ipld.Kind_Bool:
		l, _ := a.AsBool()
		r, _ := b.AsBool()
		return l == r
	case ipld.Kind_Bytes:
		l, _ := a.AsBytes()
		r, _ := b.AsBytes()
		return bytes.Equal(l, r)
	case ipld.Kind_Float:
		l, _ := a.AsFloat()
		r, _ := b.AsFloat()
		return l == r
	case ipld.Kind_Int:
		l, _ := a.AsInt()
		r, _ := b.AsInt()
		return l == r
	case ipld.Kind_Invalid:
		return true
	case ipld.Kind_Null:
		return true
	case ipld.Kind_Link:
		l, _ := a.AsLink()
		r, _ := b.AsLink()
		if l.String() != r.String() {
			p := a.Prototype()

			if _, ok := s.expanded[p]; !ok {
				s.expanded[p] = 1
			} else {
				s.expanded[p] = s.expanded[p] + 1
			}
			return false
		}
		return true
	case ipld.Kind_String:
		l, _ := a.AsString()
		r, _ := b.AsString()
		return l == r
	case ipld.Kind_List:
		return false
	case ipld.Kind_Map:
		return false
	}
	return true
}

func (s *state) nodeTerminator(n ipld.Node) interface{} {
	switch n.Kind() {
	case ipld.Kind_Bool:
		v, _ := n.AsBool()
		return v
	case ipld.Kind_Bytes:
		v, _ := n.AsBytes()
		if _, ok := n.(types.Address); ok {
			a, err := address.NewFromBytes(v)
			if err != nil {
				return v
			}
			return a.String()
		} else if _, ok := n.(types.BigInt); ok {
			i := big.NewInt(0)
			i.SetBytes(v)
			return i.String()
		} else if _, ok := n.(types.BitField); ok {
			b, err := bitfield.NewFromBytes(v)
			if err != nil {
				return v
			}
			buf := bytes.NewBuffer([]byte{})
			b.MarshalCBOR(buf)
			return hex.EncodeToString(buf.Bytes())
		}
		return v
	case ipld.Kind_Float:
		v, _ := n.AsFloat()
		return v
	case ipld.Kind_Int:
		v, _ := n.AsInt()
		return v
	case ipld.Kind_Invalid:
		return "<Invalid>"
	case ipld.Kind_Link:
		v, _ := n.AsLink()
		return v.String()
	case ipld.Kind_Null:
		return nil
	case ipld.Kind_String:
		v, _ := n.AsString()
		return v
	case ipld.Kind_List:
		items := make([]ipld.Node, 0)
		v := n.ListIterator()
		for !v.Done() {
			_, val, _ := v.Next()
			items = append(items, val)
		}
		return items
	case ipld.Kind_Map:
		items := make(map[string]ipld.Node)
		v := n.MapIterator()
		for !v.Done() {
			k, val, _ := v.Next()
			ks, err := k.AsString()
			if err != nil {
				ks = fmt.Sprintf("%v", k)
			}
			if _, ok := k.(types.RawAddress); ok {
				a, err := address.NewFromBytes([]byte(ks))
				if err == nil {
					ks = a.String()
				}
			} else if _, ok := k.(types.CidString); ok {
				c := cid.Undef
				if err := c.UnmarshalBinary([]byte(ks)); err == nil {
					ks = c.String()
				}
			}
			items[ks] = val
		}
		return items
	}
	return fmt.Sprintf("%v", n)
}

// Ref is an indirection to help goCmp deal with the multiple transformations involved in resolving an ipld Link
type Ref struct {
	ipld.Node
}

func (s *state) nodeLoader(n ipld.Node) interface{} {
	if n.Kind() == ipld.Kind_Link {
		l, _ := n.AsLink()
		cidlink, ok := l.(cidlink.Link)
		if !ok {
			return l.String()
		}
		if cnt, ok := s.expanded[n.Prototype()]; ok && cnt >= s.maxExpand {
			return l.String()
		}
		c := cidlink.Cid
		proto := LinkDest(n)
		if proto == nil {
			return fmt.Sprintf("Unresolved CID %T: %s", n, l.String())
		}
		builder := proto.NewBuilder()
		err := Load(s.ctx, c, s.store, builder)
		if err != nil {
			return fmt.Sprintf("load of %s failed: %v", c, err)
		}
		dest := builder.Build()
		return Ref{dest}
	}
	// A special interpretation for lotus actors.
	if n.Prototype() == types.Type.LotusActors {
		m := s.nodeTerminator(n).(map[string]ipld.Node)
		typed, err := TypeActorHead(n)
		if err != nil {
			return m
		}
		m["Head"] = typed
		return m
	}

	return s.nodeTerminator(n)
}

func mustCid(n ipld.Node) cid.Cid {
	l, err := n.AsLink()
	if err != nil {
		return cid.Undef
	}
	asCid, ok := l.(cidlink.Link)
	if !ok {
		return cid.Undef
	}
	return asCid.Cid
}

// Diff provides a human readable difference between two ipld dags
func Diff(ctx context.Context, store blockstore.Blockstore, left, right ipld.Node) string {
	s := state{
		ctx:       ctx,
		maxExpand: 5,
		expanded:  make(map[ipld.NodePrototype]int),
		store:     store,
	}
	return cmp.Diff(left, right, cmp.CompareExpander("Node", s.matchingNode, s.nodeTerminator, s.nodeLoader))
}
