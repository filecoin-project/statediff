package main

import (
	"fmt"
	"os"

	messages "github.com/filecoin-project/statediff/types/gen/messages"
	v0 "github.com/filecoin-project/statediff/types/gen/v0"
	v2 "github.com/filecoin-project/statediff/types/gen/v2"

	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"
	gengraphql "github.com/ipld/go-ipld-prime/schema/gen/graphql"
	gengraphqlserver "github.com/ipld/go-ipld-prime/schema/gen/graphql/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Must specify destination directory")
		os.Exit(1)
	}

	ts := schema.TypeSystem{}
	ts.Init()
	adjCfg := &gengo.AdjunctCfg{
		CfgUnionMemlayout: map[schema.TypeName]string{
			"Any":            "interface",
			"LotusActorHead": "interface",
		},
	}

	// Prelude.  (This is boilerplate; it will be injected automatically by the schema libraries in the future, but isn't yet.)
	ts.Accumulate(schema.SpawnBool("Bool"))
	ts.Accumulate(schema.SpawnInt("Int"))
	ts.Accumulate(schema.SpawnFloat("Float"))
	ts.Accumulate(schema.SpawnString("String"))
	ts.Accumulate(schema.SpawnBytes("Bytes"))
	ts.Accumulate(schema.SpawnMap("Map",
		"String", "Any", true,
	))
	ts.Accumulate(schema.SpawnList("List",
		"Any", true,
	))
	ts.Accumulate(schema.SpawnLink("Link"))
	ts.Accumulate(schema.SpawnUnion("Any",
		[]schema.TypeName{
			"Bool",
			"Int",
			"Float",
			"String",
			"Bytes",
			"Map",
			"List",
			"Link",
		},
		schema.SpawnUnionRepresentationKinded(map[ipld.ReprKind]schema.TypeName{
			ipld.ReprKind_Bool:   "Bool",
			ipld.ReprKind_Int:    "Int",
			ipld.ReprKind_Float:  "Float",
			ipld.ReprKind_String: "String",
			ipld.ReprKind_Bytes:  "Bytes",
			ipld.ReprKind_Map:    "Map",
			ipld.ReprKind_List:   "List",
			ipld.ReprKind_Link:   "Link",
		}),
	))

	accumulateABI(ts)
	ts.Accumulate(schema.SpawnBytes("Address"))
	ts.Accumulate(schema.SpawnString("RawAddress"))
	ts.Accumulate(schema.SpawnString("CidString"))
	ts.Accumulate(schema.SpawnList("List__Address", "Address", true))
	ts.Accumulate(schema.SpawnList("List__Link", "Link", true))
	ts.Accumulate(schema.SpawnBytes("BigInt"))
	accumulateCrypto(ts)
	accumulateLotus(ts)
	v0.Accumulate(ts)
	v2.Accumulate(ts)
	messages.Accumulate(ts)

	if errs := ts.ValidateGraph(); errs != nil {
		for _, err := range errs {
			fmt.Printf("- %s\n", err)
		}
		panic("not happening")
	}

	gengo.Generate(os.Args[1], "types", ts, adjCfg)
	if len(os.Args) > 2 {
		gengraphql.Generate(os.Args[2], ts)
		gengraphqlserver.Generate(os.Args[2], "lib", ts, "types", "github.com/filecoin-project/statediff/types")
	}
}
