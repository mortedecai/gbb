package main

import (
	"github.com/mortedecai/gbb/config"
	"github.com/mortedecai/gbb/gbb"
	"go.uber.org/zap"
	"os"
)

var (
	version string = "<unknown>"
	logger  *zap.SugaredLogger
)

func greetings() string {
	return "Go Burn Bits"
}

func Version() string {
	return version
}

func getLogger() *zap.SugaredLogger {
	if l, err := zap.NewDevelopment(); err == nil {
		logger = l.Sugar().Named("gbb")
	}
	return logger
}

func loadConfig() *config.GoBurnBits {
	var cfg *config.GoBurnBits
	var err error

	if cfg, err = config.FromConfig(); err != nil {
		logger.Errorw("Load Configuration", "Error", err)
		cfg = config.Default()
	}
	return cfg
}

func main() {
	logger = getLogger()
	logger.Infow("Go Burn Bits", "Version", Version())
	cfg := loadConfig()

	err := gbb.New(cfg).Run(os.Args[1:])
	if err != nil {
		logger.Errorw("Command Result", "Result", "Error", "Details", err)
		return
	}
	logger.Infow("Command Result", "Result", "Success")
}
