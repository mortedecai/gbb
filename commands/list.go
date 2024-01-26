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
