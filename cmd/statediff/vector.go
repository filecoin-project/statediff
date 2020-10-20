package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	bs "github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
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

	d, err := PrintDiff(c.Context, preCid, postCid, store)
	if err != nil {
		return err
	}
	fmt.Printf("%s", d)

	return nil
}

func PrintDiff(ctx context.Context, pre, post cid.Cid, store blockstore.Blockstore) (string, error) {
	l, err := statediff.Transform(ctx, pre, store, "stateRoot")
	if err != nil {
		l, err = statediff.Transform(ctx, pre, store, "versionedStateRoot")
		if err != nil {
			return "", fmt.Errorf("Could not load left root: %s", err)
		}
	}
	r, err := statediff.Transform(ctx, post, store, "stateRoot")
	if err != nil {
		r, err = statediff.Transform(ctx, post, store, "versionedStateRoot")
		if err != nil {
			return "", fmt.Errorf("Could not load right root: %s", err)
		}
	}

	s := fmt.Sprintf("--- %s\n+++ %s\n@@ -1,1 +1,1 @@\n", pre, post)
	s += fmt.Sprintf("%v\n", statediff.Diff(
		ctx,
		store,
		l,
		r))
	return s, nil
}
