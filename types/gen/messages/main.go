package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

// Accumulate defines schema for filecoin message parameters
func Accumulate(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsInitExecParams",
		[]schema.StructField{
			schema.SpawnStructField("CodeCID", "Link", false, false),
			schema.SpawnStructField("ConstructorParams", "Bytes", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	accumulateMarket(ts)
	accumulateMiner(ts)
	accumulateMultisig(ts)
	accumulatePaych(ts)
	accumulatePower(ts)
	accumulateReward(ts)
	accumulateVerifreg(ts)
}
