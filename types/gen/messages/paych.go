package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulatePaych(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsPaychConstructor",
		[]schema.StructField{
			schema.SpawnStructField("From", "Address", false, false),
			schema.SpawnStructField("To", "Address", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("SignedVoucher",
		[]schema.StructField{
			schema.SpawnStructField("ChannelAddr", "Address", false, false),
			schema.SpawnStructField("TimeLockMin", "ChainEpoch", false, false),
			schema.SpawnStructField("TimeLockMax", "ChainEpoch", false, false),
			schema.SpawnStructField("SecretPreimage", "Bytes", true, false),
			schema.SpawnStructField("Extra", "ModVerifyParams", true, true),
			schema.SpawnStructField("Lane", "Int", false, false),
			schema.SpawnStructField("Nonce", "Int", false, false),
			schema.SpawnStructField("Amount", "BigInt", false, false),
			schema.SpawnStructField("MinSettleHeight", "ChainEpoch", true, false),
			schema.SpawnStructField("Merges", "List__Merge", true, false),
			schema.SpawnStructField("Signature", "Signature", false, true),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__Merge", "Merge", false))
	ts.Accumulate(schema.SpawnStruct("ModVerifyParams",
		[]schema.StructField{
			schema.SpawnStructField("Method", "MethodNum", false, false),
			schema.SpawnStructField("Params", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnStruct("Merge",
		[]schema.StructField{
			schema.SpawnStructField("Lane", "Int", false, false),
			schema.SpawnStructField("Nonce", "Int", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsPaychUpdateChannelState",
		[]schema.StructField{
			schema.SpawnStructField("Sv", "SignedVoucher", false, false),
			schema.SpawnStructField("Secret", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))
}
