package lib

import (
	"github.com/ipfs/go-cid"
)

type TestVector struct {
	CAR string `json:"car"`

	Pre  *Preconditions  `json:"preconditions"`
	Post *Postconditions `json:"postconditions"`
}

type StateTree struct {
	RootCID cid.Cid `json:"root_cid"`
}

type Preconditions struct {
	StateTree *StateTree `json:"state_tree"`
}

// Postconditions contain a representation of VM state at th end of the test
type Postconditions struct {
	StateTree *StateTree `json:"state_tree"`
}
