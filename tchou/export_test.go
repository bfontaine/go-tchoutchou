package tchou

import (
	"testing"

	"github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/franela/goblin"
	o "github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/onsi/gomega"
)

func TestExports(t *testing.T) {

	g := goblin.Goblin(t)
	o.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Station", func() {
		s := Station{
			Name:    "name",
			Address: "address",
			Lat:     42.0,
			Long:    17.0,
			Ref:     "FOO",
			slug:    "/gare-de-foo-FOO",
		}

		g.Describe(".records()", func() {
			g.It("Should return name/address/lat/long/ref", func() {
				r := s.records()

				o.Expect(len(r)).To(o.Equal(5))

				o.Expect(r).To(o.Equal([]string{
					"name", "address", "42", "17", "FOO",
				}))
			})
		})
	})

	g.Describe("Stations", func() {
		g.Describe(".WriteCSV(io.Writer)", func() {
			// TODO
		})

		g.Describe(".WriteJSON(io.Writer)", func() {
			// TODO
		})
	})
}
