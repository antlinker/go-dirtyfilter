package store_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDirtyFilterMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Store Suite Test.")
}
