package gbb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/mortedecai/go-burn-bits/gbb/mocks"
	"github.com/mortedecai/go-burn-bits/gbberror"
)

var _ = Describe("Gbb", func() {
	It("should be possible to create a new gbb instance", func() {
		host := "localhost:9990"
		token := "abc"
		g := New(host, token)
		Expect(g).ToNot(BeNil())
		Expect(g.Host).To(Equal(host))
		Expect(g.AuthToken).To(Equal(token))
	})
	Describe("Run", func() {
		var (
			instance GoBurnBits
		)
		const (
			localhost = "localhost"
			token     = "abc"
		)
		BeforeEach(func() {
			instance = New(localhost, token)
		})
		It("should fail with bad arguments if none given", func() {
			Expect(instance.Run([]string{})).Should(MatchError(gbberror.ErrBadArguments))
		})
	})
	Describe("Server Calls", func() {
		req := createGenericRequest()
		entries := []struct {
			context          string
			outcome          string
			req              *http.Request
			mockExpectations func(*gomock.Controller, *GBB)
			errChecker       func(error)
		}{
			{
				context: "client call fails",
				outcome: "A RequestFailed error should be received",
				mockExpectations: func(gc *gomock.Controller, gbb *GBB) {
					client := mocks.NewMockGBBClient(gc)
					gbb.Client = client
					client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(nil, errors.New("doh"))
				},
				errChecker: func(err error) {
					Expect(err).Should(MatchError(gbberror.ErrRequestFailed))
				},
			},
			{
				context: "response has incorrect status",
				outcome: "A RequestFailed error should be received",
				mockExpectations: func(gc *gomock.Controller, gbb *GBB) {
					client := mocks.NewMockGBBClient(gc)
					gbb.Client = client
					client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(createResponse(http.StatusTeapot), nil)
				},
				errChecker: func(err error) {
					Expect(err).Should(MatchError(gbberror.ErrUnexpectedResponse))
				},
			},
		}

		for _, e := range entries {
			entry := e
			instance := New("localhost", "abc")
			Context(entry.context, func() {
				It(entry.outcome, func() {
					mockCtrl := gomock.NewController(GinkgoT())
					defer mockCtrl.Finish()
					entry.mockExpectations(mockCtrl, instance)
					err := instance.handleServerCall(req, http.StatusOK, nil)
					entry.errChecker(err)
				})
			})
		}
	})

	Describe("Run", func() {
		var (
			instance GoBurnBits
		)
		const (
			localhost = "localhost"
			token     = "abc"
		)
		BeforeEach(func() {
			instance = New(localhost, token)
		})
		It("should return an error with insufficient arguments", func() {
			Expect(instance.Run([]string{})).Should(MatchError(gbberror.ErrBadArguments))
		})
	})

	Describe("Download file", func() {
		var (
			instance GoBurnBits
		)
		const (
			localhost = "localhost"
			token     = ""
		)
		BeforeEach(func() {
			instance = New(localhost, token)
		})
		It("should fail with an error if no auth token is supplied", func() {
			Expect(instance.Run([]string{"download"})).Should(MatchError(gbberror.ErrNoAuthToken))
		})
		It("should fail if no output direcotry is supplied to download", func() {
			Expect(instance.Run([]string{"download", "--authToken", "abc"})).Should(MatchError(gbberror.ErrNoOutputDir))
		})
		It("should fail if only an empty output dir is supplied", func() {
			Expect(instance.Run([]string{"download", "--authToken", "abc", "--outputDir", ""})).Should(MatchError(gbberror.ErrNoOutputDir))
		})
		It("should fail if output dir is outside of the current directory (eg. ../)", func() {
			Expect(instance.Run([]string{"download", "--authToken", "abc", "--outputDir", "../jetsons"})).Should(MatchError(gbberror.ErrBadOutputDir))
		})

		It("should call to the server and handle the response", func() {
			wd, err := os.Getwd()
			Expect(err).ToNot(HaveOccurred())
			mockCtrl := gomock.NewController(GinkgoT())
			defer mockCtrl.Finish()
			dir, err := os.MkdirTemp(wd, "gbb")
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dir)

			relativeDir := dir[len(wd)+1:]

			client := mocks.NewMockGBBClient(mockCtrl)
			instance.(*GBB).Client = client

			req, err := http.NewRequest(http.MethodGet, "http://localhost", bytes.NewBuffer([]byte("{}")))
			Expect(err).ToNot(HaveOccurred())
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			data, err := os.ReadFile("testdata/sample_download_files.json")
			Expect(err).ToNot(HaveOccurred())
			mockResp := httptest.NewRecorder()
			mockResp.WriteHeader(http.StatusOK)
			mockResp.Write(data)

			client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(mockResp.Result(), nil)
			args := []string{"download", "--authToken", "abc", "--outputDir", relativeDir}
			Expect(instance.Run(args)).ToNot(HaveOccurred())

			var filesResponse GBBDownloadFilesResponse
			Expect(json.Unmarshal(data, &filesResponse)).ToNot(HaveOccurred())

			entries, err := os.ReadDir(dir)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(entries)).To(Equal(len(filesResponse.Data.Files)))

			for _, v := range filesResponse.Data.Files {
				writtenData, err := os.ReadFile(v.Filename.ToAbsolutePath(dir))
				Expect(err).ToNot(HaveOccurred())
				Expect(string(writtenData)).To(Equal(v.Code))
			}
		})
	})

	Describe("WriteFiles", func() {
		entries := []struct {
			context         string
			outcome         string
			files           []GBBDownloadFile
			errCheck        func(err error)
			expFilesWritten []int
		}{
			{
				context:  "empty file list",
				outcome:  "no files written",
				files:    []GBBDownloadFile{},
				errCheck: func(err error) { Expect(err).ToNot(HaveOccurred()) },
			},
			{
				context: "single file lis - zero entriest",
				outcome: "one files written",
				files: []GBBDownloadFile{
					{
						Filename: "",
						Code:     "",
						RamUsage: 0,
					},
				},
				errCheck: func(err error) { Expect(err).ToNot(HaveOccurred()) },
			},
			{
				context: "single file list - no directory",
				outcome: "one files written",
				files: []GBBDownloadFile{
					{
						Filename: "testFile1.js",
						Code:     "// Hi\n// This is a file.",
						RamUsage: 0,
					},
				},
				errCheck: func(err error) { Expect(err).ToNot(HaveOccurred()) },
			},
			{
				context: "single file list - in directory",
				outcome: "one files written",
				files: []GBBDownloadFile{
					{
						Filename: "foo/testFile1.js",
						Code:     "// Hi\n// This is a file.",
						RamUsage: 0,
					},
				},
				errCheck: func(err error) { Expect(err).ToNot(HaveOccurred()) },
			},
		}
		const localhost = "localhost"
		for _, e := range entries {
			entry := e
			outputDir, _ := os.MkdirTemp("", "gbb")
			Context(entry.context, func() {
				It(entry.outcome, func() {
					instance := New(localhost, "")
					entry.errCheck(instance.WriteFiles(outputDir, entry.files))

					for _, idx := range entry.expFilesWritten {
						path := entry.files[idx].Filename.ToAbsolutePath(outputDir)
						fi, err := os.Stat(path)
						Expect(err).ToNot(HaveOccurred())
						Expect(fi.IsDir()).To(BeFalse())
						data, err := os.ReadFile(path)
						Expect(err).ToNot(HaveOccurred())
						Expect(string(data)).To(Equal(entry.files[idx].Code))
					}
				})
			})
		}
	})
})

func createGenericRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost", bytes.NewBuffer([]byte("{}")))
	return req
}

func createResponse(status int) *http.Response {
	resp := httptest.NewRecorder()
	resp.WriteHeader(status)
	return resp.Result()
}
