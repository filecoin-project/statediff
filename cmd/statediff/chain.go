package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/filecoin-project/lotus/api"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/lib"
)

var chainCmd = &cli.Command{
	Name:        "chain",
	Description: "Examine the state delta of an API object",
	Action:      runChainCmd,
	Flags: []cli.Flag{
		&lib.ApiFlag,
		&expandActorsFlag,
	},
}

func objToStateTree(ctx context.Context, api api.FullNode, obj string, after bool) (cid.Cid, error) {
	var height abi.ChainEpoch
	var baseKey types.TipSetKey
	toCompute := []*types.Message{}

	objCid, err := cid.Parse(obj)
	if err != nil {
		return cid.Undef, err
	}

	if msg, err := api.StateSearchMsg(ctx, objCid); err == nil {
		ts, err := api.ChainGetTipSet(ctx, msg.TipSet)
		if err != nil {
			return cid.Undef, err
		}

		if !after {
			return ts.ParentState(), nil
		}

		msgData, err := api.ChainGetMessage(ctx, objCid)
		if err != nil {
			return cid.Undef, err
		}
		height = ts.Height()
		baseKey = ts.Parents()
		toCompute = append(toCompute, msgData)
	} else if block, err := api.ChainGetBlock(ctx, objCid); err == nil {
		if !after {
			return block.ParentStateRoot, nil
		}

		msgs, err := api.ChainGetBlockMessages(ctx, objCid)
		if err != nil {
			return cid.Undef, err
		}
		toCompute = append(toCompute, msgs.BlsMessages...)
		height = block.Height
		baseKey = types.NewTipSetKey(block.Parents...)
	} else {
		return cid.Undef, fmt.Errorf("Could not parse '%s'", obj)
	}

	compute, err := api.StateCompute(ctx, height, toCompute, baseKey)
	if err != nil {
		return cid.Undef, err
	}
	fmt.Printf("after: %s\n", compute.Root)
	return compute.Root, nil
}

func runChainCmd(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return fmt.Errorf("no descriptor provided")
	}

	client, err := lib.GetAPI(c)
	if err != nil {
		return err
	}

	store := statediff.StoreFor(c.Context, client)

	var preCid, postCid cid.Cid

	obj := c.Args().Get(0)
	parts := strings.Split(obj, "..")

	if len(parts) == 1 {
		preCid, err = objToStateTree(c.Context, client, parts[0], false)
		if err != nil {
			return err
		}
		postCid, err = objToStateTree(c.Context, client, parts[0], true)
	} else if len(parts) == 2 {
		preCid, err = objToStateTree(c.Context, client, parts[0], false)
		if err != nil {
			return err
		}
		postCid, err = objToStateTree(c.Context, client, parts[1], true)
	} else {
		return fmt.Errorf("invalid descriptor: %s", obj)
	}
	if err != nil {
		return err
	}

	fmt.Printf("comparing state trees at %s-%s\n", preCid, postCid)

	if c.IsSet(expandActorsFlag.Name) {
		interestCids := c.String(expandActorsFlag.Name)
		opt := statediff.ExpandActors
		if len(interestCids) > 0 {
			opt, err = statediff.WithActorExpansionFromUser(interestCids)
			if err != nil {
				return err
			}
		}
		fmt.Printf("%v\n", statediff.Diff(
			c.Context,
			store,
			preCid,
			postCid,
			opt))
	} else {
		fmt.Printf("%v\n", statediff.Diff(
			c.Context,
			store,
			preCid,
			postCid))
	}

	return nil
}
