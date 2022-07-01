package main

import (
	"io/ioutil"

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

	m.PUT("./hello", hello)

	m.POST("./hello", hello)
	m.ServerStart(":9999")
	ioutil.ReadAll

}

func RateLimiterWithConfig(config RateLimiterConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
		}
	}
}
