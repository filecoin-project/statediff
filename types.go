package statediff

import (
	"encoding/json"

	storageMinerActor "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	"github.com/ipfs/go-cid"
)

type MinerPartitionJSONBitfield struct {
	storageMinerActor.Partition
}

// TODO: keep in sync w/ 
// https://github.com/filecoin-project/specs-actors/blob/master/actors/builtin/miner/partition_state.go#L17
type jsonMinerPartition struct {
	Sectors JSONBitField
	Faults JSONBitField
	Recoveries JSONBitField
	Terminated JSONBitField
	ExpirationsEpochs cid.Cid
	EarlyTerminated cid.Cid
	LivePower storageMinerActor.PowerPair
	FaultyPower storageMinerActor.PowerPair
	RecoveringPower storageMinerActor.PowerPair
}

func (m MinerPartitionJSONBitfield) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonMinerPartition{
		Sectors: JSONBitField{m.Sectors},
		Faults: JSONBitField{m.Faults},
		Recoveries: JSONBitField{m.Recoveries},
		Terminated: JSONBitField{m.Terminated},
		ExpirationsEpochs: m.ExpirationsEpochs,
		EarlyTerminated: m.EarlyTerminated,
		LivePower: m.LivePower,
		FaultyPower: m.FaultyPower,
		RecoveringPower: m.RecoveringPower,
	})
}

type MinerDeadlineJSONBitfield struct {
	storageMinerActor.Deadline
}

// TODO: keep in sync w/ 
// https://github.com/filecoin-project/specs-actors/blob/v0.9.6/actors/builtin/miner/deadline_state.go#L29
type jsonMinerDeadline struct {
	Partitions cid.Cid
	ExpirationsEpochs cid.Cid
	PostSubmissions JSONBitField
	EarlyTerminations JSONBitField
	LiveSectors uint64
	TotalSectors uint64
	FaultyPower storageMinerActor.PowerPair
}

func (m MinerDeadlineJSONBitfield) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonMinerDeadline{
		Partitions: m.Partitions,
		ExpirationsEpochs: m.ExpirationsEpochs,
		PostSubmissions: JSONBitField{m.PostSubmissions},
		EarlyTerminations: JSONBitField{m.EarlyTerminations},
		LiveSectors: m.LiveSectors,
		TotalSectors: m.TotalSectors,
		FaultyPower: m.FaultyPower,
	})
}
