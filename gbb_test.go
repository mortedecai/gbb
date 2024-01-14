package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

var _ = Describe("Go Burn Bits", func() {

	Describe("Initial basic test for coverage and script setup", func() {
		It("should return 'Go Burn Bits'", func() {
			Expect(greetings()).To(Equal("Go Burn Bits"))
		})
		It("should  print 'Go Burn Bits'", func() {
			observedZapCore, observedLogs := observer.New(zap.InfoLevel)
			observedLogger := zap.New(observedZapCore)
			l := logger
			logger = observedLogger.Sugar()
			version = "v0.0.0-test"
			main()
			Expect(observedLogs.All()[0].Message).To(HavePrefix("Go Burn Bits [v0.0.0-test]"))
			logger = l
		})
	})
})
