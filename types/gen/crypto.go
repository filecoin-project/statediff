package main

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateCrypto(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnBytes("Signature"))
	// TODO: better decomposition of signature once BytePrefix unions implemented
	/*
		ts.Accumulate(schema.SpawnUnion("Signature",
			[]schema.TypeName{
				"SignatureSecp256k1",
				"SignatureBLS",
			},
			schema.UnionRepresentation_BytePrefix{}))
	*/
}
