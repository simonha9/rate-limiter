package pkg

import (

)

// Initially rate limiter will be done through token bucket and will do the rest later

// RateLimiter is a struct that holds the rate limit information
type RateLimiter struct {
	// Rate is the number of requests allowed per second
	Rate int
	// Burst is the number of requests allowed in a single burst
	Burst int
	// LastRequest is the time of the last request
	LastRequest int64
}