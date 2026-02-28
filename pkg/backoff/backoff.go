package backoff

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"time"
)

type JitterType int

const (
	FullJitter JitterType = iota

	EqualJiter
)

type ExponentialBackoff struct {
	BaseDelay  time.Duration
	MaxDelay   time.Duration
	Multiplier float64
	Jitter     JitterType
	Rand       *rand.Rand
	Attempts   int
}

type RetryFunc func() error
type RetryableFunc func(err error) bool

func NewBackoff(maxAttempts int) ExponentialBackoff {
	// You can tune based on maxAttempts if you want
	return ExponentialBackoff{
		BaseDelay:  100 * time.Millisecond,
		MaxDelay:   5 * time.Second,
		Multiplier: 2.0,
		Jitter:     FullJitter,
		Attempts:   maxAttempts,
	}
}

func (b ExponentialBackoff) Retry(ctx context.Context, bo ExponentialBackoff, isRetryable RetryableFunc, fn RetryFunc) error {
	var lastErr error
	for attempt := 0; attempt < b.Attempts; attempt++ {
		if err := fn(); err != nil {
			lastErr = err
			if isRetryable != nil && !isRetryable(err) {
				return err
			}
			if attempt == b.Attempts-1 {
				return err
			}
			if err := bo.Sleep(ctx, attempt); err != nil {
				return err
			}
			continue
		}
		if lastErr == nil {
			lastErr = errors.New("retry: exhausted attempts")
		}
	}
	return lastErr
}

func (b ExponentialBackoff) Next(attempt int) time.Duration {

	if attempt < 0 {
		attempt = 0
	}
	if b.Multiplier < 0 {
		b.Multiplier = 2.0
	}
	if b.BaseDelay == 0 {
		b.BaseDelay = 100 * time.Millisecond
	}
	if b.MaxDelay == 0 {
		b.MaxDelay = 10 * time.Millisecond
	}

	pow := math.Pow(b.Multiplier, float64(attempt))
	capDelay := time.Duration(float64(b.BaseDelay) * pow)
	if capDelay > b.MaxDelay {
		capDelay = b.MaxDelay
	}

	r := b.Rand
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().Unix()))
	}
	switch b.Jitter {
	case FullJitter:
		half := capDelay / 2
		return half + time.Duration(r.Int63n(int64(half))+1)
	case EqualJiter:
		fallthrough
	default:
		return time.Duration(r.Int63n(int64(capDelay) + 1))
	}
}

func (b ExponentialBackoff) Sleep(ctx context.Context, attempt int) error {
	d := b.Next(attempt)
	t := time.NewTicker(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
