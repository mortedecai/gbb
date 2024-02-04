package client

import (
	"github.com/mortedecai/gbb/client/mocks"
	"github.com/mortedecai/gbb/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/mortedecai/gbb/gbberror"
)

var _ = Describe("Upload", func() {
	Describe("Not Yet Implemented", func() {
		var (
			uo *mocks.MockUploadOption
		)
		BeforeEach(func() {
			ctrl := gomock.NewController(GinkgoT())
			uo = mocks.NewMockUploadOption(ctrl)
		})
		It("should return not yet implemented until implementation is complete", func() {
			uo.EXPECT().Server().Return("http://localhost:9990").Times(1)
			uo.EXPECT().ToUpload().Return([]models.GBBFileName{"testdata/upload-test-1.js"})
			Expect(HandleUpload(uo)).To(MatchError(gbberror.ErrNotYetImplemented))
		})
	})
})
