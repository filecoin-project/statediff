package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"

	bs "github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/ipld/go-car"
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

func runVectorCmd(c *cli.Context) error {
	file, err := os.Open(vectorFlags.file)
	if err != nil {
		return err
	}

	var tv TestVector
	if err := json.NewDecoder(file).Decode(&tv); err != nil {
		return err
	}

	buf := bytes.NewBuffer(tv.CAR)
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
