package v2

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMiner(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MinerV2State",
		[]schema.StructField{
			schema.SpawnStructField("Info", "Link__MinerV2Info", false, false),
			schema.SpawnStructField("PreCommitDeposits", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("LockedFunds", "BigInt", false, false),       //TokenAmount
			schema.SpawnStructField("VestingFunds", "Link__MinerV0VestingFunds", false, false),
			schema.SpawnStructField("FeeDebt", "BigInt", false, false),                                    //TokenAmount
			schema.SpawnStructField("InitialPledge", "BigInt", false, false),                              //TokenAmount
			schema.SpawnStructField("PreCommittedSectors", "Link__MinerV0SectorPreCommits", false, false), // Map, HAMT[SectorNumber]SectorPreCommitOnChainInfo
			schema.SpawnStructField("PreCommittedSectorsExpiry", "Link", false, false),                    // BitFieldQueue (AMT[Epoch]*BitField)
			schema.SpawnStructField("AllocatedSectors", "Link__BitField", false, false),
			schema.SpawnStructField("Sectors", "Link__MinerV2SectorInfo", false, false), // Array, AMT[SectorNumber]SectorOnChainInfo (sparse)
			schema.SpawnStructField("ProvingPeriodStart", "ChainEpoch", false, false),
			schema.SpawnStructField("CurrentDeadline", "Int", false, false),
			schema.SpawnStructField("Deadlines", "Link__MinerV2Deadlines", false, false),
			schema.SpawnStructField("EarlyTerminations", "BitField", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV2State", "MinerV2State"))

	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV2Info",
		"MinerV2Info",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV2Info",
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
			schema.SpawnStructField("ConsensusFaultElapsed", "ChainEpoch", false, false),
			schema.SpawnStructField("PendingOwnerAddress", "Address", false, true),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV2Deadlines",
		"MinerV2Deadlines",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV2Deadlines",
		[]schema.StructField{
			schema.SpawnStructField("Due", "List__MinerV2DeadlineLink", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnList("List__MinerV2DeadlineLink", "Link__MinerV2Deadline", true))
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV2Deadline",
		"MinerV2Deadline",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV2Deadline",
		[]schema.StructField{
			schema.SpawnStructField("Partitions", "Link__MinerV2Partition", false, false), // AMT[PartitionNumber]Partition
			schema.SpawnStructField("ExpirationEpochs", "Link", false, false),             // AMT[ChainEpoch]BitField
			schema.SpawnStructField("PostSubmissions", "BitField", false, false),
			schema.SpawnStructField("EarlyTerminations", "BitField", false, false),
			schema.SpawnStructField("LiveSectors", "Int", false, false),
			schema.SpawnStructField("TotalSectors", "Int", false, false),
			schema.SpawnStructField("FaultyPower", "MinerV0PowerPair", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MinerV2Partition",
		[]schema.StructField{
			schema.SpawnStructField("Sectors", "BitField", false, false),
			schema.SpawnStructField("Unproven", "BitField", false, false),
			schema.SpawnStructField("Faults", "BitField", false, false),
			schema.SpawnStructField("Recoveries", "BitField", false, false),
			schema.SpawnStructField("Terminated", "BitField", false, false),
			schema.SpawnStructField("ExpirationsEpochs", "Link__MinerV0ExpirationSet", false, false), // AMT[ChainEpoch]ExpirationSet
			schema.SpawnStructField("EarlyTerminated", "Link", false, false),                         // AMT[ChainEpoch]BitField
			schema.SpawnStructField("LivePower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("UnprovenPower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("FaultyPower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("RecoveringPower", "MinerV0PowerPair", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__MinerV2Partition", "String", "MinerV2Partition", true)) // Key=PartitionNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV2Partition", "Map__MinerV2Partition"))

	ts.Accumulate(schema.SpawnStruct("MinerV2SectorOnChainInfo",
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
			schema.SpawnStructField("ReplacedSectorAge", "ChainEpoch", false, false),
			schema.SpawnStructField("ReplacedDayReward", "BigInt", false, false), //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__SectorV2OnChainInfo", "String", "MinerV2SectorOnChainInfo", true)) // Key=SectorNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV2SectorInfo", "Map__SectorV2OnChainInfo"))

}
