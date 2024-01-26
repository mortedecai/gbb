package client

import (
	"bytes"
	"errors"
	"github.com/mortedecai/gbb/response"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/mortedecai/gbb/client/mocks"
	"github.com/mortedecai/gbb/gbberror"
)

var _ = Describe("Gbb", func() {
	Describe("Server Calls", func() {
		req := createGenericRequest()
		entries := []struct {
			context          string
			outcome          string
			req              *http.Request
			mockExpectations func(*gomock.Controller)
			errChecker       func(error)
		}{
			{
				context: "client call fails",
				outcome: "A RequestFailed error should be received",
				mockExpectations: func(gc *gomock.Controller) {
					client := mocks.NewMockGBBClient(gc)
					Client = client
					client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(nil, errors.New("doh"))
				},
				errChecker: func(err error) {
					Expect(err).Should(MatchError(gbberror.ErrRequestFailed))
				},
			},
			{
				context: "response has incorrect status",
				outcome: "A RequestFailed error should be received",
				mockExpectations: func(gc *gomock.Controller) {
					client := mocks.NewMockGBBClient(gc)
					Client = client
					client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(createResponse(http.StatusTeapot), nil)
				},
				errChecker: func(err error) {
					Expect(err).Should(MatchError(gbberror.ErrUnexpectedResponse))
				},
			},
		}

		for _, e := range entries {
			entry := e
			Context(entry.context, func() {
				It(entry.outcome, func() {
					mockCtrl := gomock.NewController(GinkgoT())
					defer mockCtrl.Finish()
					entry.mockExpectations(mockCtrl)
					err := handleServerCall(req, http.StatusOK, nil)
					entry.errChecker(err)
				})
			})
		}
	})

	Describe("WriteFiles", func() {
		entries := []struct {
			context         string
			outcome         string
			files           []response.GBBDownloadFile
			errCheck        func(err error)
			expFilesWritten []int
		}{
			{
				context:  "empty file list",
				outcome:  "no files written",
				files:    []response.GBBDownloadFile{},
				errCheck: func(err error) { Expect(err).ToNot(HaveOccurred()) },
			},
			{
				context: "single file lis - zero entriest",
				outcome: "one files written",
				files: []response.GBBDownloadFile{
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
				files: []response.GBBDownloadFile{
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
				files: []response.GBBDownloadFile{
					{
						Filename: "foo/testFile1.js",
						Code:     "// Hi\n// This is a file.",
						RamUsage: 0,
					},
				},
				errCheck: func(err error) { Expect(err).ToNot(HaveOccurred()) },
			},
		}
		for _, e := range entries {
			entry := e
			outputDir, _ := os.MkdirTemp("", "client")
			Context(entry.context, func() {
				It(entry.outcome, func() {
					entry.errCheck(WriteFiles(outputDir, entry.files))
					defer os.RemoveAll(outputDir)

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
