package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"github.com/simonha9/rate-limiter/pkg"
)

// need to implement gin handlerfunc interface for it to be a middleware
func (t LeakyBucketLimiter) TokenLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if (c.Request.URL.Path == "/ping") {

			if (t.buckets[0].NumTokens > 0) {
				t.buckets[0].NumTokens--
			} else if (t.buckets[1].NumTokens > 0) {
				t.buckets[1].NumTokens--
			} else {
				// raise a rate limit error
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
					"message": "Rate limit exceeded",
					"retry-after": "1",
				})
				return
			}
		}
		c.Next()
	}
}

type LeakyBucketLimiter struct {
	refillRate int
	buckets []pkg.Bucket
	numBuckets int
}

func NewLeakyBucketLimiter() LeakyBucketLimiter {
	return LeakyBucketLimiter{
		refillRate: 1,
		buckets: []pkg.Bucket{
			pkg.Bucket{
				Capacity: 10,
				NumTokens: 0,
			},
			pkg.Bucket{
				Capacity: 10,
				NumTokens: 0,
			},
		},
	}
}

func (t LeakyBucketLimiter) refill() {
	for i := 0; i < t.numBuckets; i++ {
		t.buckets[i].NumTokens = t.buckets[i].Capacity
	}
}

func (t LeakyBucketLimiter) Run() {
	for {
		t.refill()
		time.Sleep(time.Duration(t.refillRate) * time.Second)
	}
}