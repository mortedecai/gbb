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
