package main

import (
	"github.com/mortedecai/gbb/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	version string = "<unknown>"
	logger  *zap.SugaredLogger
)

func init() {
	logger = getLogger()
}

func getLogger() *zap.SugaredLogger {
	if l, err := zap.NewProduction(zap.IncreaseLevel(zapcore.ErrorLevel)); err == nil {
		logger = l.Sugar().Named("gbb")
	}
	return logger
}

func main() {
	var a *app.App
	var err error

	if a, err = app.New(version); err != nil {
		logger.Errorw("Command Result", "Result", "Error", "Details", err)
		return
	}
	if err = a.Run(); err != nil {
		logger.Errorw("Command Result", "Result", "Error", "Error", err)
	}
	logger.Debugw("Command Result", "Result", "Success")
}
