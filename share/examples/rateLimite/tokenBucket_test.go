package rateLimite

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"testing"
	"time"
)

// 使用 go-stress-testing 工具来压测
// go-stress-testing -c 20 -u http://localhost:8080/ping
// 观察压测结果和服务方日志
func TestTokenBucket(t *testing.T) {
	engine := gin.Default()
	// 限速 3rps，允许突发 10 个并发，超过 500ms 就不再等待
	engine.Use(LimiterByTokenBucket(3, 10, 500*time.Millisecond))
	engine.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	_ = engine.Run(":8080")
}

func LimiterByTokenBucket(r rate.Limit, b int, t time.Duration) gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(ctx *gin.Context) {
		key := ctx.ClientIP()
		l, _ := limiters.LoadOrStore(key, rate.NewLimiter(r, b))
		c, cancel := context.WithTimeout(ctx, t)
		defer cancel()
		if err := l.(*rate.Limiter).Wait(c); err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": err})
		}
		ctx.Next()
	}
}
