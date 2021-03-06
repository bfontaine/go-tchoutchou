package matchers_test

import (
	. "github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/onsi/gomega"
	. "github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/onsi/gomega/matchers"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("BeTrue", func() {
	It("should handle true and false correctly", func() {
		Ω(true).Should(BeTrue())
		Ω(false).ShouldNot(BeTrue())
	})

	It("should only support booleans", func() {
		success, err := (&BeTrueMatcher{}).Match("foo")
		Ω(success).Should(BeFalse())
		Ω(err).Should(HaveOccurred())
	})
})
