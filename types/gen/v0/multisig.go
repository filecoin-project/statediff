package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMultisig(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MultisigV0State",
		[]schema.StructField{
			schema.SpawnStructField("Signers", "List__Address", false, false),
			schema.SpawnStructField("NumApprovalsThreshold", "Int", false, false),
			schema.SpawnStructField("NextTxnID", "Int", false, false),
			schema.SpawnStructField("InitialBalance", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("StartEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("UnlockDuration", "ChainEpoch", false, false),
			schema.SpawnStructField("PendingTxns", "Link__MultisigV0Transaction", false, false), //hamt[TxnID]Multisigv0Transaction
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MultisigV0State", "MultisigV0State"))
	ts.Accumulate(schema.SpawnMap("Map__MultisigV0Transaction", "String", "MultisigV0Transaction", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MultisigV0Transaction", "Map__MultisigV0Transaction"))
	ts.Accumulate(schema.SpawnStruct("MultisigV0Transaction",
		[]schema.StructField{
			schema.SpawnStructField("To", "Address", false, false),
			schema.SpawnStructField("Value", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("Method", "MethodNum", false, false),
			schema.SpawnStructField("Params", "Bytes", false, false),
			schema.SpawnStructField("Approved", "List__Address", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))
}
