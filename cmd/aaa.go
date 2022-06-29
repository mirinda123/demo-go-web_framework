package main

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

func main() {
	c := Context
	h = applyMiddleware()
	h(c)
}
