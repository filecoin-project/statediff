package messages

import (
	"github.com/ipld/go-ipld-prime/schema"
)

// Accumulate defines schema for filecoin message parameters
func Accumulate(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnStruct("MessageParamsInitExecParams",
		[]schema.StructField{
			schema.SpawnStructField("CodeCID", "Link", false, false),
			schema.SpawnStructField("ConstructorParams", "Bytes", false, false),
		},
		schema.StructRepresentation_Tuple{},
	))

	accumulateMarket(ts)
	accumulateMiner(ts)
	accumulateMultisig(ts)
	accumulatePaych(ts)
	accumulatePower(ts)
	accumulateReward(ts)
	accumulateVerifreg(ts)

	ts.Accumulate(schema.SpawnUnion("LotusMessageV2Params", []schema.TypeName{
		"MessageParamsInitExecParams",
		"MinerV0SectorPreCommitInfo",
		"MessageParamsMarketWithdrawBalance",
		"MessageParamsMarketPublishDeals",
		"MessageParamsMarketVerifyDeals",
		"MessageParamsMarketActivateDeals",
		"MessageParamsMarketTerminateDeals",
		"MessageParamsMarketComputeCommitment",
		"MessageParamsMinerConstructor",
		"MessageParamsMinerChangeAddress",
		"MessageParamsMinerChangeMultiaddrs",
		"MessageParamsMinerChangePeerID",
		"MessageParamsMinerSubmitWindowedPoSt",
		"MessageParamsMinerProveCommitSector",
		"MessageParamsMinerCheckSectorProven",
		"MessageParamsMinerConfirmSectorProofs",
		"MessageParamsMinerExtendSectorExpiration",
		"MessageParamsMinerTerminateSectors",
		"MessageParamsMinerDeclareFaults",
		"MessageParamsMinerDeclareFaultsRecovered",
		"MessageParamsMinerCompactPartitions",
		"MessageParamsMinerCompactSectorNumbers",
		"ApplyRewardParams",
		"MessageParamsMinerReportFault",
		"MessageParamsMinerWithdrawBalance",
		"MessageParamsMinerDeferredCron",
		"MessageParamsMultisigConstructor",
		"MessageParamsMultisigPropose",
		"MessageParamsMultisigTxnID",
		"MessageParamsMultisigAddSigner",
		"MessageParamsMultisigRemoveSigner",
		"MessageParamsMultisigSwapSigner",
		"MessageParamsMultisigChangeThreshold",
		"MessageParamsMultisigLockBalance",
		"MessageParamsPaychConstructor",
		"MessageParamsPaychUpdateChannelState",
		"MessageParamsPowerCreateMiner",
		"MessageParamsPowerUpdateClaimed",
		"MessageParamsPowerEnrollCron",
		"MessageParamsPowerCurrentTotal",
		"MessageParamsRewardAwardBlock",
		"MessageParamsVerifregAddVerifier",
		"MessageParamsVerifregUseBytes",
	}, schema.UnionRepresentation_Kinded{}))
}
