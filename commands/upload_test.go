package commands

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Upload", func() {
	It("should create an upload command object", func() {
		rootCmd, err := Root(version)
		Expect(err).ToNot(HaveOccurred())
		Expect(rootCmd).ToNot(BeNil())

		cmd, err := Upload(rootCmd)
		Expect(err).ToNot(HaveOccurred())
		Expect(cmd).ToNot(BeNil())
	})
})

var _ = Describe("Upload Option", func() {
	var (
		opt *uploadOption
	)
	BeforeEach(func() {
		opt = &uploadOption{
			rootOption: &rootOption{host: localhost, port: 9990, authToken: authToken},
		}
	})
	It("should not be valid if there are no files to upload", func() {
		Expect(opt.ToUpload()).To(BeEmpty())
		Expect(opt.Valid()).To(BeFalse())
	})
	It("should become valid if there is a file to upload", func() {
		opt.toUpload = "./foo.js"
		Expect(opt.ToUpload()).ToNot(BeEmpty())
		Expect(opt.Valid()).To(BeTrue())
	})
})
