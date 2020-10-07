package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateCron(ts schema.TypeSystem) {
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

	ts.Accumulate(schema.SpawnStruct("CronV0State",
		[]schema.StructField{
			schema.SpawnStructField("Entries", "List__CronV0Entry", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__CronV0State", "CronV0State"))
}
