package peloton_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
)

var _ = Describe(".Me()", func() {
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
	It("Returns a first name", func() {
		m, _ := client.Me()
		Expect(m.FirstName).To(Equal("John"))
	})
	It("Returns a last name", func() {
		m, _ := client.Me()
		Expect(m.LastName).To(Equal("Doe"))
	})
	It("Returns a username", func() {
		m, _ := client.Me()
		Expect(m.Username).To(Equal("jdoe"))
	})
})
