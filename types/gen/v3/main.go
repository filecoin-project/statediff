package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

// Accumulate defines schema for the filecoin v0 state types.
func Accumulate(ts schema.TypeSystem) {
	accumulateMarket(ts)
	accumulateMiner(ts)
	accumulatePower(ts)
	accumulateInit(ts)
	accumulateVerifiedRegistry(ts)
	accumulatePaych(ts)
	accumulateMultisig(ts)

	ts.Accumulate(schema.SpawnUnion("LotusActorV3Head", []schema.TypeName{
		"MarketV3State",
		"MinerV3State",
		"PowerV3State",
		"InitV3State",
		"VerifregV3State",
		"PaychV3State",
		"MultisigV3State",

		"MarketV2State",
		"MinerV2State",
		"PowerV2State",
		"RewardV2State",
		"AccountV0State",
		"CronV0State",
		"InitV0State",
		"MarketV0State",
		"MinerV0State",
		"MultisigV0State",
		"PaychV0State",
		"PowerV0State",
		"RewardV0State",
		"VerifregV0State",
	}, schema.UnionRepresentation_Kinded{}))
	ts.Accumulate(schema.SpawnLinkReference("Link__LotusActorV3Head", "LotusActorV3Head"))
}
