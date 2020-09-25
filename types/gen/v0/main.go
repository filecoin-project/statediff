package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

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
