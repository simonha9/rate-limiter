# rate-limiter
Rate limiter built in Golang

- Server-side rate limiter
- work in a distributed environment that will be emulated through go routines
- separate service as middleware instead of in the application code
- implemented algorithms:
    token bucket
    leaking bucket
    fixed window counter
    sliding window log
    sliding window counter