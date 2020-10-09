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
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/lib"
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
	},
}

func runVectorCmd(c *cli.Context) error {
	file, err := os.Open(vectorFlags.file)
	if err != nil {
		return err
	}

	var tv lib.TestVector
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

	statediff.AllowDegradedADLNodes = true

	preCid := tv.Pre.StateTree.RootCID
	postCid := tv.Post.StateTree.RootCID
	l, err := statediff.Transform(c.Context, preCid, store, "stateRoot")
	if err != nil {
		return fmt.Errorf("Could not load pre root: %s", err)
	}
	r, err := statediff.Transform(c.Context, postCid, store, "stateRoot")
	if err != nil {
		return fmt.Errorf("Could not load postroot: %s", err)
	}

	fmt.Printf("--- %s\n+++ %s\n@@ -1,1 +1,1 @@\n", preCid, postCid)
	fmt.Printf("%v\n", statediff.Diff(
		c.Context,
		store,
		l,
		r))

	return nil
}
