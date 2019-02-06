package main

import (
	"fmt"
	"os"

	"github.com/bitconch/bus"
	"github.com/bitconch/bus/gobin/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	// LocalMode can TRUE of FALSE, when in local mode, process will fetch configuration from local machine
	LocalMode string
	// KeypairFile is the client keypair file, which stores
	KeypairFile string
	// PublicMode: TRUE or FALSE,  is the location of the ledger file
	PublicMode string
	// BindPortNum the port number needed to be binded with
	BindPortNum string

	gitCommit = ""
	// add new App with description
	app = utils.NewApp(gitCommit, "Bitconch chain fullnode config CLI")
)

// Flags to be used in the cli
var (
	localModeFlag = cli.StringFlag{
		Name:        "local,l",
		Usage:       "Local mode or not, fetch configuration from local machine(127.0.0.1)",
		Destination: &LocalMode,
	}

	keypairFlag = cli.StringFlag{
		Name:        "keypair,k",
		Usage:       "Client keypair file",
		Destination: &KeypairFile,
	}

	publicModeFlag = cli.StringFlag{
		Name:        "public,p",
		Usage:       "Ledger file location",
		Destination: &PublicMode,
	}

	bindFlag = cli.StringFlag{
		Name:        "bind,b",
		Usage:       "Bind to port number (local mode) or an address (public mode)",
		Destination: &BindPortNum,
	}
)

//init define subcommand and flags linked to cli
func init() {
	// clapcli is the action function
	app.Action = fullnodeConfigCli

	// define the sub commands
	app.Commands = []cli.Command{
		//commandGenerate,
	}

	// define the flags
	app.Flags = []cli.Flag{
		localModeFlag,
		keypairFlag,
		publicModeFlag,
		bindFlag,
	}
}

func main() {

	//bus.CallFullnode()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func fullnodeConfigCli(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}

	// handle the arguments
	fmt.Println("Do some stuff")
	// start the full node instance
	bus.CallFullnodeConfig(
		LocalMode,
		KeypairFile,
		PublicMode,
		BindPortNum,
	)
	return nil
}
