package package main

import ()

type TokenLimiter struct {
	refillRate float64
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
			}
		}
	}
}

func (t TokenLimiter) Run() {
	while true {
		// refill at the refill rate, need to use a ticker

		// for each request, check if there are enough tokens
		// if so, decrement the tokens and return
		// if not, wait until there are enough tokens (drop the current request)
		

	}
}