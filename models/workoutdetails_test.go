package models_test

import (
	"encoding/json"
	"fmt"
	. "github.com/mickfeech/peloton/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("WorkoutDetails Model", func() {
	var myWorkouts WorkOutDetails
	BeforeEach(func() {
		testData, err := ioutil.ReadFile("../testdata/workoutdetail.json")
		if err != nil {
			fmt.Printf("Error: %v", err)
			Fail("Error retrieving json test data")
		}
		_ = json.Unmarshal([]byte(testData), &myWorkouts)
	})
	It("Has a duration attribute", func() {
		Expect(myWorkouts.Duration).To(Equal(900))
	})
	It("Has an Average Summary attribute", func() {
		Expect(myWorkouts.AverageSummaries).NotTo(Equal(nil))
	})
	It("Has an Summary attribute", func() {
		Expect(myWorkouts.Summaries).NotTo(Equal(nil))
	})
	It("Has an Metrics attribute", func() {
		Expect(myWorkouts.Metrics).NotTo(Equal(nil))
	})
	It("Has an TargetPerformanceMetrics attribute", func() {
		Expect(myWorkouts.TargetPerformanceMetrics).NotTo(Equal(nil))
	})
	It("Has proper Average Summary Data", func() {
		Expect(myWorkouts.AverageSummaries[0].DisplayName).To(Equal("Avg Output"))
		Expect(myWorkouts.AverageSummaries[0].DisplayUnit).To(Equal("watts"))
		Expect(myWorkouts.AverageSummaries[0].Value).To(Equal(71))
	})
	It("Has proper Summary Data", func() {
		Expect(myWorkouts.Summaries[0].DisplayName).To(Equal("Total Output"))
		Expect(myWorkouts.Summaries[0].DisplayUnit).To(Equal("kj"))
		Expect(myWorkouts.Summaries[0].Value).To(Equal(64))
	})
	It("Has proper Metrics Data", func() {
		Expect(myWorkouts.Metrics[0].DisplayName).To(Equal("Output"))
		Expect(myWorkouts.Metrics[0].DisplayUnit).To(Equal("watts"))
		Expect(myWorkouts.Metrics[0].MaxValue).To(Equal(150))
	})
	It("Has proper TargetPerformanceMetrics Data", func() {
		Expect(myWorkouts.TargetPerformanceMetrics.CadenceTimeInRange).To(Equal(90))
		Expect(myWorkouts.TargetPerformanceMetrics.ResistanceTimeInRange).To(Equal(63))
	})
})
