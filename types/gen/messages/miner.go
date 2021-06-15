package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMiner(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerConstructor",
		[]schema.StructField{},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerChangeAddress",
		[]schema.StructField{
			schema.SpawnStructField("NewWorker", "Address", false, false),
			schema.SpawnStructField("NewControlAddrs", "List__Address", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerChangeMultiaddrs",
		[]schema.StructField{
			schema.SpawnStructField("NewMultiaddrs", "List__Multiaddrs", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerChangePeerID",
		[]schema.StructField{
			schema.SpawnStructField("NewID", "PeerID", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerSubmitWindowedPoSt",
		[]schema.StructField{
			schema.SpawnStructField("Deadline", "Int", false, false),
			schema.SpawnStructField("Partitions", "List__MinerPostPartition", false, false),
			schema.SpawnStructField("Proofs", "List__MinerPoStProof", false, false),
			schema.SpawnStructField("ChainCommitEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("ChainCommitRand", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnStruct("MinerPostPartition",
		[]schema.StructField{
			schema.SpawnStructField("Index", "Int", false, false),
			schema.SpawnStructField("Skipped", "BitField", false, false),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__MinerPostPartition", "MinerPostPartition", false))
	ts.Accumulate(schema.SpawnStruct("MinerPostProof",
		[]schema.StructField{
			schema.SpawnStructField("PoStProof", "Int", false, false),
			schema.SpawnStructField("ProofBytes", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__MinerPoStProof", "MinerPostProof", false))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerProveCommitSector",
		[]schema.StructField{
			schema.SpawnStructField("SectorNumber", "SectorNumber", false, false),
			schema.SpawnStructField("Proof", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerCheckSectorProven",
		[]schema.StructField{
			schema.SpawnStructField("SectorNumber", "SectorNumber", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerConfirmSectorProofs",
		[]schema.StructField{
			schema.SpawnStructField("Sectors", "List__SectorNumber", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerExtendSectorExpiration",
		[]schema.StructField{
			schema.SpawnStructField("Extension", "List__MinerExpirationExtend", false, false),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__MinerExpirationExtend", "MinerExpirationExtend", false))
	ts.Accumulate(schema.SpawnStruct("MinerExpirationExtend",
		[]schema.StructField{
			schema.SpawnStructField("Deadline", "Int", false, false),
			schema.SpawnStructField("Partition", "Int", false, false),
			schema.SpawnStructField("Sectors", "BitField", false, false),
			schema.SpawnStructField("NewExpiration", "ChainEpoch", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerTerminateSectors",
		[]schema.StructField{
			schema.SpawnStructField("Terminations", "List__MinerTerminationDecl", false, false),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__MinerTerminationDecl", "MinerTerminationDecl", false))
	ts.Accumulate(schema.SpawnStruct("MinerTerminationDecl",
		[]schema.StructField{
			schema.SpawnStructField("Deadline", "Int", false, false),
			schema.SpawnStructField("Partition", "Int", false, false),
			schema.SpawnStructField("Sectors", "BitField", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerDeclareFaults",
		[]schema.StructField{
			schema.SpawnStructField("Faults", "List__MinerTerminationDecl", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerDeclareFaultsRecovered",
		[]schema.StructField{
			schema.SpawnStructField("Recoveries", "List__MinerTerminationDecl", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerCompactPartitions",
		[]schema.StructField{
			schema.SpawnStructField("Deadline", "Int", false, false),
			schema.SpawnStructField("Partitions", "BitField", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerCompactSectorNumbers",
		[]schema.StructField{
			schema.SpawnStructField("MaskSectorNumbers", "BitField", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("ApplyRewardParams",
		[]schema.StructField{
			schema.SpawnStructField("Reward", "BigInt", false, false),  // TokenAmount
			schema.SpawnStructField("Penalty", "BigInt", false, false), // TokenAmount
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerReportFault",
		[]schema.StructField{
			schema.SpawnStructField("BlockHeader1", "Bytes", false, false),
			schema.SpawnStructField("BlockHeader2", "Bytes", false, false),
			schema.SpawnStructField("BlockHeaderExtra", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerWithdrawBalance",
		[]schema.StructField{
			schema.SpawnStructField("AmountRequested", "BigInt", false, false), // TokenAmount
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerDeferredCron",
		[]schema.StructField{
			schema.SpawnStructField("EventType", "Int", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMinerDisputeWindowedPoSt",
		[]schema.StructField{
			schema.SpawnStructField("Deadline", "Int", false, false),
			schema.SpawnStructField("PoStIndex", "Int", false, false),
		}, schema.StructRepresentation_Tuple{}))
}
