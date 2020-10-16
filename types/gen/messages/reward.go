package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateReward(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsRewardAwardBlock",
		[]schema.StructField{
			schema.SpawnStructField("Miner", "Address", false, false),
			schema.SpawnStructField("Penalty", "BigInt", false, false),
			schema.SpawnStructField("GasReward", "BigInt", false, false),
			schema.SpawnStructField("WinCount", "Int", false, false),
		}, schema.StructRepresentation_Tuple{}))
}
