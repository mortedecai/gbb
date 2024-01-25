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

const (
	cmdResMsg  = "GBB Result"
	resMsg     = "Result"
	errResult  = "Error"
	detailsMsg = "Details"
)

func init() {
	logger = getLogger()
}

func getLogger() *zap.SugaredLogger {
	if l, err := zap.NewProduction(zap.IncreaseLevel(zapcore.ErrorLevel)); err == nil {
		logger = l.Sugar().Named("client")
	}
	return logger
}

func main() {
	var a *app.App
	var err error

	a = app.New(version)
	if err = a.Run(); err != nil {
		logger.Errorw(cmdResMsg, resMsg, errResult, detailsMsg, err)
	}
	logger.Debugw(cmdResMsg, resMsg, "Success")
}
