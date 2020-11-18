module github.com/filecoin-project/statediff

go 1.14

require (
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/evanw/esbuild v0.7.18
	github.com/filecoin-project/go-address v0.0.4
	github.com/filecoin-project/go-bitfield v0.2.2
	github.com/filecoin-project/go-fil-markets v1.0.0
	github.com/filecoin-project/go-hamt-ipld/v2 v2.0.0
	github.com/filecoin-project/go-multistore v0.0.3
	github.com/filecoin-project/go-state-types v0.0.0-20201013222834-41ea465f274f
	github.com/filecoin-project/lotus v1.1.2
	github.com/filecoin-project/specs-actors v0.9.13
	github.com/filecoin-project/specs-actors/v2 v2.2.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/graphql-go/graphql v0.7.9
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.5
	github.com/ipfs/go-hamt-ipld v0.1.1
	github.com/ipfs/go-ipfs-blockstore v1.0.1
	github.com/ipfs/go-ipld-cbor v0.0.5-0.20200428170625-a0bd04d3cbdf
	github.com/ipld/go-car v0.1.1-0.20201015032735-ff6ccdc46acc
	github.com/ipld/go-ipld-prime v0.5.1-0.20200910124733-350032422383
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multiaddr-net v0.2.0
	github.com/multiformats/go-multihash v0.0.14
	github.com/polydawn/refmt v0.0.0-20190809202753-05966cbd336a
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/urfave/cli/v2 v2.3.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20200826160007-0b9f6c5fb163
	github.com/willscott/carbs v0.0.4-0.20201024182220-b40626ac657b
	github.com/willscott/go-cmp v0.5.2-0.20200812183318-8affb9542345
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
)

replace github.com/supranational/blst => github.com/filecoin-project/statediff/extern/fil-blst v0.0.0-20201112005514-7550fe42bb9a

replace github.com/filecoin-project/filecoin-ffi => github.com/filecoin-project/statediff/extern/filecoin-ffi v0.0.0-20201112005514-7550fe42bb9a

replace github.com/whyrusleeping/ledger-filecoin-go => github.com/filecoin-project/statediff/extern/fil-ledger v0.0.0-20201112214200-3592b9922dcc

replace github.com/ipld/go-ipld-prime => github.com/willscott/go-ipld-prime v0.0.4-0.20201117233610-e20af13830d8
//replace github.com/ipld/go-ipld-prime => ../go-ipld-prime

replace github.com/filecoin-project/go-hamt-ipld/v2 => github.com/willscott/go-hamt-ipld/v2 v2.0.1-0.20201115233521-38d62e432eed
