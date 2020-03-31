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
})
