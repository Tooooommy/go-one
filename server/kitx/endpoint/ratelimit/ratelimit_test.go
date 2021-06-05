package ratelimit

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func TestRatelimit(t *testing.T) {
	limiter := rate.NewLimiter(5, 1)
	var i = 1
	for {
		if limiter.Allow() {
			i++
			fmt.Println("allow ", i, time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
