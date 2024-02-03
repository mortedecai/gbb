package commands

import (
	"fmt"
	"github.com/mortedecai/gbb/client"
	"github.com/spf13/cobra"
	"strings"
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

type uploadOption struct {
	*rootOption
	toUpload string
}

func (opt *uploadOption) ToUpload() []string {
	var uploads []string
	if strings.TrimSpace(opt.toUpload) != "" {
		uploads = []string{opt.toUpload}
	}
	return uploads
}

func (opt *uploadOption) Valid() bool {
	return len(opt.ToUpload()) > 0 && opt.rootOption.Valid()
}

func handleUpload(cmd *cobra.Command, args []string) error {
	var err error
	opt := &uploadOption{rootOption: &rootOption{}}
	if opt.host, opt.port, opt.authToken, err = handleCommonFlags(cmd); err != nil {
		return err
	}
	if opt.toUpload, err = flagReader.GetString(cmd, "file"); err != nil {
		return err
	}

	fmt.Printf("\nUploading %s to http://%s:%d with token len %d.\n", opt.toUpload, opt.host, opt.port, len(opt.authToken))
	return client.HandleUpload(opt)
}
