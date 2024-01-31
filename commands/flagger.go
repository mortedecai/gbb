package commands

import "github.com/spf13/cobra"

var (
	flagReader Flagger
)

func init() {
	flagReader = &flagger{}
}

type Flagger interface {
	GetString(*cobra.Command, string) (string, error)
	GetInt(*cobra.Command, string) (int, error)
}

//go:generate mockgen -destination=./mocks/mock_flagger.go -package=mocks github.com/mortedecai/gbb/commands Flagger

type flagger struct{}

func (f *flagger) GetString(cmd *cobra.Command, flag string) (string, error) {
	return cmd.Flags().GetString(flag)
}

func (f *flagger) GetInt(cmd *cobra.Command, flag string) (int, error) {
	return cmd.Flags().GetInt(flag)
}
