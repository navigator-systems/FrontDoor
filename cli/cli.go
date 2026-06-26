package cli

import (
	"context"
	"frontdoor/server"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func InitCli() {

	if err := serverCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

var serverCmd = &cli.Command{
	Name:    "server",
	Aliases: []string{"s", "serve"},
	Usage:   "Start the FrontDoor service",
	Action: func(ctx context.Context, serverCmd *cli.Command) error {
		server.Server()
		return nil
	},
}
