package Legacy_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLegacy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Legacy Suite")
}
