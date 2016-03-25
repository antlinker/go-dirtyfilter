package filter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDirtyFilterMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filter Suite Test.")
}
