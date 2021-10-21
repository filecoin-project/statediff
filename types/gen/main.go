package main

import (
	"fmt"
	"os"

	messages "github.com/filecoin-project/statediff/types/gen/messages"
	v0 "github.com/filecoin-project/statediff/types/gen/v0"
	v2 "github.com/filecoin-project/statediff/types/gen/v2"
	v3 "github.com/filecoin-project/statediff/types/gen/v3"
	v4 "github.com/filecoin-project/statediff/types/gen/v4"
	v5 "github.com/filecoin-project/statediff/types/gen/v5"
	v6 "github.com/filecoin-project/statediff/types/gen/v6"

	gengraphql "github.com/ipld/go-ipld-graphql/gen"
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"
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
		schema.SpawnUnionRepresentationKinded(map[ipld.Kind]schema.TypeName{
			ipld.Kind_Bool:   "Bool",
			ipld.Kind_Int:    "Int",
			ipld.Kind_Float:  "Float",
			ipld.Kind_String: "String",
			ipld.Kind_Bytes:  "Bytes",
			ipld.Kind_Map:    "Map",
			ipld.Kind_List:   "List",
			ipld.Kind_Link:   "Link",
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
	v3.Accumulate(ts)
	v4.Accumulate(ts)
	v5.Accumulate(ts)
	v6.Accumulate(ts)
	messages.Accumulate(ts)

	if errs := ts.ValidateGraph(); errs != nil {
		for _, err := range errs {
			fmt.Printf("- %s\n", err)
		}
		panic("not happening")
	}

	gengo.Generate(os.Args[1], "types", ts, adjCfg)
	if len(os.Args) > 2 {
		gengraphql.Generate(os.Args[2], "lib", ts, "types", "github.com/filecoin-project/statediff/types")
	}
}
