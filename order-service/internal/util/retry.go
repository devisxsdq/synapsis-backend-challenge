package util

import (
	"context"
	"time"

	bo "github.com/cenkalti/backoff/v4"
)

func WithRetry(ctx context.Context, fn func() error) error {
	b := bo.NewExponentialBackOff()
	b.InitialInterval = 100 * time.Millisecond
	b.MaxElapsedTime = 3 * time.Second
	return bo.Retry(func() error { return fn() }, bo.WithContext(b, ctx))
}
