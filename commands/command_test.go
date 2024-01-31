package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mortedecai/gbb/client"
	"github.com/mortedecai/gbb/client/mocks"
	"github.com/mortedecai/gbb/response"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"os"
)

var _ = Describe("Command Integration Test", func() {
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
	entries := map[string][]struct {
		context     string
		outcome     string
		args        []string
		shouldPanic bool
		addDir      bool
	}{
		"download": {
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
			{
				context:     "string for port",
				outcome:     "should panic",
				args:        []string{"client", "download", "-a", "abc", "-p", "a23a4", "-d"},
				shouldPanic: true,
				addDir:      true,
			},
		},
	}
	for cmd, testEntries := range entries {
		Describe(fmt.Sprintf("Command: %s", cmd), func() {
			for _, e := range testEntries {
				entry := e
				Context(entry.context, func() {
					It(entry.outcome, func() {
						var req *http.Request
						response := createEmptyDownloadResponse()

						os.Args = entry.args
						if entry.addDir {
							os.Args = append(os.Args, tempDir)
						}
						rootCmd, err := Root(version)
						Expect(err).ToNot(HaveOccurred())
						_, err = Download(rootCmd)
						Expect(err).ToNot(HaveOccurred())

						f := func() {
							if err := rootCmd.Execute(); err != nil {
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
