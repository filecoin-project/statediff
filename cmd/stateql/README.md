StateQL
---

This binary provides a graphql server query endpoint.
The schema is generated from the filecoin schema defined in `/types/gen` in this repo.

Heavy lifting of the graphql tree generation in `schema.go` is performed through [go-ipld-prime](https://github.com/willscott/go-ipld-prime/tree/graphql/schema/gen/graphql/server) code generation.

(the code generation is triggered through `go generate ./...` at the root of this repo.)

Example Queries
---

Check that things are working (Get latest available tipset height)
```graphql
query Query {
  Height(at: -1) {
    Height
  }
}
```

Get the Init Actor's mapping of key->ID addresses
```graphql
query Query {
  Height(at: -1) {
    Height
    ParentStateRoot {
      Version
      Actors{
        At(key: "f01") {
        	Head {
            ... on InitV0State {
              NextID
              AddressMap{
                All {
                  Key
                  Value
                }
              }
            }
          }
        }
      }
    }
  }
}
```

Pull a parameter out of messages of a specific type
```graphql
fragment pci on MinerV0SectorPreCommitInfo {
  Expiration
  SectorNumber
}

query Query {
  Height(at: -1) {
    Height
    Messages {
      BlsMessages {
        AllOf(method: 6) {
          InterpretedParams(actorType: "storageMinerActorV2") {
            ...pci
          }
        }
      }
    }
  }
}
```
Note: the `actorType` string is defined as the value in the table of [mapped types](https://github.com/filecoin-project/statediff/blob/master/transform.go#L35)

A parameterized query for a Miner's Owner
```graphql
query Query($addr: RawAddress!) {
  Height(at: -1) {
    ParentStateRoot{
      Actors{
        At(key: $addr) {
          Head {
            ... on MinerV2State {
              Info {
                Owner
              }
            }
          }
        }
      }
    }
  }
}

with Query variable:
{
  "addr": "f02770"
}
```
