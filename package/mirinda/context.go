package mirinda

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Errors
var (
	ErrUnsupportedMediaType        = NewHTTPError(http.StatusUnsupportedMediaType)
	ErrNotFound                    = NewHTTPError(http.StatusNotFound)
	ErrUnauthorized                = NewHTTPError(http.StatusUnauthorized)
	ErrForbidden                   = NewHTTPError(http.StatusForbidden)
	ErrMethodNotAllowed            = NewHTTPError(http.StatusMethodNotAllowed)
	ErrStatusRequestEntityTooLarge = NewHTTPError(http.StatusRequestEntityTooLarge)
	ErrTooManyRequests             = NewHTTPError(http.StatusTooManyRequests)
	ErrBadRequest                  = NewHTTPError(http.StatusBadRequest)
	ErrBadGateway                  = NewHTTPError(http.StatusBadGateway)
	ErrInternalServerError         = NewHTTPError(http.StatusInternalServerError)
	ErrRequestTimeout              = NewHTTPError(http.StatusRequestTimeout)
	ErrServiceUnavailable          = NewHTTPError(http.StatusServiceUnavailable)
	ErrValidatorNotRegistered      = errors.New("validator not registered")
	ErrRendererNotRegistered       = errors.New("renderer not registered")
	ErrInvalidRedirectCode         = errors.New("invalid redirect status code")
	ErrCookieNotFound              = errors.New("cookie not found")
	ErrInvalidCertOrKeyType        = errors.New("invalid cert or key type, must be string or []byte")
	ErrInvalidListenerNetwork      = errors.New("invalid listener network")
)

func NewHTTPError(code int) *HTTPError {
	he := &HTTPError{Code: code, Message: http.StatusText(code)}
	return he
}

// 实现error 接口.
func (he *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
}

// 对error的包装
type HTTPError struct {
	Code    int
	Message interface{}
}

//包装了对writer和rea的操作
type Context struct {
	HttpWriter http.ResponseWriter
	HttpReq    *http.Request
	M          *Mirinda
}

func (m *Mirinda) NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		HttpWriter: w,
		HttpReq:    req,
		M:          m,
	}
}
func (c *Context) HttpString(s string) error {
	_, err := fmt.Fprint(c.HttpWriter, s)
	return err
}

func (c *Context) responseJSON(s interface{}) error {
	enc := json.NewEncoder(c.HttpWriter)
	return enc.Encode(s)
}
