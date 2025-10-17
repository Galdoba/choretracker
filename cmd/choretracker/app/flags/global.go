package flags

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

const (
	GLOBAL_MODE       = "run-mode"
	SERVER_PORT       = "port"
	VALUE_MODE_CLI    = "cli"
	VALUE_MODE_TUI    = "tui"
	VALUE_MODE_SERVER = "server"
	CHORE_TITLE       = "title"
	CHORE_DESCRIPTION = "description"
	CHORE_AUTHOR      = "author"
	CHORE_SCHEDULE    = "schedule"
	CHORE_COMMENT     = "comment"
	CHORE_ID          = "id"
)

var MODE = cli.StringFlag{
	Name:        GLOBAL_MODE,
	Usage:       "'tui', 'cli' or 'server'",
	DefaultText: "tui",
	Value:       "tui",
}

var PORT = cli.IntFlag{
	Name:        SERVER_PORT,
	DefaultText: "42000",
	Usage:       "port to start server on",
	Value:       42000,
	Aliases:     []string{"p"},
	Validator:   validatePort,
}

func validatePort(port int) error {
	if port < 42000 || port > 65000 {
		return fmt.Errorf("port must be between 42000 and 65000")
	}
	return nil
}

var TITLE = cli.StringFlag{
	Name:    CHORE_TITLE,
	Usage:   "use chore title variable",
	Aliases: []string{"t"},
}

var DESCRIPTION = cli.StringFlag{
	Name:    CHORE_DESCRIPTION,
	Usage:   "use chore description variable",
	Aliases: []string{"d"},
}

var AUTHOR = cli.StringFlag{
	Name:    CHORE_AUTHOR,
	Usage:   "use chore author variable",
	Aliases: []string{"a"},
}

var SHEDULE = cli.StringFlag{
	Name:    CHORE_SCHEDULE,
	Usage:   "use chore schedule variable",
	Aliases: []string{"s"},
}

var COMMENT = cli.StringFlag{
	Name:    CHORE_COMMENT,
	Usage:   "use chore comment variable",
	Aliases: []string{"c"},
}

var ID = cli.Int64Flag{
	Name:        CHORE_ID,
	Usage:       "use chore id variable (required for cli-mode)",
	Aliases:     []string{"i"},
	HideDefault: true,
}
