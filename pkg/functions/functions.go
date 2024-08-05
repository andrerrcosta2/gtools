// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package functions

// Runnable represents a function that returns nothing.
type Runnable func()

// Function represents a function that takes a value of type T and returns a value of type R.
type Function[T any, R any] func(T) R

// Function2 represents a function that takes a value of type T and returns two values of type U and V.
type Function2[T any, U any, V any] func(T) (U, V)

// Consumer represents a function that takes a value of type T and returns nothing.
type Consumer[T any] func(T)

// Supplier represents a function that returns a value of type T.
type Supplier[T any] func() T

// BiPredicate represents a function that takes two values of type T and U and returns a boolean.
type BiPredicate[T any, U any] func(T, U) bool

// BiFunction represents a function that takes two values of type T and U and returns a value of type R.
type BiFunction[T any, U any, R any] func(T, U) R

// BiFunction2 represents a function that takes two values of type T and U and returns two values of type R and R2.
type BiFunction2[T any, U any, R any, R2 any] func(T, U) (R, R2)

// BiConsumer represents a function that takes two values of type T and U and returns nothing.
type BiConsumer[T any, U any] func(T, U)

// Identity returns the input value unchanged.
// This function serves as an identity function for any type.
func Identity[T any](x T) T {
	return x
}

// Constant returns a function that always returns the provided value.
// This is useful for creating a constant function that can be used as a callback or in other scenarios.
//
// Parameters:
// value - The value to be returned by the function.
//
// Returns:
// a function that always returns the provided value.
func Constant[T any](value T) func() T {
	// This inner function is the actual implementation of the constant function.
	// It returns the provided value.
	return func() T {
		return value
	}
}

// Flip reverses the order of arguments for a binary function.
// Given a function f that takes arguments of types A and B and returns a value of type C,
// Flip returns a new function that takes arguments of types B and A and returns a value of type C.
func Flip[A, B, C any](f func(A, B) C) func(B, A) C {
	return func(b B, a A) C {
		return f(a, b)
	}
}

// Compose is a higher-order function that takes two functions, `f` and `g`, and returns a new function.
// The new function applies `g` to its input, and then applies `f` to the result.
// It is useful for composing multiple functions together.
//
// The type parameters `A`, `B`, and `C` represent the types of the input and output values of the functions.
//
// Parameters:
// - f: a function that takes a value of type `B` and returns a value of type `C`.
// - g: a function that takes a value of type `A` and returns a value of type `B`.
//
// Returns:
// - a new function that takes a value of type `A` and returns a value of type `C`.
func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	// The returned function applies `g` to its input, and then applies `f` to the result.
	return func(a A) C {
		return f(g(a))
	}
}

// Curry is a higher-order function that takes a function `f` that takes two arguments of types `A` and `B` and returns a value of type `C`.
// It returns a new function that takes only one argument of type `A` and returns a new function that takes one argument of type `B` and returns a value of type `C`.
// This is known as currying, where a function with multiple arguments is converted into a sequence of functions each with a single argument.
// This allows for partial application of arguments, where only some of the arguments are provided at the time of function call.
// The curried function can then be called with the remaining arguments at a later time.
func Curry[A, B, C any](f func(A, B) C) func(A) func(B) C {
	// Return a function that takes an argument of type `A`
	return func(a A) func(B) C {
		// Return a function that takes an argument of type `B`
		return func(b B) C {
			// Call the original function `f` with the provided arguments `a` and `b`
			return f(a, b)
		}
	}
}

// AndThen chains two functions together. It takes two functions `f` and `g` as arguments.
// Function `f` takes an argument of type `A` and returns a value of type `B` along with an error.
// Function `g` takes an argument of type `B` and returns a value of type `A` along with an error.
// AndThen returns a new function that takes an argument of type `A` and returns a value of type `A` along with an error.
func AndThen[A, B any](f func(A) (B, error), g func(B) (A, error)) func(A) (A, error) {
	return func(a A) (A, error) {
		// Call function `f` with the input argument `a`
		b, err := f(a)

		// If there was an error from function `f`, return the input argument `a` and the error
		if err != nil {
			return a, err
		}

		// Call function `g` with the result from function `f`
		return g(b)
	}
}

// OrElse chains two functions together. It takes two functions `f` and `g` as arguments.
// Function `f` takes an argument of type `A` and returns a value of type `B` along with an error.
// Function `g` takes an argument of type `A` and returns a value of type `B` along with an error.
// OrElse returns a new function that takes an argument of type `A` and returns a value of type `B` along with an error.
func OrElse[A, B any](f func(A) (B, error), g func(A) (B, error)) func(A) (B, error) {
	return func(a A) (B, error) {
		// Call function `f` with the input argument `a`
		b, err := f(a)

		// If there was an error from function `f`, call function `g` with the input argument `a`
		if err != nil {
			return g(a)
		}

		// Return the result from function `f`
		return b, nil
	}
}

// Tap applies the given function `f` to the input value `x` and then returns `x`.
// It is useful for performing side effects or debugging purposes.
//
// Parameters:
//
//	x: The input value of type `T`.
//	f: The function of type `func(T)` that will be applied to `x`.
//
// Returns:
//
//	The input value `x` of type `T`.
func Tap[T any](x T, f func(T)) T {
	// Apply the function `f` to the input value `x`
	f(x)

	// Return the input value `x`
	return x
}

// Tap2 applies the given function `f` to the input values `x` and `y` and then returns `x` and `y`.
// It is useful for performing side effects or debugging purposes.
//
// Parameters:
//
//	x: The first input value of type `T`.
//	y: The second input value of type `U`.
//	f: The function of type `func(T, U)` that will be applied to `x` and `y`.
//
// Returns:
//
//	The first input value `x` of type `T` and the second input value `y` of type `U`.
func Tap2[T, U any](x T, y U, f func(T, U)) (T, U) {
	// Apply the function `f` to the input values `x` and `y`
	f(x, y)

	// Return the input values `x` and `y`
	return x, y
}

// Apply applies the function f to the input values a and b and returns the result.
//
// Parameters:
// - f: The function to be applied. It takes two parameters of type A and B and returns a value of type B.
// - a: The first input value of type A.
// - b: The second input value of type B.
//
// Returns:
// - The result of applying the function f to the input values a and b.
func Apply[A, B any](f func(A, B) B, a A, b B) B {
	return f(a, b)
}

// Apply3 applies the function f to the input values a, b, and c and returns the result.
//
// Parameters:
// - f: The function to be applied. It takes three parameters of type A, B, and C and returns a value of type C.
// - a: The first input value of type A.
// - b: The second input value of type B.
// - c: The third input value of type C.
//
// Returns:
// - The result of applying the function f to the input values a, b, and c.
func Apply3[A, B, C any](f func(A, B, C) C, a A, b B, c C) C {
	return f(a, b, c)
}

// Apply4 applies the function f to the input values a, b, c, and d and returns the result.
//
// Parameters:
// - f: The function to be applied. It takes four parameters of type A, B, C, and D and returns a value of type D.
// - a: The first input value of type A.
// - b: The second input value of type B.
// - c: The third input value of type C.
// - d: The fourth input value of type D.
//
// Returns:
// - The result of applying the function f to the input values a, b, c, and d.
func Apply4[A, B, C, D any](f func(A, B, C, D) D, a A, b B, c C, d D) D {
	return f(a, b, c, d)
}

// Memoize is a higher-order function that takes a function f and returns a new function that caches the results of f.
// The new function checks if the result for a given input x is already cached, and if so, returns the cached result instead of calling f again.
// This can be useful for functions that have expensive computations or I/O operations.
//
// Parameters:
// - f: The function to be memoized. It takes a parameter of type T and returns a value of type R.
//
// Returns:
//   - A new function that takes a parameter of type T and returns a value of type R.
//     The new function checks if the result for the input x is already cached, and if so, returns the cached result.
//     If the result is not cached, it calls f with the input x, stores the result in the cache, and returns it.
func Memoize[T comparable, R any](f func(T) R) func(T) R {
	cache := make(map[T]R) // Create a map to store the cached results

	return func(x T) R { // Return a new function that takes an input x and returns the result of f(x)
		if result, ok := cache[x]; ok { // Check if the result for x is already cached
			return result // If it is, return the cached result
		}

		result := f(x) // If the result is not cached, call f with the input x and store the result in the cache
		cache[x] = result

		return result // Return the result of f(x)
	}
}

// Once returns a new function that, when called, will call the provided function `f` only once.
// The result of the first call to `f` is stored and returned on subsequent calls.
//
// Parameters:
// - f: The function to be called only once. It takes no parameters and returns a value of type `T`.
//
// Returns:
//   - A new function that, when called, will call the provided function `f` only once.
//     The result of the first call to `f` is stored and returned on subsequent calls.
func Once[T any](f Supplier[T]) Supplier[T] {
	// Initialize a flag to track if the function has been called
	var once bool

	// Initialize a variable to store the result of the first call to `f`
	var result T

	// Return a new function that will call `f` only once
	return func() T {
		// If the function has not been called yet, call it and store the result
		if !once {
			result = f()
			once = true
		}

		// Return the stored result
		return result
	}
}

// Partial is a higher-order function that takes a function `f` and a value `a`, and returns a new function that takes a value `b`.
// When the new function is called with `b`, it calls `f` with `a` and `b` and returns the result.
// This is useful for partial application of functions.
//
// Parameters:
// - f: The function to be partially applied. It takes two parameters of types `A` and `B`, and returns a value of type `C`.
// - a: The value of type `A` to be partially applied.
//
// Returns:
// - A new function that takes a value of type `B` and returns a value of type `C`.
func Partial[A, B, C any](f func(A, B) C, a A) func(B) C {
	// Return a new function that calls `f` with `a` and the provided `b` value.
	return func(b B) C {
		return f(a, b)
	}
}
