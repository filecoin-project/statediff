package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePower(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("PowerV3State",
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
			schema.SpawnStructField("CronEventQueue", "Link__PowerV3CronEvent", false, false), // Multimap, (HAMT[ChainEpoch]AMT[CronEvent]
			schema.SpawnStructField("FirstCronEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("Claims", "Link__PowerV3Claim", false, false), // Map, HAMT[address]Claim
			schema.SpawnStructField("ProofValidationBatch", "Link", false, true),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__PowerV3State", "PowerV3State"))

	ts.Accumulate(schema.SpawnMap("Map__PowerV3Claim", "RawAddress", "PowerV2Claim", true))
	ts.Accumulate(schema.SpawnLinkReference("Link__PowerV3Claim", "Map__PowerV3Claim"))

	ts.Accumulate(schema.SpawnMap("Map__PowerV3CronEvent", "String", "PowerV0CronEvent", true))
	ts.Accumulate(schema.SpawnMap("Multimap__PowerV3CronEvent", "String", "Map__PowerV3CronEvent", true)) //ChainEpoch keys
	ts.Accumulate(schema.SpawnLinkReference("Link__PowerV3CronEvent", "Multimap__PowerV3CronEvent"))
}
