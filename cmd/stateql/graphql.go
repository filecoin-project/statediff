package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	gqlib "github.com/filecoin-project/statediff/cmd/stateql/lib"
	"github.com/filecoin-project/statediff/lib"
	"github.com/filecoin-project/statediff/lib/debug"

	"github.com/gorilla/handlers"

	"github.com/urfave/cli/v2"
)

var bindFlag = cli.StringFlag{
	Name:  "bind",
	Usage: "Bind to a specific local host:port. Specified as [address]:port",
	Value: ":0",
}

var crtFlag = cli.StringFlag{
	Name:  "tls",
	Usage: "root name (files of <name>.crt and <name>.key expected) for tls cert",
	Value: "",
}

var pprofFlag = cli.StringFlag{
	Name:  "pprof",
	Usage: "include a pprof handler at a given mount point",
}

var graphCmd = &cli.Command{
	Name:        "serve",
	Description: "Serve GraphQL endpoint",
	Action:      runGraphCmd,
	Flags: []cli.Flag{
		&lib.ApiFlag,
		&lib.CarFlag,
		&lib.SqlFlag,
		&lib.NewSqlFlag,
		&lib.VectorFlag,
		&lib.NoCacheFlag,
		&bindFlag,
		&crtFlag,
		&pprofFlag,
	},
}

func runGraphCmd(c *cli.Context) error {
	ds, err := lib.Lazy(c)
	if err != nil {
		return err
	}
	mux := gqlib.GetGraphQL(c, ds)
	if c.IsSet(pprofFlag.Name) {
		mux.Handle(fmt.Sprintf("/%s/", c.String(pprofFlag.Name)), debug.WithProfile())
	}

	lis, err := net.Listen("tcp", c.String(bindFlag.Name))
	if err != nil {
		return err
	}
	s := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, mux),
	}

	fmt.Printf("Listening at %v\n", lis.Addr())
	if c.IsSet(crtFlag.Name) {
		b := c.String(crtFlag.Name)
		return s.ServeTLS(lis, b+".crt", b+".key")
	} else {
		return s.Serve(lis)
	}
}
