package gbberror

import "errors"

var (
	// Development errors

	ErrNotYetImplemented = errors.New("not yet implemented")

	// Misc driver errors

	ErrBadArguments = errors.New("bad arguments")
	ErrBadHost      = errors.New("bad host")
	ErrBadPort      = errors.New("bad port")
	ErrBadTarget    = errors.New("bad target")

	// Auth'n & Auth'z errors

	ErrNoAuthToken = errors.New("no auth token")

	// File & Directory errors

	ErrNoOutputDir  = errors.New("no output directory supplied")
	ErrBadOutputDir = errors.New("bad output directory")
	ErrFileIssue    = errors.New("file error")

	// HTTP Errors

	ErrUnexpectedResponse = errors.New("unexpected status code")
	ErrResponseReadFailed = errors.New("response read failed")
	ErrBadJSON            = errors.New("bad JSON for marshal/unmarshal")
	ErrRequestFailed      = errors.New("issue creating request")

	// BitBurner Errors

	ErrBitBurnerFailure = errors.New("failed response from BitBurner")
)

const (
	StandardWrapper = "%w: %s"
)
