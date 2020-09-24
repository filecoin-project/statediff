package main

import (
	"github.com/ipld/go-ipld-prime/schema"
)

func accumulateABI(ts schema.TypeSystem) {
	ts.Accumulate(schema.SpawnInt("MethodNum"))
}
