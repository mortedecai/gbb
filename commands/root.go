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

func (ro *rootOption) Server() string {
	return fmt.Sprintf("http://%s:%d", ro.host, ro.port)
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
