package v2

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMarket(ts schema.TypeSystem) {
	// Note: what's changed in market v2 is not the schema types, but rather the hamt format.
	ts.Accumulate(schema.SpawnStruct("MarketV2State",
		[]schema.StructField{
			schema.SpawnStructField("Proposals", "Link__MarketV2RawDealProposal", false, false),     //AMT[DealID]DealProposal
			schema.SpawnStructField("States", "Link__MarketV0DealState", false, false),              //AMT[DealID]DealState
			schema.SpawnStructField("PendingProposals", "Link__MarketV2DealProposal", false, false), // HAMT[DealCid]DealProposal
			schema.SpawnStructField("EscrowTable", "Link__BalanceTable", false, false),              // BalanceTable
			schema.SpawnStructField("LockedTable", "Link__BalanceTable", false, false),              // BalanceTable
			schema.SpawnStructField("NextID", "DealID", false, false),
			schema.SpawnStructField("DealOpsByEpoch", "Link__MultimapDealID", false, false), // SetMultimap, HAMT[epoch]Set
			schema.SpawnStructField("LastCron", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalClientLockedCollateral", "BigInt", false, false),   //TokenAmount
			schema.SpawnStructField("TotalProviderLockedCollateral", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("TotalClientStorageFee", "BigInt", false, false),         //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV2State", "MarketV2State"))

	ts.Accumulate(schema.SpawnStruct("MarketV2DealProposal",
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
	ts.Accumulate(schema.SpawnMap("Map__MarketV2DealProposal", "CidString", "MarketV2DealProposal", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV2DealProposal", "Map__MarketV2DealProposal"))
	ts.Accumulate(schema.SpawnMap("Map__MarketV2RawDealProposal", "String", "MarketV2DealProposal", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV2RawDealProposal", "Map__MarketV2RawDealProposal"))
}
