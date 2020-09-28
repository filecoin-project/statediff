package main

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateLotus(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("LotusBlockHeader",
		[]schema.StructField{
			schema.SpawnStructField("Miner", "Address", false, false),
			schema.SpawnStructField("Ticket", "LotusTicket", false, true),
			schema.SpawnStructField("ElectionProof", "LotusElectionProof", false, true),
			schema.SpawnStructField("BeaconEntries", "List__LotusBeaconEntry", false, false),
			schema.SpawnStructField("WinPoStProof", "List__PoStProof", false, false),
			schema.SpawnStructField("Parents", "List__Link", false, false),
			schema.SpawnStructField("ParentWeight", "BigInt", false, false),
			schema.SpawnStructField("Height", "ChainEpoch", false, false),
			schema.SpawnStructField("ParentStateRoot", "Link__LotusStateRoot", false, false),
			schema.SpawnStructField("ParentMessageReceipts", "Link", false, false),
			schema.SpawnStructField("Messages", "Link", false, false),
			schema.SpawnStructField("BLSAggregate", "Signature", false, true),
			schema.SpawnStructField("Timestamp", "Int", false, false),
			schema.SpawnStructField("BlockSig", "Signature", false, true),
			schema.SpawnStructField("ForkSignaling", "Int", false, false),
			schema.SpawnStructField("ParentBaseFee", "BigInt", false, false), //TokenAmount
		},
		schema.StructRepresentation_Tuple{},
	))

	ts.Accumulate(schema.SpawnStruct("LotusTicket",
		[]schema.StructField{
			schema.SpawnStructField("VRFProof", "Bytes", false, false),
		},
		schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("LotusBeaconEntry",
		[]schema.StructField{
			schema.SpawnStructField("Round", "Int", false, false),
			schema.SpawnStructField("Data", "Bytes", false, false),
		},
		schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__LotusBeaconEntry", "LotusBeaconEntry", true))

	ts.Accumulate(schema.SpawnStruct("LotusElectionProof",
		[]schema.StructField{
			schema.SpawnStructField("WinCount", "Int", false, false),
			schema.SpawnStructField("VRFProof", "Bytes", false, false),
		},
		schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("PoStProof",
		[]schema.StructField{
			// https://github.com/filecoin-project/go-state-types/blob/master/abi/sector.go
			schema.SpawnStructField("PoStProof", "Int", false, false),
			schema.SpawnStructField("ProofBytes", "Bytes", false, false),
		},
		schema.StructRepresentation_Tuple{}))
	ts.Accumulate(schema.SpawnList("List__PoStProof", "PoStProof", false))

	ts.Accumulate(schema.SpawnLinkReference("Link__LotusStateRoot", "LotusStateRoot"))

	ts.Accumulate(schema.SpawnStruct("LotusStateRoot",
		[]schema.StructField{
			schema.SpawnStructField("Version", "Int", false, false),
			schema.SpawnStructField("Actors", "Link", false, false),
			schema.SpawnStructField("Info", "Link", false, false),
		},
		schema.StructRepresentation_Tuple{}))

	ts.Accumulate(schema.SpawnStruct("LotusActors",
		[]schema.StructField{
			schema.SpawnStructField("Code", "Link", false, false),
			schema.SpawnStructField("Head", "Link", false, false),
			schema.SpawnStructField("Nonce", "Int", false, false),
			schema.SpawnStructField("Balance", "BigInt", false, false),
		},
		schema.StructRepresentation_Tuple{}))
}
