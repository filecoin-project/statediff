package lib

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/types"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
)

func AddFields() {
	LotusMessage__type.AddFieldConfig("InterpretedParams", &graphql.Field{
		Name: "InterpretParams",
		Args: graphql.FieldConfigArgument{
			"actorType": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Type: LotusMessageV2Params__type,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ts := p.Source.(types.LotusMessage)
			t := (p.Args["actorType"]).(string)
			at := statediff.LotusType(t)
			paramNode, _, err := statediff.ParseParams(ts.Params.Bytes(), ts.Method.Int(), at)
			if err != nil {
				// Check if it's in the union.
				// We return 'Any' on unkown but they aren't part of the union.
				// In these cases. we'll return instead an empty option to prevent errors.
				if LotusMessageV2Params__type.ResolveType(graphql.ResolveTypeParams{Value: paramNode}) == nil {
					paramNode = types.Type.MessageParamsMinerConstructor__Repr.NewBuilder().Build()
				}
			}
			return paramNode, err
		},
	})

	LotusBlockHeader__type.AddFieldConfig("Time", &graphql.Field{
		Name:        "Time",
		Description: "ISO8601 timestamp of the time at which this block appears. derived from epoch height.",
		Type:        graphql.NewNonNull(graphql.String),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ts := p.Source.(types.LotusBlockHeader)
			t := epochToTime(int(ts.Height.Int()))
			return ISO8601(&t), nil
		},
	})

	List__LinkLotusMessage__type.AddFieldConfig("CountOf", &graphql.Field{
		Name: "CountOf",
		Args: graphql.FieldConfigArgument{
			"method": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Type: graphql.NewNonNull(graphql.Int),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ts := p.Source.(types.List__LinkLotusMessage)
			m := p.Args["method"].(int64)

			store := p.Context.Value(storeCtx).(blockstore.Blockstore)

			b := types.Type.LotusMessage__Repr.NewBuilder()
			c := 0
			li := ts.ListIterator()
			for !li.Done() {
				_, lm, err := li.Next()
				if err != nil {
					return 0, err
				}
				mcid, err := lm.AsLink()
				if err != nil {
					return 0, err
				}
				mcl := mcid.(cidlink.Link)
				if err = statediff.Load(p.Context, mcl.Cid, &statediff.LotusBS{store}, b); err != nil {
					return 0, err
				}
				msg := b.Build()
				b.Reset()

				method, err := msg.LookupByString("Method")
				if err != nil {
					return 0, err
				}
				mn, err := method.AsInt()
				if err != nil {
					return 0, err
				}
				if mn == m {
					c++
				}
			}

			return c, nil
		},
	})

	List__LinkLotusMessage__type.AddFieldConfig("AllOf", &graphql.Field{
		Name: "AllOf",
		Type: graphql.NewList(LotusMessage__type),
		Args: graphql.FieldConfigArgument{
			"method": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			m := p.Args["method"].(int64)
			ts, ok := p.Source.(types.List__LinkLotusMessage)
			if !ok {
				return nil, errNotNode
			}
			it := ts.ListIterator()
			children := make([]ipld.Node, 0)
			for !it.Done() {
				_, node, err := it.Next()
				if err != nil {
					return nil, err
				}

				targetCid, err := node.AsLink()
				if err != nil {
					return nil, err
				}

				if cl, ok := targetCid.(cidlink.Link); ok {
					v := p.Context.Value(nodeLoaderCtxKey)
					if v == nil {
						return cl.Cid, nil
					}
					loader, ok := v.(func(context.Context, cidlink.Link, ipld.NodeBuilder) (ipld.Node, error))
					if !ok {
						return nil, errInvalidLoader
					}

					builder := types.Type.LotusMessage__Repr.NewBuilder()
					n, err := loader(p.Context, cl, builder)
					if err != nil {
						return nil, err
					}
					node = n
				} else {
					return nil, errInvalidLink
				}

				method, err := node.LookupByString("Method")
				if err != nil {
					return nil, err
				}
				mn, err := method.AsInt()
				if err != nil {
					return 0, err
				}
				if mn == m {
					children = append(children, node)
				}
			}
			return children, nil
		},
	})

	Map__LotusActors__type.AddFieldConfig("AllOf", &graphql.Field{
		Name: "AllOf",
		Args: graphql.FieldConfigArgument{
			"type": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Type: graphql.NewList(Map__LotusActors__type__entry),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			match := (p.Args["type"]).(string)
			ts, ok := p.Source.(ipld.Node)
			if !ok {
				return nil, errNotNode
			}
			it := ts.MapIterator()
			children := make([][]ipld.Node, 0)

			for !it.Done() {
				k, v, err := it.Next()
				if err != nil {
					return nil, err
				}
				codeType, err := v.LookupByString("Code")
				if err != nil {
					return nil, err
				}
				codeLink, err := codeType.AsLink()
				if err != nil {
					return nil, err
				}
				if statediff.LotusActorCodes[codeLink.(cidlink.Link).String()] == statediff.LotusType(match) {
					children = append(children, []ipld.Node{k, v})
				}
			}
			return children, nil
		},
	})
}

// RawAddress__type__serialize - and undo the presentation of the "bytes" of address for presentation.
func RawAddress__type__serialize(value interface{}) interface{} {
	switch value := value.(type) {
	case ipld.Node:
		b, err := value.AsString()
		if err != nil {
			return err
		}
		a, err := address.NewFromBytes([]byte(b))
		if err != nil {
			return b
		}
		return a.String()
	default:
		return nil
	}
}

// RawAddress__type__parse - massage incoming strings into their "bytes-in-string" on disk form
func RawAddress__type__parse(value interface{}) interface{} {
	builder := types.Type.RawAddress__Repr.NewBuilder()
	switch v2 := value.(type) {
	case string:
		a, err := address.NewFromString(v2)
		if err != nil {
			// todo: log
			return nil
		}
		builder.AssignString(string(a.Bytes()))
	case *string:
		a, err := address.NewFromString(*v2)
		if err != nil {
			// todo: log
			return nil
		}
		builder.AssignString(string(a.Bytes()))
	default:
		return nil
	}
	return builder.Build()
}

// RawAddress__type__parseLiteral - massage incoming strings into their "bytes-in-string" on disk form
func RawAddress__type__parseLiteral(valueAST ast.Value) interface{} {
	builder := types.Type.RawAddress__Repr.NewBuilder()
	switch valueAST := valueAST.(type) {
	case *ast.StringValue:
		a, err := address.NewFromString(valueAST.Value)
		if err != nil {
			// todo: log
			return nil
		}
		builder.AssignString(string(a.Bytes()))
	default:
		return nil
	}

	return builder.Build()
}

// Address__type__serialize - and undo the presentation of the "bytes" of addres for presentation.
func Address__type__serialize(value interface{}) interface{} {
	switch value := value.(type) {
	case ipld.Node:
		b, err := value.AsBytes()
		if err != nil {
			return err
		}
		a, err := address.NewFromBytes(b)
		if err != nil {
			return b
		}
		return a.String()
	default:
		return nil
	}
}

// Print big-ints in a decimal / human readable string form.
func BigInt__type__serialize(value interface{}) interface{} {
	switch value := value.(type) {
	case ipld.Node:

		b, err := value.AsBytes()
		if err != nil {
			return err
		}
		bi := big.NewInt(0)
		return bi.SetBytes(b).String()
	default:
		return nil
	}
}

// LotusActors__Head__resolve resolves head cid based on code.
func LotusActors__Head__resolve(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(types.LotusActors)
	if !ok {
		return nil, fmt.Errorf("not node")
	}

	targetCid := ts.FieldHead().Link()
	if cl, ok := targetCid.(cidlink.Link); ok {
		v := p.Context.Value(nodeLoaderCtxKey)
		if v == nil {
			return cl.Cid, nil
		}

		loader, ok := v.(func(context.Context, cidlink.Link, ipld.NodeBuilder) (ipld.Node, error))
		if !ok {
			return nil, fmt.Errorf("invalid Loader provided")
		}

		typedTarget, err := statediff.TypeActorHead(ts)
		if err != nil {
			return nil, err
		}

		builder := statediff.LinkDest(typedTarget).NewBuilder()
		return loader(p.Context, cl, builder)
	}
	return nil, fmt.Errorf("Invalid link")
}

// For hamts we use the hamt node rather than the expected builder.
func InitV0State__AddressMap__resolve(p graphql.ResolveParams) (interface{}, error) {
	ts, ok := p.Source.(types.InitV0State)
	if !ok {
		return nil, errNotNode
	}

	targetCid := ts.FieldAddressMap().Link()

	if cl, ok := targetCid.(cidlink.Link); ok {
		v := p.Context.Value(nodeLoaderCtxKey)
		if v == nil {
			return cl.Cid, nil
		}
		loader, ok := v.(func(context.Context, cidlink.Link, ipld.NodeBuilder) (ipld.Node, error))
		if !ok {
			return nil, errInvalidLoader
		}

		builder := statediff.LotusPrototypes[statediff.InitActorAddresses].NewBuilder()
		return loader(p.Context, cl, builder)
	}
	return nil, errInvalidLink
}

// CidString__type__serialize is internally going to be cid bytes, but faked through a string.
func CidString__type__serialize(value interface{}) interface{} {
	switch value := value.(type) {
	case ipld.Node:
		s, err := value.AsString()
		if err != nil {
			return err
		}
		c, err := cid.Parse([]byte(s))
		if err != nil {
			return err
		}
		return c.String()
	default:
		return nil
	}
}

// CidString__type__parse is internally going to be cid bytes, but faked through a string.
func CidString__type__parse(value interface{}) interface{} {
	builder := types.Type.CidString__Repr.NewBuilder()
	switch v2 := value.(type) {
	case string:
		c, err := cid.Decode(v2)
		if err != nil {
			return err
		}
		builder.AssignString(string(c.Bytes()))
	case *string:
		c, err := cid.Decode(*v2)
		if err != nil {
			return err
		}
		builder.AssignString(string(c.Bytes()))
	default:
		return nil
	}
	return builder.Build()
}

// CidString__type__parseLiteral is internally going to be cid bytes, but faked through a string.
func CidString__type__parseLiteral(valueAST ast.Value) interface{} {
	builder := types.Type.CidString__Repr.NewBuilder()
	switch valueAST := valueAST.(type) {
	case *ast.StringValue:
		c, err := cid.Decode(valueAST.Value)
		if err != nil {
			return err
		}
		builder.AssignString(string(c.Bytes()))
	default:
		return nil
	}
	return builder.Build()
}

func ISO8601(t *time.Time) string {
	var tz string
	if z, offset := t.Zone(); z == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d:00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
		t.Second(), tz)
}
