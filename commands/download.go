package commands

import (
	"fmt"
	"github.com/mortedecai/gbb/client"

	"github.com/spf13/cobra"
)

type downloadOption struct {
	*rootOption
	destination string
}

func (do *downloadOption) Destination() string {
	return do.destination
}

func (do *downloadOption) Valid() bool {
	return do.destination != "" && do.rootOption.Valid()
}

func Download(rootCmd *cobra.Command) (*cobra.Command, error) {

	// downloadCmd represents the download command
	var downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Allows the user to download files from BitBurner",
		Long:  `The download command allows users to retrieve files from inside the Bit Burner game via the Remote API Server.`,
		//	Args:  cobra.MinimumNArgs(1),
		RunE: handleDownload,
	}
	downloadCmd.Flags().StringP("destination", "d", "./", "The base directory to download files into (Default: './').")
	rootCmd.AddCommand(downloadCmd)
	if err := downloadCmd.MarkFlagDirname("destination"); err != nil {
		return nil, err
	}
	return downloadCmd, downloadCmd.MarkFlagRequired("destination")
}

func handleDownload(cmd *cobra.Command, args []string) error {
	opt := &downloadOption{rootOption: &rootOption{}}
	var err error
	opt.host, opt.port, opt.authToken, err = handleCommonFlags(cmd)
	if err != nil {
		return err
	}
	opt.destination, err = flagReader.GetString(cmd, "destination")
	if err != nil {
		return err
	}

	fmt.Printf("Downloading from http://%s:%d with token len %d to %s\n", opt.host, opt.port, len(opt.authToken), opt.destination)
	return client.HandleDownload(opt)
}
