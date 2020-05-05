package peloton_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
)

var _ = Describe("Schedules Models", func() {
	BeforeEach(func() {
		testData, err := ioutil.ReadFile("testdata/schedule.json")
		if err != nil {
			fmt.Printf("Error: %v", err)
			Fail("Error retrieving json test data")
		}
		gock.New("https://api.onepeloton.com/api/v3/ride/live").
			Reply(200).
			JSON(testData)
		gock.InterceptClient(client.Client.GetClient())
	})
	It("Should not return an error", func() {
		_, err := client.GetSchedule(1588680000, 1588698000)
		Expect(err).To(BeNil())
	})
	It("Should not return an error", func() {
		m, _ := client.GetSchedule(1588680000, 1588698000)
		Expect(m.Count).To(Equal(4))
	})
})
