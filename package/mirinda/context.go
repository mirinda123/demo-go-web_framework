package mirinda

import (
	"fmt"
	"net/http"
)

//包装了对writer和rea的操作
type Context struct {
	httpWriter http.ResponseWriter
	httpReq    *http.Request
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		httpWriter: w,
		httpReq:    req,
	}
}
func (c *Context) HttpString(s string) {
	fmt.Fprint(c.httpWriter, s)
}
