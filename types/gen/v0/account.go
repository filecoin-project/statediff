package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateAccount(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("AccountV0State",
		[]schema.StructField{
			schema.SpawnStructField("Address", "Address", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__AccountV0State", "AccountV0State"))
}
