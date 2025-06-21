package advanced // rate_limiting.go // 17-mavzu

import (
  "fmt"
  "sync"
  "time"
)

type RateLimiter struct {
  tokens     chan struct{}
  refillTime time.Duration
}

type RateLimiter2 struct {
  mu        sync.Mutex
  count     int
  limit     int
  window    time.Duration
  resetTime time.Time
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {

  rl := &RateLimiter{
    tokens:     make(chan struct{}, rateLimit),
    refillTime: refillTime,
  }

  for range rateLimit {
    rl.tokens <- struct{}{}
  }
  go rl.startRefill()
  return rl
}

func NewRateLimiter2(limit int, window time.Duration) *RateLimiter2 {
  return &RateLimiter2{
    limit:  limit,
    window: window,
  }
}

func (rl *RateLimiter) startRefill() {
  ticker := time.NewTicker(rl.refillTime)
  defer ticker.Stop()

  for {
    select {
    case <-ticker.C:
      select {
      case rl.tokens <- struct{}{}:
      default:
      }
    }
  }
}

func (rl *RateLimiter) allow() bool {
  select {
  case <-rl.tokens:
    return true
  default:
    return false
  }
}

func (rl *RateLimiter2) allow2() bool {
  rl.mu.Lock()
  defer rl.mu.Unlock()

  now := time.Now()
  if now.After(rl.resetTime) {
    rl.resetTime = now.Add(rl.window)
    rl.count = 0
  }

  if rl.count < rl.limit {
    rl.count++
    return true
  }
  return false
}

func WorkWithRateLimiting() {
  // WorkWithRateLimiting1()
  WorkWithRateLimiting2()
}

func WorkWithRateLimiting1() {

  rl := NewRateLimiter(5, time.Second)

  for range 10 {
    if rl.allow() {
      fmt.Println("Request allowed!")
    } else {
      fmt.Println("request denied!")
    }
    time.Sleep(200 * time.Millisecond)
  }
}
func WorkWithRateLimiting2() {
  rl := NewRateLimiter2(5, 2*time.Second)

  for range 10 {
    if rl.allow2() {
      fmt.Println("Request allowed!")
    } else {
      fmt.Println("request denied!")
    }
    time.Sleep(200 * time.Millisecond)
  }
}
