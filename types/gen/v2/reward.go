package v2

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateReward(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("RewardV2State",
		[]schema.StructField{
			schema.SpawnStructField("CumsumBaseline", "BigInt", false, false), //Spacetime
			schema.SpawnStructField("CumsumRealized", "BigInt", false, false), //Spacetime
			schema.SpawnStructField("EffectiveNetworkTime", "ChainEpoch", false, false),
			schema.SpawnStructField("EffectiveBaselinePower", "BigInt", false, false), //StoragePower
			schema.SpawnStructField("ThisEpochReward", "BigInt", false, false),        //TokenAmount
			schema.SpawnStructField("ThisEpochRewardSmoothed", "V0FilterEstimate", false, false),
			schema.SpawnStructField("ThisEpochBaselinePower", "BigInt", false, false), //StoragePower
			schema.SpawnStructField("Epoch", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalStoragePowerReward", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("SimpleTotal", "BigInt", false, false),
			schema.SpawnStructField("BaselineTotal", "BigInt", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__RewardV2State", "RewardV2State"))
}
