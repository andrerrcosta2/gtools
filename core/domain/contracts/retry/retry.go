// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package retry

import "time"

type Should[T any] func(T, error) bool

// Backoff controls retry pacing.
type Backoff func(retry int) time.Duration

func Always[T any](T, error) bool { return true }

func NoBackoff(_ int) time.Duration { return 0 }
