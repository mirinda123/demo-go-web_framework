package main

import (
	"github.com/mirinda123/mirinda-goweb/package/mirinda"
)

// Handler
// 这个自己定义的handler需要被注册到一个map里面

func hello(c *mirinda.Context) {
	c.HttpString("helloworld")
}
func main() {
	m := mirinda.New()

	m.GET("/hello", hello)
	m.ServerStart(":9999")

}
