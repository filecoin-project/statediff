package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"path"

	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/lib"
	"github.com/filecoin-project/statediff/build"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

var assetsFlag = cli.StringFlag{
	Name:  "assets",
	Usage: "Loads assets from disk at provided root (looking in '$assets/npm/app' and '$assets/cmd/stateexplorer/static'), rather than in-binary generated bundle",
	Value: ".",
}

var bindFlag = cli.StringFlag{
	Name:  "bind",
	Usage: "Bind to a specific local host:port. Specified as [address]:port",
	Value: ":0",
}

var exploreCmd = &cli.Command{
	Name:        "explore",
	Description: "Examine a state tree in a browser",
	Action:      runExploreCmd,
	Flags: []cli.Flag{
		&lib.ApiFlag,
		&assetsFlag,
		&bindFlag,
	},
}

func runExploreCmd(c *cli.Context) error {
	client, err := lib.GetAPI(c)
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

		as, ok := r.URL.Query()["as"]
		if !ok || len(as[0]) < 1 {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: no type")))
			return
		}

		parsed, err := cid.Parse(keys[0])
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %w", err)))
			return
		}

		transformed, err := statediff.Transform(r.Context(), parsed, store, as[0])
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		transformedBytes, err := json.Marshal(transformed)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(transformedBytes))
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
		mux.Handle("/", http.FileServer(http.Dir(path.Join(c.String(assetsFlag.Name), "cmd", "stateexplorer", "static"))))
	} else {
		mux.Handle("/", http.FileServer(AssetFile()))
	}

	lis, err := net.Listen("tcp", c.String(bindFlag.Name))
	if err != nil {
		return err
	}
	s := &http.Server{
		Handler: mux,
	}
	fmt.Printf("Listening at %v\n", lis.Addr())
	return s.Serve(lis)
}
