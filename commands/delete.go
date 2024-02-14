package commands

import (
	"github.com/mortedecai/gbb/gbberror"
	"github.com/mortedecai/gbb/models"
	"github.com/spf13/cobra"
	"strings"
)

func Delete(rootCmd *cobra.Command) (*cobra.Command, error) {
	return nil, gbberror.ErrNotYetImplemented
}

type deleteOption struct {
	*rootOption
	toDelete string
}

func (opt *deleteOption) ToDelete() []models.GBBFileName {
	if strings.TrimSpace(opt.toDelete) != "" {
		return []models.GBBFileName{models.GBBFileName(opt.toDelete)}
	}
	return []models.GBBFileName{}
}

func (opt *deleteOption) Valid() bool {
	return len(strings.TrimSpace(opt.toDelete)) > 0 && opt.rootOption.Valid()
}
