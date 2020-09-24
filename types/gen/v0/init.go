package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateInit(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("CronV0Entry",
		[]schema.StructField{
			schema.SpawnStructField("Receiver", "Address", false, false),
			schema.SpawnStructField("MethodNum", "MethodNum", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnList("List__CronV0Entry",
		"CronV0Entry", false,
	))

	ts.Accumulate(schema.SpawnStruct("InitV0State",
		[]schema.StructField{
			schema.SpawnStructField("AddressMap", "Any", false, false),
			schema.SpawnStructField("NextID", "ActorID", false, false),
			schema.SpawnStructField("NetworkName", "String", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
}
