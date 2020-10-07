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
}
