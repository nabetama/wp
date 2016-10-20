package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	search,
}

var search = cli.Command{
	Name:  "search",
	Usage: "search word(s)",
	Description: `
	Search word.
`,
	Action: doSearch,
	Flags:  searchFlags,
}

var searchFlags = []cli.Flag{
	cli.BoolFlag{Name: "ignorecase, i", Usage: "ignore match of the word(s)"},
}

func doSearch(c *cli.Context) error {
	isIgnore := c.Args().Get(0)

	return nil
}
