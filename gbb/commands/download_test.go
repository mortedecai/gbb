package commands

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Download", func() {
	It("should create a download command object", func() {
		root, err := Root(version)
		Expect(err).ToNot(HaveOccurred())
		cmd, err1 := Download(root)
		Expect(cmd).ToNot(BeNil())
		Expect(err1).ToNot(HaveOccurred())
	})
})

var _ = Describe("Download Option", func() {
	var (
		baseDo *downloadOption
	)
	BeforeEach(func() {
		baseDo = &downloadOption{
			rootOption: &rootOption{
				host:      "example.com",
				port:      1234,
				authToken: "abc",
			},
			destination: "./",
		}
	})
	It("should be possible to create and use a download option", func() {
		do := baseDo
		Expect(do.Host()).To(Equal(do.host))
		Expect(do.Port()).To(Equal(do.port))
		Expect(do.AuthToken()).To(Equal(do.authToken))
		Expect(do.Destination()).To(Equal(do.destination))
		Expect(do.Valid()).To(BeTrue())
	})
	It("should not be valid if destination is empty", func() {
		do := baseDo
		do.destination = ""
		Expect(do.Valid()).To(BeFalse())
	})
	It("should not be valid if the root option isn't valid", func() {
		do := baseDo
		do.host = ""
		Expect(do.Valid()).To(BeFalse())
	})
})
