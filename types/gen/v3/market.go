package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMarket(ts schema.TypeSystem) {
	// Note: what's changed in market v3 is the hamt format.
	ts.Accumulate(schema.SpawnStruct("MarketV3State",
		[]schema.StructField{
			schema.SpawnStructField("Proposals", "Link__MarketV3RawDealProposal", false, false),     //AMTv3[DealID]DealProposal
			schema.SpawnStructField("States", "Link__MarketV3DealState", false, false),              //AMTv3[DealID]DealState
			schema.SpawnStructField("PendingProposals", "Link__MarketV3DealProposal", false, false), // HAMTv3[DealCid]DealProposal
			schema.SpawnStructField("EscrowTable", "Link__V3BalanceTable", false, false),            // BalanceTable
			schema.SpawnStructField("LockedTable", "Link__V3BalanceTable", false, false),            // BalanceTable
			schema.SpawnStructField("NextID", "DealID", false, false),
			schema.SpawnStructField("DealOpsByEpoch", "Link__MarketV3MultimapDealID", false, false), // SetMultimap, HAMTv3[epoch]Set
			schema.SpawnStructField("LastCron", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalClientLockedCollateral", "BigInt", false, false),   //TokenAmount
			schema.SpawnStructField("TotalProviderLockedCollateral", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("TotalClientStorageFee", "BigInt", false, false),         //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV3State", "MarketV3State"))

	ts.Accumulate(schema.SpawnMap("Map__MarketV3DealProposal", "CidString", "MarketV2DealProposal", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV3DealProposal", "Map__MarketV3DealProposal"))
	ts.Accumulate(schema.SpawnMap("Map__MarketV3RawDealProposal", "String", "MarketV2DealProposal", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV3RawDealProposal", "Map__MarketV3RawDealProposal"))
	ts.Accumulate(schema.SpawnMap("Map__MarketV3DealState", "String", "MarketV0DealState", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV3DealState", "Map__MarketV3DealState"))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV3MultimapDealID", "MapV3__List__DealID"))
	ts.Accumulate(schema.SpawnMap("MapV3__List__DealID", "String", "List__DealID", false))
	ts.Accumulate(schema.SpawnMap("Map__V3BalanceTable", "RawAddress", "BigInt", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__V3BalanceTable", "Map__V3BalanceTable"))

}
