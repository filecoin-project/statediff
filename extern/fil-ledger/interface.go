package ledger_filecoin_go

import (
	"fmt"
)

type FL struct{}

func FindLedgerFilecoinApp() (*FL, error) {
	return nil, fmt.Errorf("Stubbed out")
}

func (f *FL) Close() {
}

func (f *FL) GetAddressPubKeySECP256K1(bip44Path []uint32) (pubkey []byte, addrByte []byte, addrString string, err error) {
	return nil, nil, "", nil
}

func (f *FL) ShowAddressPubKeySECP256K1(bip44Path []uint32) (pubkey []byte, addrByte []byte, addrString string, err error) {
	return nil, nil, "", nil
}

func (f *FL) SignSECP256K1(bip44Path []uint32, transaction []byte) (*SignatureAnswer, error) {
	return nil, nil
}

type SignatureAnswer struct{}

func (sa *SignatureAnswer) SignatureBytes() []byte { return nil }
