package gbb_test

import (
	"github.com/mortedecai/gbb/config"
	"github.com/mortedecai/gbb/gbb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	It("should be possible to create a non-nil app from default config", func() {
		app := gbb.New(config.Default())
		Expect(app).ToNot(BeNil())
	})
})
