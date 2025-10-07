package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/urfave/cli/v3"
)

func Start(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		fmt.Println("start Start action")
		return nil
	}
}
