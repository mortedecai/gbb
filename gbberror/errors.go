package gbberror

import "errors"

var (
	NotYetImplemented = errors.New("not yet implemented")
	NoAuthToken       = errors.New("no auth token")
)
