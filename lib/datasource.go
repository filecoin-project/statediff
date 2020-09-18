// establish connection to a 'FullNode' filecoin API with cli configuration.
// Mostly with an eye towards ease of use for lotus users.
package lib

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/filecoin-project/statediff"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"github.com/mitchellh/go-homedir"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
	"github.com/urfave/cli/v2"
)

type StateRootFunc func(context.Context) []cid.Cid

var ApiFlag = cli.StringFlag{
	Name:    "api",
	Usage:   "api endpoint, formatted as token:multiaddr",
	Value:   "",
	EnvVars: []string{"FULLNODE_API_INFO"},
}

var CarFlag = cli.StringFlag{
	Name:  "car",
	Usage: "car file location for data source",
	Value: "",
}

var VectorFlag = cli.StringFlag{
	Name:  "vector",
	Usage: "test-vector.json file location for data source",
	Value: "",
}

func tryGetAPIFromHomeDir() ([]string, error) {
	p, err := homedir.Expand("~/.lotus")
	if err != nil {
		return nil, fmt.Errorf("Could not find API from file system. specify explicitly")
	}
	tokenPath := filepath.Join(p, "token")
	apiPath := filepath.Join(p, "api")

	token, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	apiData, err := ioutil.ReadFile(apiPath)
	if err != nil {
		return nil, err
	}
	return []string{string(token), string(apiData)}, nil
}

func GetAPI(c *cli.Context) (api.FullNode, error) {
	var err error
	api := c.String(ApiFlag.Name)
	sp := strings.SplitN(api, ":", 2)
	if len(sp) != 2 {
		sp, err = tryGetAPIFromHomeDir()
		if err != nil {
			return nil, fmt.Errorf("invalid api value, missing token or address: %s", api)
		}
	}

	token := sp[0]
	ma, err := multiaddr.NewMultiaddr(sp[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse provided multiaddr: %w", err)
	}

	_, dialAddr, err := manet.DialArgs(ma)
	if err != nil {
		return nil, fmt.Errorf("invalid api multiAddr: %w", err)
	}

	addr := "ws://" + dialAddr + "/rpc/v0"
	headers := http.Header{}
	if len(token) != 0 {
		headers.Add("Authorization", "Bearer "+token)
	}

	node, _, err := client.NewFullNodeRPC(c.Context, addr, headers)
	if err != nil {
		return nil, fmt.Errorf("could not connect to api: %w", err)
	}
	return node, nil
}

func GetHead(client api.FullNode) StateRootFunc {
	return func(ctx context.Context) []cid.Cid {
		head, err := client.ChainHead(ctx)
		if err != nil {
			return []cid.Cid{}
		}
		return head.Cids()
	}
}

func GetCar(c *cli.Context) (StateRootFunc, blockstore.Blockstore, error) {
	file, err := os.Open(c.String(CarFlag.Name))
	if err != nil {
		return nil, nil, err
	}

	store := blockstore.NewTemporary()
	root, err := car.LoadCar(store, file)
	if err != nil {
		return nil, nil, err
	}
	return func(_ context.Context) []cid.Cid { return root.Roots }, store, nil
}

func GetVector(c *cli.Context) (StateRootFunc, blockstore.Blockstore, error) {
	file, err := os.Open(c.String(VectorFlag.Name))
	if err != nil {
		return nil, nil, err
	}

	var tv TestVector
	if err := json.NewDecoder(file).Decode(&tv); err != nil {
		return nil, nil, err
	}

	b, err := base64.StdEncoding.DecodeString(tv.CAR)
	if err != nil {
		return nil, nil, err
	}
	buf := bytes.NewBuffer(b)
	gr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, nil, err
	}
	defer gr.Close()

	store := blockstore.NewTemporary()
	_, err = car.LoadCar(store, gr)
	if err != nil {
		return nil, nil, err
	}
	cids := []cid.Cid{tv.Pre.StateTree.RootCID, tv.Post.StateTree.RootCID}

	return func(_ context.Context) []cid.Cid { return cids }, store, nil
}

func GetBlockstore(c *cli.Context) (api.FullNode, StateRootFunc, blockstore.Blockstore, error) {
	if c.IsSet(CarFlag.Name) {
		srf, store, err := GetCar(c)
		return nil, srf, store, err
	}
	if c.IsSet(VectorFlag.Name) {
		srf, store, err := GetVector(c)
		return nil, srf, store, err
	}

	client, err := GetAPI(c)
	if err != nil {
		return nil, nil, nil, err
	}
	return client, GetHead(client), statediff.StoreFor(c.Context, client), nil
}
