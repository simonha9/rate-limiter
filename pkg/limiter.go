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

}