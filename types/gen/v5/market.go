package v5

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMarket(ts schema.TypeSystem) {
	// Note: state is identical to v3 but method signature of ComputeDataCommitment has changed
	ts.Accumulate(schema.SpawnStruct("MarketV5State",
		[]schema.StructField{
			schema.SpawnStructField("Proposals", "Link__MarketV3RawDealProposal", false, false),     // AMTv3[DealID]DealProposal
			schema.SpawnStructField("States", "Link__MarketV3DealState", false, false),              // AMTv3[DealID]DealState
			schema.SpawnStructField("PendingProposals", "Link__MarketV3DealProposal", false, false), // HAMTv3[DealCid]DealProposal
			schema.SpawnStructField("EscrowTable", "Link__V3BalanceTable", false, false),            // BalanceTable
			schema.SpawnStructField("LockedTable", "Link__V3BalanceTable", false, false),            // BalanceTable
			schema.SpawnStructField("NextID", "DealID", false, false),
			schema.SpawnStructField("DealOpsByEpoch", "Link__MarketV3MultimapDealID", false, false), // SetMultimap, HAMTv3[epoch]Set
			schema.SpawnStructField("LastCron", "ChainEpoch", false, false),
			schema.SpawnStructField("TotalClientLockedCollateral", "BigInt", false, false),   // TokenAmount
			schema.SpawnStructField("TotalProviderLockedCollateral", "BigInt", false, false), // TokenAmount
			schema.SpawnStructField("TotalClientStorageFee", "BigInt", false, false),         // TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MarketV5State", "MarketV5State"))
}
