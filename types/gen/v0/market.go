package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMarket(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MarketV0State",
		[]schema.StructField{
			schema.SpawnStructField("Proposals", "Link", false, false),        //AMT[DealID]DealProposal
			schema.SpawnStructField("States", "Link", false, false),           //AMT[DealID]DealState
			schema.SpawnStructField("PendingProposals", "Link", false, false), // HAMT[DealCid]DealProposal
			schema.SpawnStructField("EscrowTable", "Link", false, false),      // BalanceTable
			schema.SpawnStructField("LockedTable", "Link", false, false),      // BalanceTable
			schema.SpawnStructField("NextID", "DealID", false, false),
			schema.SpawnStructField("DealOpsByEpoch", "Link", false, false), // SetMultimap, HAMT[epoch]Set
			schema.SpawnStructField("LastCron", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalClientLockedCollateral", "BigInt", false, false),   //TokenAmount
			schema.SpawnStructField("TotalProviderLockedCollateral", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("TotalClientStorageFee", "BigInt", false, false),         //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MarketV0DealProposal",
		[]schema.StructField{
			schema.SpawnStructField("PieceCID", "Link", false, false),
			schema.SpawnStructField("PieceSize", "PaddedPieceSize", false, false),
			schema.SpawnStructField("VerifiedDeal", "Bool", false, false),
			schema.SpawnStructField("Client", "Address", false, false),
			schema.SpawnStructField("Provider", "Address", false, false),
			schema.SpawnStructField("Label", "String", false, false),
			schema.SpawnStructField("StartEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("EndEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("StoragePricePerEpoch", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("ProviderCollateral", "BigInt", false, false),   //TokenAmount
			schema.SpawnStructField("ClientCollateral", "BigInt", false, false),     //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MarketV0DealState",
		[]schema.StructField{
			schema.SpawnStructField("SectorStartEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("LastUpdatedEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("SlashEpoch", "ChainEpoch", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
}
