package lib

import (
	"context"
	"fmt"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/types"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
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

		loader, ok := v.(func(context.Context, cidlink.Link, ipld.NodeBuilder) error)
		if !ok {
			return nil, fmt.Errorf("invalid Loader provided")
		}

		typedTarget, err := statediff.TypeActorHead(ts)
		if err != nil {
			return nil, err
		}

		builder := statediff.LinkDest(typedTarget).NewBuilder()
		if err := loader(p.Context, cl, builder); err != nil {
			return nil, err
		}
		return builder.Build(), nil
	}
	return nil, fmt.Errorf("Invalid link")

}
