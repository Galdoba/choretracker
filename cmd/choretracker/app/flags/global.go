package flags

import (
	"github.com/urfave/cli/v3"
)

const (
	GLOBAL_MODE       = "run-mode"
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
