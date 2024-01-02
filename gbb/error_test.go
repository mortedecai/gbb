package gbb_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/go-burn-bits/gbb"
)

var _ = Describe("Error", func() {
	const (
		method1 = "hello"
		method2 = "Goodbye"

		msg1 = "Simple error message"
	)
	entries := []struct {
		context string
		expectation string
		method string
		message string
		details interface{}
		expErrorJSON string
	}{
		{
			context: "no details",
			expectation: "should create an error with no details",
			method: method1,
			message: msg1,
			expErrorJSON: fmt.Sprintf(`{"method":"%s","message":"%s"}`, method1, msg1),
		},
		{
			context: "no details",
			expectation: "should create an error with no details",
			method: method1,
			message: msg1,
			details: struct{
				File string `json:"file"`
				Line int `json:"line"`
			}{
				File: "/home/foo",
				Line: 10,
			},
			expErrorJSON: fmt.Sprintf(`{"method":"%s","message":"%s","details":{"file":"/home/foo","line":10}}`, method1, msg1),
		},
	}
	for _, e := range entries {
		entry := e
		Context(entry.context, func() {
			It(entry.expectation, func() {
				err := gbb.CreateError(entry.method, entry.message, entry.details)
				Expect(err).To(HaveOccurred())
				Expect(err.JSON()).To(Equal(entry.expErrorJSON))
			})
		})
	}
})
