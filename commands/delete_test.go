package commands

import (
	"github.com/mortedecai/gbb/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mortedecai/gbb/gbberror"
)

var _ = Describe("Delete", func() {
	It("should return gbberror.NotYetImplemented until implementation is complete", func() {
		_, err := Delete(nil)
		Expect(err).To(MatchError(gbberror.ErrNotYetImplemented))
	})
})

var _ = Describe("DeleteOption", func() {
	Describe("ToDelete", func() {
		It("should create a single item list if the filename isn't empty", func() {
			do := &deleteOption{
				rootOption: &rootOption{
					host:      "localhost",
					port:      9990,
					authToken: "abc",
				},
				toDelete: "deleteMe.ns",
			}
			Expect(do.ToDelete()).To(Equal([]models.GBBFileName{"deleteMe.ns"}))
		})
		It("should create an empty list if the filename is empty", func() {
			do := &deleteOption{
				rootOption: &rootOption{
					host:      "localhost",
					port:      9990,
					authToken: "abc",
				},
			}
			Expect(do.ToDelete()).To(Equal([]models.GBBFileName{}))

		})
	})
	Describe("Validation", func() {
		entries := []struct {
			context     string
			outcome     string
			do          *deleteOption
			expValidity bool
		}{
			{
				context:     "zero option",
				outcome:     "should be invalid",
				do:          &deleteOption{},
				expValidity: false,
			},
			{
				context: "zero value root option",
				outcome: "should be invalid",
				do: &deleteOption{
					rootOption: &rootOption{},
					toDelete:   "deleteMe.ns",
				},
				expValidity: false,
			},
			{
				context: "fully qualified option",
				outcome: "should be valid",
				do: &deleteOption{
					rootOption: &rootOption{
						host:      "localhost",
						port:      9990,
						authToken: "abc",
					},
					toDelete: "deleteMe.ns",
				},
				expValidity: true,
			},
		}
		for _, e := range entries {
			entry := e
			Context(entry.context, func() {
				It(entry.outcome, func() {
					Expect(entry.do.Valid()).To(Equal(entry.expValidity))
				})
			})
		}
	})

})
