package middleware

import (
	"io"

	"github.com/mirinda123/mirinda-goweb/package/mirinda"
)

//body_limit 限制请求体的大小
//首先判断ContentLength有没有超过
//然后判断实际的内容读取有没有超过，这样做双重保险

//MiddlewareFunc是中间件函数，用来包装，接收一个HandlerFunc，包装好后，返回一个经过包装的HandlerFunc
type MiddlewareFunc func(h mirinda.HandlerFunc) mirinda.HandlerFunc

//包装一个Reader, 代替初始的Reader
//最终的目标是把BodyLimitReader给req.Body
//所以BodyLimitReader要实现io.ReadCloser
type BodyLimitReader struct {

	//结构体嵌套，自动获得其字段
	BodyLimitConfig

	reader io.ReadCloser

	//计算读取了多少
	count int64
}

//配置类
type BodyLimitConfig struct {
	// Maximum allowed size for a request body, it can be specified
	// as `4x` or `4xB`, where x is one of the multiple from K, M, G, T or P.
	Limit int64
}

//返回一个中间件函数
func BodyLimit(limit int64) MiddlewareFunc {
	return func(handler mirinda.HandlerFunc) mirinda.HandlerFunc {
		return func(c *mirinda.Context) (err error) {
			// content length 是否超标
			config := BodyLimitConfig{Limit: limit}
			config.Limit = limit
			if c.HttpReq.ContentLength > config.Limit {
				return mirinda.ErrStatusRequestEntityTooLarge
			}
			//研究一下源码里的pool是干什么的
			err = handler(c)

			return
		}
	}
}

//实现Read接口
func (blr *BodyLimitReader) Read(b []byte) (n int, err error) {
	n, err = blr.reader.Read(b)
	//每次读取的时候count计数
	blr.count += int64(n)

	if blr.count > blr.Limit {
		return n, mirinda.ErrStatusRequestEntityTooLarge
	}
	return
}

func (r *BodyLimitReader) Close() error {
	return r.reader.Close()
}
