package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePower(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("PowerV0State",
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
			schema.SpawnStructField("CronEventQueue", "Link", false, false), // Multimap, (HAMT[ChainEpoch]AMT[CronEvent]
			schema.SpawnStructField("FirstCronEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("LastProcessedCronEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("Claims", "Link", false, false), // Map, HAMT[address]Claim
			schema.SpawnStructField("ProofValidationBatch", "Link", false, true),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("PowerV0CronEvent",
		[]schema.StructField{
			schema.SpawnStructField("MinerAddr", "Address", false, false),
			schema.SpawnStructField("CallbackPayload", "Bytes", false, false),
		},
		schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("PowerV0Claim",
		[]schema.StructField{
			schema.SpawnStructField("RawBytePower", "BigInt", false, false),    //StoragePower
			schema.SpawnStructField("QualityAdjPower", "BigInt", false, false), //StoragePower
		},
		schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("V0FilterEstimate",
		[]schema.StructField{
			schema.SpawnStructField("PositionEstimate", "BigInt", false, false),
			schema.SpawnStructField("VelocityEstimate", "BigInt", false, false),
		},
		schema.StructRepresentation_Tuple{}))
}
