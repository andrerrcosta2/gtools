package retry

import (
	"context"
	"time"
)

type Config[T any] struct {
	Ctx context.Context
	Att uint
	Tmo time.Duration
	Dl  time.Duration
	Bf  float64
	Prm []T
	Srt func(T, T) bool
}
