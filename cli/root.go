package cli

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func Exec() {
	if err := chaosCmd.Run(context.Background(), os.Args); err != nil {
		log.Println(err)
	}
}

var chaosCmd = &cli.Command{
	Name:  "go-chaos",
	Usage: "A terminal based app that injects chaos into your cloud infrastructure",
	Commands: []*cli.Command{
		validateCmd,
		destroyCmd,
		targetCmd,
	},
}
