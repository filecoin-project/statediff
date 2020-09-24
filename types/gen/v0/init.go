package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateInit(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("InitV0State",
		[]schema.StructField{
			schema.SpawnStructField("AddressMap", "Link", false, false),
			schema.SpawnStructField("NextID", "ActorID", false, false),
			schema.SpawnStructField("NetworkName", "String", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
}
