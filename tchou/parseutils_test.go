package tchou

import (
	"testing"

	"github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/franela/goblin"
	o "github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/onsi/gomega"
)

func TestParseutils(t *testing.T) {

	g := goblin.Goblin(t)
	o.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("softParseFloat", func() {
		g.It("Should return 0.0 with an empty string", func() {
			o.Expect(softParseFloat("")).To(o.Equal(0.0))
		})

		g.It("Should parse floats in base 10", func() {
			o.Expect(softParseFloat("10.10")).To(o.Equal(10.10))
		})

		g.It("Should parse integers as float", func() {
			o.Expect(softParseFloat("42")).To(o.Equal(42.0))
		})

		g.It("Should parse floats", func() {
			o.Expect(softParseFloat("42.3")).To(o.Equal(42.3))
		})
	})
}
