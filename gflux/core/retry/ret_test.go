package retry

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

func defaultConfig[T any](dl time.Duration, att uint, bf float64) Config[T] {
	return Config[T]{
		Dl:  dl,
		Att: att,
		Bf:  bf,
		Ctx: context.Background(),
	}
}

func TestOf_SuccessOnFirstAttempt(t *testing.T) {
	config := defaultConfig[any](time.Millisecond, 3, 0)

	result, err := Of(config, func() (int, error) {
		return 42, nil
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("expected result 42, got %v", result)
	}
}

func TestOf_SuccessOnSecondAttempt(t *testing.T) {
	config := defaultConfig[any](time.Millisecond, 3, 0)
	attempts := 0

	result, err := Of(config, func() (int, error) {
		attempts++
		if attempts == 2 {
			return 42, nil
		}
		return 0, fmt.Errorf("error on attempt %d", attempts)
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("expected result 42, got %v", result)
	}
	if attempts != 2 {
		t.Errorf("expected 2 attempts, got %d", attempts)
	}
}

func TestOf_AllAttemptsFail(t *testing.T) {
	config := defaultConfig[any](time.Millisecond, 3, 0)
	attempts := 0

	_, err := Of(config, func() (int, error) {
		attempts++
		return 0, fmt.Errorf("error on attempt %d", attempts)
	})

	if err == nil {
		t.Error("expected error, got nil")
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
	if !strings.Contains(err.Error(), "all 3 attempts failed") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestOff_SuccessOnFirstAttempt(t *testing.T) {
	config := defaultConfig[int](time.Millisecond, 3, 0)
	config.Prm = []int{1, 2, 3}

	result, err := Each(config, func(param int) (int, error) {
		if param == 1 {
			return 42, nil
		}
		return 0, fmt.Errorf("failed")
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("expected result 42, got %v", result)
	}
}

func TestOff_SuccessOnSecondParam(t *testing.T) {
	config := defaultConfig[int](time.Millisecond, 3, 0)
	config.Prm = []int{1, 2, 3}

	result, err := Each(config, func(param int) (int, error) {
		if param == 2 {
			return 42, nil
		}
		return 0, fmt.Errorf("failed with param %d", param)
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("expected result 42, got %v", result)
	}
}

func TestOff_AllParamsFail(t *testing.T) {
	config := defaultConfig[int](time.Millisecond, 3, 0)
	config.Prm = []int{1, 2, 3}

	_, err := Each(config, func(param int) (int, error) {
		return 0, fmt.Errorf("failed with param %d", param)
	})

	if err == nil {
		t.Error("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "all 3 attempts failed") {
		t.Errorf("unexpected error message: %v", err)
	}
}
