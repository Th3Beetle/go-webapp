package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    
    main "go-webapp/webapp"
	"time"
	"io/ioutil"
	"net/http"

)

func TestWebapp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Webapp Suite")
}

var _ = Describe("Webapp", func() {
	It("serves Hello World", func() {
		
		go main.StartSrv("8000")
        time.Sleep(2)

		resp, err := http.Get("http://127.0.0.1:8000")
		if err != nil {
			panic("Error request: " + err.Error())
		}

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
        	panic("error body")
        }

		Expect(string(body)).To(Equal("<h1>Hello World!</h1>"))
	})

})
