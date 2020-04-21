package models_test

import (
	"encoding/json"
	"fmt"
	. "github.com/mickfeech/peloton/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("Workout Model", func() {
	var myWorkouts Workouts
	BeforeEach(func() {
		testData, err := ioutil.ReadFile("../testdata/workouts.json")
		if err != nil {
			fmt.Printf("Error: %v", err)
			Fail("Error retrieving json test data")
		}
		_ = json.Unmarshal([]byte(testData), &myWorkouts)
	})
	It("Has a count attribute", func() {
		Expect(myWorkouts.Count).To(Equal(20))
	})
	It("Has a data attribute", func() {
		Expect(myWorkouts.Data).NotTo(Equal(nil))
	})
	It("Has a FitnessDiscipline attribute", func() {
		Expect(myWorkouts.Data[0].FitnessDiscipline).To(Equal("stretching"))
	})
	It("Has a Created attribute", func() {
		Expect(myWorkouts.Data[0].Created).To(Equal(1586615278))
	})
	It("Has a EndTime attribute", func() {
		Expect(myWorkouts.Data[0].EndTime).To(Equal(1586615579))
	})
	It("Has a StartTime attribute", func() {
		Expect(myWorkouts.Data[0].StartTime).To(Equal(1586615281))
	})
	It("Has a StartTime attribute", func() {
		Expect(myWorkouts.Data[0].UserID).To(Equal("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	})
	It("Has a ID attribute", func() {
		Expect(myWorkouts.Data[0].ID).To(Equal("f813e2143150469c9aed108a40ca61ef"))
	})
	It("Has a status attribute", func() {
		Expect(myWorkouts.Data[0].Status).To(Equal("COMPLETE"))
	})
	It("Has a PersonalRecord attribute", func() {
		Expect(myWorkouts.Data[0].PersonalRecord).To(Equal(false))
	})
})
