package mirinda

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type Mirinda struct {
	Server  *http.Server
	routers map[string]HandlerFunc
}

// 创建一个Mirinda 实例
func New() (e *Mirinda) {
	m := &Mirinda{
		Server:  new(http.Server),
		routers: make(map[string]HandlerFunc),
	}

	return m
}
func (m *Mirinda) ServerStart(port string) error {
	m.Server.Addr = port

	// m实现了ServeHTTP接口，接口中进行一些handle
	return http.ListenAndServe(port, m)
}
func (m *Mirinda) GET(path string, h HandlerFunc) error {
	//注册handler
	m.routers[path] = h
	return nil
}

//这个接口很重要
//用户自定义的接口会写入map中，然后ServeHTTP读取map，绑定自定义接口
func (m *Mirinda) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	if handler, ok := m.routers[req.URL.Path]; ok {
		handler(c)
	} else {
		fmt.Fprint(w, "404 NOT FOUND")
	}

}
