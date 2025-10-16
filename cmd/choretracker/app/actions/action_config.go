package actions

import (
	"context"
	"fmt"
	"os"

	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/urfave/cli/v3"
)

func ConfigAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		logger := actx.GetLogger()
		logger.Infof("configuration data requested")
		cfg := actx.Config()
		cfgPath := actx.ConfigPath()
		if cfgPath == "" {
			return utils.LogError(logger, "config filepath not found", nil)
		}
		fmt.Fprintf(os.Stderr, "configuration file can be found at:\n%v\n", cfgPath)

		logpath := actx.LogfilePath()
		if logpath == "" {
			return utils.LogError(logger, "log filepath not found", nil)
		}
		fmt.Fprintf(os.Stderr, "logfile can be found at:\n%v\n\n", logpath)
		fmt.Fprintf(os.Stderr, "%v\n", cfg.String())
		return nil
	}
}
