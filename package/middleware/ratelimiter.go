package middleware

import (
	"sync"
	"time"

	"github.com/mirinda123/mirinda-goweb/package/mirinda"
	"golang.org/x/time/rate"
)

// 用于限制一个时间段内从特定 IP 或 id 到服务器的请求数量。
type RateLimiterConfig struct {

	// IdentifierExtractor uses echo.Context to extract the identifier for a visitor
	IdentifierExtractor Extractor
	// Store 是rate limiter的具体策略的接口
	Store RateLimiterStore
}

// RateLimiterMemoryStore 是  RateLimiterStore 接口的实现
type RateLimiterMemoryStore struct {
	visitors map[string]*Visitor
	mutex    sync.Mutex
	rate     rate.Limit //for more info check out Limiter docs - https://pkg.go.dev/golang.org/x/time/rate#Limit.
	//Limit defines the maximum frequency of some events. Limit is represented as number of events per second. A zero Limit allows no events.
	burst       int           // 令牌桶的大小
	expiresIn   time.Duration //超时时间，控制map中visitor的存活
	lastCleanup time.Time     //上次回收垃圾的时间
}

// Visitor signifies a unique user's limiter details
//每个Visitor有一个Limiter
type Visitor struct {
	*rate.Limiter
	lastSeen time.Time
}

// Extractor 用来从 echo.Context中提取身份
type Extractor func(context *mirinda.Context) (string, error)

type RateLimiterStore interface {
	// Stores for the rate limiter have to implement the Allow method
	Allow(identifier string) (bool, error)
}

type RateLimiterStoreConfig struct {
	Rate      rate.Limit    // 使用官方的令牌桶库，Rate of requests allowed to pass as req/s. 参考 https://pkg.go.dev/golang.org/x/time/rate#Limit.
	Burst     int           // Burst additionally allows a number of requests to pass when rate limit is reached
	ExpiresIn time.Duration // ExpiresIn is the duration after that a rate limiter is cleaned up
}

func CreateRateLimiterWithConfig(config RateLimiterConfig) MiddlewareFunc {
	return func(handler mirinda.HandlerFunc) mirinda.HandlerFunc {
		return func(c *mirinda.Context) error {

			//提取身份
			identifier, err := config.IdentifierExtractor(c)
			if err != nil {
				c.M.HTTPErrorHandler(err, c)
				return nil
			}
			//如果拒绝访问
			//Allow 对AllowN(time.Now(),1) 进行封装
			if allow, err := config.Store.Allow(identifier); !allow {
				c.M.HTTPErrorHandler(err, c)
				return nil
			}
			return handler(c)
		}
	}
}

// Allow implements RateLimiterStore.Allow
//Allow 对AllowN(time.Now(),1) 进行封装
func (store *RateLimiterMemoryStore) Allow(identifier string) (bool, error) {
	store.mutex.Lock()
	limiter, exists := store.visitors[identifier]
	if !exists {
		limiter = new(Visitor)
		limiter.Limiter = rate.NewLimiter(store.rate, store.burst)
		store.visitors[identifier] = limiter
	}
	limiter.lastSeen = time.Now()

	//扫描整个表，进行一次回收
	if time.Now().Sub(store.lastCleanup) > store.expiresIn {
		store.cleanupStaleVisitors()
	}
	store.mutex.Unlock()

	//AllowN 方法表示，截止到某一时刻，目前桶中数目是否至少为 n 个，
	//满足则返回 true，同时从桶中消费 n 个 token。反之不消费桶中的Token，返回false。
	//Allow 实际上就是对 AllowN(time.Now(),1) 进行简化的函数
	return limiter.AllowN(time.Now(), 1), nil
}

/*
cleanupStaleVisitors helps manage the size of the visitors map by removing stale records
of users who haven't visited again after the expiry time has elapsed
*/
func (store *RateLimiterMemoryStore) cleanupStaleVisitors() {

	//遍历每一个visitor
	for id, visitor := range store.visitors {
		if time.Now().Sub(visitor.lastSeen) > store.expiresIn {
			//删除
			delete(store.visitors, id)
		}
	}
	store.lastCleanup = time.Now()
}
