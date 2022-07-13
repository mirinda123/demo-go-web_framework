package mirinda

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context) error

type routersValue struct {
	f      HandlerFunc
	method string
}
type Mirinda struct {
	Server          *http.Server
	routers         map[string]*routersValue
	middlewareSlice []HandlerFunc
}

// 创建一个Mirinda 实例
func New() (e *Mirinda) {
	m := &Mirinda{
		Server:          new(http.Server),
		routers:         make(map[string]*routersValue),
		middlewareSlice: make([]HandlerFunc, 0),
	}

	return m
}

func (m *Mirinda) Use(middleware ...HandlerFunc) {
	m.middlewareSlice = append(m.middlewareSlice, middleware...)
}
func (m *Mirinda) ServerStart(port string) error {
	m.Server.Addr = port

	// m实现了ServeHTTP接口，接口中进行一些handle
	return http.ListenAndServe(port, m)
}
func (m *Mirinda) GET(path string, h HandlerFunc) error {
	//注册handler
	v := &routersValue{
		f:      h,
		method: "GET",
	}
	m.routers[path] = v
	return nil
}
func (m *Mirinda) addRouterTable(path string, method string, h HandlerFunc) {
	v := &routersValue{
		f:      h,
		method: method,
	}
	m.routers[path] = v
}
func (m *Mirinda) PUT(path string, h HandlerFunc) {
	//注册handler
	m.addRouterTable(path, "PUT", h)

}

func (m *Mirinda) POST(path string, h HandlerFunc) {
	//注册handler
	m.addRouterTable(path, "POST", h)
}

func (m *Mirinda) DELETE(path string, h HandlerFunc) {
	//注册handler
	m.addRouterTable(path, "DELETE", h)
}

//这个接口很重要
//用户自定义的接口会写入map中，然后ServeHTTP读取map，绑定自定义接口
func (m *Mirinda) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	k, ok := m.routers[req.URL.Path]

	// 方法必须配对上
	if ok && k.method == req.Method {
		if err := k.f(c); err != nil {
			m.HTTPErrorHandler(err, c)
		}
	} else {
		fmt.Fprint(w, "404 NOT FOUND")
	}

}

//返回错误信息给客户端
func (m *Mirinda) HTTPErrorHandler(err error, c *Context) {

	he, ok := err.(*HTTPError)

	//如果类型断言不成功
	if !ok {
		he = &HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	_ = c.responseJSON(he)

}
