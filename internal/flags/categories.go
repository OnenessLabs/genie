package flags

import "github.com/urfave/cli/v2"

const (
	LoggingCategory = "LOGGING AND DEBUGGING"
	MiscCategory    = "OTHERS"
)

func init() {
	cli.HelpFlag.(*cli.BoolFlag).Category = MiscCategory
	cli.VersionFlag.(*cli.BoolFlag).Category = MiscCategory
}
