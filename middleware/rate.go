package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

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
