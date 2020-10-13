package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMultisig(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigConstructor",
		[]schema.StructField{
			schema.SpawnStructField("Signers", "List__Address", false, false),
			schema.SpawnStructField("NumApprovalsThreshold", "Int", false, false),
			schema.SpawnStructField("UnlockDuration", "ChainEpoch", false, false),
			schema.SpawnStructField("StartEpoch", "ChainEpoch", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigPropose",
		[]schema.StructField{
			schema.SpawnStructField("To", "Address", false, false),
			schema.SpawnStructField("Value", "BigInt", false, false),
			schema.SpawnStructField("Method", "MethodNum", false, false),
			schema.SpawnStructField("Params", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigTxnID",
		[]schema.StructField{
			schema.SpawnStructField("ID", "Int", false, false), //TxnID
			schema.SpawnStructField("ProposeHash", "Bytes", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigAddSigner",
		[]schema.StructField{
			schema.SpawnStructField("Signer", "Address", false, false),
			schema.SpawnStructField("Increase", "Bool", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigRemoveSigner",
		[]schema.StructField{
			schema.SpawnStructField("Signer", "Address", false, false),
			schema.SpawnStructField("Decrease", "Bool", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigSwapSigner",
		[]schema.StructField{
			schema.SpawnStructField("From", "Address", false, false),
			schema.SpawnStructField("To", "Address", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigChangeThreshold",
		[]schema.StructField{
			schema.SpawnStructField("NewThreshold", "Int", false, false),
		}, schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("MessageParamsMultisigLockBalance",
		[]schema.StructField{
			schema.SpawnStructField("StartEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("UnlockDuration", "ChainEpoch", false, false),
			schema.SpawnStructField("Amount", "BigInt", false, false),
		}, schema.StructRepresentation_Tuple{}))

}
