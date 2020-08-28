//go:generate go run github.com/go-bindata/go-bindata/go-bindata -fs -prefix "static/" static/

package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var expandActorsFlag = cli.StringFlag{
	Name:  "expand-actors",
	Usage: "Control which actors state is expanded. pass with no argument or 'all' to expand all states, provide a class, like 'InitActor', or specific state root cid(s) separate by commas",
	Value: "",
}

func main() {
	app := &cli.App{
		Name:        "statediff",
		Usage:       "State Inspector 🕵️‍♂️",
		Description: "State Inspector 🕵️‍♂️",
		Commands: []*cli.Command{
			vectorCmd,
			chainCmd,
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
