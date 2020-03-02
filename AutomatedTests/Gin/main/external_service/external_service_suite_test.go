package external_service_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestExternalService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExternalService Suite")
}
