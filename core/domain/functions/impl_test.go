// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package functions

import (
	"errors"
	"fmt"
	"testing"
)

func TestIdentity(t *testing.T) {
	id1 := Identity[int](1)
	if id1 != 1 {
		t.Errorf("Identity() = %v, want %v", id1, 1)
	}

	id2 := Identity[string]("hello")
	if id2 != "hello" {
		t.Errorf("Identity() = %v, want %v", id2, "hello")
	}
}

func TestConstant(t *testing.T) {
	expectedValue := 42
	constantFunc := Constant(expectedValue)

	for i := 0; i < 10; i++ {
		if result := constantFunc(); result != expectedValue {
			t.Errorf("Constant() = %v, want %v", result, expectedValue)
		}
	}

}

func TestFlip(t *testing.T) {
	flipFunc := Flip(func(a, b int) float32 {
		return float32(a / b)
	})
	if flipFunc(3, 9) != 3.0 {
		t.Errorf("Flip() = %v, want %v", flipFunc(1, 2), 3)
	}
	if flipFunc(5, 10) != 2.0 {
		t.Errorf("Flip() = %v, want %v", flipFunc(5, 10), 2.0)
	}
	if flipFunc(2, 10) != 5.0 {
		t.Errorf("Flip() = %v, want %v", flipFunc(2, 10), 5.0)
	}
}

func TestCompose(t *testing.T) {
	double := func(x int) int {
		return x * 2
	}
	toString := func(x int) string {
		return fmt.Sprint(x) // Convert int to string by adding '0'
	}
	composed := Compose(toString, double)

	input := 3
	expected := "6"

	result := composed(input)
	if result != expected {
		t.Errorf("Compose() = %v, want %v", result, expected)
	}
}

func TestCurry(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}
	curriedSum := Curry(sum)
	r1 := curriedSum(1)(2)
	if r1 != 3 {
		t.Errorf("Curry() = %v, want %v", r1, 3)
	}

	r2 := curriedSum(3)
	r3 := r2(4)
	if r3 != 7 {
		t.Errorf("Curry() = %v, want %v", r3, 10)
	}

}

func TestAndThen(t *testing.T) {
	f := func(x int) (int, error) {
		if x > 0 {
			return x * 2, nil
		}
		return 0, errors.New("negative number")
	}
	g := func(y int) (int, error) {
		if y < 10 {
			return y + 1, nil
		}
		return 0, errors.New("value too large")
	}
	chain := AndThen(f, g)

	result, err := chain(3)
	if err != nil || result != 7 {
		t.Errorf("AndThen(3) = (%v, %v), want (7, nil)", result, err)
	}

	result, err = chain(-1)
	if err == nil || result != -1 {
		t.Errorf("AndThen(-1) = (%v, %v), want (-1, error)", result, err)
	}
}

func TestOrElse(t *testing.T) {
	f := func(x int) (int, error) {
		if x > 0 {
			return x * 2, nil
		}
		return 0, errors.New("negative number")
	}
	g := func(x int) (int, error) {
		return x + 10, nil // Fallback function returns x + 10 regardless of the input
	}

	orElse := OrElse(f, g)

	result, err := orElse(3)
	if err != nil || result != 6 {
		t.Errorf("OrElse(3) = (%v, %v), want (6, nil)", result, err)
	}

	result, err = orElse(-1)
	if err != nil || result != 9 {
		t.Errorf("OrElse(-1) = (%v, %v), want (9, nil)", result, err)
	}
}

func TestMemoize(t *testing.T) {
	f := func(x int) int {
		return x * x
	}

	mem := Memoize(f)

	r1 := mem(4)
	e1 := 16
	if r1 != e1 {
		t.Errorf("Memoize(4) = %v, want %v", r1, e1)
	}

	r2 := mem(4)
	if r2 != e1 {
		t.Errorf("Memoize(4) = %v, want %v", r2, e1)
	}

	r3 := mem(5)
	e2 := 25
	if r3 != e2 {
		t.Errorf("Memoize(5) = %v, want %v", r3, e2)
	}
}
