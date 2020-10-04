package statediff

import (
	"bytes"
	"context"
	"fmt"

	"github.com/filecoin-project/statediff/types"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/willscott/go-cmp/cmp"
)

func matchingNode(a, b ipld.Node) bool {
	if a.Prototype() != b.Prototype() {
		return false
	}
	switch a.ReprKind() {
	case ipld.ReprKind_Bool:
		l, _ := a.AsBool()
		r, _ := b.AsBool()
		return l == r
	case ipld.ReprKind_Bytes:
		l, _ := a.AsBytes()
		r, _ := b.AsBytes()
		return bytes.Equal(l, r)
	case ipld.ReprKind_Float:
		l, _ := a.AsFloat()
		r, _ := b.AsFloat()
		return l == r
	case ipld.ReprKind_Int:
		l, _ := a.AsInt()
		r, _ := b.AsInt()
		return l == r
	case ipld.ReprKind_Invalid:
		return true
	case ipld.ReprKind_Null:
		return true
	case ipld.ReprKind_Link:
		l, _ := a.AsLink()
		r, _ := b.AsLink()
		return l.String() == r.String()
	case ipld.ReprKind_String:
		l, _ := a.AsString()
		r, _ := b.AsString()
		return l == r
	case ipld.ReprKind_List:
		return false
	case ipld.ReprKind_Map:
		return false
	}
	return true
}

func nodeTerminator(n ipld.Node) interface{} {
	switch n.ReprKind() {
	case ipld.ReprKind_Bool:
		v, _ := n.AsBool()
		return v
	case ipld.ReprKind_Bytes:
		v, _ := n.AsBytes()
		return v
	case ipld.ReprKind_Float:
		v, _ := n.AsFloat()
		return v
	case ipld.ReprKind_Int:
		v, _ := n.AsInt()
		return v
	case ipld.ReprKind_Invalid:
		return "<Invalid>"
	case ipld.ReprKind_Link:
		v, _ := n.AsLink()
		return v.String()
	case ipld.ReprKind_Null:
		return nil
	case ipld.ReprKind_String:
		v, _ := n.AsString()
		return v
	case ipld.ReprKind_List:
		items := make([]ipld.Node, 0)
		v := n.ListIterator()
		for !v.Done() {
			_, val, _ := v.Next()
			items = append(items, val)
		}
		return items
	case ipld.ReprKind_Map:
		items := make(map[string]ipld.Node)
		v := n.MapIterator()
		for !v.Done() {
			k, val, _ := v.Next()
			kv := nodeTerminator(k)
			if as, ok := kv.(string); ok {
				items[as] = val
			} else {
				items[fmt.Sprintf("%v", kv)] = val
			}
		}
		return items
	}
	return fmt.Sprintf("%v", n)
}

// Ref is an indirection to help goCmp deal with the multiple transformations involved in resolving an ipld Link
type Ref struct {
	ipld.Node
}

func makeNodeLoader(ctx context.Context, store blockstore.Blockstore) func(n ipld.Node) interface{} {
	nodeLoader := func(n ipld.Node) interface{} {
		if n.ReprKind() == ipld.ReprKind_Link {
			l, _ := n.AsLink()
			cidlink, ok := l.(cidlink.Link)
			if !ok {
				return l.String()
			}
			c := cidlink.Cid
			proto := LinkDest(n)
			if proto == nil {
				return fmt.Sprintf("Unresolved CID %T: %s", n, l.String())
			}
			builder := proto.NewBuilder()
			err := Load(ctx, c, store, builder)
			if err != nil {
				return fmt.Sprintf("load of %s failed: %v", c, err)
			}
			dest := builder.Build()
			return Ref{dest}
		}
		// A special interpretation for lotus actors.
		if n.Prototype() == types.Type.LotusActors {
			m := nodeTerminator(n).(map[string]ipld.Node)
			typed, err := TypeActorHead(n)
			if err != nil {
				return m
			}
			m["Head"] = typed
			return m
		}

		return nodeTerminator(n)
	}
	return nodeLoader
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
	loader := makeNodeLoader(ctx, store)
	return cmp.Diff(left, right, cmp.CompareExpander("Node", matchingNode, nodeTerminator, loader))
}
