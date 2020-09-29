package v0

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateVerifiedRegistry(ts schema.TypeSystem) {
	//ts.Accumulate(schema.SpawnBytes("DataCap")) // = StoragePower = big.int
	ts.Accumulate(schema.SpawnStruct("VerifregV0State",
		[]schema.StructField{
			schema.SpawnStructField("RootKey", "Address", false, false),
			schema.SpawnStructField("Verifiers", "Link__DataCap", false, false),       //HAMT[addr.Address]DataCap
			schema.SpawnStructField("VerifiedClients", "Link__DataCap", false, false), //HAMT[addr.Address]DataCap
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnMap("Map__DataCap", "RawAddress", "BigInt", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__DataCap", "Map__DataCap"))
}
