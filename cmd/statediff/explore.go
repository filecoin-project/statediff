package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"path"

	"github.com/urfave/cli/v2"
	"github.com/ipfs/go-cid"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/build"
)

var assetsFlag = cli.StringFlag{
	Name:  "assets",
	Usage: "Loads assets from disk at provided root (looking in '$assets/npm/app' and '$assets/cmd/statediff/static'), rather than in-binary generated bundle",
	Value: ".",
}

var exploreCmd = &cli.Command{
	Name:        "explore",
	Description: "Examine a state tree in a browser",
	Action:      runExploreCmd,
	Flags: []cli.Flag{
		&apiFlag,
		&assetsFlag,
	},
}

func runExploreCmd(c *cli.Context) error {
	client, err := GetAPI(c)
	if err != nil {
		return err
	}

	store := statediff.StoreFor(c.Context, client)

	cidResolver := func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["cid"]
		if !ok || len(keys[0]) < 1 {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: no cid")))
			return
		}

		parsed, err := cid.Parse(keys[0])
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %w", err)))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = store.Get(parsed)
	}

	headResolver := func(w http.ResponseWriter, r *http.Request) {
		head, err := client.ChainHead(r.Context())
		w.Header().Set("Content-Type", "text/plain")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error: %w", err)))
		} else {
			w.Write([]byte(head.Key().String()))
		}
	}


	mux := http.NewServeMux()
	mux.HandleFunc("/cid", cidResolver)
	mux.HandleFunc("/head", headResolver)
	if c.IsSet(assetsFlag.Name) {
		scriptResolver := func(w http.ResponseWriter, r *http.Request) {
			data := build.Compile(path.Join(c.String(assetsFlag.Name), "npm", "app"))
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, data)
		}
		mux.HandleFunc("/app.js", scriptResolver)
		mux.Handle("/", http.FileServer(http.Dir(path.Join(c.String(assetsFlag.Name),"cmd","statediff","static"))))
	} else {
		mux.Handle("/", http.FileServer(AssetFile()))
	}

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
