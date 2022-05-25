module github.com/filecoin-project/statediff

go 1.14

require (
	github.com/evanw/esbuild v0.12.10
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-amt-ipld/v2 v2.1.1-0.20201006184820-924ee87a1349 // indirect
	github.com/filecoin-project/go-bitfield v0.2.4
	github.com/filecoin-project/go-bs-postgres-chainnotated v0.0.0-20210421230102-321bccf8b648
	github.com/filecoin-project/go-fil-markets v1.13.1
	github.com/filecoin-project/go-hamt-ipld/v2 v2.0.0
	github.com/filecoin-project/go-hamt-ipld/v3 v3.1.0
	github.com/filecoin-project/go-state-types v0.1.1-0.20210915140513-d354ccf10379
	github.com/filecoin-project/lotus v1.13.0
	github.com/filecoin-project/specs-actors v0.9.14
	github.com/filecoin-project/specs-actors/v2 v2.3.5
	github.com/filecoin-project/specs-actors/v3 v3.1.1
	github.com/filecoin-project/specs-actors/v6 v6.0.0
	github.com/gorilla/handlers v1.5.1
	github.com/graphql-go/graphql v0.7.9
	github.com/hashicorp/golang-lru v0.5.4
	github.com/ipfs/go-block-format v0.0.3
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-datastore v0.4.6
	github.com/ipfs/go-ipfs-blockstore v1.0.4
	github.com/ipfs/go-ipld-cbor v0.0.5
	github.com/ipfs/go-log/v2 v2.3.0
	github.com/ipld/go-car v0.3.2-0.20211001225732-32d0d9933823
	github.com/ipld/go-ipld-graphql v0.0.0-20211021213353-9727002b9c62
	github.com/ipld/go-ipld-prime v0.12.4-0.20211019203807-9ddc75a0c4ae
	github.com/jackc/pgx/v4 v4.11.0
	github.com/lib/pq v1.7.0 // indirect
	github.com/libp2p/go-libp2p-core v0.16.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.4.1
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multihash v0.0.15
	github.com/polydawn/refmt v0.0.0-20201211092308-30ac6d18308e
	github.com/urfave/cli/v2 v2.3.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20210713220151-be142a5ae1a8
	github.com/willscott/carbs v0.0.4
	github.com/willscott/go-cmp v0.5.2-0.20200812183318-8affb9542345
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/whyrusleeping/ledger-filecoin-go => github.com/filecoin-project/statediff/extern/fil-ledger v0.0.0-20201112214200-3592b9922dcc

replace github.com/filecoin-project/go-hamt-ipld/v2 => github.com/willscott/go-hamt-ipld/v2 v2.0.1-0.20211020201145-202c9a6a5241

replace (
	github.com/filecoin-project/fil-blst => ./extern/fil-blst
	github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
	github.com/supranational/blst => ./extern/fil-blst/blst
)
