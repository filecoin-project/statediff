package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMarket(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MarketV0State",
		[]schema.StructField{
			schema.SpawnStructField("Proposals", "Link__MarketV0RawDealProposal", false, false),     //AMT[DealID]DealProposal
			schema.SpawnStructField("States", "Link__MarketV0DealState", false, false),              //AMT[DealID]DealState
			schema.SpawnStructField("PendingProposals", "Link__MarketV0DealProposal", false, false), // HAMT[DealCid]DealProposal
			schema.SpawnStructField("EscrowTable", "Link__BalanceTable", false, false),              // BalanceTable
			schema.SpawnStructField("LockedTable", "Link__BalanceTable", false, false),              // BalanceTable
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
	ts.Accumulate(schema.SpawnMap("Map__MarketV0DealProposal", "CidString", "MarketV0DealProposal", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV0DealProposal", "Map__MarketV0DealProposal"))
	ts.Accumulate(schema.SpawnMap("Map__MarketV0RawDealProposal", "String", "MarketV0DealProposal", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV0RawDealProposal", "Map__MarketV0RawDealProposal"))

	ts.Accumulate(schema.SpawnStruct("MarketV0DealState",
		[]schema.StructField{
			schema.SpawnStructField("SectorStartEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("LastUpdatedEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("SlashEpoch", "ChainEpoch", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__MarketV0DealState", "String", "MarketV0DealState", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV0DealState", "Map__MarketV0DealState"))

	ts.Accumulate(schema.SpawnMap("Map__BalanceTable", "RawAddress", "BigInt", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__BalanceTable", "Map__BalanceTable"))
}
