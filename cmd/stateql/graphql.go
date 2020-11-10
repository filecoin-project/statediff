package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/statediff"
	"github.com/filecoin-project/statediff/lib"
	"github.com/filecoin-project/statediff/types"

	//"github.com/filecoin-project/statediff/types"
	"github.com/gorilla/handlers"
	"github.com/graphql-go/graphql"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"

	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/urfave/cli/v2"
)

var client api.FullNode
var head statediff.StateRootFunc
var store blockstore.Blockstore

const storeCtx = "store"

var tipsetMap map[int]cid.Cid

func index(ctx context.Context, store blockstore.Blockstore, head statediff.StateRootFunc) {
	fmt.Printf("running jank tipset index\n")
	n := head(ctx)
	if len(n) == 0 {
		fmt.Printf("... or not: no head provided.\n")
		return
	}
	tipsetMap = make(map[int]cid.Cid)
	c := n[0]
	for true {
		p := types.Type.LotusBlockHeader__Repr.NewBuilder()
		if err := statediff.Load(ctx, c, store, p); err != nil {
			fmt.Printf("exiting index. no header @ %s\n", c)
			return
		}
		n := p.Build()
		h, err := n.LookupByString("Height")
		if err != nil {
			fmt.Printf("exiting index. no Height on header\n")
			return
		}
		hn, err := h.AsInt()
		if err != nil {
			fmt.Printf("exiting index. height not interpreted as int\n")
			return
		}
		tipsetMap[hn] = c

		parents, err := n.LookupByString("Parents")
		if err != nil {
			fmt.Printf("exiting index. no parents on %s\n", c)
			return
		}
		_, parentCid, err := parents.ListIterator().Next()
		if err != nil {
			fmt.Printf("exiting index. parents on %s empty\n", c)
			return
		}
		pcc, err := parentCid.AsLink()
		if err != nil {
			fmt.Printf("exiting index. parents on %s not link\n", c)
			return
		}
		c = pcc.(cidlink.Link).Cid
	}
}

func lazy(c *cli.Context) error {
	if client != nil {
		return nil
	}
	var err error
	client, head, store, err = lib.GetBlockstore(c)
	if err != nil {
		return err
	}
	if client == nil && tipsetMap == nil {
		index(c.Context, store, head)
		fmt.Printf("fake metadata thing can serve %d tipsets\n", len(tipsetMap))
	}
	return nil
}

var bindFlag = cli.StringFlag{
	Name:  "bind",
	Usage: "Bind to a specific local host:port. Specified as [address]:port",
	Value: ":0",
}

var graphCmd = &cli.Command{
	Name:        "serve",
	Description: "Serve GraphQL endpoint",
	Action:      runGraphCmd,
	Flags: []cli.Flag{
		&lib.ApiFlag,
		&lib.CarFlag,
		&lib.VectorFlag,
		&bindFlag,
	},
}

func CorsMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next(w, r)
	})
}

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func runGraphCmd(c *cli.Context) error {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"Tipset": &graphql.Field{
					Type: LotusBlockHeader__type,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.ID,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						s, ok := p.Context.Value(storeCtx).(blockstore.Blockstore)
						if !ok {
							return nil, fmt.Errorf("no datastore provided")
						}
						id := p.Args["id"]
						idCid, err := cid.Parse(id)
						if err != nil {
							return nil, err
						}
						return statediff.Transform(p.Context, idCid, s, string(statediff.LotusTypeTipset))
					},
				},
				"Heights": &graphql.Field{
					Type: graphql.NewList(LotusBlockHeader__type),
					Args: graphql.FieldConfigArgument{
						"from": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"to": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						s, ok := p.Context.Value(storeCtx).(blockstore.Blockstore)
						if !ok {
							return nil, fmt.Errorf("no datastore provided")
						}
						from := p.Args["from"].(int)
						to := p.Args["to"].(int)

						out := make([]ipld.Node, 0, to-from)
						for i := from; i < to; i++ {
							ch := tipsetMap[i]
							n, err := statediff.Transform(p.Context, ch, s, string(statediff.LotusTypeTipset))
							if err != nil {
								return nil, err
							}
							out = append(out, n)
						}
						return out, nil
					},
				},
			},
		}),
	})
	if err != nil {
		log.Fatal(err)
	}

	loader := func(ctx context.Context, cl cidlink.Link, builder ipld.NodeBuilder) error {
		return statediff.Load(ctx, cl.Cid, store, builder)
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", CorsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if err := lazy(c); err != nil {
			log.Printf("failed to connect to backend: %v", err)
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("backend error:%v", err)))
			return
		}

		sctx := context.WithValue(r.Context(), storeCtx, store)
		ctx := context.WithValue(sctx, nodeLoaderCtxKey, loader)
		var result *graphql.Result
		if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
			var p postData
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			defer r.Body.Close()

			result = graphql.Do(graphql.Params{
				Context:        ctx,
				Schema:         schema,
				RequestString:  p.Query,
				VariableValues: p.Variables,
				OperationName:  p.Operation,
			})
		} else if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Printf("failed to read req: %v", err)
				return
			}
			result = graphql.Do(graphql.Params{
				Context:       ctx,
				Schema:        schema,
				RequestString: r.Form.Get("query"),
			})
		} else {
			result = graphql.Do(graphql.Params{
				Context:       ctx,
				Schema:        schema,
				RequestString: r.URL.Query().Get("query"),
			})
		}

		if len(result.Errors) > 0 {
			log.Printf("Query had errors: %s, %v", r.URL.Query().Get("query"), result.Errors)
		}
		json.NewEncoder(w).Encode(result)
	}))
	mux.Handle("/", http.FileServer(AssetFile()))

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
