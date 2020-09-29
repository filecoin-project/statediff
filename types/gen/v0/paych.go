package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePaych(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("PaychV0State",
		[]schema.StructField{
			schema.SpawnStructField("From", "Address", false, false),
			schema.SpawnStructField("To", "Address", false, false),
			schema.SpawnStructField("ToSend", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("SettlingAt", "ChainEpoch", false, false),
			schema.SpawnStructField("MinSettleHeight", "ChainEpoch", false, false),
			schema.SpawnStructField("LaneStates", "Link__PaychV0LaneState", false, false), //AMT<LaneState>
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnMap("Map__PaychV0LaneState", "String", "PaychV0LaneState", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__PaychV0LaneState", "Map__PaychV0LaneState"))

	ts.Accumulate(schema.SpawnStruct("PaychV0LaneState",
		[]schema.StructField{
			schema.SpawnStructField("Redeemed", "BigInt", false, false), // big.int
			schema.SpawnStructField("Nonce", "Int", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
}
