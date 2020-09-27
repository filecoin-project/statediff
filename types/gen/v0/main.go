package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

// Accumulate defines schema for the filecoin v0 state types.
func Accumulate(ts schema.TypeSystem) {
	accumulateAccount(ts)
	accumulateCron(ts)
	accumulateInit(ts)
	accumulateMarket(ts)
	accumulateMiner(ts)
	accumulateMultisig(ts)
	accumulatePaych(ts)
	accumulatePower(ts)
	accumulateReward(ts)
	accumulateVerifiedRegistry(ts)
}
