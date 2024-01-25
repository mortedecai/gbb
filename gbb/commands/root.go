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
	"github.com/mortedecai/gbb/gbberror"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
)

type rootOption struct {
	host      string
	port      int
	authToken string
}

func (ro *rootOption) Host() string {
	return ro.host
}

func (ro *rootOption) Port() int {
	return ro.port
}

func (ro *rootOption) AuthToken() string {
	return ro.authToken
}

func (ro *rootOption) Valid() bool {
	return ro.host != "" && ro.port != 0 && ro.authToken != ""
}

func (ro *rootOption) AddAuth(req *http.Request) *http.Request {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ro.authToken))
	return req
}

func Root(version string) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "gbb",
		Short: "gbb provides file transfer capabilties for BitBurner",
		Long: `gbb is a suite of tools to help in the development of BitBurner scripts locally and
transfer them into the BitBurner game via the API Server.

Usage examples to come`,
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       fmt.Sprintf("Go Burn Bits [%s]", version),
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(commands *cobra.Command, args []string) { },
	}
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "host to connect to (default is `localhost`)")
	rootCmd.PersistentFlags().IntP("port", "p", 9990, "port to connect to host on (default is 9990)")
	rootCmd.PersistentFlags().StringP("authToken", "a", "", "BitBurner Server API Key")
	return rootCmd, rootCmd.MarkPersistentFlagRequired("authToken")
}

func handleCommonFlags(cmd *cobra.Command) (host string, port int, authToken string, err error) {
	//if host, err = cmd.Flags().GetString("host"); err != nil {
	if host, err = flagReader.GetString(cmd, "host"); err != nil {
		return
	}
	if port, err = flagReader.GetInt(cmd, "port"); err != nil {
		return
	}
	if authToken, err = flagReader.GetString(cmd, "authToken"); err != nil {
		return
	}
	if host = strings.TrimSpace(host); host == "" {
		err = gbberror.ErrBadHost
		return
	}
	if port < 0 || port > 65535 {
		err = gbberror.ErrBadPort
		return
	}
	if authToken = strings.TrimSpace(authToken); authToken == "" {
		err = gbberror.ErrNoAuthToken
	}
	return
}
