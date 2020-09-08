package main

import (
	"fmt"
	"os"

	bs "github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/ipld/go-car"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/statediff"
)

var carFlags struct {
	file string
	preCid string
	postCid string
}

var carCmd = &cli.Command{
	Name:        "car",
	Description: "Examine the state delta from a CAR",
	Action:      runCarCmd,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Usage:       "car store of cids",
			Destination: &carFlags.file,
			Required: true,
		},
		&cli.StringFlag{
			Name:        "preCid",
			Usage:       "initial state root",
			Destination: &carFlags.preCid,
			Required: true,
		},
		&cli.StringFlag{
			Name:        "postCid",
			Usage:       "final state root",
			Destination: &carFlags.postCid,
			Required: true,
		},
		&expandActorsFlag,
	},
}

func runCarCmd(c *cli.Context) error {
	file, err := os.Open(carFlags.file)
	if err != nil {
		return err
	}

	store := bs.NewTemporary()
	_, err = car.LoadCar(store, file)
	if err != nil {
		return err
	}

	preCid, err := cid.Parse(carFlags.preCid)
	if err != nil {
		return err
	}
	postCid, err := cid.Parse(carFlags.postCid)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", statediff.Diff(
		c.Context,
		store,
		preCid,
		postCid,
		statediff.ExpandActors))

	return nil
}
