package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "adm"
	app.Version = Version
	app.Usage = ""
	app.Author = "Uchio KONDO"
	app.Email = "udzura@pepabo.com"
	app.Commands = Commands

	app.Run(os.Args)
}
