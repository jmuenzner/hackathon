package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/autopilothq/lg"
)

var (
	log      lg.Log
	logLevel string
)

// SetupLog is called once, after config has loaded to setup a logger
func SetupLog() error {
	conf := Reader()

	logFormat := conf.GetString("Banks.Log.Format")
	logLevel = conf.GetString("Banks.Log.Level")

	if logLevel == "" {
		return errors.New("Missing config key Banks.Log.Level")
	}

	lvl, err := lg.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	switch logFormat {
	case "json":
		lg.SetOutput(os.Stdout, lg.JSON(), lg.MinLevel(lvl))

	case "text":
		lg.SetOutput(os.Stdout, lg.PlainText(), lg.MinLevel(lvl))

	case "":
		return errors.New("Missing config key Banks.Log.Format")

	default:
		return fmt.Errorf(`
					config key Banks.Log.Format value '%s' is invalid.
					Valid values are 'text' and 'json'`, logFormat)
	}

	log = lg.Extend()

	return nil
}

// Log returns a banks logger
func Log() lg.Log {
	return log
}

// LogLevel returns the banks logger loglevel
func LogLevel() string {
	return logLevel
}
