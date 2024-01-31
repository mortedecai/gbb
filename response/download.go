package response

import (
	"os"
	"path"
	"strings"
)

type GBBDownloadFilesResponse struct {
	Success bool                `json:"success"`
	Data    GBBFileDownloadData `json:"data,omitempty"`
}

type GBBFileDownloadData struct {
	Files []GBBDownloadFile `json:"files,omitempty"`
}

type GBBFileName string

func (gfn GBBFileName) IsValid() bool {
	return gfn != ""
}

func (gfn GBBFileName) HasDir() bool {
	return strings.Contains(gfn.String(), "/")
}

func (gfn GBBFileName) String() string {
	return string(gfn)
}

func (gfn GBBFileName) ToAbsolutePath(outputDir string) string {
	return path.Join(outputDir, path.Clean(gfn.String()))
}

func (gfn GBBFileName) CreatePath(outputDir string) error {
	absPath := gfn.ToAbsolutePath(outputDir)
	dirList := path.Dir(absPath)
	return os.MkdirAll(dirList, 0750)
}

type GBBDownloadFile struct {
	Filename GBBFileName `json:"filename"`
	Code     string      `json:"code,omitempty"`
	RamUsage int         `json:"ramUsage,omitempty"`
}
