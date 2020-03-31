package peloton

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
)

var _ = Describe("ApiClient", func() {
	var client *ApiClient

	BeforeSuite(func() {
		client = NewApiClient("", "")
	})
	BeforeEach(func() {
		defer gock.Off()
	})

	Describe("Me", func() {
		BeforeEach(func() {
			testData, err := ioutil.ReadFile("testdata/me.json")
			if err != nil {
				fmt.Printf("Error: %v", err)
				Fail("Error retrieving json test data")
			}
			gock.New("https://api.onepeloton.com").
				Get("/api/me").
				Reply(200).
				JSON(testData)
		})

		It("Returns a first name", func() {
			gock.InterceptClient(client.Client.GetClient())
			m, _ := client.Me()
			Expect(m.FirstName).To(Equal("John"))
		})
	})
})
