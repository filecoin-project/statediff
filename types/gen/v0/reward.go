package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateReward(ts schema.TypeSystem) {

	ts.Accumulate(schema.SpawnBytes("V0Spacetime"))
	ts.Accumulate(schema.SpawnStruct("RewardV0State",
		[]schema.StructField{
			schema.SpawnStructField("CumsumBaseline", "V0Spacetime", false, false),
			schema.SpawnStructField("CumsumRealized", "V0Spacetime", false, false),
			schema.SpawnStructField("EffectiveNetworkTime", "ChainEpoch", false, false),
			schema.SpawnStructField("EffectiveBaselinePower", "StoragePower", false, false),
			schema.SpawnStructField("ThisEpochReward", "TokenAmount", false, false),
			schema.SpawnStructField("ThisEpochRewardSmoothed", "V0FilterEstimate", false, true),
			schema.SpawnStructField("ThisEpochBaselinePower", "StoragePower", false, false),
			schema.SpawnStructField("Epoch", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalMined", "TokenAmount", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
}
