package adm

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandInit,
	commandAdd,
	commandSync,
	commandDiscard,
	commandPush,
	commandStatus,
	commandInstall,
	commandSecret,
}

var commandInit = cli.Command{
	Name:  "init",
	Usage: "",
	Description: `
`,
	Action: doInit,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "",
	Description: `
`,
	Action: doAdd,
}

var commandSync = cli.Command{
	Name:  "sync",
	Usage: "",
	Description: `
`,
	Action: doSync,
}

var commandDiscard = cli.Command{
	Name:  "discard",
	Usage: "",
	Description: `
`,
	Action: doDiscard,
}

var commandPush = cli.Command{
	Name:  "push",
	Usage: "",
	Description: `
`,
	Action: doPush,
}

var commandStatus = cli.Command{
	Name:  "status",
	Usage: "",
	Description: `
`,
	Action: doStatus,
}

var commandInstall = cli.Command{
	Name:  "install",
	Usage: "",
	Description: `
`,
	Action: doInstall,
}

var commandSecret = cli.Command{
	Name:  "secret",
	Usage: "",
	Description: `
`,
	Action: doSecret,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doInit(c *cli.Context) {
}

func doAdd(c *cli.Context) {
}

func doSync(c *cli.Context) {
}

func doDiscard(c *cli.Context) {
}

func doPush(c *cli.Context) {
}

func doStatus(c *cli.Context) {
}

func doInstall(c *cli.Context) {
}

func doSecret(c *cli.Context) {
}
