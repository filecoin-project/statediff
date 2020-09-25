package main

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateABI(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnInt("MethodNum"))
	ts.Accumulate(schema.SpawnInt("ActorID"))
	ts.Accumulate(schema.SpawnInt("ChainEpoch"))
	ts.Accumulate(schema.SpawnInt("TokenAmount"))
	ts.Accumulate(schema.SpawnInt("UnpaddedPieceSize"))
	ts.Accumulate(schema.SpawnInt("PaddedPieceSize"))
	ts.Accumulate(schema.SpawnBytes("PeerID"))
	ts.Accumulate(schema.SpawnInt("SectorSize"))
	ts.Accumulate(schema.SpawnInt("SectorNumber"))
	ts.Accumulate(schema.SpawnLinkReference("Link__BitField", "BitField"))
	ts.Accumulate(schema.SpawnBytes("BitField"))
	ts.Accumulate(schema.SpawnLink("RegisteredSealProof"))
	ts.Accumulate(schema.SpawnBytes("StoragePower")) // big.Int

	ts.Accumulate(schema.SpawnList("List__DealID", "DealID", true))
	ts.Accumulate(schema.SpawnInt("DealID"))
	ts.Accumulate(schema.SpawnBytes("DealWeight")) // big.Int

	ts.Accumulate(schema.SpawnList("List__Multiaddrs", "Multiaddr", true))
	ts.Accumulate(schema.SpawnBytes("Multiaddr"))
}
