package gbb_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/mortedecai/go-burn-bits/gbb"
	"github.com/mortedecai/go-burn-bits/gbb/mocks"
	"github.com/mortedecai/go-burn-bits/gbberror"
)

var _ = Describe("Gbb", func() {
	It("should be possible to create a new gbb instance", func() {
		host := "localhost:9990"
		token := "abc"
		g := gbb.New(host, token)
		Expect(g).ToNot(BeNil())
		Expect(g.Host).To(Equal(host))
		Expect(g.AuthToken).To(Equal(token))
	})
	Describe("Run", func() {
		var (
			instance gbb.GoBurnBits
		)
		const (
			localhost = "localhost"
			token     = "abc"
		)
		BeforeEach(func() {
			instance = gbb.New(localhost, token)
		})
		It("should fail initially with not yet implemented", func() {
			Expect(instance.Run([]string{""})).Should(MatchError(gbberror.NotYetImplemented))
		})
	})
	Describe("Download file", func() {
		var (
			instance gbb.GoBurnBits
		)
		const (
			localhost = "localhost"
			token     = "abc"
		)
		BeforeEach(func() {
			instance = gbb.New(localhost, token)
		})
		It("should fail with an error if no auth token is supplied", func() {
			Expect(instance.HandleDownload([]string{})).Should(MatchError(gbberror.NoAuthToken))
		})
		It("should fail if no output direcotry is supplied to download", func() {
			Expect(instance.HandleDownload([]string{"--authToken", "abc"})).Should(MatchError(gbberror.NoOutputDir))
		})
		It("should fail if only an empty output dir is supplied", func() {
			Expect(instance.HandleDownload([]string{"--authToken", "abc", "--outputDir", ""})).Should(MatchError(gbberror.NoOutputDir))
		})
		It("should fail if output dir is outside of the current directory (eg. ../)", func() {
			Expect(instance.HandleDownload([]string{"--authToken", "abc", "--outputDir", "../jetsons"})).Should(MatchError(gbberror.BadOutputDir))
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
			instance.(*gbb.GBB).Client = client

			req, err := http.NewRequest(http.MethodGet, "http://localhost", bytes.NewBuffer([]byte("{}")))
			Expect(err).ToNot(HaveOccurred())
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			data, err := os.ReadFile("testdata/sample_download_files.json")
			Expect(err).ToNot(HaveOccurred())
			mockResp := httptest.NewRecorder()
			mockResp.WriteHeader(http.StatusOK)
			mockResp.Write(data)

			client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(mockResp.Result(), nil)
			args := []string{"--authToken", "abc", "--outputDir", relativeDir}
			Expect(instance.HandleDownload(args)).ToNot(HaveOccurred())

			var filesResponse gbb.GBBDownloadFilesResponse
			Expect(json.Unmarshal(data, &filesResponse)).ToNot(HaveOccurred())

			entries, err := os.ReadDir(dir)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(entries)).To(Equal(len(filesResponse.Data.Files)))

			for _, v := range filesResponse.Data.Files {
				writtenData, err := os.ReadFile(path.Join(dir, v.Filename))
				Expect(err).ToNot(HaveOccurred())
				Expect(string(writtenData)).To(Equal(v.Code))
			}

		})
	})
})
