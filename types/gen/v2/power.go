package v2

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePower(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("PowerV2State",
		[]schema.StructField{
			schema.SpawnStructField("TotalRawBytePower", "BigInt", false, false),         //StoragePower
			schema.SpawnStructField("TotalBytesCommitted", "BigInt", false, false),       //StoragePower
			schema.SpawnStructField("TotalQualityAdjPower", "BigInt", false, false),      //StoragePower
			schema.SpawnStructField("TotalQABytesCommitted", "BigInt", false, false),     //StoragePower
			schema.SpawnStructField("TotalPledgeCollateral", "BigInt", false, false),     //TokenAmount
			schema.SpawnStructField("ThisEpochRawBytePower", "BigInt", false, false),     //StoragePower
			schema.SpawnStructField("ThisEpochQualityAdjPower", "BigInt", false, false),  //StoragePower
			schema.SpawnStructField("ThisEpochPledgeCollateral", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("ThisEpochQAPowerSmoothed", "V0FilterEstimate", false, true),
			schema.SpawnStructField("MinerCount", "Int", false, false),
			schema.SpawnStructField("MinerAboveMinPowerCount", "Int", false, false),
			schema.SpawnStructField("CronEventQueue", "Link__PowerV0CronEvent", false, false), // Multimap, (HAMT[ChainEpoch]AMT[CronEvent]
			schema.SpawnStructField("FirstCronEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("Claims", "Link__PowerV2Claim", false, false), // Map, HAMT[address]Claim
			schema.SpawnStructField("ProofValidationBatch", "Link", false, true),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__PowerV2State", "PowerV2State"))

	ts.Accumulate(schema.SpawnStruct("PowerV2Claim",
		[]schema.StructField{
			schema.SpawnStructField("SealProofType", "Int", false, false),      //RegisteredSealProof
			schema.SpawnStructField("RawBytePower", "BigInt", false, false),    //StoragePower
			schema.SpawnStructField("QualityAdjPower", "BigInt", false, false), //StoragePower
		},
		schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnMap("Map__PowerV2Claim", "RawAddress", "PowerV2Claim", true))
	ts.Accumulate(schema.SpawnLinkReference("Link__PowerV2Claim", "Map__PowerV2Claim"))

}
