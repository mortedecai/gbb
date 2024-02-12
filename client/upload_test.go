package client

import (
	"encoding/json"
	"fmt"
	"github.com/mortedecai/gbb/client/mocks"
	"github.com/mortedecai/gbb/models"
	"github.com/mortedecai/gbb/response"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Upload", func() {
	Describe("Upload single file", func() {
		var (
			uo         *mocks.MockUploadOption
			mockClient *mocks.MockGBBClient
			req        *http.Request
			origClient GBBClient
		)
		BeforeEach(func() {
			ctrl := gomock.NewController(GinkgoT())
			uo = mocks.NewMockUploadOption(ctrl)
			mockClient = mocks.NewMockGBBClient(ctrl)

			origClient = Client
			Client = mockClient
		})
		AfterEach(func() {
			Client = origClient
		})
		It("should return not yet implemented until implementation is complete", func() {
			successResponse := response.GBBUploadFileResponse{Success: true}
			successData, err := json.Marshal(successResponse)
			Expect(err).ToNot(HaveOccurred())

			mockResp := httptest.NewRecorder()
			mockResp.Header().Add("Content-Type", "application/json")
			mockResp.WriteHeader(200)
			mockResp.Write(successData)

			uo.EXPECT().Server().Return("http://localhost:9990").Times(1)
			uo.EXPECT().ToUpload().Return([]models.GBBFileName{"testdata/upload-test-1.js"})
			uo.EXPECT().AuthToken().Return(token).Times(1)
			addAuth := uo.EXPECT().AddAuth(gomock.AssignableToTypeOf(req))
			addAuth.Do(func(r *http.Request) {
				r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", uo.AuthToken()))
				addAuth.Return(r)
			})

			mockClient.EXPECT().Do(gomock.AssignableToTypeOf(req)).Return(mockResp.Result(), nil)
			Expect(HandleUpload(uo)).ToNot(HaveOccurred())
		})
	})
})
