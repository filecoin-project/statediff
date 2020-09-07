// establish connection to a 'FullNode' filecoin API with cli configuration.
// Mostly with an eye towards ease of use for lotus users.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/mitchellh/go-homedir"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
	"github.com/urfave/cli/v2"
)

var apiFlag = cli.StringFlag{
	Name:    "api",
	Usage:   "api endpoint, formatted as token:multiaddr",
	Value:   "",
	EnvVars: []string{"FULLNODE_API_INFO"},
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
	api := c.String(apiFlag.Name)
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
