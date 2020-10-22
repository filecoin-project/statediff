module github.com/filecoin-project/statediff

go 1.14

require (
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/evanw/esbuild v0.7.19
	github.com/filecoin-project/go-address v0.0.4
	github.com/filecoin-project/go-bitfield v0.2.1
	github.com/filecoin-project/go-state-types v0.0.0-20200928172055-2df22083d8ab
	github.com/filecoin-project/lotus v0.9.0
	github.com/filecoin-project/specs-actors v0.9.12
	github.com/filecoin-project/specs-actors/v2 v2.0.1
	github.com/gorilla/handlers v1.5.1
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.5
	github.com/ipfs/go-hamt-ipld v0.1.1
	github.com/ipfs/go-ipfs-blockstore v1.0.1
	github.com/ipfs/go-ipld-cbor v0.0.5-0.20200428170625-a0bd04d3cbdf
	github.com/ipld/go-car v0.1.1-0.20200923150018-8cdef32e2da4
	github.com/ipld/go-ipld-prime v0.5.1-0.20200910124733-350032422383
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multiaddr-net v0.2.0
	github.com/polydawn/refmt v0.0.0-20190809202753-05966cbd336a
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/urfave/cli/v2 v2.2.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20200826160007-0b9f6c5fb163
	github.com/willscott/go-cmp v0.5.2-0.20200812183318-8affb9542345
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)

//replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

replace github.com/filecoin-project/filecoin-ffi => github.com/filecoin-project/statediff/extern/filecoin-ffi v0.0.0-20201008211259-9fe031c6c661
