package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePower(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("PowerV0State",
		[]schema.StructField{
			schema.SpawnStructField("TotalRawBytePower", "StoragePower", false, false),
			schema.SpawnStructField("TotalBytesCommitted", "StoragePower", false, false),
			schema.SpawnStructField("TotalQualityAdjPower", "StoragePower", false, false),
			schema.SpawnStructField("TotalQABytesCommitted", "StoragePower", false, false),
			schema.SpawnStructField("TotalPledgeCollateral", "TokenAmount", false, false),
			schema.SpawnStructField("ThisEpochRawBytePower", "StoragePower", false, false),
			schema.SpawnStructField("ThisEpochQualityAdjPower", "StoragePower", false, false),
			schema.SpawnStructField("ThisEpochPledgeCollateral", "TokenAmount", false, false),
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
			schema.SpawnStructField("RawBytePower", "StoragePower", false, false),
			schema.SpawnStructField("QualityAdjPower", "StoragePower", false, false),
		},
		schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("V0FilterEstimate",
		[]schema.StructField{
			schema.SpawnStructField("PositionEstimate", "Bytes", false, false), // Big.Int
			schema.SpawnStructField("VelocityEstimate", "Bytes", false, false), // Big.Int
		},
		schema.StructRepresentation_Tuple{}))
}
