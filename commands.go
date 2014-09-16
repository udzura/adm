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
	Usage: "Initialize a dotfile manager",
	Description: `
	--repo specify repo name
`,
	Action: doInit,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "Add dotfiles to files managed",
	Description: `
`,
	Action: doAdd,
}

var commandSync = cli.Command{
	Name:  "sync",
	Usage: "Sync dotfiles current status to managed files",
	Description: `
`,
	Action: doSync,
}

var commandDiscard = cli.Command{
	Name:  "discard",
	Usage: "Discard dotfiles status and reset to be last saved one",
	Description: `
`,
	Action: doDiscard,
}

var commandPush = cli.Command{
	Name:  "push",
	Usage: "Push managed dotfiles to remote repositories",
	Description: `
`,
	Action: doPush,
}

var commandStatus = cli.Command{
	Name:  "status",
	Usage: "Show status of managed dotfiles",
	Description: `
`,
	Action: doStatus,
}

var commandInstall = cli.Command{
	Name:  "install",
	Usage: "Install dotfiles from remote config, load from `Dotfiles.source' by default",
	Description: `
`,
	Action: doInstall,
}

var commandSecret = cli.Command{
	Name:  "secret",
	Usage: "Register/unregister secret string values to file",
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
	log.Println("Going to implement first!")
}

func doAdd(c *cli.Context) {
	panic("Not yet implemented.")
}

func doSync(c *cli.Context) {
	panic("Not yet implemented.")
}

func doDiscard(c *cli.Context) {
	panic("Not yet implemented.")
}

func doPush(c *cli.Context) {
	panic("Not yet implemented.")
}

func doStatus(c *cli.Context) {
	panic("Not yet implemented.")
}

func doInstall(c *cli.Context) {
	panic("Not yet implemented.")
}

func doSecret(c *cli.Context) {
	panic("Not yet implemented.")
}
