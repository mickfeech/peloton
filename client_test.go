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
	Describe(".Instructors()", func() {
		BeforeEach(func() {
			testData, err := ioutil.ReadFile("testdata/instructors.json")
			if err != nil {
				fmt.Printf("Error: %v", err)
				Fail("Error retrieving json test data")
			}
			gock.New("https://api.onepeloton.com").
				Get("/api/instructor").
				Reply(200).
				JSON(testData)
			gock.InterceptClient(client.Client.GetClient())
		})

		It("Should not return an error", func() {
			_, err := client.Instructors()
			Expect(err).To(BeNil())
		})
		It("Returns a count of instructors", func() {
			m, _ := client.Instructors()
			Expect(m.Total).To(Equal(34))
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
})
