package main

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateABI(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnInt("MethodNum"))
	ts.Accumulate(schema.SpawnInt("ActorID"))
	ts.Accumulate(schema.SpawnInt("ChainEpoch"))
	//ts.Accumulate(schema.SpawnBytes("TokenAmount")) // big.Int
	ts.Accumulate(schema.SpawnInt("UnpaddedPieceSize"))
	ts.Accumulate(schema.SpawnInt("PaddedPieceSize"))
	ts.Accumulate(schema.SpawnBytes("PeerID"))
	ts.Accumulate(schema.SpawnInt("SectorSize"))
	ts.Accumulate(schema.SpawnInt("SectorNumber"))
	ts.Accumulate(schema.SpawnList("List__SectorNumber", "SectorNumber", false))
	ts.Accumulate(schema.SpawnLinkReference("Link__BitField", "BitField"))
	ts.Accumulate(schema.SpawnBytes("BitField"))
	ts.Accumulate(schema.SpawnMap("Map__BitField", "String", "BitField", true))

	ts.Accumulate(schema.SpawnLink("RegisteredSealProof"))
	//ts.Accumulate(schema.SpawnBytes("StoragePower")) // big.Int

	ts.Accumulate(schema.SpawnMap("Map__List__DealID", "String", "List__DealID", false))
	ts.Accumulate(schema.SpawnList("List__DealID", "DealID", true))
	ts.Accumulate(schema.SpawnInt("DealID"))
	//ts.Accumulate(schema.SpawnBytes("DealWeight")) // big.Int

	ts.Accumulate(schema.SpawnList("List__Multiaddrs", "Multiaddr", true))
	ts.Accumulate(schema.SpawnBytes("Multiaddr"))

	ts.Accumulate(schema.SpawnStruct("SealVerifyInfo",
		[]schema.StructField{
			schema.SpawnStructField("SealProof", "RegisteredSealProof", false, false),
			schema.SpawnStructField("SectorID", "Int", false, false),
			schema.SpawnStructField("DealIDs", "List__DealID", false, false),
			schema.SpawnStructField("Randomness", "Bytes", false, false),
			schema.SpawnStructField("InteractiveRandomness", "Bytes", false, false),
			schema.SpawnStructField("Proof", "Bytes", false, false),
			schema.SpawnStructField("SealedCID", "Link", false, false),
			schema.SpawnStructField("UnsealedCID", "Link", false, false),
		}, schema.StructRepresentation_Tuple{}))
}
