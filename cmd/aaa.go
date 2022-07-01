package main

import "io/ioutil"

func handler(c *Context) {
	process(c)
}

func middlewareFunc(h HandlerFunc) HandlerFunc {
	return func(c echo.Context) err {
		enhance_before()
		h(c)
		enhance_after()
	}
}

func applyMiddleware(h HandlerFunc, middleware ...MiddlewareFunc) HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}

	//聚合成一个HandlerFunc
	return h
}

type ABC interface {
	foo(int)
}

type X struct {
	abc ABC
	cde int
}

type Y struct {
	a int
}

func (Y) foo(i int) {
	i++
}

func fee(abc ABC) {
	abc.foo(1)
}
func main2() {
	y := Y{
		a: 5,
	}
	x := X{
		abc: y,
	}
	fee(x.abc)
	ioutil.ReadAll()
}
