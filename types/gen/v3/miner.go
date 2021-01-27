package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMiner(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MinerV3State",
		[]schema.StructField{
			schema.SpawnStructField("Info", "Link__MinerV2Info", false, false),
			schema.SpawnStructField("PreCommitDeposits", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("LockedFunds", "BigInt", false, false),       //TokenAmount
			schema.SpawnStructField("VestingFunds", "Link__MinerV0VestingFunds", false, false),
			schema.SpawnStructField("FeeDebt", "BigInt", false, false),                                    //TokenAmount
			schema.SpawnStructField("InitialPledge", "BigInt", false, false),                              //TokenAmount
			schema.SpawnStructField("PreCommittedSectors", "Link__MinerV3SectorPreCommits", false, false), // Map, HAMT[SectorNumber]SectorPreCommitOnChainInfo
			schema.SpawnStructField("PreCommittedSectorsExpiry", "Link", false, false),                    // BitFieldQueue (AMT[Epoch]*BitField)
			schema.SpawnStructField("AllocatedSectors", "Link__BitField", false, false),
			schema.SpawnStructField("Sectors", "Link__MinerV3SectorInfo", false, false), // Array, AMT[SectorNumber]SectorOnChainInfo (sparse)
			schema.SpawnStructField("ProvingPeriodStart", "ChainEpoch", false, false),
			schema.SpawnStructField("CurrentDeadline", "Int", false, false),
			schema.SpawnStructField("Deadlines", "Link__MinerV3Deadlines", false, false),
			schema.SpawnStructField("EarlyTerminations", "BitField", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV3State", "MinerV3State"))

	ts.Accumulate(schema.SpawnMap("Map__V3SectorPreCommitOnChainInfo", "String", "MinerV0SectorPreCommitOnChainInfo", true)) // Key=SectorNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV3SectorPreCommits", "Map__V3SectorPreCommitOnChainInfo"))

	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV3Deadlines",
		"MinerV3Deadlines",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV3Deadlines",
		[]schema.StructField{
			schema.SpawnStructField("Due", "List__MinerV3DeadlineLink", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnList("List__MinerV3DeadlineLink", "Link__MinerV3Deadline", true))
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV3Deadline",
		"MinerV3Deadline",
	))
	ts.Accumulate(schema.SpawnStruct("MinerV3Deadline",
		[]schema.StructField{
			schema.SpawnStructField("Partitions", "Link__MinerV3Partition", false, false), // AMT[PartitionNumber]Partition
			schema.SpawnStructField("ExpirationEpochs", "Link", false, false),             // AMT[ChainEpoch]BitField
			schema.SpawnStructField("PostSubmissions", "BitField", false, false),
			schema.SpawnStructField("EarlyTerminations", "BitField", false, false),
			schema.SpawnStructField("LiveSectors", "Int", false, false),
			schema.SpawnStructField("TotalSectors", "Int", false, false),
			schema.SpawnStructField("FaultyPower", "MinerV0PowerPair", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MinerV3Partition",
		[]schema.StructField{
			schema.SpawnStructField("Sectors", "BitField", false, false),
			schema.SpawnStructField("Unproven", "BitField", false, false),
			schema.SpawnStructField("Faults", "BitField", false, false),
			schema.SpawnStructField("Recoveries", "BitField", false, false),
			schema.SpawnStructField("Terminated", "BitField", false, false),
			schema.SpawnStructField("ExpirationsEpochs", "Link__MinerV3ExpirationSet", false, false), // AMT[ChainEpoch]ExpirationSet
			schema.SpawnStructField("EarlyTerminated", "Link", false, false),                         // AMT[ChainEpoch]BitField
			schema.SpawnStructField("LivePower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("UnprovenPower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("FaultyPower", "MinerV0PowerPair", false, false),
			schema.SpawnStructField("RecoveringPower", "MinerV0PowerPair", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__MinerV3Partition", "String", "MinerV3Partition", true)) // Key=PartitionNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV3Partition", "Map__MinerV3Partition"))

	ts.Accumulate(schema.SpawnMap("Map__SectorV3OnChainInfo", "String", "MinerV2SectorOnChainInfo", true)) // Key=SectorNumber
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV3SectorInfo", "Map__SectorV3OnChainInfo"))

	ts.Accumulate(schema.SpawnMap("Map__MinerV3ExpirationSet", "String", "MinerV0ExpirationSet", true)) // Key=ChainEpcoh
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV3ExpirationSet", "Map__MinerV3ExpirationSet"))
}
