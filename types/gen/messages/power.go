package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePower(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsPowerCreateMiner",
		[]schema.StructField{
			schema.SpawnStructField("Owner", "Address", false, false),
			schema.SpawnStructField("Worker", "Address", false, false),
			schema.SpawnStructField("SealProofType", "Int", false, false), //RegisteredSealProof
			schema.SpawnStructField("Peer", "PeerID", false, false),
			schema.SpawnStructField("Multiaddrs", "List__Multiaddrs", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsPowerUpdateClaimed",
		[]schema.StructField{
			schema.SpawnStructField("RawByteDelta", "BigInt", false, false),
			schema.SpawnStructField("QualityAdjustedDelta", "BigInt", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsPowerEnrollCron",
		[]schema.StructField{
			schema.SpawnStructField("EventEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("Payload", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsPowerCurrentTotal",
		[]schema.StructField{
			schema.SpawnStructField("RawBytePower", "BigInt", false, false),
			schema.SpawnStructField("QualityAdjPower", "BigInt", false, false),
			schema.SpawnStructField("PledgeCollateral", "BigInt", false, false),
			schema.SpawnStructField("QualityAdjPowerSmoothed", "V0FilterEstimate", false, false),
		}, schema.StructRepresentation_Tuple{}))
}
