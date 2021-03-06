package peloton_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"net/http"
	"reflect"
)

var _ = Describe("ApiClient", func() {
	Describe(".NewApiClient", func() {
		It("Will not have a cookie set with a blank username and password", func() {
			Expect(client.Client.Cookies).To(BeEmpty())
		})
		It("Will not return nil", func() {
			Expect(client.Client).NotTo(BeNil())
		})
		It("Allows you to re-set username", func() {
			client.Username = "TestUser"
			Expect(client.Username).To(Equal("TestUser"))
			Expect(client.Username).NotTo(Equal(""))
		})
		It("Allows you to re-set password", func() {
			client.Password = "TestPass"
			Expect(client.Password).To(Equal("TestPass"))
			Expect(client.Password).NotTo(Equal(""))
		})
	})
	Describe("UpdateAuthCookie", func() {
		It("Returns a cookie", func() {
			cookie, _ := client.UpdateAuthCookie()
			var httpCookie *http.Cookie
			Expect(reflect.TypeOf(cookie)).To(Equal(reflect.TypeOf(httpCookie)))
		})
	})
	Describe(".Me()", func() {
		BeforeEach(func() {
			testData, err := ioutil.ReadFile("testdata/me.json")
			if err != nil {
				fmt.Printf("Error: %v", err)
				Fail("Error retrieving json test data")
			}
			gock.New("https://api.onepeloton.com").
				Get("/api/me").
				Reply(200).
				JSON(testData)
			gock.InterceptClient(client.Client.GetClient())
		})
		It("Should not return an error", func() {
			_, err := client.Me()
			Expect(err).To(BeNil())
		})
		It("Returns a first name", func() {
			m, _ := client.Me()
			Expect(m.FirstName).To(Equal("John"))
		})
		It("Returns a last name", func() {
			m, _ := client.Me()
			Expect(m.LastName).To(Equal("Doe"))
		})
		It("Returns a username", func() {
			m, _ := client.Me()
			Expect(m.Username).To(Equal("jdoe"))
		})
	})
})
