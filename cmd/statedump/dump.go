package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/lib"
	"github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

var byActorFlag = cli.BoolFlag{
	Name: "byActor",
	Usage: "When set, one line of json is output for each actor",
	Value: false,
}

var profileFlag = cli.StringFlag{
	Name: "profile",
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

	accts := c.Args().Slice()
	if len(accts) > 0 {
		accountFilter = accts
	}

	stateTree, err := ToJSON(c.Context, store, as, h)
	if err != nil {
		return err
	}
	if c.Bool(byActorFlag.Name) {
		base := mustMap(stateTree)
		root := mustMap(base["ParentStateRoot"])
		for acct, m := range root {
			acctMap := mustMap(m)
			acctMap["Address"] = acct
			for _, k := range denormalizedTipsetKeys {
				acctMap[k] = base[k]
			}
			acctBytes, err := json.Marshal(acctMap)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", string(acctBytes))
		}
	} else {
		bytes, err := json.Marshal(stateTree)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(bytes))
	}
	return nil
}


func ToJSON(ctx context.Context, store blockstore.Blockstore, as string, c cid.Cid) (interface{}, error) {
	// Load Data
	transformed, err := statediff.Transform(ctx, c, store, as)
	if err != nil {
		return nil, err
	}
	
	// Destructure to untyped.
	flat, err := json.Marshal(transformed)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(flat, &m); err != nil {
		a := make([]interface{}, 0)
		if err := json.Unmarshal(flat, &a); err == nil {
			return a, nil
		}

		s := ""
		if err := json.Unmarshal(flat, &s); err == nil {
			return s, nil
		}
		return nil, err
	}

	realAs := statediff.ResolveType(as)
	// special case hacks.
	if realAs == statediff.LotusTypeTipset {
		unfollow(m, "ParentMessageReceipts")
		unfollow(m, "Messages")
		unfollow(m, "Parents")
	} else if realAs == statediff.LotusTypeStateroot {
		if accountFilter != nil {
			m2 := make(map[string]interface{})
			for _, a := range accountFilter {
				m2[a] = m[a]
			}
			m = m2
		}
		for _, val := range m {
			vm := mustMap(val)
			unfollow(vm, "Code")
			t := statediff.LotusActorCodes[asString(vm, "Code")]
			vm["HeadCID"] = vm["Head"]
			unfollow(vm, "HeadCID")
			if !noAccountExpand {
				vm["Head"] = followObj(ctx, store, string(t), mustMap(vm["Head"]))
			}
			vm["Type"] = t
		}
		return m, nil
	}

	// recursive expansions.
	followObj(ctx, store, string(realAs), m)

	return m, nil
}

func asString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if n, ok := v.(string); ok {
			return n
		}
	}
	return ""
}

func mustList(m interface{}) []interface{} {
	if ms, ok := m.([]interface{}); ok {
		return ms
	}
	fmt.Printf("%v not list!\n", m)
	return []interface{}{}
}

func mustMap(m interface{}) map[string]interface{} {
	if b, ok := m.(map[string]interface{}); ok {
		return b
	}
	return make(map[string]interface{})
}

func unfollow(m map[string]interface{}, k string) {
	if a, ok := m[k].(map[string]interface{}); ok {
		m[k] = unfollowObj(a)
	} else if b, ok := m[k].([]interface{}); ok {
		m[k] = unfollowSlice(b)
	}
}

func unfollowObj(m map[string]interface{}) interface{} {
	if c, ok := m["/"]; ok && len(m) == 1 {
		return c;
	}
	for k2, _ := range m {
		unfollow(m, k2)
	}
	return m
}

func unfollowSlice(l []interface{}) []interface{} {
	n := make([]interface{}, 0, len(l))
	for i := range l {
		if a, ok := l[i].(map[string]interface{}); ok {
			n = append(n, unfollowObj(a))
		} else if b, ok := l[i].([]interface{}); ok {
			n = append(n, unfollowSlice(b))
		} else {
			n = append(n, l[i])
		}
	}
	return n
}

func follow(ctx context.Context, store blockstore.Blockstore, as string, m map[string]interface{}, k string) {
	if a, ok := m[k].(map[string]interface{}); ok {
		m[k] = followObj(ctx, store, as + "." + k, a)
	} else if b, ok := m[k].([]interface{}); ok {
		m[k] = followSlice(ctx, store, as + "." + k, b)
	}
}

func followObj(ctx context.Context, store blockstore.Blockstore, as string, m map[string]interface{}) interface{} {
	if c, ok := m["/"]; ok && len(m) == 1 {
		asCid, err := cid.Parse(c)
		if err != nil {
			return fmt.Sprintf("%s", err)
		}
		if asCid.Type() != cid.DagCBOR {
			return c
		}
		j, err := ToJSON(ctx, store, as, asCid)
		if err != nil {
			return fmt.Sprintf("%s", err)
		}
		return j;
	}
	for k2, _ := range m {
		follow(ctx, store, as, m, k2)
	}
	return m
}

func followSlice(ctx context.Context, store blockstore.Blockstore, as string, l []interface{}) []interface{} {
	n := make([]interface{}, 0, len(l))
	for i := range l {
		if a, ok := l[i].(map[string]interface{}); ok {
			n = append(n, followObj(ctx, store, as, a))
		} else if b, ok := l[i].([]interface{}); ok {
			n = append(n, followSlice(ctx, store, as, b))
		} else {
			n = append(n, l[i])
		}
	}
	return n
}
