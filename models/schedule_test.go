package models_test

import (
	"encoding/json"
	"fmt"
	. "github.com/mickfeech/peloton/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("Schedule Model", func() {
	var mySchedule ScheduleResponse
	BeforeEach(func() {
		testData, err := ioutil.ReadFile("../testdata/schedule.json")
		if err != nil {
			fmt.Printf("Error: %v", err)
			Fail("Error retrieving json test data")
		}
		_ = json.Unmarshal([]byte(testData), &mySchedule)
	})
	It("Has a Count attribute", func() {
		Expect(mySchedule.Count).To(Equal(4))
	})
	It("Has a Total attribute", func() {
		Expect(mySchedule.Total).To(Equal(4))
	})
	It("Has a data attribute", func() {
		Expect(mySchedule.Data).NotTo(Equal(nil))
	})
	It("Has an Instructors attribute", func() {
		Expect(mySchedule.Instructors).NotTo(Equal(nil))
	})
	It("Has an RideTypes attribute", func() {
		Expect(mySchedule.RideTypes).NotTo(Equal(nil))
	})
	It("Has an ID attribute", func() {
		Expect(mySchedule.Data[0].ID).To(Equal("d7fca24717014685aa9461d8b64e147a"))
	})
	It("Has a RideID attribute", func() {
		Expect(mySchedule.Data[0].RideID).To(Equal("6228e8bfda184914b38138cd8ea463d1"))
	})
	It("Has an StartTime attribute", func() {
		Expect(mySchedule.Data[0].StartTime).To(Equal(1588690800))
	})
	It("Has an EndTime attribute", func() {
		Expect(mySchedule.Data[0].EndTime).To(Equal(1588692772))
	})
	It("Has an Status attribute", func() {
		Expect(mySchedule.Data[0].Status).To(Equal("PRE_START"))
	})
	It("Has an IsLive attribute", func() {
		Expect(mySchedule.Data[0].IsLive).To(Equal(true))
	})
	It("Has an IsStudio attribute", func() {
		Expect(mySchedule.Data[0].IsStudio).To(Equal(false))
	})
	It("Has an IsEncore attribute", func() {
		Expect(mySchedule.Data[0].IsEncore).To(Equal(true))
	})
	It("Has an TotalHomeReservations attribute", func() {
		Expect(mySchedule.Data[0].TotalHomeReservations).To(Equal(395))
	})
})
