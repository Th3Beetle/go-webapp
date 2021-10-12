package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    
    main "go-webapp/webapp"
	"time"
	"io/ioutil"
	"net/http"
	"fmt"

)

func TestWebapp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Webapp Suite")
}

var _ = BeforeSuite(func() {
	port:="8000"
	go main.StartSrv(port)
	time.Sleep(1000 * time.Millisecond)
})

var _ = Describe("Webapp", func() {
	It("serves apiHandler string", func() {
        port:="8000"

		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/api", port))
		if err != nil {
			panic("Error request: " + err.Error())
		}
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
        	panic("error body")
        }

		Expect(string(body)).To(Equal("<h1>apiHandler</h1>"))
	})

})

var _ = Describe("Webapp", func() {
	It("serves loginHandler string", func() {
        port:="8000"

		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/login", port))
		if err != nil {
			panic("Error request: " + err.Error())
		}
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
        	panic("error body")
        }

		Expect(string(body)).To(Equal("<h1>loginHandler</h1>"))
	})

})

var _ = Describe("Webapp", func() {
	It("serves keyHandler string", func() {
        port:="8000"

		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/key", port))
		if err != nil {
			panic("Error request: " + err.Error())
		}
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
        	panic("error body")
        }

		Expect(string(body)).To(Equal("<h1>keyHandler</h1>"))
	})

})

var _ = Describe("Webapp", func() {
	It("serves statisticsHandler string", func() {
        port:="8000"

		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/statistics", port))
		if err != nil {
			panic("Error request: " + err.Error())
		}
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
        	panic("error body")
        }

		Expect(string(body)).To(Equal("<h1>statisticsHandler</h1>"))
	})

})

var _ = Describe("Webapp", func() {
	It("serves indexHandler string", func() {
        port:="8000"

		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s", port))
		if err != nil {
			panic("Error request: " + err.Error())
		}
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
        	panic("error body")
        }

		Expect(string(body)).To(Equal("<h1>indexHandler</h1>"))
	})

})
