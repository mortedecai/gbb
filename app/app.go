package app

import (
	"github.com/mortedecai/gbb/commands"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type App struct {
	cmd    *cobra.Command
	logger *zap.SugaredLogger
}

func New(version string, l *zap.SugaredLogger) *App {
	app := &App{logger: l}
	app.SetupCommands(version)
	return app
}

func (a *App) Run() error {
	return a.cmd.Execute()
}

func (a *App) SetupCommands(version string) {
	a.logger.Debugw("SetupCommands", "Status", "Starting", "Version", version)
	var err error
	if a.cmd, err = commands.Root(version); err != nil {
		a.logger.Debugw("SetupCommands", "Status", "Error", "Command", "root", "Error", err)
		panic(err)
	}
	if _, err = commands.Download(a.cmd); err != nil {
		a.logger.Debugw("SetupCommands", "Status", "Error", "Command", "download", "Error", err)
		panic(err)
	}
	if _, err = commands.List(a.cmd); err != nil {
		a.logger.Debugw("SetupCommands", "Status", "Error", "Command", "list", "Error", err)
		panic(err)
	}
	a.logger.Debugw("SetupCommands", "Status", "Completed")
}
