package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMarket(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MarketV0State",
		[]schema.StructField{
			schema.SpawnStructField("Proposals", "Link", false, false), //AMT[DealID]DealProposal
			schema.SpawnStructField("States", "Link", false, false), //AMT[DealID]DealState
			schema.SpawnStructField("PendingProposals", "Link", false, false), // HAMT[DealCid]DealProposal
			schema.SpawnStructField("EscrowTable", "Link", false, false),// BalanceTable
			schema.SpawnStructField("LockedTable", "Link", false, false),// BalanceTable
			schema.SpawnStructField("NextID", "DealID", false, false),
			schema.SpawnStructField("DealOpsByEpoch", "Link", false, false),// SetMultimap, HAMT[epoch]Set
			schema.SpawnStructField("LastCron", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalClientLockedCollateral", "TokenAmount", false, false),
			schema.SpawnStructField("TotalProviderLockedCollateral", "TokenAmount", false, false),
			schema.SpawnStructField("TotalClientStorageFee", "TokenAmount", false, false),
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
			schema.SpawnStructField("StoragePricePerEpoch", "TokenAmount", false, false),
			schema.SpawnStructField("ProviderCollateral", "TokenAmount", false, false),
			schema.SpawnStructField("ClientCollateral", "TokenAmount", false, false),
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
