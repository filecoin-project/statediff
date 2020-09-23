package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/build"
	"github.com/filecoin-project/statediff/lib"
	"github.com/gorilla/handlers"
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
		&lib.CarFlag,
		&lib.VectorFlag,
		&assetsFlag,
		&bindFlag,
	},
}

func runExploreCmd(c *cli.Context) error {
	client, head, store, err := lib.GetBlockstore(c)
	if err != nil {
		return err
	}

	cidResolver := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "public, max-age=604800, immutable")
		w.Header().Set("Expires", time.Now().Add(7*24*time.Hour).Format(http.TimeFormat))
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		cid := head(r.Context())
		cidBytes, _ := json.Marshal(cid)
		w.Header().Set("Content-Type", "application/json")
		w.Write(cidBytes)
	}

	heightResolver := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h, ok := r.URL.Query()["h"]
		w.Header().Set("Content-Type", "text/plain")
		if !ok || len(h[0]) < 1 || client == nil {
			w.Write([]byte(fmt.Sprintf("error: no height")))
			return
		}

		asNumber, err := strconv.Atoi(h[0])
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error: invalid height")))
			return
		}
		tipset, err := client.ChainGetTipSetByHeight(r.Context(), abi.ChainEpoch(asNumber), types.EmptyTSK)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		w.Write([]byte(tipset.Key().String()))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cid", cidResolver)
	mux.HandleFunc("/head", headResolver)
	mux.HandleFunc("/height", heightResolver)
	if c.IsSet(assetsFlag.Name) {
		scriptResolver := func(w http.ResponseWriter, r *http.Request) {
			data := build.Compile(path.Join(c.String(assetsFlag.Name), "npm", "app"), false)
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
		Handler: handlers.LoggingHandler(os.Stdout, mux),
	}
	fmt.Printf("Listening at %v\n", lis.Addr())
	return s.Serve(lis)
}
