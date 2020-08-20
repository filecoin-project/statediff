package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	bs "github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/ipld/go-car"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/statediff"
)

var vectorFlags struct {
	file string
}

var vectorCmd = &cli.Command{
	Name:        "vector",
	Description: "Examine the state delta of a test vector",
	Action:      runVectorCmd,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Usage:       "compare the pre and post states of a test vector",
			Destination: &vectorFlags.file,
		},
		&expandActorsFlag,
	},
}

type TestVector struct {
	CAR string `json:"car"`

	Pre           *Preconditions  `json:"preconditions"`
	Post          *Postconditions `json:"postconditions"`
}

type StateTree struct {
	RootCID cid.Cid `json:"root_cid"`
}

type Preconditions struct {
	StateTree *StateTree     `json:"state_tree"`
}

// Postconditions contain a representation of VM state at th end of the test
type Postconditions struct {
	StateTree *StateTree `json:"state_tree"`
}


func runVectorCmd(c *cli.Context) error {
	file, err := os.Open(vectorFlags.file)
	if err != nil {
		return err
	}

	var tv TestVector
	if err := json.NewDecoder(file).Decode(&tv); err != nil {
		return err
	}

	b, err := base64.StdEncoding.DecodeString(tv.CAR)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)
	gr, err := gzip.NewReader(buf)
	if err != nil {
		return err
	}
	defer gr.Close()

	store := bs.NewTemporary()
	_, err = car.LoadCar(store, gr)
	if err != nil {
		return err
	}

	fmt.Printf("Diff: %v\n", statediff.Diff(
		c.Context,
		store,
		tv.Pre.StateTree.RootCID,
		tv.Post.StateTree.RootCID,
		statediff.ExpandActors))

	return nil
}
