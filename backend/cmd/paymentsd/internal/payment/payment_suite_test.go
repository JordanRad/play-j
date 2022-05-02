package payments_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPayment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payment Suite")
}
