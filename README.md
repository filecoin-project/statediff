# State exploration for filecoin (lotus centric)

This tool provides functionality for exploration and comparing changes to filecoin state.

**Contents**

- [Installation](#installation)
- [Usage](#usage)
- [API](#API)
- [License](#license)


## Installation

```bash
go get github.com/filecoin-project/statediff/cmd/statediff
```

## Usage

See what state change is expected by a [test vector](https://github.com/filecoin-project/test-vectors):

```
statediff vector --file vector.json 
```

See what state chaned on the local lotus chain across a block or message:

```
statediff chain --expand-actors all bafy...
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

## License

Dual-licensed under [MIT](https://github.com/filecoin-project/statediff/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/statediff/blob/master/LICENSE-APACHE)
