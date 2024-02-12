package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mortedecai/gbb/client"
	"github.com/mortedecai/gbb/client/mocks"
	"github.com/mortedecai/gbb/models"
	"github.com/mortedecai/gbb/response"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"os"
)

const (
	localhost = "localhost"
	localport = "9990"
	localfile = "./command_test.go"
	authToken = "abc"
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
	Describe("Upload File", func() {
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
			context      string
			outcome      string
			args         []string
			outputLogs   []string
			errorMatcher func(err error)
			callTimes    int
		}{
			{
				context: "no flags",
				outcome: "should error due to missing flags authToken and file",
				args:    []string{"gbb", "upload"},
				errorMatcher: func(err error) {
					Expect(err).To(MatchError(errors.New(`required flag(s) "authToken", "file" not set`)))
				},
				callTimes: 0,
			},
			{
				context:      "no file",
				outcome:      "should error due to file not present",
				args:         []string{"gbb", "upload", "-H", localhost, "-p", localport, "-a", authToken},
				errorMatcher: func(err error) { Expect(err).To(MatchError(errors.New(`required flag(s) "file" not set`))) },
				callTimes:    0,
			},
			{
				context:      "auth token missing",
				outcome:      "should error due to auth token not present",
				args:         []string{"gbb", "upload", "-H", localhost, "-p", localport, "-f", localfile},
				errorMatcher: func(err error) { Expect(err).To(MatchError(errors.New(`required flag(s) "authToken" not set`))) },
				callTimes:    0,
			},
			//
			// TODO: TDD Usage, the following need to be corrected upon implementation completion:
			//   * outcome: "should not error",
			//   * errorMatcher: func(err) error { Expect(err).ToNot(HaveOccurred()),
			//   * callTimes: 1,
			//
			{
				context:      "no port",
				outcome:      "should error due to port not yet implemented",
				args:         []string{"gbb", "upload", "-H", localhost, "-f", localfile, "-a", authToken},
				errorMatcher: func(err error) { Expect(err).ToNot(HaveOccurred()) },
				callTimes:    1,
			},
			{
				context:      "no host",
				outcome:      "should error due to not yet implemented",
				args:         []string{"gbb", "upload", "-p", localport, "-f", localfile, "-a", authToken},
				errorMatcher: func(err error) { Expect(err).ToNot(HaveOccurred()) },
				callTimes:    1,
			},
			{
				context:      "all present",
				outcome:      "should have error 'not yet implemented'",
				args:         []string{"gbb", "upload", "-H", localhost, "-p", localport, "-f", localfile, "-a", authToken},
				errorMatcher: func(err error) { Expect(err).ToNot(HaveOccurred()) },
				callTimes:    1,
			},
		}
		for _, v := range entries {
			entry := v
			Context(entry.context, func() {
				It(entry.outcome, func() {
					var req *http.Request
					response := createEmptyDownloadResponse()

					os.Args = entry.args
					rootCmd, err := Root(version)
					Expect(err).ToNot(HaveOccurred())
					_, err = Upload(rootCmd)
					Expect(err).ToNot(HaveOccurred())
					mockClient.EXPECT().Do(gomock.AssignableToTypeOf(req)).Return(response, nil).Times(entry.callTimes)
					err = rootCmd.Execute()
					entry.errorMatcher(err)
				})
			})
		}
	})
	Describe("Download Files", func() {

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
			outputLogs  []string
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
				args:        []string{"gbb", "download", "-H", localhost},
				shouldPanic: true,
			},
			{
				context:     "all supplied",
				outcome:     "should not panic",
				args:        []string{"gbb", "download", "-H", localhost, "-p", localport, "-a", "abc", "-d"},
				shouldPanic: false,
				addDir:      true,
			},
			{
				context:     "no download dir",
				outcome:     "should panic",
				args:        []string{"gbb", "download", "-H", localhost, "-p", localport, "-a", "abc"},
				shouldPanic: true,
			},
			{
				context:     "no host",
				outcome:     "should not panic",
				args:        []string{"gbb", "download", "-p", localport, "-a", "abc", "-d"},
				shouldPanic: false,
				addDir:      true,
			},
			{
				context:     "only auth token & output dir",
				outcome:     "should not panic",
				args:        []string{"gbb", "download", "-a", "abc", "-d"},
				shouldPanic: false,
				addDir:      true,
			},
			{
				context:     "string for port",
				outcome:     "should panic",
				args:        []string{"gbb", "download", "-a", "abc", "-p", "a23a4", "-d"},
				shouldPanic: true,
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
})

func createSuccessfulUploadResponse() *http.Response {
	basicResponse := response.GBBUploadFileResponse{Success: true}
	data, err := json.Marshal(basicResponse)
	Expect(err).ToNot(HaveOccurred())
	resp := &http.Response{
		Body:       io.NopCloser(bytes.NewReader(data)),
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}
	return resp
}

func createEmptyDownloadResponse() *http.Response {
	basicResponse := response.GBBDownloadFilesResponse{
		Success: true,
		Data: response.GBBFileDownloadData{
			Files: make([]models.GBBFileData, 0),
		},
	}
	data, err := json.Marshal(basicResponse)
	Expect(err).ToNot(HaveOccurred())
	resp := &http.Response{
		Body:       io.NopCloser(bytes.NewReader(data)),
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}
	return resp
}
