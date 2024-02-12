package models

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"
)

var _ = Describe("File", func() {
	It("should be valid if it has a file name", func() {
		filename := GBBFileName("")
		Expect(filename.IsValid()).To(BeFalse())
		filename = GBBFileName("foo.txt")
		Expect(filename.IsValid()).To(BeTrue())
	})
	It("should have a directory if the filename contains /", func() {
		filename := GBBFileName("foo.txt")
		Expect(filename.HasDir()).To(BeFalse())
		filename = GBBFileName("/foo.txt")
		Expect(filename.HasDir()).To(BeFalse())
		filename = GBBFileName("test/foo.txt")
		Expect(filename.HasDir()).To(BeTrue())
	})
	It("should return the string representation", func() {
		var str string
		filename := GBBFileName("foo.txt")
		Expect(filename.String()).To(BeAssignableToTypeOf(str))
	})
	It("should return the absolute path of the file", func() {
		filename := GBBFileName("foo.txt")
		Expect(filename.Path("/")).To(Equal("/foo.txt"))
		Expect(filename.Path("")).To(Equal("foo.txt"))
		Expect(filename.Path("../../../")).To(Equal("../../../foo.txt"))
	})
	It("should create the path for the file", func() {
		dir, err := os.MkdirTemp("", "gbb-test-*")
		Expect(err).ToNot(HaveOccurred())
		defer os.RemoveAll(dir)
		filename := GBBFileName("test/path/create/foo.txt")
		filename.CreatePath(dir)
		fi, err := os.Stat(path.Join(dir))
		Expect(err).ToNot(HaveOccurred())
		Expect(fi.IsDir()).To(BeTrue())
		fi, err = os.Stat(path.Join(dir + "/test"))
		Expect(err).ToNot(HaveOccurred())
		Expect(fi.IsDir()).To(BeTrue())
		fi, err = os.Stat(path.Join(dir + "/test/path"))
		Expect(err).ToNot(HaveOccurred())
		Expect(fi.IsDir()).To(BeTrue())
		fi, err = os.Stat(path.Join(dir + "/test/path/create"))
		Expect(err).ToNot(HaveOccurred())
		Expect(fi.IsDir()).To(BeTrue())
		fi, err = os.Stat(path.Join(dir + "/test/path/create/foo.txt"))
		Expect(err).To(HaveOccurred())
	})
})
