package limiter

import (
	"sync"
	"time"
)

type BucketLimiter struct {
	defaultCapacity uint32
	period          int

	userBuckets sync.Map
}

type tokenBucket struct {
	mu       sync.Mutex
	capacity uint32

	balance uint32
}

func newBucket(period int, capacity uint32) *BucketLimiter {
	return &BucketLimiter{
		defaultCapacity: capacity,
		period:          period,
		userBuckets:     sync.Map{},
	}
}

func (l *BucketLimiter) TakeToken(id string) error {

	value, _ := l.userBuckets.LoadOrStore(id, &tokenBucket{
		mu:       sync.Mutex{},
		capacity: l.defaultCapacity,
		balance:  l.defaultCapacity,
	})

	bucket, ok := value.(*tokenBucket)
	if !ok {
		return ErrNotFound
	}

	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	if bucket.balance == 0 {
		return ErrLimitExceeded
	}
	bucket.balance--
	return nil
}

func (l *BucketLimiter) StartRefillTokens() {
	go func() {

		ticker := time.NewTicker(time.Duration(l.period) * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			go func() {
				l.userBuckets.Range(func(key, value any) bool {
					if bucket, ok := value.(*tokenBucket); ok {
						bucket.balance += bucket.capacity - bucket.balance
					}
					return true
				})
			}()

		}

	}()
}
