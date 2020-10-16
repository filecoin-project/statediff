package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateVerifreg(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsVerifregAddVerifier",
		[]schema.StructField{
			schema.SpawnStructField("Address", "Address", false, false),
			schema.SpawnStructField("Allowance", "BigInt", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsVerifregUseBytes",
		[]schema.StructField{
			schema.SpawnStructField("Address", "Address", false, false),
			schema.SpawnStructField("DealSize", "BigInt", false, false),
		}, schema.StructRepresentation_Tuple{}))

}
