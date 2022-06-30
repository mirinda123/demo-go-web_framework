package middleware

import "github.com/mirinda123/mirinda-goweb/package/mirinda"

//一个默认的全局配置类
var (
	// DefaultBodyLimitConfig is the default BodyLimit middleware config.
	DefaultBodyLimitConfig = BodyLimitConfig{
		Limit: "4MB",
	}
)

//配置类
type BodyLimitConfig struct {
	// Maximum allowed size for a request body, it can be specified
	// as `4x` or `4xB`, where x is one of the multiple from K, M, G, T or P.
	Limit string
}

//返回一个中间件函数
func BodyLimit(limit string) MiddlewareFunc {
	config := DefaultBodyLimitConfig
	config.Limit = limit
	return func(handler *mirinda.HandlerFunc) mirinda.HandlerFunc {
		return func(c *mirinda.Context) {
			// content length 是否超标
			if c.HttpReq.ContentLength > config.Limit {
				return echo.ErrStatusRequestEntityTooLarge
			}
			//研究一下源码里的pool是干什么的
			handler(c)
			after(li.limit)
		}
	}
}

type MiddlewareFunc func(h *mirinda.HandlerFunc) mirinda.HandlerFunc

//中间件函数用来包装，接收一个HandlerFunc，包装好后，返回一个经过包装的HandlerFunc
