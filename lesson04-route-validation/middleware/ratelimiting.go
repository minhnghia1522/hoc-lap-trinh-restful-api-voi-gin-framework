package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()

	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip

}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(5, 10) // 5request/s, brust 10

		newClient := &Client{
			limiter,
			time.Now(),
		}

		clients[ip] = newClient
		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

func CleanUpClients() {
	for {
		time.Sleep((time.Millisecond))
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// hey -n 20 -c 1 -H "X-API-Key:5bcd1d01-373f-4ae3-b745-c2c7f11935a2"  http://localhost:8080/api/v1/categories/
func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)

		limiter := getRateLimiter(ip)

		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many request",
			})
			return
		}
		ctx.Next()
	}
}
