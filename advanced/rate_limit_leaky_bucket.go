package advanced // rate_limiting_leaky_bucket.go

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)

	fmt.Printf("Tokens added: %d, Tokens subtracted: %d, Total tokens: %d\n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("Last leak time: %v\n", lb.lastLeak)

	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false
}

func WorkWithRateLimitingLeakyBucket() {
	WorkWithRateLimitingLeakyBucket1()
}

func WorkWithRateLimitingLeakyBucket1() {
	var wg sync.WaitGroup
	LBI := NewLeakyBucket(5, 500*time.Millisecond)
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if LBI.Allow() {
				fmt.Println("Current time:", time.Now())
				fmt.Println("Request accepted!")
			} else {
				fmt.Println("Current time:", time.Now())
				fmt.Println("XXX------Request rejected!")
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}

	time.Sleep(500 * time.Second)

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if LBI.Allow() {
				fmt.Println("Current time:", time.Now())
				fmt.Println("Request accepted!")
			} else {
				fmt.Println("Current time:", time.Now())
				fmt.Println("XXX------Request rejected!")
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}
