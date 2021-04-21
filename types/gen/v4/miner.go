package v4

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMiner(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MinerV4State",
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
			schema.SpawnStructField("DeadlineCronActive", "Bool", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MinerV4State", "MinerV4State"))
}
