package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

//Example : rate control

type RateLimiter struct {
	mutex        sync.Mutex
	requests     map[string]int64
	requestLimit int
	duration     time.Duration
}

func NewRateLimiter(requestLimit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		requests:     make(map[string]int64),
		requestLimit: requestLimit,
		duration:     duration,
	}
}

func (rl *RateLimiter) CheckLimit(ip string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now().Unix()
	lastTime, exists := rl.requests[ip]

	if !exists || now-lastTime >= int64(rl.duration.Seconds()) {
		rl.requests[ip] = now
		return true
	}

	if rl.requests[ip] > 0 && now-lastTime < int64(rl.duration.Seconds()) {
		rl.requests[ip]++
		return rl.requests[ip] <= int64(rl.requestLimit)
	}

	return false
}

func RateLimiterMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !rl.CheckLimit(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}

//TODO This rate controller can add a blacklist through an interface
//TODO Change access rate
//TODO Interface View Blacklist

func main() {
	// 创建限流器，限制每秒钟每个 IP 的请求次数为 100
	rateLimiter := NewRateLimiter(100, time.Second)

	// 创建 Gin 引擎
	router := gin.Default()

	// 使用限流中间件
	router.Use(RateLimiterMiddleware(rateLimiter))

	// 其他路由和处理程序
	// ...
	router.Run(":8080")
}
