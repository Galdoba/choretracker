package appcontext

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Galdoba/appcontext/configmanager"
	"github.com/Galdoba/appcontext/logmanager"
	"github.com/Galdoba/appcontext/xdg"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/Galdoba/choretracker/internal/core/services"
	"github.com/Galdoba/choretracker/internal/infrastructure"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage"
)

const (
	config  = "config"
	logfile = "logfile"
)

type AppContext struct {
	config  *Config
	log     Logger
	paths   map[string]string
	service *services.TaskService
}

func (actx *AppContext) Config() *Config {
	return actx.config
}

func (actx *AppContext) GetLogger() Logger {
	return actx.log
}

func (actx *AppContext) ConfigPath() string {
	return actx.paths[config]
}

func (actx *AppContext) LogfilePath() string {
	return actx.paths[logfile]
}

func (actx *AppContext) GetService() *services.TaskService {
	return actx.service
}

func InitCli(appname string) (*AppContext, error) {
	actx := AppContext{}
	actx.paths = make(map[string]string)

	paths := xdg.New(constants.AppName)
	actx.paths[config] = paths.ConfigFile()
	actx.paths[logfile] = paths.LogFile()

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
					logmanager.WithLevelTag(true),
					logmanager.WithColor(true),
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
				logmanager.WithLevelTag(true),
				logmanager.WithColor(false),
			),
		),
	)
	actx.log = logmanager.New(logmanager.WithHandlers(logHandlers...))
	if err := confirmFile(actx.ConfigPath()); err != nil {
		return nil, err
	}
	if err := confirmFile(actx.LogfilePath()); err != nil {
		return nil, err
	}

	validator := infrastructure.DefaultValidator()
	store, err := storage.NewStorage(storage.JsonStorage, actx.config.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("failed to setup storage: %v", err)
	}
	actx.service = services.NewTaskService(store, validator, actx.log)

	return &actx, nil
}

func confirmFile(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer f.Close()
	return nil
}
