package gbberror

import "errors"

var (
	// Development errors
	NotYetImplemented = errors.New("not yet implemented")

	// Auth'n & Auth'z errors
	NoAuthToken = errors.New("no auth token")

	// File & Directory errors
	NoOutputDir  = errors.New("no output direcotry supplied")
	BadOutputDir = errors.New("bad output directory")
	FileIssue    = errors.New("file error")

	// HTTP Errors
	UnexpectedResponse = errors.New("unexpected status code")
	ResponseReadFailed = errors.New("response read failed")
	BadJSON            = errors.New("bad JSON for marshal/unmarshal")
	RequestFailed      = errors.New("issue creating request")

	// BitBurner Errors
	BitBurnerFailure = errors.New("failed response from BitBurner")
)

const (
	StandardWrapper = "%w: %s"
)
