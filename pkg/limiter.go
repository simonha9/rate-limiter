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