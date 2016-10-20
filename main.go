package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var Version string = "0.0.1"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "wp"
	app.Usage = "Looking forward the word which isn't on a dictionary."
	app.Version = Version
	app.Author = "nabetama"
	app.Author = "mao.nabeta@gmail.com"
	app.Commands = Commands
	return app
}
