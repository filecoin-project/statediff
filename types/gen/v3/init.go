package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateInit(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("InitV3State",
		[]schema.StructField{
			schema.SpawnStructField("AddressMap", "Link__V3MapActorID", false, false),
			schema.SpawnStructField("NextID", "ActorID", false, false),
			schema.SpawnStructField("NetworkName", "String", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__InitV3State", "InitV3State"))
	ts.Accumulate(schema.SpawnLinkReference("Link__V3MapActorID", "Map__V3ActorID"))
	ts.Accumulate(schema.SpawnMap("Map__V3ActorID", "RawAddress", "ActorID", false))
}
