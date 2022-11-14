package rateLimite

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"net/http"
	"sync"
	"testing"
)

// 使用 go-stress-testing 工具来压测
// go-stress-testing -c 20 -u http://localhost:8080/ping
// -c 指定并发数为 20，-u 指定压测地址
func TestLeakyBucket(t *testing.T) {
	engine := gin.Default()
	engine.Use(LimiterByLeakyBucket(5))
	engine.GET("ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	_ = engine.Run(":8080")
}

func LimiterByLeakyBucket(rps int) gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(ctx *gin.Context) {
		key := ctx.ClientIP() // key 除了 IP，还可以用 username 等
		l, _ := limiters.LoadOrStore(key, ratelimit.New(rps))
		now := l.(ratelimit.Limiter).Take()
		fmt.Printf("now: %s\n", now)
		ctx.Next()
	}
}
