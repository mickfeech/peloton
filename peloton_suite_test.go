package peloton_test

import (
	"github.com/mickfeech/peloton"
	"gopkg.in/h2non/gock.v1"
	"net/http"
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
	cookie := &http.Cookie{
		Name:     "peloton_session_id",
		Value:    "1b2448ae7cf64f48b68ee34d3d031115",
		Path:     "/",
		Domain:   "onepeloton.com",
		MaxAge:   36000,
		HttpOnly: true,
		Secure:   true,
	}
	client.Cookie = cookie
})

var _ = BeforeEach(func() {
	defer gock.Off()
})
