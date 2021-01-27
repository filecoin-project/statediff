package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateMultisig(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MultisigV3State",
		[]schema.StructField{
			schema.SpawnStructField("Signers", "List__Address", false, false),
			schema.SpawnStructField("NumApprovalsThreshold", "Int", false, false),
			schema.SpawnStructField("NextTxnID", "Int", false, false),
			schema.SpawnStructField("InitialBalance", "BigInt", false, false), //TokenAmount
			schema.SpawnStructField("StartEpoch", "ChainEpoch", false, false),
			schema.SpawnStructField("UnlockDuration", "ChainEpoch", false, false),
			schema.SpawnStructField("PendingTxns", "Link__MultisigV3Transaction", false, false), //hamt[TxnID]Multisigv0Transaction
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__MultisigV3State", "MultisigV3State"))
	ts.Accumulate(schema.SpawnMap("Map__MultisigV3Transaction", "String", "MultisigV0Transaction", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__MultisigV3Transaction", "Map__MultisigV3Transaction"))
}
