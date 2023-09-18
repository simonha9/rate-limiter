package middleware

import (
	"github.com/gin-gonic/gin"
	_ "time"
	"fmt"
	"net/http"
)

// need to implement gin handlerfunc interface for it to be a middleware
func TokenLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if (c.Request.URL.Path == "/ping") {
			// raise a rate limit error
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate limit exceeded",
				"retry-after": "1",
			})
			return
		}
		c.Next()
	}
}

type TokenLimiter struct {
	refillRate int64
	buckets []Bucket
	numBuckets int
}

type Bucket struct {
	capacity int
	numTokens int
}

func NewTokenLimiter() TokenLimiter {
	return TokenLimiter{
		refillRate: 1.0,
		buckets: []Bucket{
			Bucket{
				capacity: 10,
				numTokens: 0,
			},
			Bucket{
				capacity: 10,
				numTokens: 0,
			},
		},
	}
}

func (t TokenLimiter) refill() {
	for i := 0; i < t.numBuckets; i++ {
		t.buckets[i].numTokens = t.buckets[i].capacity
	}
}

func (t TokenLimiter) Run() {
	fmt.Println("Ticker run")
}