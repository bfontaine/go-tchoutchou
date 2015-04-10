package format_test

import (
	. "github.com/bfontaine/go-tchoutchou/Godeps/_workspace/src/github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"testing"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}
