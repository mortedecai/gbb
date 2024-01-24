/*
Copyright © 2024 Dan Taylor (github.com/mortedecai)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package commands

import (
	"fmt"
	"github.com/mortedecai/gbb/gbb"

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
	opt.destination, err = cmd.Flags().GetString("destination")
	if err != nil {
		return err
	}

	fmt.Printf("Downloading from http://%s:%d with token len %d to %s\n", opt.host, opt.port, len(opt.authToken), opt.destination)
	return gbb.HandleDownload(opt)
}
