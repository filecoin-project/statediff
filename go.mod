module github.com/filecoin-project/statediff

go 1.14

require (
	github.com/evanw/esbuild v0.6.28
	github.com/filecoin-project/chain-validation v0.0.6-0.20200907020853-f4e4e7417fea // indirect
	github.com/filecoin-project/go-address v0.0.3
	github.com/filecoin-project/go-bitfield v0.2.0
	github.com/filecoin-project/go-state-types v0.0.0-20200911004822-964d6c679cfc
	github.com/filecoin-project/lotus v0.7.2
	github.com/filecoin-project/specs-actors v0.9.10
	github.com/filecoin-project/test-vectors v0.0.0-20200903223506-84da0a5ea125 // indirect
	github.com/go-bindata/go-bindata v3.1.2+incompatible // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.4
	github.com/ipfs/go-hamt-ipld v0.1.1
	github.com/ipfs/go-ipfs-blockstore v1.0.1
	github.com/ipfs/go-ipld-cbor v0.0.5-0.20200428170625-a0bd04d3cbdf
	github.com/ipld/go-car v0.1.1-0.20200923150018-8cdef32e2da4
	github.com/ipld/go-ipld-prime v0.5.1-0.20200910124733-350032422383
	github.com/ipld/go-ipld-prime-proto v0.0.0-20200922192210-9a2bfd4440a6 // indirect
	github.com/mailru/easyjson v0.0.0-20190312143242-1de009706dbe
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multiaddr-net v0.2.0
	github.com/polydawn/refmt v0.0.0-20190809202753-05966cbd336a
	github.com/urfave/cli/v2 v2.2.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20200814224545-656e08ce49ee
	github.com/willscott/go-cmp v0.5.2-0.20200812183318-8affb9542345
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

//replace github.com/filecoin-project/filecoin-ffi => github.com/filecoin-project/statediff/extern/filecoin-ffi v0.0.0-20200904233626-6a3c8611ff64
