package main

import (
	"fmt"
	"github.com/filecoin-project/statediff/lib"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

var changedFlag = cli.BoolFlag{
	Name:  "changed",
	Usage: "When set, only actors which have changed since previous root.",
	Value: false,
}

var accountCmd = &cli.Command{
	Name:        "account",
	Description: "Dump accounts",
	Action:      runAccountCmd,
	Flags: []cli.Flag{
		&lib.ApiFlag,
		&lib.CarFlag,
		&lib.VectorFlag,
		&changedFlag,
	},
}

func runAccountCmd(c *cli.Context) error {
	_, head, store, err := lib.GetBlockstore(c)
	if err != nil {
		return err
	}

	heads := head(c.Context)
	// TODO: support alternate root than '1st head'
	h := heads[0]
	// TODO: support alternate start 'as'
	as := "tipset"

	// Load Data
	noAccountExpand = true
	stateTree, err := ToJSON(c.Context, store, as, h)
	if err != nil {
		return err
	}
	root := mustMap(stateTree)
	sr := mustMap(root["ParentStateRoot"])

	if !c.Bool(changedFlag.Name) {
		for a, _ := range sr {
			fmt.Printf("%s\n", a)
		}

		return nil
	}
	pcs := mustList(root["Parents"])
	pcids, ok := pcs[0].(string)
	if !ok {
		return fmt.Errorf("%v not cid", pcs[0])
	}
	pcid, err := cid.Parse(pcids)
	if err != nil {
		return err
	}
	prev, err := ToJSON(c.Context, store, as, pcid)
	if err != nil {
		return err
	}
	pr := mustMap(prev)
	psr := mustMap(pr["ParentStateRoot"])
	for a, acct := range sr {
		pacct, ok := psr[a]
		if !ok {
			fmt.Printf("%s\n", a)
			continue
		}
		acm := mustMap(acct)
		pacm := mustMap(pacct)
		if acm["HeadCID"] != pacm["HeadCID"] {
			fmt.Printf("%s\n", a)
		}
	}
	return nil
}
