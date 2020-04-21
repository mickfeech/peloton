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
		testData, err := ioutil.ReadFile("testdata/workouts.json")
		if err != nil {
			fmt.Printf("Error: %v", err)
			Fail("Error retrieving json test data")
		}
		gock.New("https://api.onepeloton.com").
			Get("/api/user/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa/workouts").
			Reply(200).
			JSON(testData)
		gock.InterceptClient(client.Client.GetClient())
	})
	It("Should not return an error", func() {
		_, err := client.UserWorkouts("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		Expect(err).To(BeNil())
	})
	It("Returns a count of workouts", func() {
		m, _ := client.UserWorkouts("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		Expect(m.Count).To(Equal(20))
	})
	It("Returns a data of workouts", func() {
		m, _ := client.UserWorkouts("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		Expect(m.Data).NotTo(Equal(nil))
	})
})
