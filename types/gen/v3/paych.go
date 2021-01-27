package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePaych(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("PaychV3State",
		[]schema.StructField{
			schema.SpawnStructField("From", "Address", false, false),
			schema.SpawnStructField("To", "Address", false, false),
			schema.SpawnStructField("ToSend", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("SettlingAt", "ChainEpoch", false, false),
			schema.SpawnStructField("MinSettleHeight", "ChainEpoch", false, false),
			schema.SpawnStructField("LaneStates", "Link__PaychV3LaneState", false, false), //AMT<LaneState>
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__PaychV3State", "PaychV3State"))

	ts.Accumulate(schema.SpawnMap("Map__PaychV3LaneState", "String", "PaychV0LaneState", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__PaychV3LaneState", "Map__PaychV3LaneState"))

}
