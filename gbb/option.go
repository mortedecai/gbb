package gbb

import "net/http"

// CommandOption represents the available actions on a base CommandOption
type CommandOption interface {
	AuthToken() string
	Host() string
	Port() int
	Valid() bool
	AddAuth(req *http.Request) *http.Request
}

// DownloadOption represents the options available to the Download command.
type DownloadOption interface {
	CommandOption
	Destination() string
}
