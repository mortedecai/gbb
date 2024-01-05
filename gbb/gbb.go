package gbb

import (
	"strings"

	"github.com/mortedecai/go-burn-bits/gbberror"
)

type GoBurnBits interface {
	Run([]string) error
	HandleUpload([]string) error
}

type gbb struct {
	Host      string
	AuthToken string
}

const (
	CMD_UPLOAD = "upload"
)

func New(host string, token string) *gbb {
	return &gbb{Host: host, AuthToken: token}
}

func (g *gbb) Run(_ []string) error {
	return gbberror.NotYetImplemented
}

func (g *gbb) HandleUpload(args []string) error {
	var authToken string
	const (
		argAuthToken = "--authToken"
	)

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case argAuthToken:
			i++
			authToken = args[i]
		}
	}
	if strings.TrimSpace(authToken) == "" {
		return gbberror.NoAuthToken
	}
	return gbberror.NotYetImplemented

}
