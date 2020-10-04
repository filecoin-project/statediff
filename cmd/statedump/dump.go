package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/codec/fcjson"
	"github.com/filecoin-project/statediff/lib"
	"github.com/ipfs/go-cid"
	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/urfave/cli/v2"
)

var byActorFlag = cli.BoolFlag{
	Name:  "byActor",
	Usage: "When set, one line of json is output for each actor",
	Value: false,
}

var profileFlag = cli.StringFlag{
	Name:  "profile",
	Usage: "collect cpu profile",
}

var dumpCmd = &cli.Command{
	Name:        "dump",
	Description: "Dump a state tree to json",
	Action:      runDumpCmd,
	Flags: []cli.Flag{
		&lib.ApiFlag,
		&lib.CarFlag,
		&lib.VectorFlag,
		&profileFlag,
		&byActorFlag,
	},
}

var accountFilter []string
var noAccountExpand bool

var denormalizedTipsetKeys = []string{"BeaconEntries", "ElectionProof", "ForkSignaling", "Height", "Miner", "ParentBaseFee", "Timestamp"}

func runDumpCmd(c *cli.Context) error {
	if c.String(profileFlag.Name) != "" {
		f, err := os.Create(c.String(profileFlag.Name))
		if err != nil {
			return err
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	_, head, store, err := lib.GetBlockstore(c)
	if err != nil {
		return err
	}

	heads := head(c.Context)
	// TODO: support alternate root than '1st head'
	h := heads[0]
	// TODO: support alternate start 'as'
	as := "tipset"

	tipset, err := statediff.Transform(c.Context, h, store, as)
	if err != nil {
		return err
	}

	follow := func(n ipld.Node, to ipld.Path, as string) ipld.Node {
		for _, s := range to.Segments() {
			node, err := n.LookupBySegment(s)
			if err != nil {
				return nil
			}
			n = node
		}
		l, err := n.AsLink()
		if err != nil {
			return nil
		}
		asCid, ok := l.(cidlink.Link)
		if !ok {
			return nil
		}
		followed, err := statediff.Transform(c.Context, asCid.Cid, store, as)
		if err != nil {
			return nil
		}
		return followed
	}

	stateroot := follow(tipset, ipld.ParsePath("ParentStateRoot"), "stateRoot")

	for _, acct := range c.Args().Slice() {
		asAddr, err := address.NewFromString(acct)
		if err != nil {
			return err
		}
		t, err := statediff.TypeOfActor(stateroot, string(asAddr.Bytes()))
		if err != nil {
			return err
		}
		actor := follow(stateroot, ipld.NewPath([]ipld.PathSegment{ipld.PathSegmentOfString(string(asAddr.Bytes())), ipld.PathSegmentOfString("Head")}), t)
		if actor == nil {
			return fmt.Errorf("invalid actor %s", acct)
		}

		p := fcjson.DagMarshaler{
			Path: ipld.ParsePath(t),
			Loader: func(ci cid.Cid, p ipld.Path) ipld.Node {
				if ci.Prefix().Codec == cid.FilCommitmentSealed || ci.Prefix().Codec == cid.FilCommitmentUnsealed {
					return nil
				}
				node, err := statediff.Transform(c.Context, ci, store, p.String())
				if err == nil {
					return node
				}
				fmt.Printf("err: %v\n", err)
				return nil
			},
		}

		ob := bytes.NewBuffer(nil)
		if err := p.Encoder(actor, ob); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}
		_, err = ob.WriteTo(os.Stdout)
		if err != nil {
			return err
		}
	}
	return nil
}
