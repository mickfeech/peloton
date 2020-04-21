package peloton_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
)

var _ = Describe("Instructors Model", func() {
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
		_, err := client.GetInstructors()
		Expect(err).To(BeNil())
	})
	It("Returns a count of instructors", func() {
		m, _ := client.GetInstructors()
		Expect(m.Total).To(Equal(34))
	})
})
