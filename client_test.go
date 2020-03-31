package peloton_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
)

var _ = Describe("ApiClient", func() {
	Describe(".NewApiClient", func() {
		It("Will not have a cookie set with a blank username and password", func() {
			Expect(client.Client.Cookies).To(BeEmpty())
		})
		It("Will not return nil", func() {
			Expect(client.Client).NotTo(BeNil())
		})
	})
	Describe(".Me()", func() {
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
			gock.InterceptClient(client.Client.GetClient())
		})
		It("Should not return an error", func() {
			_, err := client.Me()
			Expect(err).To(BeNil())
		})
	})
})
