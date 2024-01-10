package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Go Burn Bits", func() {

	Describe("Initial basic test for coverage and script setup", func() {
		It("should return 'Go Burn Bits'", func() {
			Expect(greetings()).To(Equal("Go Burn Bits"))
		})
		It("should  print 'Go Burn Bits'", func() {
			reader, writer, err := os.Pipe()
			Expect(err).ToNot(HaveOccurred())

			stdout := os.Stdout
			stderr := os.Stderr
			defer func() {
				os.Stdout = stdout
				os.Stderr = stderr
			}()

			os.Stdout = writer
			os.Stderr = writer

			out := make(chan string)

			wg := new(sync.WaitGroup)
			wg.Add(1)
			version = "v0.0.0-test"
			main()

			go func() {
				var buf bytes.Buffer
				wg.Done()
				io.Copy(&buf, reader)
				out <- buf.String()
			}()

			wg.Wait()

			writer.Close()

			str := strings.TrimSpace(<-out)
			Expect(str).To(HavePrefix("Go Burn Bits [v0.0.0-test]"))
		})
	})
})
