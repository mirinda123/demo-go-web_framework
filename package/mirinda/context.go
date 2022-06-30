package mirinda

import (
	"fmt"
	"net/http"
)

//包装了对writer和rea的操作
type Context struct {
	HttpWriter http.ResponseWriter
	HttpReq    *http.Request
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		HttpWriter: w,
		HttpReq:    req,
	}
}
func (c *Context) HttpString(s string) {
	fmt.Fprint(c.HttpWriter, s)
}
