//go:generate go run github.com/filecoin-project/statediff/build/gen "../../npm/app" static/app.js
//go:generate go run github.com/go-bindata/go-bindata/go-bindata -fs -prefix "static/" static/

package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "stateexplorer",
		Usage:       "State Explorer 🗺",
		Description: "State Explorer 🗺",
		Commands: []*cli.Command{
			exploreCmd,
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))
	for _, c := range app.Commands {
		sort.Sort(cli.FlagsByName(c.Flags))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
