package v2

import (
	"github.com/ipld/go-ipld-prime/schema"
)

// Accumulate defines schema for the filecoin v0 state types.
func Accumulate(ts schema.TypeSystem) {
	accumulateMiner(ts)
	accumulateMarket(ts)
	accumulatePower(ts)
	accumulateReward(ts)

	ts.Accumulate(schema.SpawnUnion("LotusActorV2Head", []schema.TypeName{
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
	ts.Accumulate(schema.SpawnLinkReference("Link__LotusActorV2Head", "LotusActorV2Head"))
}
