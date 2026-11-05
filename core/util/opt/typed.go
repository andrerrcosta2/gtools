// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package opt

import "github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"

// Cast returns an Option of type T that contains the value if it is of type T.
// If the value is not of type T, it returns a None Option.
//
// Parameters:
// - value: the value to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value if it is of type T.
// - A pointer to a None Option if the value is not of type T.
func Cast[T any](value any) *Option[T] {
	cast, ok := value.(T)
	if !ok {
		return None[T]()
	}
	return Of(cast)
}

// ZeroString returns an Option of type string that contains the value if it isn't zero.
// If the value is zero, it returns a None Option.
// It is used to represent the absence of a value from a point of view of zero-values.
// That means it considers as nil a zero-value, which might be a false-positive depending
// on the requirements of your application.
//
// Parameters:
// - value: the value to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value if it isn't zero.
// - A pointer to a None Option if the value is zero.
func ZeroString(s string) *Option[string] {
	if s == "" {
		return None[string]()
	}
	return Of(s)
}

// ZeroNumber returns an Option of type nums.Any that contains the value if it isn't zero.
// If the value is zero, it returns a None Option.
// It is used to represent the absence of a value from a point of view of zero-values.
// That means it considers as nil a zero-value, which might be a false-positive depending
// on the requirements of your application.
//
// Parameters:
// - value: the value to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value if it isn't zero.
// - A pointer to a None Option if the value is zero.
func ZeroNumber[T nums.Any](i T) *Option[T] {
	if i == 0 {
		return None[T]()
	}
	return Of(i)
}

// ZeroSlice returns an Option of type S that contains the value if it isn't zero.
// If the value is zero, it returns a None Option.
// It is used to represent the absence of a value from a point of view of zero-values.
// That means it considers as nil a zero-value, which might be a false-positive depending
// on the requirements of your application.
//
// Parameters:
// - value: the value to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value if it isn't zero.
// - A pointer to a None Option if the value is zero.
func ZeroSlice[S ~[]T, T any](s S) *Option[S] {
	if len(s) == 0 {
		return None[S]()
	}
	return Of(s)
}

// ZeroMap returns an Option of type M that contains the value if it isn't zero.
// If the value is zero, it returns a None Option.
// It is used to represent the absence of a value from a point of view of zero-values.
// That means it considers as nil a zero-value, which might be a false-positive depending
// on the requirements of your application.
//
// Parameters:
// - value: the value to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value if it isn't zero.
// - A pointer to a None Option if the value is zero.
func ZeroMap[M ~map[K]V, K comparable, V any](m M) *Option[M] {
	if len(m) == 0 {
		return None[M]()
	}
	return Of(m)
}
