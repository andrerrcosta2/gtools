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

// Supplier2 represents a function that returns two values of type T and K.
type Supplier2[T any, K any] func() (T, K)

// Predicate represents a function that takes a value of type T and returns a boolean.
type Predicate[T any] func(T) bool

// BiPredicate represents a function that takes two values of type T and U and returns a boolean.
type BiPredicate[T any, U any] func(T, U) bool

// BiPredicate2 represents a function that takes two values of type T and U and returns two booleans.
type BiPredicate2[T any, U any] func(T, U) (bool, bool)

// BiFunction represents a function that takes two values of type T and U and returns a value of type R.
type BiFunction[T any, U any, R any] func(T, U) R

// BiFunction2 represents a function that takes two values of type T and U and returns two values of type R and R2.
type BiFunction2[T any, U any, R any, R2 any] func(T, U) (R, R2)

// BiConsumer represents a function that takes two values of type T and U and returns nothing.
type BiConsumer[T any, U any] func(T, U)

// TriConsumer represents a function that takes three values of type T, U, and V and returns nothing.
type TriConsumer[T any, U any, V any] func(T, U, V)
