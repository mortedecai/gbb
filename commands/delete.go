package commands

import (
	"fmt"
	"strings"

	"github.com/mortedecai/gbb/client"
	"github.com/mortedecai/gbb/models"
	"github.com/spf13/cobra"
)

func Delete(rootCmd *cobra.Command) (*cobra.Command, error) {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Allows the user to delete files from BitBurner",
		Long:  `The delete command allows users to remove files from the BitBurner game via the Remote API Server.`,
		RunE:  handleDelete,
	}
	deleteCmd.Flags().StringP("file", "f", "", "The file of the file to delete from a BitBurner server")
	rootCmd.AddCommand(deleteCmd)
	return deleteCmd, deleteCmd.MarkFlagRequired("file")
}

type deleteOption struct {
	*rootOption
	toDelete string
}

func (opt *deleteOption) ToDelete() models.GBBFileName {
	return models.GBBFileName(strings.TrimSpace(opt.toDelete))
}

func (opt *deleteOption) Valid() bool {
	return len(strings.TrimSpace(opt.toDelete)) > 0 && opt.rootOption.Valid()
}

func handleDelete(cmd *cobra.Command, args []string) error {
	var err error
	opt := &deleteOption{rootOption: &rootOption{}}
	if opt.host, opt.port, opt.authToken, err = handleCommonFlags(cmd); err != nil {
		return err
	}
	if opt.toDelete, err = flagReader.GetString(cmd, "file"); err != nil {
		return err
	}

	fmt.Printf("\nDeleting %s from http://%s:%d with token len %d.\n", opt.toDelete, opt.host, opt.port, len(opt.authToken))
	return client.HandleDelete(opt)
}
