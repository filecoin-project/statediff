package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePaych(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("PaychV0State",
		[]schema.StructField{
			schema.SpawnStructField("To", "Address", false, false),
			schema.SpawnStructField("ToSend", "TokenAmount", false, false),
			schema.SpawnStructField("SettlingAt", "ChainEpoch", false, false),
			schema.SpawnStructField("MinSettleHeight", "ChainEpoch", false, false),
			schema.SpawnStructField("LateStates", "Link", false, false), //AMT<LaneState>
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("PaychV0LaneState",
		[]schema.StructField{
			schema.SpawnStructField("Redeemed", "Bytes", false, false), // big.int
			schema.SpawnStructField("Nonce", "Int", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
}
