package main

import (
	"fmt"
	"os"

	bs "github.com/filecoin-project/lotus/blockstore"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/statediff"
)

var carFlags struct {
	file string
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
			Required:    true,
		},
	},
}

func runCarCmd(c *cli.Context) error {
	file, err := os.Open(carFlags.file)
	if err != nil {
		return err
	}

	store := bs.NewMemory()
	_, err = car.LoadCar(store, file)
	if err != nil {
		return err
	}

	if c.Args().Len() != 2 {
		return fmt.Errorf("Usage: statediff car --file <file> <pre CID> <post CID>")
	}
	preCid, err := cid.Parse(c.Args().Get(0))
	if err != nil {
		return err
	}
	postCid, err := cid.Parse(c.Args().Get(1))
	if err != nil {
		return err
	}

	statediff.AllowDegradedADLNodes = true

	d, err := PrintDiff(c.Context, preCid, postCid, store)
	if err != nil {
		return err
	}
	fmt.Printf("%s", d)

	return nil
}
