package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/filecoin-project/statediff/codec/fcjson"

	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/build"
	gqlib "github.com/filecoin-project/statediff/cmd/stateql/lib"
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

var qlFlag = cli.BoolFlag{
	Name:  "graphql",
	Usage: "enable graphql endpoint",
	Value: true,
}

var exploreCmd = &cli.Command{
	Name:        "explore",
	Description: "Examine a state tree in a browser",
	Action:      runExploreCmd,
	Flags: []cli.Flag{
		&lib.ApiFlag,
		&lib.CarFlag,
		&lib.SqlFlag,
		&lib.VectorFlag,
		&lib.NoCacheFlag,
		&assetsFlag,
		&bindFlag,
		&qlFlag,
	},
}

func runExploreCmd(c *cli.Context) error {
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
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}

		ds, err := lib.Lazy(c)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		txCtx := context.WithValue(r.Context(), statediff.StateRootKey, ds.Head(r.Context()))
		transformed, err := statediff.Transform(txCtx, parsed, ds.Store(), as[0])
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = fcjson.Encoder(transformed, w)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
		}
	}

	headResolver := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ds, err := lib.Lazy(c)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		blkHdr := ds.Head(r.Context())
		cidBytes, _ := json.Marshal([]cid.Cid{blkHdr})
		w.Header().Set("Content-Type", "application/json")
		w.Write(cidBytes)
	}

	heightResolver := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h, ok := r.URL.Query()["h"]
		ds, err := lib.Lazy(c)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		if !ok || len(h[0]) < 1 || ds.Store() == nil {
			w.Write([]byte(fmt.Sprintf("error: no height")))
			return
		}

		asNumber, err := strconv.Atoi(h[0])
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error: invalid height")))
			return
		}
		tipset, err := ds.CidAtHeight(r.Context(), int64(asNumber))
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}
		w.Write([]byte(tipset.String()))
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

	if c.Bool(qlFlag.Name) {
		ds, err := lib.Lazy(c)
		if err != nil {
			return err
		}
		gqMux := gqlib.GetGraphQL(c, ds)
		mux.Handle("/graphql", gqMux)
		mux.Handle("/graphql.html", gqMux)
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
