package test

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

func TestMiddlewareBodyLimit(t *testing.T) {
	m := mirinda.New()

	m.GET("/hello", hello)
	m.ServerStart(":9999")

	// 这里应该改成POST,请求体多发点东西
	resp, _ := http.Get("http://localhost:9999/hello")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	if string(body) != "helloworld" {
		t.Errorf("WRONG")
	}

}
