package commands

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {
	It("should create a list command object", func() {
		root, err := Root(version)
		Expect(err).ToNot(HaveOccurred())
		cmd, err1 := Download(root)
		Expect(cmd).ToNot(BeNil())
		Expect(err1).ToNot(HaveOccurred())
	})
})

var _ = Describe("Download Option", func() {
	var (
		baseLo *listOption
	)
	BeforeEach(func() {
		baseLo = &listOption{
			rootOption: &rootOption{
				host:      "example.com",
				port:      1234,
				authToken: "abc",
			},
		}
	})
	It("should be possible to create and use a download option", func() {
		lo := baseLo
		Expect(lo.Host()).To(Equal(lo.host))
		Expect(lo.Port()).To(Equal(lo.port))
		Expect(lo.AuthToken()).To(Equal(lo.authToken))
		Expect(lo.Valid()).To(BeTrue())
	})
	It("should not be valid if the root option isn't valid", func() {
		lo := baseLo
		lo.host = ""
		Expect(lo.Valid()).To(BeFalse())
	})
	It("should be possible to create the list command", func() {
		rootCmd, rootErr := Root("v0.0.0")
		Expect(rootCmd).ToNot(BeNil())
		Expect(rootErr).ToNot(HaveOccurred())
		listCmd, listErr := List(rootCmd)
		Expect(listCmd).ToNot(BeNil())
		Expect(listErr).ToNot(HaveOccurred())
	})
})
