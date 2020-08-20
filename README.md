# State exploration for filecoin (lotus centric)

This tool provides functionality for exploration and comparing changes to filecoin state.

** Contents **

- [Installation](#installation)
- [Usage](#usage)
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

## License

Dual-licensed under [MIT](https://github.com/filecoin-project/statediff/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/statediff/blob/master/LICENSE-APACHE)
