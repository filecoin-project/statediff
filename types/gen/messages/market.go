package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMarket(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsMarketWithdrawBalance",
		[]schema.StructField{
			schema.SpawnStructField("ProviderOrClientAmount", "Address", false, false),
			schema.SpawnStructField("Amount", "BigInt", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMarketPublishDeals",
		[]schema.StructField{
			schema.SpawnStructField("Deals", "List__ClientDealProposal", false, false),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__ClientDealProposal", "MarketClientDealProposal", false))
	ts.Accumulate(schema.SpawnStruct("MarketClientDealProposal",
		[]schema.StructField{
			schema.SpawnStructField("Proposal", "MarketV2DealProposal", false, false),
			schema.SpawnStructField("ClientSignature", "Signature", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMarketVerifyDeals",
		[]schema.StructField{
			schema.SpawnStructField("DealIDs", "List__DealID", false, false),
			schema.SpawnStructField("SectorExpiry", "ChainEpoch", false, false),
			schema.SpawnStructField("SectorStart", "ChainEpoch", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMarketActivateDeals",
		[]schema.StructField{
			schema.SpawnStructField("DealIDs", "List__DealID", false, false),
			schema.SpawnStructField("SectorExpiry", "ChainEpoch", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMarketTerminateDeals",
		[]schema.StructField{
			schema.SpawnStructField("Epoch", "ChainEpoch", false, false),
			schema.SpawnStructField("DealIDs", "List__DealID", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMarketComputeCommitment",
		[]schema.StructField{
			schema.SpawnStructField("DealIDs", "List__DealID", false, false),
			schema.SpawnStructField("SectorType", "RegisteredSealProof", false, false),
		}, schema.StructRepresentation_Tuple{}))

}
