# State exploration for filecoin (lotus centric)

This tool provides functionality for exploration and comparing changes to filecoin state.

**Contents**

- [Installation](#installation)
- [Usage](#usage)
- [API](#API)
- [License](#license)


## CLI

```bash
go get github.com/filecoin-project/statediff/cmd/statediff
```

### Usage

See what state change is expected by a [test vector](https://github.com/filecoin-project/test-vectors):

```
statediff vector --file vector.json 
```

See what state changed on the local lotus chain across a block or message:

```
statediff chain --expand-actors all bafy...
```

Compare two roots in a CAR
```
statediff car --file export.car <left CID> <right CID>
```

## API

Other tools working with state trees can access statediff by importing `github.com/filecoin-project/statediff`.

* `Diff(context.Context, blockstore.Blockstore, a, b cid.Cid, ...Option) string`
Diff generates a textual unified diff between stateroots `a` and `b`.
State objects are retreived out of the provided [blockstore](https://github.com/ipfs/go-ipfs-blockstore). 
  * Options can be used to control the amount of recursion / expansion performed.
    In particular, `ExpandActors` will perform recursive introspection into each
    individual actor account with a differing HEAD state, and `ExpandActorByCid`
    will selectively expand actor accounts based on provided CIDs.

## Web

The web viewer (`stateexplorer`) provides a JSON transformation layer and interactive
state tree exploration.

```bash
git clone https://github.com/filecoin-project/statediff
go generate ./...
```
Note: you need `npm` installed for `go generate` to properly bundle frontend assets.

When running in production, either use the provided docker image,
or run the explorer in its self-contained-binary form
(assets are bundled via the `go generate` above.)

```
go run ./cmd/stateexplorer explore --bind 0.0.0.0:33333 --api $(cat ~/.lotus/token):$(cat ~/.lotus/api)
```

If not explicitly provided as an argument, statediff/stateexplorer will attempt to locate a lotus instance running on the same host by probing your home directory.

In development mode, assets can be loaded from disk so that changes are reflected on reload.

```
go run ./cmd/stateexplorer explore --bind 0.0.0.0:33333 --assets .
```

## License

Dual-licensed under [MIT](https://github.com/filecoin-project/statediff/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/statediff/blob/master/LICENSE-APACHE)
