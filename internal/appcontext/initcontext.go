package appcontext

import (
	"fmt"
	"os"

	"github.com/Galdoba/appcontext/configmanager"
	"github.com/Galdoba/appcontext/logmanager"
	"github.com/Galdoba/appcontext/xdg"
	"github.com/Galdoba/choretracker/internal/constants"
)

type AppContext struct {
	config *Config
	log    Logger
}

func Init(appname string) (*AppContext, error) {
	actx := AppContext{}

	paths := xdg.New(constants.AppName)

	cfgman, err := configmanager.New(appname, defaultConfig(paths))
	if err != nil {
		return nil, fmt.Errorf("config init failed: %v", err)
	}
	if err := cfgman.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "config loading failed: %v\nfallback to default configuration...\n", err)
	}
	actx.config = cfgman.Config()
	logHandlers := []*logmanager.MessageHandler{}
	if actx.config.Log.ConsoleOutput {
		logHandlers = append(logHandlers,
			logmanager.NewHandler(
				logmanager.Stderr,
				logmanager.StringToLevel(actx.config.Log.Level),
				logmanager.NewTextFormatter(
					logmanager.WithTimePrecision(0),
					logmanager.WithLevelTag(false),
					logmanager.WithColor(actx.config.Log.ConsoleColors),
				),
			),
		)
	}
	logHandlers = append(logHandlers,
		logmanager.NewHandler(
			paths.LogFile(),
			logmanager.LevelDebug,
			logmanager.NewTextFormatter(
				logmanager.WithTimePrecision(0),
				logmanager.WithLevelTag(false),
				logmanager.WithColor(false),
			),
		),
	)
	actx.log = logmanager.New(logmanager.WithHandlers(logHandlers...))

	return &actx, nil
}
