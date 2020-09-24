package main

import (
	"fmt"

	"github.com/filecoin-project/statediff/types/gen/v0"

	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
	gengo "github.com/ipld/go-ipld-prime/schema/gen/go"
)

func main() {
	ts := schema.TypeSystem{}
	ts.Init()
	adjCfg := &gengo.AdjunctCfg{
		CfgUnionMemlayout: map[schema.TypeName]string{
			"Any": "interface",
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
	v0.Accumulate(ts)
	

	if errs := ts.ValidateGraph(); errs != nil {
		for _, err := range errs {
			fmt.Printf("- %s\n", err)
		}
		panic("not happening")
	}

	gengo.Generate("..", "types", ts, adjCfg)
}
