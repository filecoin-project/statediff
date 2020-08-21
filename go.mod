module github.com/filecoin-project/statediff

go 1.14

require (
	github.com/filecoin-project/go-address v0.0.3
	github.com/filecoin-project/go-bitfield v0.2.0
	github.com/filecoin-project/lotus v0.4.3-0.20200820203717-d1718369a182
	github.com/filecoin-project/specs-actors v0.9.3
	github.com/filecoin-project/test-vectors v0.0.0-20200819133914-e20cc29cc926 // indirect
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.4
	github.com/ipfs/go-hamt-ipld v0.1.1
	github.com/ipfs/go-ipld-cbor v0.0.5-0.20200428170625-a0bd04d3cbdf
	github.com/ipld/go-car v0.1.1-0.20200526133713-1c7508d55aae
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.2.2
	github.com/multiformats/go-multiaddr-net v0.1.5
	github.com/urfave/cli/v2 v2.2.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20200814224545-656e08ce49ee
	github.com/willscott/go-cmp v0.5.2-0.20200812183318-8affb9542345
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
replace github.com/filecoin-project/sector-storage => ./extern/sector-storage
