package app

import (
	"github.com/spf13/cobra"

	"github.com/mortedecai/gbb/gbb/commands"
)

type App struct {
	cmd *cobra.Command
}

func New(version string) (*App, error) {
	app := &App{}
	app.SetupCommands(version)
	return app, nil
}

func (a *App) Run() error {
	return a.cmd.Execute()
}

func (a *App) SetupCommands(version string) {
	var err error
	if a.cmd, err = commands.Root(version); err != nil {
		panic(err)
	}
	if _, err = commands.Download(a.cmd); err != nil {
		panic(err)
	}
}
