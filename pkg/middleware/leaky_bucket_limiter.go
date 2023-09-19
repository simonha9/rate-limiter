package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

// need to implement gin handlerfunc interface for it to be a middleware
func (t LeakyBucketLimiter) LeakyBucketLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if queue is full
		if (len(t.queue) == t.capacity) {
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

type LeakyBucketLimiter struct {
	processRate int
	queue chan int
	capacity int
}

func NewLeakyBucketLimiter(capacity int) LeakyBucketLimiter {
	return LeakyBucketLimiter{
		processRate: 1,
		capacity: capacity,
		queue: make(chan int, capacity),
	}
}

func (t LeakyBucketLimiter) process() {
	<- t.queue
}

func (t LeakyBucketLimiter) Run() {
	for {
		t.process()
		time.Sleep(time.Duration(t.processRate) * time.Second)
	}
}