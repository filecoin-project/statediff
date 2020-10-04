package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "statediff",
		Usage:       "State Diff 🕵️‍♂️",
		Description: "State Diff 🕵️‍♂️",
		Commands: []*cli.Command{
			vectorCmd,
			carCmd,
			chainCmd,
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
