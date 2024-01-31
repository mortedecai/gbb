package main

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"io"
	"os"
	"strings"
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Go Burn Bits", func() {

	Describe("Initial basic test for coverage and script setup", func() {
		var (
			originalArgs   []string
			originalLogger *zap.SugaredLogger
			core           zapcore.Core
			logs           *observer.ObservedLogs
		)
		BeforeEach(func() {
			originalLogger = logger
			originalArgs = os.Args
			core, logs = observer.New(zap.ErrorLevel)
			logger = zap.New(core).Sugar()
		})
		AfterEach(func() {
			os.Args = originalArgs
			logger = originalLogger
		})
		It("should print 'Go Burn Bits'", func() {
			os.Args = []string{"client", "-v"}
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
			Expect(str).To(HaveSuffix("Go Burn Bits [v0.0.0-test]"))
		})
		// Note: This test is done differently as we are checking the zap logs as opposed to std out / std err
		It("should log an error", func() {
			os.Args = []string{"client", "hoooah"}
			main()
			entry := logs.All()[0]
			fmt.Printf("entry: %v\ncontext map: %v\n", entry, entry.ContextMap())
			Expect(entry.Message).To(Equal(cmdResMsg))
			contextMap := entry.ContextMap()
			Expect(contextMap[resMsg]).To(Equal(errResult))
			Expect(contextMap[detailsMsg]).To(ContainSubstring(`unknown command "hoooah"`))
		})
	})
})
