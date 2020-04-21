package peloton_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
)

var _ = Describe("Workouts Model", func() {
	BeforeEach(func() {
		testData, err := ioutil.ReadFile("testdata/workoutdetail.json")
		if err != nil {
			fmt.Printf("Error: %v", err)
			Fail("Error retrieving json test data")
		}
		gock.New("https://api.onepeloton.com").
			Get("/api/workout/f813e2143150469c9aed108a40ca61ef/performance_graph").
			Reply(200).
			JSON(testData)
		gock.InterceptClient(client.Client.GetClient())
	})
	It("Should not return an error", func() {
		_, err := client.GetWorkoutDetails("f813e2143150469c9aed108a40ca61ef", "5")
		Expect(err).To(BeNil())
	})
	It("Should not return an error", func() {
		m, _ := client.GetWorkoutDetails("f813e2143150469c9aed108a40ca61ef", "5")
		Expect(m.Duration).To(Equal(900))
	})
})
