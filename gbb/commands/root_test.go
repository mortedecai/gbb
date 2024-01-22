package commands

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const version = "v0.0.0-test"

var _ = Describe("Root", func() {
	It("should create a root command object", func() {
		cmd, err := Root(version)
		Expect(cmd).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
	})
})

var _ = Describe("Root Option", func() {
	It("should be possible to create and use a root option", func() {
		ro := &rootOption{
			host:      "example.com",
			port:      1234,
			authToken: "abc",
		}
		Expect(ro.Host()).To(Equal(ro.host))
		Expect(ro.Port()).To(Equal(ro.port))
		Expect(ro.AuthToken()).To(Equal(ro.authToken))
	})
})
var _ = Describe("Root Option", func() {
	var (
		baseRo *rootOption
	)
	BeforeEach(func() {
		baseRo = &rootOption{
			host:      "example.com",
			port:      1234,
			authToken: "abc",
		}
	})
	It("should be possible to create and use a download option", func() {
		ro := baseRo
		Expect(ro.Host()).To(Equal(ro.host))
		Expect(ro.Port()).To(Equal(ro.port))
		Expect(ro.AuthToken()).To(Equal(ro.authToken))
		Expect(ro.Valid()).To(BeTrue())
	})
	It("should not be valid if the root option isn't valid", func() {
		ro := baseRo
		ro.host = ""
		Expect(ro.Valid()).To(BeFalse())
	})
	It("should not be valid if the root option isn't valid", func() {
		ro := baseRo
		ro.authToken = ""
		Expect(ro.Valid()).To(BeFalse())
	})
	It("should not be valid if the root option isn't valid", func() {
		ro := baseRo
		ro.port = 0
		Expect(ro.Valid()).To(BeFalse())
	})
})
