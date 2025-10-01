// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterables

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"math/rand"
	"sort"
	"sync"
)

// OfSlice creates a new Slice from a given variadic list of values.
// It takes a variable number of arguments of type N and returns a pointer to a Slice[N].
func OfSlice[G any](values ...G) *Slice[G] {
	// Create a new slice with the given values
	s := Slice[G](values)
	// Return a pointer to the newly created slice
	return &s
}

type Slice[G any] []G

// After returns a new slice with all elements after the given index.
// It takes an integer i as an argument and returns a slice of type Slice[N].
func (s *Slice[G]) After(i int) *Slice[G] {
	*s = (*s)[i:]
	return s
}

// Append appends the given values to the end of the slice.
// It takes a variable number of arguments of type N and appends them to the slice.
// It returns the same slice after the append op has finished.
func (s *Slice[G]) Append(v ...G) *Slice[G] {
	// Append the values to the slice
	*s = append(*s, v...)
	// Return the same slice
	return s
}

func (s *Slice[G]) Async(fn func(i int, v G), maxParallels int) *Slice[G] {
	// Create a Semaphore to limit the number of concurrent ops
	smp := make(chan struct{}, maxParallels)

	// Iterate over the slice and call the function for each element in a goroutine
	for i, v := range *s {
		// Call the function in a goroutine
		go func(i int, v G) {
			// Acquire the Semaphore
			smp <- struct{}{}
			// Call the function with the index and value
			fn(i, v)
			// Release the Semaphore
			<-smp
		}(i, v)
	}

	// Return the same slice
	return s
}

// At returns the element at the given index i.
// If the index is out of bounds, it returns the zero value of type N.
// It takes an integer i as an argument and returns a value of type N.
func (s *Slice[G]) At(i int) G {
	if i < 0 || i >= len(*s) {
		var zeroValue G
		return zeroValue
	}
	return (*s)[i]
}

// Before returns a new slice with all elements before the given index.
// It takes an integer i as an argument and returns a slice of type Slice[N].
func (s *Slice[G]) Before(i int) *Slice[G] {
	*s = (*s)[:i]
	return s
}

// Between returns a new slice with all elements between the given indices.
// It takes two integers i and j as arguments and returns a slice of type Slice[N].
func (s *Slice[G]) Between(i, j int) *Slice[G] {
	*s = (*s)[i:j]
	return s
}

// Duplicate returns two slices, one with the same elements as the original slice,
// and the other with a clone of the original slice.
// It takes no arguments and returns two slices of type Slice[N].
func (s *Slice[G]) Duplicate() (*Slice[G], *Slice[G]) {
	// Create a new slice with the same length as the original
	newSlice := make(Slice[G], s.Len())

	// Copy the elements from the original slice to the new slice
	copy(newSlice, *s)

	// Return the original slice and the new slice
	return &newSlice, s
}

// Each calls the given function for each element in the slice, passing the value to the function.
// It takes a Consumer function that takes a value of type N and returns nothing.
// It iterates over the slice and calls the function for each element.
// It returns the same slice after all ops have finished.
func (s *Slice[G]) Each(fn functions.Consumer[G]) *Slice[G] {
	// Iterate over the slice and call the function for each element
	for _, v := range *s {
		// Call the function with the current value
		fn(v)
	}
	// Return the same slice
	return s
}

// EachN calls the given function for each element in the slice, passing the index and value to the function.
//
// It takes a BiConsumer function that takes two parameters: the index of the element and the element value.
// It iterates over the slice and calls the function for each element.
// It returns the same slice after all ops have finished.
func (s *Slice[G]) EachN(fn functions.BiConsumer[int, G]) *Slice[G] {
	// Iterate over the slice and call the function for each element
	for i, v := range *s {
		fn(i, v)
	}
	return s
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
// It takes a predicate function `fn` as an argument. The predicate function takes a value of type `N` and returns a boolean.
// The function returns a new slice with all elements from the original slice that satisfy the predicate function.
// The element order in the new slice is the same as in the original slice.
func (s *Slice[G]) Filter(fn functions.Predicate[G]) *Slice[G] {
	// Create a new slice with enough capacity to store all elements from the original slice that satisfy the predicate
	newSlice := make([]G, 0, s.Len())
	// Iterate over the original slice
	for _, v := range *s {
		// Check if the current element satisfies the predicate
		if fn(v) {
			// If it does, append it to the new slice
			newSlice = append(newSlice, v)
		}
	}
	// Return the new slice
	return &Slice[G]{}
}

// FilterN creates a new slice with all elements that pass the test implemented by the provided function.
// It takes a BiPredicate function `fn` as an argument. The BiPredicate function takes two parameters: the index of the element and the element value.
// It returns a new slice with all elements from the original slice that satisfy the predicate function.
// The element order in the new slice is the same as in the original slice.
func (s *Slice[G]) FilterN(fn functions.BiPredicate[int, G]) *Slice[G] {
	// Create a new slice with enough capacity to store all elements from the original slice that satisfy the predicate
	newSlice := make(Slice[G], 0, len(*s))
	// Iterate over the original slice
	for i, v := range *s {
		// Check if the current element satisfies the predicate
		if fn(i, v) {
			// If it does, append it to the new slice
			newSlice = append(newSlice, v)
		}
	}
	// Return the new slice
	return &newSlice
}

// First returns the first element of the slice and whether the value exists.
func (s *Slice[G]) First() (first G, ok bool) {
	if s.IsEmpty() {
		return
	}
	return (*s)[0], true
}

// IsEmpty returns true if the slice is empty, false otherwise.
func (s *Slice[G]) IsEmpty() bool {
	return len(*s) == 0
}

// Last returns the last element of the slice and whether the value exists.
func (s *Slice[G]) Last() (last G, ok bool) {
	if s.IsEmpty() {
		return
	}
	return (*s)[len(*s)-1], true
}

// Len returns the length of the slice.
func (s *Slice[G]) Len() int {
	return len(*s)
}

// Map applies the given function to each element in the slice and returns a new slice with the results.
// It takes a single argument of type Function[N, N] and returns a pointer to a Slice[N].
// The function is called for each element in the slice, and the result is appended to the new slice.
// The order of the elements in the new slice is the same as the order of the elements in the original slice.
func (s *Slice[G]) Map(fn functions.Function[G, G]) *Slice[G] {
	// Create a new slice with enough capacity to store all elements from the original slice that satisfy the predicate
	newSlice := make(Slice[G], 0, len(*s))
	// Iterate over the original slice
	for _, v := range *s {
		// Check if the current element satisfies the predicate
		newSlice = append(newSlice, fn(v))
	}
	// Return the new slice
	return &newSlice
}

// Operation calls the given function for each element in the slice, passing the index and a pointer to the slice
// itself to the function.
// It returns the same slice after all ops have finished.
func (s *Slice[G]) Operation(fn functions.BiConsumer[int, *Slice[G]]) *Slice[G] {
	// Iterate over the slice and call the function for each element
	for i := range *s {
		// Call the function with the current index and a pointer to the slice
		fn(i, s)
	}
	// Return the same slice
	return s
}

// Parallel calls the given function for each element in the slice, passing the index and value to the function,
// concurrently. It takes a BiConsumer function that takes two parameters: the index of the element and the element value.
// It uses a WaitGroup to wait for all goroutines to finish.
// It takes a Semaphore to limit the number of concurrent ops.
// The function returns the same slice after all concurrent ops have finished.
func (s *Slice[G]) Parallel(fn functions.BiConsumer[int, G], maxParallels int) *Slice[G] {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(s.Len())

	// Create a Semaphore to limit the number of concurrent ops
	smp := make(chan struct{}, maxParallels)

	// Iterate over the slice and call the function for each element in a goroutine
	for i, v := range *s {
		// Call the function in a goroutine
		go func(i int, v G) {
			defer wg.Done()
			// Acquire the Semaphore
			smp <- struct{}{}
			// Call the function with the index and value
			fn(i, v)
			// Release the Semaphore
			<-smp
		}(i, v)
	}

	// SetWaitingPoint for all goroutines to finish
	wg.Wait()

	// Return the same slice
	return s
}

// Rand returns a random element from the slice.
// It takes no arguments and returns a value of type N.
func (s *Slice[G]) Rand() G {
	// Generate a random index between 0 and the length of the slice
	randIndex := rand.Intn(s.Len())
	// Return the element at the generated index
	return (*s)[randIndex]
}

// RemoveAt returns a new slice with the element at the given index removed.
// It takes an integer i as an argument and returns a slice of type Slice[N].
func (s *Slice[G]) RemoveAt(i int) *Slice[G] {
	a := (*s)[:i]
	*s = append(a, (*s)[i+1:]...)
	return s
}

// Remove returns a new slice with the first element that satisfies the given predicate removed.
// It takes a Predicate function as an argument and returns a slice of type Slice[N].
func (s *Slice[G]) Remove(v G, compare functions.BiFunction[G, G, int]) (*Slice[G], bool) {
	for i := 0; i < s.Len(); i++ {
		if compare((*s)[i], v) == 0 {
			return s.RemoveAt(i), true
		}
	}
	return s, false
}

// Some returns a new slice with n random elements from the original slice.
// It takes an integer n as an argument and returns a slice of type Slice[N].
// If 'n' is greater than the length of the original slice,
// it'll return a slice with the same length as the original slice.
func (s *Slice[G]) Some(n int) *Slice[G] {
	// Check if the requested length is greater than the original slice
	if n > s.Len() {
		// If it is, set the requested length to the same length as the original slice
		n = s.Len()
	}

	// Create a clone of the original slice
	// We use a clone to avoid modifying the original slice
	copySlice := make(Slice[G], len(*s))
	copy(copySlice, *s)

	// This is the Fisher-Yates shuffle algorithm. It is used for generating a random permutation of a finite sequence
	// It works by looping through the slice and swapping each element with a random element from the remaining elements
	for i := len(copySlice) - 1; i > 0; i-- {
		ii := rand.Intn(i + 1)
		// Swap elements
		copySlice[i], copySlice[ii] = copySlice[ii], copySlice[i]
	}

	// Return the first n elements of the shuffled slice
	result := copySlice[:n]
	return &result
}

// Sort sorts the slice using the given less function.
func (s *Slice[G]) Sort(less functions.BiPredicate[G, G]) *Slice[G] {
	sort.SliceStable(*s, func(i, j int) bool {
		return less((*s)[i], (*s)[j])
	})
	return s
}

// ToSet returns a new slice with all unique elements from the original slice.
// It requires a comparators.KeyTyped[G, string] function as an argument.
func (s *Slice[G]) ToSet(cmp comparators.KeyTyped[G, string]) *Slice[G] {
	// Sort the slice using the given less function
	values := make(map[string]bool)
	set := make(Slice[G], 0, s.Len())
	for _, v := range *s {
		hash := cmp.Hash(v)
		if !values[hash] {
			values[hash] = true
			set = append(set, v)
		}
	}
	return &set
}

// Values returns the underlying slice of values.
// It returns a slice of a type []N.
func (s *Slice[G]) Values() []G {
	return *s
}

// OfMap creates a new Map from the given pairs of values.
// It takes a variable number of arguments of type tuple.Pair[N, V] and returns a pointer to a Map[N, V].
// The map is initialized with the given key-value pairs.
func OfMap[K prim.Ordered, V any](values ...str.Entry[K, V]) *Map[K, V] {
	// Initialize the map with the given key-value pairs.
	m := Map[K, V]{}
	for _, entry := range values {
		m[entry.Key()] = entry.Value()
	}
	// Return a pointer to the map.
	return &m
}

type Map[K prim.Ordered, V any] map[K]V

// At returns the value associated with the given key k.
// It takes a single argument of type N and returns a value of type V.
func (m *Map[K, V]) At(k K) V {
	// Return the value associated with the given key
	return (*m)[k]
}

// Contains checks if the map contains the given key k.
// It takes a single argument of type N and returns a boolean indicating if the key is present in the map.
func (m *Map[K, V]) Contains(k K) bool {
	// Check if the map contains the given key
	_, ok := (*m)[k]
	return ok
}

// Each calls the given function for each key-value pair in the map.
// It takes a BiConsumer function that takes two parameters: the key of type N and a pointer to the value of type V.
// It iterates over the map and calls the function for each key-value pair.
// It returns a pointer to the map.
func (m *Map[K, V]) Each(fn functions.BiConsumer[K, V]) *Map[K, V] {
	// Iterate over the map and call the function for each key-value pair
	for k, v := range *m {
		// Pass the value to the function
		fn(k, v)
	}
	// Return a pointer to the map
	return m
}

// IsEmpty checks if the map is empty.
// It returns a boolean indicating if the map is empty.
func (m *Map[K, V]) IsEmpty() bool {
	return len(*m) == 0
}

// Len returns the number of key-value pairs in the map.
// It returns the number of entries in the map.
func (m *Map[K, V]) Len() int {
	// Return the number of entries in the map
	return len(*m)
}

// Operation calls the given function for each key in the map, passing the key and a pointer to the map itself to the function.
// It iterates over the map and calls the function for each key.
// It returns a pointer to the map.
func (m *Map[K, V]) Operation(fn functions.BiConsumer[K, *Map[K, V]]) *Map[K, V] {
	// Iterate over the map and call the function for each key
	for k := range *m {
		// Pass the map itself to the function
		fn(k, m)
	}
	// Return a pointer to the map
	return m
}

// Parallel calls the given function for each key-value pair in the map, passing the key and the value to the function.
// It takes a BiConsumer function that takes two parameters: the key of type N and the value of type V.
// It iterates over the map concurrently and calls the function for each key-value pair.
// It returns a pointer to the map.
func (m *Map[K, V]) Parallel(fn functions.BiConsumer[K, V], maxParallels int) *Map[K, V] {
	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	// ToSet the maximum number of goroutines to run concurrently
	smp := make(chan struct{}, maxParallels)
	// Increase the waiting group
	wg.Add(m.Len())
	// Iterate over the map and call the function for each key-value pair concurrently
	for k, v := range *m {
		// Start a goroutine for each key-value pair
		go func(k K, v V) {
			defer wg.Done()
			// Acquire the Semaphore
			smp <- struct{}{}
			// Call the function with the key and value
			fn(k, v)
			// Release the Semaphore
			<-smp
		}(k, v)
	}
	// SetWaitingPoint for all goroutines to finish
	wg.Wait()
	// Return a pointer to the map
	return m
}

// Put adds a new key-value pair to the map.
// It takes a single argument of type N for the key and a single argument of type V for the value.
// It returns a pointer to the map.
func (m *Map[K, V]) Put(k K, v V) *Map[K, V] {
	// Append the new key-value pair to the map
	(*m)[k] = v
	// Return a pointer to the map
	return m
}

// Remove deletes the key-value pair with the given key k from the map.
// It takes a single argument of type N and returns a pointer to the map.
func (m *Map[K, V]) Remove(k K) *Map[K, V] {
	// Delete the key-value pair with the given key from the map.
	delete(*m, k)
	// Return a pointer to the map.
	return m
}

// Values return a slice of all values in the map.
// It iterates over the map and appends each value to the slice.
// The length of the returned slice is compare to the number of entries in the map.
func (m *Map[K, V]) Values() []V {
	// Create a slice to store the values
	values := make([]V, 0, len(*m))
	// Iterate over the map and append each value to the slice
	for _, v := range *m {
		values = append(values, v)
	}
	// Return the slice of values
	return values
}

// OfSliceMap creates a new empty SliceMap.
// It returns a pointer to an empty SliceMap[N, V].
func OfSliceMap[K comparable, V any](values ...str.Entry[K, []V]) *SliceMap[K, V] {
	// Initialize the map with the given key-value pairs.
	m := SliceMap[K, V]{}
	for _, entry := range values {
		m[entry.Key()] = entry.Value()
	}
	// Return a pointer to the map.
	return &m
}

type SliceMap[K comparable, V any] map[K][]V

// Append add a value to the slice of values associated with the given key in the map.
// If the key isn't present in the map, it creates a new slice with the given value.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) Append(k K, v V) *SliceMap[K, V] {
	// Get the slice of values associated with the given key
	slice, ok := (*m)[k]
	// If the key isn't present in the map, create a new slice with the given value
	if !ok {
		(*m)[k] = []V{v}
	} else {
		// Append the value to the existing slice
		(*m)[k] = append(slice, v)
	}
	// Return a pointer to the map
	return m
}

// At returns the slice of values associated with the given key in the map.
// It returns a slice of a type []V where V is the type of the values in the map.
// If the key isn't present in the map, it returns nil.
func (m *SliceMap[K, V]) At(k K) []V {
	return (*m)[k]
}

// Contains checks if the map contains the given key k.
// It takes a single argument of type N for the key and returns a boolean indicating if the key is present in the map.
func (m *SliceMap[K, V]) Contains(k K) bool {
	// Check if the map contains the given key
	_, ok := (*m)[k]
	return ok
}

// Each calls the given function for each key-value pair in the map.
// It takes a BiConsumer function that takes two parameters: the key of type N and a pointer to the value of a type []V.
// It iterates over the map and calls the function for each key-value pair.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) Each(fn functions.BiConsumer[K, *V]) *SliceMap[K, V] {
	// Iterate over each key-value pair in the map
	for k, v := range *m {
		for i := range v {
			fn(k, &v[i])
		}
	}
	// Return a pointer to the map
	return m
}

// EachSlice calls the given function for each key-value pair in the map.
// It takes a BiConsumer function that takes two parameters: the key of type N and a pointer to the value of a type []V.
// It iterates over the map and calls the function for each key-value pair.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) EachSlice(fn functions.BiConsumer[K, *[]V]) *SliceMap[K, V] {
	// Iterate over each key-value pair in the map
	for k, v := range *m {
		// Call the function with the key and a pointer to the value
		fn(k, &v)
	}
	// Return a pointer to the map
	return m
}

// IsEmpty checks if the map is empty.
// It returns a boolean indicating if the map is empty.
func (m *SliceMap[K, V]) IsEmpty() bool {
	return len(*m) == 0
}

// Len returns the number of key-value pairs in the map.
// It returns an integer representing the number of entries in the map.
func (m *SliceMap[K, V]) Len() int {
	// Return the length of the map
	return len(*m)
}

func (m *SliceMap[K, V]) MapValues(fn functions.BiFunction[K, []V, []V]) *SliceMap[K, V] {
	for k, v := range *m {
		(*m)[k] = fn(k, v)
	}
	return m
}

func (m *SliceMap[K, V]) MapEach(fn functions.BiFunction[K, *V, V]) *SliceMap[K, V] {
	for k, slice := range *m {
		for i := range slice {

			slice[i] = fn(k, &slice[i])
		}
	}
	return m
}

// Operation calls the given function for each key in the map, passing the key and a pointer to the map itself to
// the function.
// It iterates over the map and calls the function for each key.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) Operation(fn functions.BiConsumer[K, *SliceMap[K, V]]) *SliceMap[K, V] {
	// Iterate over the map and call the function for each key
	for k := range *m {
		// Pass a pointer to the map itself to the function
		fn(k, m)
	}
	// Return a pointer to the map
	return m
}

func (m *SliceMap[K, V]) Parallel(fn functions.BiConsumer[K, []V], maxParallels int) *SliceMap[K, V] {
	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	// ToSet the maximum number of goroutines to run concurrently
	smp := make(chan struct{}, maxParallels)
	// Increase the waiting group
	wg.Add(m.Len())
	// Iterate over the map and call the function for each key-value pair concurrently
	for k, v := range *m {
		// Start a goroutine for each key-value pair
		go func(k K, v []V) {
			defer wg.Done()
			// Acquire the Semaphore
			smp <- struct{}{}
			// Call the function with the key and value
			fn(k, v)
			// Release the Semaphore
			<-smp
		}(k, v)
	}
	// SetWaitingPoint for all goroutines to finish
	wg.Wait()
	// Return a pointer to the map
	return m
}

// Put adds a new key-value pair to the map.
// It takes a single argument of type N for the key and a single argument of a type []V for the value.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) Put(k K, v []V) *SliceMap[K, V] {
	// Append the new key-value pair to the map
	(*m)[k] = v
	// Return a pointer to the map
	return m
}

// PutIfAbsent adds a new key-value pair to the map if the key isn't present.
// It takes a single argument of type N for the key and a single argument of a type []V for the value.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) PutIfAbsent(k K, v []V) *SliceMap[K, V] {
	// Check if the key isn't present in the map
	if _, ok := (*m)[k]; !ok {
		// Append the new key-value pair to the map
		// Append the new key-value pair to the map
		(*m)[k] = v
	}
	// Return a pointer to the map
	return m
}

// PutOrAppend adds a new value to the slice of values associated with the given key in the map.
// If the key isn't present in the map, it creates a new slice with the given value.
// It takes a single argument of type N for the key and a single argument of type V for the value.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) PutOrAppend(k K, v V) *SliceMap[K, V] {
	// Check if the key is present in the map
	if m.Contains(k) {
		// If the key is present, append the value to the existing slice
		m.Append(k, v)
	} else {
		// If the key isn't present, add the new key-value pair to the map
		m.Put(k, []V{v})
	}
	// Return a pointer to the map
	return m
}

// Remove removes the key-value pair associated with the given key from the map.
// If the key isn't present in the map, it does nothing.
// It takes a single argument of type N for the key.
// It returns a pointer to the map.
func (m *SliceMap[K, V]) Remove(k K) *SliceMap[K, V] {
	delete(*m, k)
	return m
}

// Values return a slice of all the values in the map.
// It returns a slice of a type []V where V is the type of the values in the map.
func (m *SliceMap[K, V]) Values() []V {
	// Create a slice to store all the values in the map
	result := make([]V, 0, len(*m))

	// Iterate over each value in the map
	for _, v := range *m {
		// Append all the values to the result slice
		result = append(result, v...)
	}

	// Return the slice of all the values in the map
	return result
}
