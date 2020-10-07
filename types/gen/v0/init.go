package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateInit(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("InitV0State",
		[]schema.StructField{
			schema.SpawnStructField("AddressMap", "Link__MapActorID", false, false),
			schema.SpawnStructField("NextID", "ActorID", false, false),
			schema.SpawnStructField("NetworkName", "String", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__InitV0State", "InitV0State"))
	ts.Accumulate(schema.SpawnLinkReference("Link__MapActorID", "Map__ActorID"))
	ts.Accumulate(schema.SpawnMap("Map__ActorID", "RawAddress", "ActorID", false))
}
