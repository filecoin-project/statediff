module github.com/filecoin-project/statediff

go 1.14

require (
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/evanw/esbuild v0.11.5
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-bitfield v0.2.4
	github.com/filecoin-project/go-bs-postgres-chainnotated v0.0.0-20210421230102-321bccf8b648
	github.com/filecoin-project/go-fil-markets v1.5.0
	github.com/filecoin-project/go-hamt-ipld/v2 v2.0.0
	github.com/filecoin-project/go-hamt-ipld/v3 v3.0.1
	github.com/filecoin-project/go-multistore v0.0.3
	github.com/filecoin-project/go-state-types v0.1.0
	github.com/filecoin-project/lotus v1.6.0
	github.com/filecoin-project/specs-actors v0.9.13
	github.com/filecoin-project/specs-actors/v2 v2.3.5-0.20210114162132-5b58b773f4fb
	github.com/filecoin-project/specs-actors/v3 v3.1.0
	github.com/gorilla/handlers v1.5.1
	github.com/graphql-go/graphql v0.7.9
	github.com/hashicorp/golang-lru v0.5.4
	github.com/ipfs/go-block-format v0.0.3
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.4.5
	github.com/ipfs/go-ipfs-blockstore v1.0.3
	github.com/ipfs/go-ipld-cbor v0.0.5
	github.com/ipfs/go-log/v2 v2.1.3
	github.com/ipld/go-car v0.1.1-0.20201119040415-11b6074b6d4d
	github.com/ipld/go-ipld-graphql v0.0.0-20210225034639-cece726cd342
	github.com/ipld/go-ipld-prime v0.7.0
	github.com/ipld/go-ipld-prime-proto v0.1.1 // indirect
	github.com/jackc/pgx/v4 v4.11.0
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multiaddr-net v0.2.0
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multihash v0.0.15
	github.com/polydawn/refmt v0.0.0-20190809202753-05966cbd336a
	github.com/urfave/cli/v2 v2.3.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20210303213153-67a261a1d291
	github.com/willscott/carbs v0.0.4
	github.com/willscott/go-cmp v0.5.2-0.20200812183318-8affb9542345
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/supranational/blst => github.com/filecoin-project/statediff/extern/fil-blst v0.0.0-20201112005514-7550fe42bb9a

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

//github.com/filecoin-project/statediff/extern/filecoin-ffi v0.0.0-20201112005514-7550fe42bb9a

replace github.com/whyrusleeping/ledger-filecoin-go => github.com/filecoin-project/statediff/extern/fil-ledger v0.0.0-20201112214200-3592b9922dcc

replace github.com/filecoin-project/go-hamt-ipld/v2 => github.com/willscott/go-hamt-ipld/v2 v2.0.1-0.20210225034344-6d6dfa9b3960
