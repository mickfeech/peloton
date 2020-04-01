package peloton_test

import (
	"github.com/mickfeech/peloton"
	"gopkg.in/h2non/gock.v1"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var client *peloton.ApiClient

func TestPeloton(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Peloton Suite")
}

var _ = BeforeSuite(func() {
	client, _ = peloton.NewApiClient("", "")
})

var _ = BeforeEach(func() {
	defer gock.Off()
})
