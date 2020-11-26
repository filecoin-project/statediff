//go:generate go run github.com/go-bindata/go-bindata/go-bindata -pkg lib -fs -prefix "static/" static/

package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/filecoin-project/statediff"
	sdlib "github.com/filecoin-project/statediff/lib"
	"github.com/graphql-go/graphql"
	"github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/urfave/cli/v2"
)

const storeCtx = "store"

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

func GetGraphQL(c *cli.Context, source sdlib.Datasource) *http.ServeMux {
	AddFields()

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
				"Height": &graphql.Field{
					Type: LotusBlockHeader__type,
					Args: graphql.FieldConfigArgument{
						"at": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						s, ok := p.Context.Value(storeCtx).(blockstore.Blockstore)
						if !ok {
							return nil, fmt.Errorf("no datastore provided")
						}
						at, ok := p.Args["at"].(int)
						if !ok {
							return nil, fmt.Errorf("invalid height %d", at)
						}
						if at == -1 {
							ch := source.Head(p.Context)
							return statediff.Transform(p.Context, ch, s, string(statediff.LotusTypeTipset))
						}
						ch, err := source.CidAtHeight(p.Context, at)
						if err != nil {
							return nil, fmt.Errorf("Have not indexed a block at height %d", at)
						}
						if ch == cid.Undef {
							return nil, nil
						}
						return statediff.Transform(p.Context, ch, s, string(statediff.LotusTypeTipset))
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
							ch, err := source.CidAtHeight(p.Context, i)
							if err != nil {
								return nil, fmt.Errorf("Have not indexed a block at height %d", i)
							}
							if ch == cid.Undef {
								continue
							}
							n, err := statediff.Transform(p.Context, ch, s, string(statediff.LotusTypeTipset))
							if err != nil {
								return nil, err
							}
							out = append(out, n)
						}
						return out, nil
					},
				},
				"Samples": &graphql.Field{
					Type: graphql.NewList(LotusBlockHeader__type),
					Args: graphql.FieldConfigArgument{
						"from": &graphql.ArgumentConfig{
							Type:        graphql.Int,
							Description: "Starting unix time stamp in seconds",
						},
						"to": &graphql.ArgumentConfig{
							Type:        graphql.Int,
							Description: "Ending unix time stamp in seconds",
						},
						"interval": &graphql.ArgumentConfig{
							Type:        graphql.Int,
							Description: "Interval in seconds. Only really makes sense in multiples of 30",
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// figure out which blocks to get.
						s, ok := p.Context.Value(storeCtx).(blockstore.Blockstore)
						if !ok {
							return nil, fmt.Errorf("no datastore provided")
						}
						from := p.Args["from"].(int)
						to := p.Args["to"].(int)
						interval := p.Args["interval"].(int)

						out := make([]ipld.Node, 0)
						// get them.
						for i := from; i < to; i += interval {
							ch, err := source.CidAtHeight(p.Context, timeStampToEpoch(time.Unix(int64(i), 0)))
							if err != nil {
								return nil, fmt.Errorf("Have not indexed a block at height %d", i)
							}
							if ch == cid.Undef {
								continue
							}
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

	loader := func(ctx context.Context, cl cidlink.Link, builder ipld.NodeBuilder) (ipld.Node, error) {
		ca := statediff.GetGlobalNodeCache()
		key := string(cl.Bytes()) + fmt.Sprintf("%T", builder.Prototype())
		if n := ca.Get(key); n != nil {
			return n, nil
		}
		if err := statediff.Load(ctx, cl.Cid, source.Store(), builder); err != nil {
			return nil, err
		}
		n := builder.Build()
		if n != nil {
			ca.Add(key, n)
		}
		return n, nil
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", CorsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		s := source.Store()
		if s == nil {
			log.Printf("failed to connect to backend")
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sctx := context.WithValue(r.Context(), storeCtx, s)
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
	mux.HandleFunc("/graphql.html", func(w http.ResponseWriter, r *http.Request) {
		fs := AssetFile()
		f, err := fs.Open("index.html")
		if err != nil {
			return
		}
		http.ServeContent(w, r, "graphql.html", time.Unix(0, 0), f)
	})
	return mux
}

func timeStampToEpoch(t time.Time) int {
	return t.Second()/30 - (1598306400 / 30)
}

func epochToTime(e int) time.Time {
	return time.Unix(int64(e*30)+1598306400, 0)
}
