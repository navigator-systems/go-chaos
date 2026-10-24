package cli

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	cmds "github.com/navigator-systems/go-chaos/cmd"
)

var validateCmd = &cli.Command{
	Name:  "validate",
	Usage: "validate file.hcl",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		file := cmd.Args().Get(0)
		err := cmds.ValidateCmd(file)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	},
}

var destroyCmd = &cli.Command{
	Name:  "destroy",
	Usage: "destroy file.hcl",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		file := cmd.Args().Get(0)
		fmt.Println("destroy", file)
		return nil
	},
}

var targetCmd = &cli.Command{
	Name:  "target",
	Usage: "target file.hcl aws.ec2.app:prod",
	Action: func(ctx context.Context, cmd *cli.Command) error {
		file := cmd.Args().Get(0)
		target := cmd.Args().Get(1)
		_, err := cmds.TargetCmd(file, target)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	},
}
