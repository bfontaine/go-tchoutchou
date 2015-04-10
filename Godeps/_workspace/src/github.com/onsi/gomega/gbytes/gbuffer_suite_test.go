package gbytes_test

import (
	. "github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"testing"
)

func TestGbytes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gbytes Suite")
}
