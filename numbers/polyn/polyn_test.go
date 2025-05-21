// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package polyn

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

type input struct {
	N            int
	Coefficients []float64
}

func TestN(t *testing.T) {
	check := func(in input) bool {
		var expected float64
		for i, cfc := range in.Coefficients {
			expected += cfc * float64(Pow(in.N, i))
		}
		actual := N(in.N, in.Coefficients...)
		return actual == expected
	}

	cfg := &quick.Config{
		Values: func(args []reflect.Value, r *rand.Rand) {
			n := r.Intn(11)
			coeffs := make([]float64, r.Intn(5)+1)
			for i := range coeffs {
				coeffs[i] = r.Float64()*200 - 100
			}
			args[0] = reflect.ValueOf(input{N: n, Coefficients: coeffs})
		},
	}

	if err := quick.Check(check, cfg); err != nil {
		t.Errorf("TestN failed: %v", err)
	}
}

// TestLin Quick Check for Lin function
func TestLin(t *testing.T) {
	check := func(a, b float64, n int) bool {
		expected := a + b*float64(n)
		actual := Lin(a, b, n)
		return actual == expected
	}

	cfg := &quick.Config{
		Values: func(args []reflect.Value, r *rand.Rand) {
			a := r.Float64()*100 - 50
			b := r.Float64()*100 - 50
			n := r.Intn(20) // [-20, 20]
			args[0] = reflect.ValueOf(a)
			args[1] = reflect.ValueOf(b)
			args[2] = reflect.ValueOf(n)
		},
	}

	if err := quick.Check(check, cfg); err != nil {
		t.Errorf("TestLin failed: %v", err)
	}
}

// Quick Check for Quad function
func TestQuad(t *testing.T) {
	check := func(a, b, c float64, n int) bool {
		expected := a + b*float64(n) + c*float64(n*n)
		actual := Quad(a, b, c, n)
		return actual == expected
	}

	cfg := &quick.Config{
		Values: func(args []reflect.Value, r *rand.Rand) {
			a := r.Float64()*200 - 100
			b := r.Float64()*200 - 100
			c := r.Float64()*200 - 100
			n := r.Intn(21) - 10 // [-10, 10]
			args[0] = reflect.ValueOf(a)
			args[1] = reflect.ValueOf(b)
			args[2] = reflect.ValueOf(c)
			args[3] = reflect.ValueOf(n)
		},
	}

	if err := quick.Check(check, cfg); err != nil {
		t.Errorf("TestQuad failed: %v", err)
	}
}

// Quick Check for Cub function
func TestCub(t *testing.T) {
	check := func(a, b, c, d float64, n int) bool {
		expected := a + b*float64(n) + c*float64(n*n) + d*float64(n*n*n)
		actual := Cub(a, b, c, d, n)
		return actual == expected
	}

	cfg := &quick.Config{
		Values: func(args []reflect.Value, r *rand.Rand) {
			a := r.Float64()*200 - 100
			b := r.Float64()*200 - 100
			c := r.Float64()*200 - 100
			d := r.Float64()*200 - 100
			n := r.Intn(21) - 10 // [-10, 10]
			args[0] = reflect.ValueOf(a)
			args[1] = reflect.ValueOf(b)
			args[2] = reflect.ValueOf(c)
			args[3] = reflect.ValueOf(d)
			args[4] = reflect.ValueOf(n)
		},
	}

	if err := quick.Check(check, cfg); err != nil {
		t.Errorf("TestCub failed: %v", err)
	}
}

// Quick Check for Pow function
func TestPow(t *testing.T) {
	check := func(base, exp int) bool {
		// Avoiding negative or huge exponents
		if exp < 0 || exp > 10 {
			return true // skip invalid cases
		}

		expected := 1
		for i := 0; i < exp; i++ {
			expected *= base
		}

		actual := Pow(base, exp)
		return actual == expected
	}

	cfg := &quick.Config{
		Values: func(args []reflect.Value, r *rand.Rand) {
			base := r.Intn(21) - 10 // [-10, 10]
			exp := r.Intn(11)       // [0, 10]
			args[0] = reflect.ValueOf(base)
			args[1] = reflect.ValueOf(exp)
		},
	}

	if err := quick.Check(check, cfg); err != nil {
		t.Errorf("TestPow failed: %v", err)
	}
}
