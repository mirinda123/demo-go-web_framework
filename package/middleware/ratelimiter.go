package middleware

import (
	"io"

	"github.com/mirinda123/mirinda-goweb/package/mirinda"
)

type RateLimitReader struct {

	//结构体嵌套，自动获得其字段
	BodyLimitConfig

	Reader io.ReadCloser

	//计算读取了多少
	count int64
}

type RateLimiterConfig struct {
	BeforeFunc BeforeFunc
	// IdentifierExtractor uses echo.Context to extract the identifier for a visitor
	IdentifierExtractor Extractor
	// Store 是rate limiter的具体策略
	Store RateLimiterStore
	// ErrorHandler provides a handler to be called when IdentifierExtractor returns an error
	ErrorHandler func(context mirinda.Context, err error) error
	// DenyHandler provides a handler to be called when RateLimiter denies access
	DenyHandler func(context mirinda.Context, identifier string, err error) error
}

// Extractor is used to extract data from echo.Context
type Extractor func(context mirinda.Context) (string, error)
