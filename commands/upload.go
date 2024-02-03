package commands

import (
	"github.com/mortedecai/gbb/gbberror"
	"github.com/spf13/cobra"
)

func Upload(rootCmd *cobra.Command) (*cobra.Command, error) {
	var uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Allows the user to upload files into BitBurner",
		Long:  `The upload command allows users to push files into the BitBurner game via the Remote API Server.`,
		RunE:  handleUpload,
	}
	uploadCmd.Flags().StringP("file", "f", "", "The file to upload into BitBurner")
	rootCmd.AddCommand(uploadCmd)
	return uploadCmd, uploadCmd.MarkFlagRequired("file")
}

func handleUpload(cmd *cobra.Command, args []string) error {
	return gbberror.ErrNotYetImplemented
}
