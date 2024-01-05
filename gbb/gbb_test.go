package gbb_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/go-burn-bits/gbb"
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
	Describe("Upload file", func() {
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
			Expect(instance.HandleUpload([]string{})).Should(MatchError(gbberror.NoAuthToken))
		})
		It("should fail initially with not yet implemented", func() {
			Expect(instance.HandleUpload([]string{"--authToken", "abc"})).Should(MatchError(gbberror.NotYetImplemented))
		})
	})
})
