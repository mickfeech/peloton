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
	var mySchedule Schedule
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
})
