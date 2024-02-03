package client

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/gbb/gbberror"
)

var _ = Describe("Upload", func() {
	Describe("Not Yet Implemented", func() {
		Expect(HandleUpload(nil)).To(MatchError(gbberror.ErrNotYetImplemented))
	})
})
