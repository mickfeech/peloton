package models_test

import (
	"encoding/json"
	"fmt"
	. "github.com/mickfeech/peloton/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("Me Model", func() {
	var me User

	BeforeEach(func() {
		testData, err := ioutil.ReadFile("../testdata/me.json")
		if err != nil {
			fmt.Printf("Error: %v", err)
			Fail("Error retrieving json test data")
		}
		_ = json.Unmarshal([]byte(testData), &me)
	})

	It("Has a FirstName attribute", func() {
		Expect(me.FirstName).To(Equal("John"))
	})
	It("Has a FirstName attribute", func() {
		Expect(me.LastName).To(Equal("Doe"))
	})
	It("Has a Username attribute", func() {
		Expect(me.Username).To(Equal("jdoe"))
	})
	It("Has a Birthday attribute", func() {
		Expect(me.Birthday).To(Equal(323358625))
	})
	It("Has a Ftp attribute", func() {
		Expect(me.Ftp).To(Equal(153))
	})
	It("Has a Gender attribute", func() {
		Expect(me.Gender).To(Equal("male"))
	})
	It("Has a Location attribute", func() {
		Expect(me.Location).To(Equal("Nowhere USA"))
	})
	It("Has a Weight attribute", func() {
		Expect(me.Weight).To(Equal(192.0))
	})
	It("Has a Height attribute", func() {
		Expect(me.Height).To(Equal(71.0))
	})
	It("Has a TotalWorkOuts attribute", func() {
		Expect(me.TotalWorkOuts).To(Equal(483))
	})
	It("Has a DefaultMaxHr attribute", func() {
		Expect(me.DefaultMaxHr).To(Equal(183))
	})
	It("Has a CustomMaxHr attribute", func() {
		Expect(me.CustomMaxHr).To(Equal(203))
	})
})
