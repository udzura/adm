package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/udzura/adm"
)

func main() {
	app := cli.NewApp()
	app.Name = "adm"
	app.Version = adm.Version
	app.Usage = "A Dotfile Manager"
	app.Author = "Uchio KONDO"
	app.Email = "udzura@udzura.jp"
	app.Commands = adm.Commands

	app.Run(os.Args)
}
