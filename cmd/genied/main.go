package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/OnenessLabs/genie/internal/debug"
	"github.com/OnenessLabs/genie/internal/flags"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/urfave/cli/v2"
	"go.uber.org/automaxprocs/maxprocs"
)

const ENV_VARS_PREFIX = "GENIE"

var app = flags.NewApp("the Oneness network node daemon")

func init() {
	// Initialize the CLI app and start
	app.Action = func(ctx *cli.Context) error {
		if args := ctx.Args().Slice(); len(args) > 0 {
			return fmt.Errorf("invalid command: %q", args[0])
		}

		// start services here, e.g.
		// someService := SomeService.New()
		// defer someService.Shutdown()
		// someService.Start()
		// someService.Wait()

		return nil
	}

	// uncomment this to add more commands, see chaincmd.go in go-ethereum
	// app.Commands = []*cli.Command{}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = flags.Merge(
		debug.Flags,
	)
	flags.AutoEnvVars(app.Flags, ENV_VARS_PREFIX)

	app.Before = func(ctx *cli.Context) error {
		maxprocs.Set() // Automatically set GOMAXPROCS to match Linux container CPU quota.
		flags.MigrateGlobalFlags(ctx)
		if err := debug.Setup(ctx); err != nil {
			return err
		}
		flags.CheckEnvVars(ctx, app.Flags, ENV_VARS_PREFIX)
		return nil
	}
	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		prompt.Stdin.Close() // Resets terminal mode.
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
