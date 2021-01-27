package v3

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateVerifiedRegistry(ts schema.TypeSystem) {
	//ts.Accumulate(schema.SpawnBytes("DataCap")) // = StoragePower = big.int
	ts.Accumulate(schema.SpawnStruct("VerifregV3State",
		[]schema.StructField{
			schema.SpawnStructField("RootKey", "Address", false, false),
			schema.SpawnStructField("Verifiers", "Link__V3DataCap", false, false),       //HAMT[addr.Address]DataCap
			schema.SpawnStructField("VerifiedClients", "Link__V3DataCap", false, false), //HAMT[addr.Address]DataCap
		},
		schema.StructRepresentation_Tuple{},
	))
	ts.Accumulate(schema.SpawnLinkReference("Link__VerifregV3State", "VerifregV3State"))

	ts.Accumulate(schema.SpawnMap("Map__V3DataCap", "RawAddress", "BigInt", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__V3DataCap", "Map__V3DataCap"))
}
