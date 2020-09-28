package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateReward(ts schema.TypeSystem) {

	//ts.Accumulate(schema.SpawnBytes("V0Spacetime"))
	ts.Accumulate(schema.SpawnStruct("RewardV0State",
		[]schema.StructField{
			schema.SpawnStructField("CumsumBaseline", "BigInt", false, false), //Spacetime
			schema.SpawnStructField("CumsumRealized", "BigInt", false, false), //Spacetime
			schema.SpawnStructField("EffectiveNetworkTime", "ChainEpoch", false, false),
			schema.SpawnStructField("EffectiveBaselinePower", "BigInt", false, false), //StoragePower
			schema.SpawnStructField("ThisEpochReward", "BigInt", false, false),        //TokenAmount
			schema.SpawnStructField("ThisEpochRewardSmoothed", "V0FilterEstimate", false, true),
			schema.SpawnStructField("ThisEpochBaselinePower", "BigInt", false, false), //StoragePower
			schema.SpawnStructField("Epoch", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalMined", "BigInt", false, false), //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))
}
