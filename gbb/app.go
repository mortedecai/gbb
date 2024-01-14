package gbb

import (
	"github.com/mortedecai/gbb/config"
	"github.com/mortedecai/gbb/gbberror"
)

type App struct {
	Config *config.GoBurnBits
}

func New(cfg *config.GoBurnBits) *App {
	return &App{Config: cfg}
}

func (a *App) Run(args []string) error {
	return gbberror.ErrNotYetImplemented
}
