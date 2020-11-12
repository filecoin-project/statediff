package ledgerwallet

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

var errBad = fmt.Errorf("stubbed out.")

type LedgerWallet struct {
}

func NewWallet(ds dtypes.MetadataDS) *LedgerWallet {
	return &LedgerWallet{}
}

type LedgerKeyInfo struct {
	Address address.Address
	Path    []uint32
}

var _ api.WalletAPI = (*LedgerWallet)(nil)

func (lw LedgerWallet) WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta api.MsgMeta) (*crypto.Signature, error) {
	return nil, errBad
}

func (lw LedgerWallet) WalletDelete(ctx context.Context, k address.Address) error {
	return errBad
}

func (lw LedgerWallet) WalletExport(ctx context.Context, k address.Address) (*types.KeyInfo, error) {
	return nil, errBad
}

func (lw LedgerWallet) WalletHas(ctx context.Context, k address.Address) (bool, error) {
	return false, errBad
}

func (lw LedgerWallet) WalletImport(ctx context.Context, kinfo *types.KeyInfo) (address.Address, error) {
	return address.Undef, errBad
}

func (lw LedgerWallet) WalletList(ctx context.Context) ([]address.Address, error) {
	return nil, errBad
}

func (lw LedgerWallet) WalletNew(ctx context.Context, t types.KeyType) (address.Address, error) {
	return address.Undef, errBad
}

func (lw *LedgerWallet) Get() api.WalletAPI {
	return nil
}
