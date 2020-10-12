package statediff

import (
	"github.com/filecoin-project/statediff/types"
	"github.com/ipld/go-ipld-prime"
)

type methodtable map[int]ipld.NodePrototype

var initTable = methodtable{
	2: types.Type.MessageParamsInitExecParams__Repr,
}

var marketTable = methodtable{
	2: types.Type.Address__Repr,
	3: types.Type.MessageParamsMarketWithdrawBalance__Repr,
	4: types.Type.MessageParamsMarketPublishDeals__Repr,
	5: types.Type.MessageParamsMarketVerifyDeals__Repr,
	6: types.Type.MessageParamsMarketActivateDeals__Repr,
	7: types.Type.MessageParamsMarketTerminateDeals__Repr,
	8: types.Type.MessageParamsMarketComputeCommitment__Repr,
	9: types.Type.Any__Repr,
}

var minerTable = methodtable{
	1:  types.Type.MessageParamsMinerConstructor__Repr,
	2:  types.Type.Any__Repr,
	3:  types.Type.MessageParamsMinerChangeAddress__Repr,
	4:  types.Type.MessageParamsMinerChangePeerID__Repr,
	5:  types.Type.MessageParamsMinerSubmitWindowedPoSt__Repr,
	6:  types.Type.MinerV0SectorPreCommitInfo__Repr,
	7:  types.Type.MessageParamsMinerProveCommitSector__Repr,
	8:  types.Type.MessageParamsMinerExtendSectorExpiration__Repr,
	9:  types.Type.MessageParamsMinerTerminateSectors__Repr,
	10: types.Type.MessageParamsMinerDeclareFaults__Repr,
	11: types.Type.MessageParamsMinerDeclareFaultsRecovered__Repr,
	12: types.Type.MessageParamsMinerDeferredCron__Repr,
	13: types.Type.MessageParamsMinerCheckSectorProven__Repr,
	14: types.Type.ApplyRewardParams__Repr,
	15: types.Type.MessageParamsMinerReportFault__Repr,
	16: types.Type.MessageParamsMinerWithdrawBalance__Repr,
	17: types.Type.MessageParamsMinerConfirmSectorProofs__Repr,
	18: types.Type.MessageParamsMinerChangeMultiaddrs__Repr,
	19: types.Type.MessageParamsMinerCompactPartitions__Repr,
	20: types.Type.MessageParamsMinerCompactSectorNumbers__Repr,
	21: types.Type.Any__Repr,
	22: types.Type.Any__Repr,
	23: types.Type.Address__Repr,
}

var messageParamTable = map[LotusType]methodtable{
	InitActorState:         initTable,
	MarketActorState:       marketTable,
	StorageMinerActorState: minerTable,
}
