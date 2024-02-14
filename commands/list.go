package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mortedecai/gbb/client"
)

type listOption struct {
	*rootOption
}

func (lo *listOption) Valid() bool {
	return lo.rootOption.Valid()
}

func List(rootCmd *cobra.Command) (*cobra.Command, error) {

	// downloadCmd represents the download command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Allows the user to list the files in BitBurner",
		Long:  `The list command allows users to list files and locations from inside the Bit Burner game via the Remote API Server.`,
		RunE:  handleList,
	}
	rootCmd.AddCommand(listCmd)
	return listCmd, nil
}

func handleList(cmd *cobra.Command, args []string) error {
	opt := &listOption{rootOption: &rootOption{}}
	var err error
	opt.host, opt.port, opt.authToken, err = handleCommonFlags(cmd)
	if err != nil {
		return err
	}

	fmt.Printf("Listing files from http://%s:%d with token len %d\n", opt.host, opt.port, len(opt.authToken))
	return client.HandleList(opt)
}
