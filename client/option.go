package client

import "net/http"

// CommandOption represents the available actions on a base CommandOption
type CommandOption interface {
	AuthToken() string
	Host() string
	Port() int
	Valid() bool
	AddAuth(req *http.Request) *http.Request
}

//go:generate mockgen -destination=./mocks/mock_root_option.go -package=mocks github.com/mortedecai/gbb/client CommandOption

// DownloadOption represents the options available to the Download command.
type DownloadOption interface {
	CommandOption
	Destination() string
}

//go:generate mockgen -destination=./mocks/mock_download_option.go -package=mocks github.com/mortedecai/gbb/client DownloadOption
