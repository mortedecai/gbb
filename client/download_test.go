package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mortedecai/gbb/client/mocks"
	"github.com/mortedecai/gbb/response"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"os"
)

const (
	localhost = "localhost"
	token     = "abc"
)

var _ = Describe("Download file", func() {
	It("should call to the server and handle the response", func() {
		var req *http.Request
		wd, err := os.Getwd()
		Expect(err).ToNot(HaveOccurred())
		mockCtrl := gomock.NewController(GinkgoT())
		defer mockCtrl.Finish()
		dir, err := os.MkdirTemp(wd, "client")
		Expect(err).ToNot(HaveOccurred())
		defer os.RemoveAll(dir)

		client := mocks.NewMockGBBClient(mockCtrl)
		opt := mocks.NewMockDownloadOption(mockCtrl)
		opt.EXPECT().Host().Return(localhost).Times(1)
		opt.EXPECT().Port().Return(9990).Times(1)
		opt.EXPECT().Destination().Return(dir).Times(1)
		opt.EXPECT().AuthToken().Return(token).Times(1)
		addAuth := opt.EXPECT().AddAuth(gomock.AssignableToTypeOf(req))
		addAuth.Do(func(r *http.Request) {
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", opt.AuthToken()))
			addAuth.Return(r)
		})
		Client = client

		req, err = http.NewRequest(http.MethodGet, "http://localhost", bytes.NewBuffer([]byte("{}")))
		Expect(err).ToNot(HaveOccurred())
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		data, err := os.ReadFile("testdata/sample_download_files.json")
		Expect(err).ToNot(HaveOccurred())
		mockResp := httptest.NewRecorder()
		mockResp.WriteHeader(http.StatusOK)
		mockResp.Write(data)

		client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(mockResp.Result(), nil)
		Expect(HandleDownload(opt)).ToNot(HaveOccurred())

		var filesResponse response.GBBDownloadFilesResponse
		Expect(json.Unmarshal(data, &filesResponse)).ToNot(HaveOccurred())

		entries, err := os.ReadDir(dir)
		Expect(err).ToNot(HaveOccurred())
		Expect(len(entries)).To(Equal(len(filesResponse.Data.Files)))

		for _, v := range filesResponse.Data.Files {
			writtenData, err := os.ReadFile(v.Filename.Path(dir))
			Expect(err).ToNot(HaveOccurred())
			Expect(string(writtenData)).To(Equal(v.Code))
		}
	})

})
