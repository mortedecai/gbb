package gbb_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/go-burn-bits/gbb"
)

type errPosition struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

var _ = Describe("Error", func() {
	const (
		method1 = "hello"
		method2 = "Goodbye"

		msg1 = "Simple error message"
	)
	var (
		fooErrPosition = errPosition{
			File: "/home/foo",
			Line: 10,
		}
		entries = []struct {
			context      string
			expectation  string
			method       string
			message      string
			details      interface{}
			expErrorJSON string
			expError     string
		}{
			{
				context:      "no details",
				expectation:  "should create an error with no details",
				method:       method1,
				message:      msg1,
				expErrorJSON: fmt.Sprintf(`{"method":"%s","message":"%s"}`, method1, msg1),
				expError:     fmt.Sprintf("[%s] %s [Details: |%v|]", method1, msg1, nil),
			},
			{
				context:      "no details",
				expectation:  "should create an error with no details",
				method:       method1,
				message:      msg1,
				details:      fooErrPosition,
				expErrorJSON: fmt.Sprintf(`{"method":"%s","message":"%s","details":{"file":"/home/foo","line":10}}`, method1, msg1),
				expError:     fmt.Sprintf("[%s] %s [Details: |%v|]", method1, msg1, fooErrPosition),
			},
		}
	)
	for _, e := range entries {
		entry := e
		Context(entry.context, func() {
			It(entry.expectation, func() {
				err := gbb.CreateError(entry.method, entry.message, entry.details)
				Expect(err).To(HaveOccurred())
				Expect(err.JSON()).To(Equal(entry.expErrorJSON))
				Expect(err.Error()).To(Equal(entry.expError))
			})
		})
	}
})
