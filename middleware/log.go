package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lancer/constant"
	"log/slog"
	"time"
)

func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成唯一请求ID
		requestId := uuid.New().String()

		// 将requestId添加到上下文
		c.Set("RequestId", requestId)

		// 记录请求开始时间
		startTime := time.Now()
		log := logger.With(constant.RequestId, requestId)

		// 处理请求
		c.Next()

		// 记录请求结束时间
		endTime := time.Now()

		// 计算请求延迟
		seconds := endTime.Sub(startTime).Seconds()

		// 记录请求日志
		log.InfoContext(c, "http request", slog.String("method", c.Request.Method), slog.String("path", c.Request.URL.Path), slog.String("query", c.Request.URL.RawQuery), slog.String("ip", c.ClientIP()), slog.String("user-agent", c.Request.UserAgent()), slog.Int("status", c.Writer.Status()), slog.Float64("seconds", seconds))
		log.Info("http request", slog.String("method", c.Request.Method), slog.String("path", c.Request.URL.Path), slog.String("query", c.Request.URL.RawQuery), slog.String("ip", c.RemoteIP()), slog.String("user-agent", c.Request.UserAgent()), slog.Int("status", c.Writer.Status()), slog.Float64("seconds", seconds))

	}
}
