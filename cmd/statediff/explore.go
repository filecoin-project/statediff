package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/urfave/cli/v2"
)

var exploreCmd = &cli.Command{
	Name:        "explore",
	Description: "Examine a state tree in a browser",
	Action:      runExploreCmd,
	Flags: []cli.Flag{
		&apiFlag,
	},
}

func cidResolver(w http.ResponseWriter, r *http.Request) {

}

func runExploreCmd(c *cli.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/cid", cidResolver)
	mux.Handle("/", http.FileServer(AssetFile()))

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	s := &http.Server{
		Handler: mux,
	}
	fmt.Printf("Listening at %v\n", lis.Addr())
	return s.Serve(lis)
}
