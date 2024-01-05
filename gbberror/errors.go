package gbberror

import "errors"

var (
	NotYetImplemented  = errors.New("not yet implemented")
	NoAuthToken        = errors.New("no auth token")
	NoOutputDir        = errors.New("no output direcotry supplied")
	BadOutputDir       = errors.New("bad output directory")
	FileIssue          = errors.New("file error")
	UnexpectedResponse = errors.New("unexpected status code")
	ResponseReadFailed = errors.New("response read failed")
	BadJSON            = errors.New("bad JSON for marshal/unmarshal")
	BitBurnerFailure   = errors.New("failed response from BitBurner")
)
