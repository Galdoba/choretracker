package appcontext

import (
	"path/filepath"

	"github.com/Galdoba/appcontext/xdg"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Version      string                    `toml:"version,omitempty"`
	StoragePath  string                    `toml:"storage_path,omitempty"`
	Log          LoggerConfiguration       `toml:"Logger"`
	Notification NotificationConfiguration `toml:"Notification"`
}

type LoggerConfiguration struct {
	Enabled       bool   `toml:"enabled"`
	Level         string `toml:"level"`
	FilePath      string `toml:"filepath,omitempty"`
	ConsoleOutput bool   `toml:"console output"`
	ConsoleColors bool   `toml:"console color"`
}

type NotificationConfiguration struct {
	Enabled  bool            `toml:"enabled"`
	NotifyAt []string        `toml:"notify_at"`
	Methods  map[string]bool `toml:"methods_enabled"`
}

func defaultConfig(paths *xdg.ProgramPaths) Config {
	return Config{
		Version:     constants.Version,
		StoragePath: filepath.Join(paths.PersistentDataDir(), constants.StorageFile),
		Log: LoggerConfiguration{
			Enabled:       true,
			Level:         "debug",
			FilePath:      "",
			ConsoleOutput: true,
			ConsoleColors: false,
		},
		Notification: NotificationConfiguration{
			Enabled:  false,
			NotifyAt: []string{"07:00", "21:30"},
			Methods: map[string]bool{
				"console":   true,
				"os_notify": true,
				"telegram":  false,
			},
		},
	}
}

func (cfg *Config) String() string {
	s := ""
	data, err := toml.Marshal(cfg)
	switch err == nil {
	case false:
		s += err.Error()
	case true:
		s = string(data)
	}
	return s
}
