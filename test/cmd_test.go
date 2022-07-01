package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/mirinda123/mirinda-goweb/package/mirinda"
)

func hello(c *mirinda.Context) error {
	return c.HttpString("helloworld")
}
func TestGET(t *testing.T) {
	m := mirinda.New()

	m.GET("/hello", hello)
	m.ServerStart(":9999")

	resp, _ := http.Get("http://localhost:9999/hello")
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "helloworld" {
		t.Errorf("WRONG")
	}

}
