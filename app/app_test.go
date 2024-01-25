package app_test

import (
	"bytes"
	"encoding/json"
	"github.com/mortedecai/gbb/client"
	"github.com/mortedecai/gbb/client/mocks"
	"github.com/mortedecai/gbb/response"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/gbb/app"
)

var _ = Describe("App", func() {
	var (
		originalArgs []string
		tempDir      string
		mockClient   *mocks.MockGBBClient
	)
	const (
		version = "0.0.0-test"
	)
	BeforeEach(func() {
		// Cobra uses the os.Args value directly. Capture the original arguments for reversion after each test
		originalArgs = os.Args
		if t, err := os.MkdirTemp("", "client-*"); err != nil {
			Fail("could not create temp directory")
		} else {
			tempDir = t
		}
		ctrl := gomock.NewController(GinkgoT())
		mockClient = mocks.NewMockGBBClient(ctrl)
		client.Client = mockClient
	})
	AfterEach(func() {
		// Restore os.Args value
		os.Args = originalArgs
		originalArgs = []string{}
		os.RemoveAll(tempDir)
	})
	entries := []struct {
		context     string
		outcome     string
		args        []string
		shouldPanic bool
		addDir      bool
	}{
		{
			context:     "no flags",
			outcome:     "should panic",
			args:        []string{},
			shouldPanic: true,
		},
		{
			context:     "no authToken",
			outcome:     "should panic",
			args:        []string{"client", "download", "-H", "localhost"},
			shouldPanic: true,
		},
		{
			context:     "all supplied",
			outcome:     "should not panic",
			args:        []string{"client", "download", "-H", "localhost", "-p", "9990", "-a", "abc", "-d"},
			shouldPanic: false,
			addDir:      true,
		},
		{
			context:     "no download dir",
			outcome:     "should panic",
			args:        []string{"client", "download", "-H", "localhost", "-p", "9990", "-a", "abc"},
			shouldPanic: true,
		},
		{
			context:     "no host",
			outcome:     "should not panic",
			args:        []string{"client", "download", "-p", "9990", "-a", "abc", "-d"},
			shouldPanic: false,
			addDir:      true,
		},
		{
			context:     "only auth token & output dir",
			outcome:     "should not panic",
			args:        []string{"client", "download", "-a", "abc", "-d"},
			shouldPanic: false,
			addDir:      true,
		},
	}

	for _, e := range entries {
		entry := e
		Context(entry.context, func() {
			It(entry.outcome, func() {
				var req *http.Request
				response := createEmptyDownloadResponse()

				os.Args = entry.args
				if entry.addDir {
					os.Args = append(os.Args, tempDir)
				}
				a := app.New(version)

				f := func() {
					if err := a.Run(); err != nil {
						panic(err)
					}
				}
				if entry.shouldPanic {
					mockClient.EXPECT().Do(gomock.AssignableToTypeOf(req)).Return(response, nil).Times(0)
					Expect(f).To(Panic())
				} else {
					mockClient.EXPECT().Do(gomock.AssignableToTypeOf(req)).Return(response, nil).Times(1)
					Expect(f).ToNot(Panic())
				}
			})
		})
	}
})

func createEmptyDownloadResponse() *http.Response {
	basicResponse := response.GBBDownloadFilesResponse{
		Success: true,
		Data: response.GBBFileDownloadData{
			Files: make([]response.GBBDownloadFile, 0),
		},
	}
	data, err := json.Marshal(basicResponse)
	Expect(err).ToNot(HaveOccurred())
	response := &http.Response{
		Body:       io.NopCloser(bytes.NewReader(data)),
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}
	return response
}
