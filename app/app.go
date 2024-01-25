package app

import (
	commands2 "github.com/mortedecai/gbb/commands"
	"github.com/spf13/cobra"
)

type App struct {
	cmd *cobra.Command
}

func New(version string) *App {
	app := &App{}
	app.SetupCommands(version)
	return app
}

func (a *App) Run() error {
	return a.cmd.Execute()
}

func (a *App) SetupCommands(version string) {
	var err error
	if a.cmd, err = commands2.Root(version); err != nil {
		panic(err)
	}
	if _, err = commands2.Download(a.cmd); err != nil {
		panic(err)
	}
}
