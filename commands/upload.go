package commands

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mortedecai/gbb/client"
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
	return validateFiles(opt.ToUpload()) && opt.rootOption.Valid()
}

func validateFiles(fns []string) bool {
	if len(fns) <= 0 {
		return false
	}
	for _, v := range fns {
		if !validateFile(v) {
			return false
		}
	}
	return true
}

func validateFile(fn string) bool {
	p := path.Clean(fn)
	if fi, err := os.Stat(p); err != nil || fi.IsDir() {
		return false
	}
	return true
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
