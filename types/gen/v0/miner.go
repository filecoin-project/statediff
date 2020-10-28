package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMiner(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MinerV0State",
		[]schema.StructField{
			schema.SpawnStructField("Info", "Link__MinerV0Info", false, false),
			schema.SpawnStructField("PreCommitDeposits", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("LockedFunds", "BigInt", false, false),       //TokenAmount
			schema.SpawnStructField("VestingFunds", "Link__MinerV0VestingFunds", false, false),
			// /schema.SpawnStructField("FeeDebt", "BigInt", false, false),                 //TokenAmount
			schema.SpawnStructField("InitialPledgeRequirement", "BigInt", false, false),                   //TokenAmount
			schema.SpawnStructField("PreCommittedSectors", "Link__MinerV0SectorPreCommits", false, false), // Map, HAMT[SectorNumber]SectorPreCommitOnChainInfo
			schema.SpawnStructField("PreCommittedSectorsExpiry", "Link", false, false),                    // BitFieldQueue (AMT[Epoch]*BitField)
			schema.SpawnStructField("AllocatedSectors", "Link__BitField", false, false),
			schema.SpawnStructField("Sectors", "Link__MinerV0SectorInfo", false, false), // Array, AMT[SectorNumber]SectorOnChainInfo (sparse)
			schema.SpawnStructField("ProvingPeriodStart", "ChainEpoch", false, false),
			schema.SpawnStructField("CurrentDeadline", "Int", false, false),
			schema.SpawnStructField("Deadlines", "Link__MinerV0Deadlines", false, false),
			schema.SpawnStructField("EarlyTerminations", "BitField", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0State", "MinerV0State"))

	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0Info",
		"MinerV0Info",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV0Info",
		[]schema.StructField{
			schema.SpawnStructField("Owner", "Address", false, false),
			schema.SpawnStructField("Worker", "Address", false, false),
			schema.SpawnStructField("ControlAddresses", "List__Address", false, true),
			schema.SpawnStructField("PendingWorkerKey", "MinerV0WorkerChangeKey", false, true),
			schema.SpawnStructField("PeerId", "PeerID", false, false),
			schema.SpawnStructField("Multiaddrs", "List__Multiaddrs", false, true),
			schema.SpawnStructField("SealProofType", "Int", false, false),
			schema.SpawnStructField("SectorSize", "SectorSize", false, false),
			schema.SpawnStructField("WindowPoStPartitionSectors", "Int", false, false),
			//schema.SpawnStructField("ConsensusFaultElapsed", "ChainEpoch", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MinerV0WorkerChangeKey",
		[]schema.StructField{
			schema.SpawnStructField("NewWorker", "Address", false, false),
			schema.SpawnStructField("EffectiveAt", "ChainEpoch", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0VestingFunds",
		"MinerV0VestingFunds",
	))

	ts.Accumulate(schema.SpawnStruct("MinerV0VestingFunds",
		[]schema.StructField{
			schema.SpawnStructField("Funds", "List__MinerV0VestingFund", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnList("List__MinerV0VestingFund", "MinerV0VestingFund", true))
	ts.Accumulate(schema.SpawnStruct("MinerV0VestingFund",
		[]schema.StructField{
			schema.SpawnStructField("Epoch", "ChainEpoch", false, false),
			schema.SpawnStructField("Amount", "BigInt", false, false), //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0Deadlines",
		"MinerV0Deadlines",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV0Deadlines",
		[]schema.StructField{
			schema.SpawnStructField("Due", "List__MinerV0DeadlineLink", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnList("List__MinerV0DeadlineLink", "Link__MinerV0Deadline", true))
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0Deadline",
		"MinerV0Deadline",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV0Deadline",
		[]schema.StructField{
			schema.SpawnStructField("Partitions", "Link__MinerV0Partition", false, false), // AMT[PartitionNumber]Partition
			schema.SpawnStructField("ExpirationEpochs", "Link", false, false),             // AMT[ChainEpoch]BitField
			schema.SpawnStructField("PostSubmissions", "BitField", false, false),
			schema.SpawnStructField("EarlyTerminations", "BitField", false, false),
			schema.SpawnStructField("LiveSectors", "Int", false, false),
			schema.SpawnStructField("TotalSectors", "Int", false, false),
			schema.SpawnStructField("FaultyPower", "MinerV0PowerPair", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnStruct("MinerV0PowerPair",
		[]schema.StructField{
			schema.SpawnStructField("Raw", "BigInt", false, false), //StoragePower
			schema.SpawnStructField("QA", "BigInt", false, false),  //StoragePower
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MinerV0Partition",
		[]schema.StructField{
			schema.SpawnStructField("Sectors", "BitField", false, false),
			//schema.SpawnStructField("Unproven", "BitField", false, false),
			schema.SpawnStructField("Faults", "BitField", false, false),
			schema.SpawnStructField("Recoveries", "BitField", false, false),
			schema.SpawnStructField("Terminated", "BitField", false, false),
			schema.SpawnStructField("ExpirationsEpochs", "Link__MinerV0ExpirationSet", false, false), // AMT[ChainEpoch]ExpirationSet
			schema.SpawnStructField("EarlyTerminated", "Link", false, false),                         // AMT[ChainEpoch]BitField
			schema.SpawnStructField("LivePower", "MinerV0PowerPair", false, false),
			//schema.SpawnStructField("UnprovenPower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("FaultyPower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("RecoveringPower", "MinerV0PowerPair", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__MinerV0Partition", "String", "MinerV0Partition", true)) // Key=PartitionNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0Partition", "Map__MinerV0Partition"))

	ts.Accumulate(schema.SpawnStruct("MinerV0ExpirationSet",
		[]schema.StructField{
			schema.SpawnStructField("OnTimeSectors", "BitField", false, false),
			schema.SpawnStructField("EarlySectors", "BitField", false, false),
			schema.SpawnStructField("OnTimePledge", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("ActivePower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("FaultyPower", "MinerV0PowerPair", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__MinerV0ExpirationSet", "String", "MinerV0ExpirationSet", true)) // Key=ChainEpcoh
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0ExpirationSet", "Map__MinerV0ExpirationSet"))

	ts.Accumulate(schema.SpawnStruct("MinerV0SectorPreCommitOnChainInfo",
		[]schema.StructField{
			schema.SpawnStructField("Info", "MinerV0SectorPreCommitInfo", false, false),
			schema.SpawnStructField("PreCommitDeposit", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("PreCommitEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("DealWeight", "BigInt", false, false),         //DealWeight
			schema.SpawnStructField("VerifiedDealWeight", "BigInt", false, false), //DealWeight
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__SectorPreCommitOnChainInfo", "String", "MinerV0SectorPreCommitOnChainInfo", true)) // Key=SectorNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0SectorPreCommits", "Map__SectorPreCommitOnChainInfo"))

	ts.Accumulate(schema.SpawnStruct("MinerV0SectorPreCommitInfo",
		[]schema.StructField{
			schema.SpawnStructField("SealProof", "Int", false, false), // RegisteredSealProof
			schema.SpawnStructField("SectorNumber", "SectorNumber", false, false),
			schema.SpawnStructField("SealedCID", "Link", false, false),
			schema.SpawnStructField("SealRandEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("DealIDs", "List__DealID", false, true),
			schema.SpawnStructField("Expiration", "ChainEpoch", false, false),
			schema.SpawnStructField("ReplaceCapacity", "Bool", false, false),
			schema.SpawnStructField("ReplaceSectorDeadline", "Int", false, false),
			schema.SpawnStructField("ReplaceSectorPartition", "Int", false, false),
			schema.SpawnStructField("ReplaceSectorNumber", "SectorNumber", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MinerV0SectorOnChainInfo",
		[]schema.StructField{
			schema.SpawnStructField("SectorNumber", "SectorNumber", false, false),
			schema.SpawnStructField("SealProof", "Int", false, false), //RegisteredSealProof
			schema.SpawnStructField("SealedCID", "Link", false, false),
			schema.SpawnStructField("DealIDs", "List__DealID", false, false),
			schema.SpawnStructField("Activation", "ChainEpoch", false, false),
			schema.SpawnStructField("Expiration", "ChainEpoch", false, false),
			schema.SpawnStructField("DealWeight", "BigInt", false, false),            //DealWeight
			schema.SpawnStructField("VerifiedDealWeight", "BigInt", false, false),    //DealWeight
			schema.SpawnStructField("InitialPledge", "BigInt", false, false),         //TokenAmount
			schema.SpawnStructField("ExpectedDayReward", "BigInt", false, false),     //TokenAmount
			schema.SpawnStructField("ExpectedStorageReward", "BigInt", false, false), //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__SectorOnChainInfo", "String", "MinerV0SectorOnChainInfo", true)) // Key=SectorNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV0SectorInfo", "Map__SectorOnChainInfo"))
}
