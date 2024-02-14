package models

import (
	"os"
	"path"
	"strings"
)

type GBBFileName string

func (gfn GBBFileName) IsValid() bool {
	return gfn != ""
}

func (gfn GBBFileName) HasDir() bool {
	return strings.LastIndex(gfn.String(), "/") > 0
}

func (gfn GBBFileName) String() string {
	return string(gfn)
}

func (gfn GBBFileName) Path(outputDir string) string {
	return path.Join(outputDir, path.Clean(gfn.String()))
}

func (gfn GBBFileName) CreatePath(outputDir string) error {
	absPath := gfn.Path(outputDir)
	dirList := path.Dir(absPath)
	return os.MkdirAll(dirList, 0750)
}

type GBBFileData struct {
	Filename GBBFileName `json:"filename"`
	Code     string      `json:"code,omitempty"`
	RamUsage int         `json:"ramUsage,omitempty"`
}
