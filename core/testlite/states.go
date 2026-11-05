// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testlite

import (
	"fmt"
	"strings"
	"sync"

	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type CounterMap interface {
	Add(delta int, keys ...string)
	Get(key ...string) int
	Len() int
	Reset()
	AssertCounts(expected int, keys ...string) bool
	AssertCountTo(key string, expected int, msgAndArgs ...any) bool
	AssertRegisteredCounts(expected int, msgAndArgs ...any) bool
	AssertRegisteredCountsBy(function.BiPredicate[string, any], int, ...any) bool
}

type FlagMap interface {
	Set(value bool, keys ...string)
	Get(key string) bool
	Len() int
	Reset()
	AssertFlags(expected bool, keys ...string) bool
	AssertFlagTo(key string, expected bool, msgAndArgs ...any) bool
	AssertRegisteredFlags(expected int, msgAndArgs ...any) bool
	AssertRegisteredFlagsBy(function.BiPredicate[string, bool], int, ...any) bool
}

func Map(t HelperTesting) CounterMap {
	return &counterMap{
		mu: sync.RWMutex{},
		m:  make(map[string]int),
		t:  t,
	}
}

type counterMap struct {
	mu sync.RWMutex
	m  map[string]int
	t  HelperTesting
}

func (c *counterMap) Add(delta int, keys ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, key := range keys {
		c.m[key] += delta
	}
}

func (c *counterMap) Get(key ...string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	sum := 0
	for _, k := range key {
		sum += c.m[k]
	}
	return sum
}

func (c *counterMap) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m = make(map[string]int)
}

func (c *counterMap) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.m)
}

func (c *counterMap) AssertCounts(expected int, keys ...string) bool {
	c.t.Helper()

	c.mu.RLock()
	defer c.mu.RUnlock()

	var failed []string

	for _, key := range keys {
		got := c.m[key]
		if got != expected {
			failed = append(failed,
				fmt.Sprintf("%q: expected %d, got %d", key, expected, got),
			)
		}
	}

	if len(failed) > 0 {
		fail(c.t, fmt.Sprintf("call assertions failed:\n\t%s",
			strings.Join(failed, "\n\t")))
		return false
	}
	return true
}

func (c *counterMap) AssertCountTo(key string, expected int, msgAndArgs ...any) bool {
	c.t.Helper()
	c.mu.RLock()
	defer c.mu.RUnlock()
	if got := c.m[key]; got != expected {
		fail(c.t, fmt.Sprintf("fail on AssertCountTo %q: expected %d, got %d", key, expected, got), msgAndArgs...)
		return false
	}
	return true
}

func (c *counterMap) AssertRegisteredCounts(expected int, msgAndArgs ...any) bool {
	c.t.Helper()
	c.mu.RLock()
	defer c.mu.RUnlock()
	if len(c.m) != expected {
		fail(c.t, fmt.Sprintf("fail on AssertRegisteredCounts: expected %d, got %d", expected, len(c.m)),
			msgAndArgs...)
		return false
	}
	return true
}

func (c *counterMap) AssertRegisteredCountsBy(f function.BiPredicate[string, any], expectedCount int, msgAndArgs ...any) bool {
	c.t.Helper()
	c.mu.RLock()
	defer c.mu.RUnlock()

	count := 0
	for k, v := range c.m {
		if f(k, v) {
			count++
		}
	}

	if count != expectedCount {
		fail(c.t, fmt.Sprintf("fail on AssertRegisteredCountsBy: expected %d, got %d", expectedCount, count),
			msgAndArgs...)
		return false
	}
	return true
}

func Flag(t HelperTesting) FlagMap {
	return &flagMap{
		mu: sync.RWMutex{},
		m:  make(map[string]bool),
		t:  t,
	}
}

type flagMap struct {
	mu sync.RWMutex
	m  map[string]bool
	t  HelperTesting
}

func (c *flagMap) Set(value bool, keys ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, key := range keys {
		c.m[key] = value
	}
}

func (c *flagMap) Get(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.m[key]
}

func (c *flagMap) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.m)
}

func (c *flagMap) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m = make(map[string]bool)
}

func (c *flagMap) AssertFlags(expected bool, keys ...string) bool {
	c.t.Helper()

	c.mu.RLock()
	defer c.mu.RUnlock()

	var failed []string

	for _, key := range keys {
		got := c.m[key]
		if got != expected {
			failed = append(failed,
				fmt.Sprintf("%q: expected %d, got %d", key, expected, got),
			)
		}
	}

	if len(failed) > 0 {
		c.t.Fatalf(
			"call assertions failed:\n\t%s",
			strings.Join(failed, "\n\t"),
		)
	}
	return true
}

func (c *flagMap) AssertFlagTo(key string, expected bool, msgAndArgs ...any) bool {
	c.t.Helper()
	c.mu.RLock()
	defer c.mu.RUnlock()
	if got := c.m[key]; got != expected {
		fail(c.t, fmt.Sprintf("fail on AssertFlagTo %q: expected '%t', got '%t'", key, expected, got), msgAndArgs...)
	}
	return true
}

func (c *flagMap) AssertRegisteredFlags(expected int, msgAndArgs ...any) bool {
	c.t.Helper()
	c.mu.RLock()
	defer c.mu.RUnlock()
	if len(c.m) != expected {
		fail(c.t, fmt.Sprintf("fail on AssertRegisteredFlags: expected %d, got %d", expected, len(c.m)), msgAndArgs...)
	}
	return true
}

func (c *flagMap) AssertRegisteredFlagsBy(f function.BiPredicate[string, bool], expectedCount int, msgAndArgs ...any) bool {
	c.t.Helper()
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Count the number of flags satisfying the f
	count := 0
	for key, value := range c.m {
		if f(key, value) {
			count++
		}
	}

	// Assert the count matches the expected value
	if count != expectedCount {
		fail(c.t, fmt.Sprintf("fail on AssertRegisteredFlagsBy: expected %d, got %d", expectedCount, count),
			msgAndArgs...)
		return false
	}
	return true
}

type ConstantsMap interface {
	Register(key string, value any)
	Get(key string) any
	Has(key string) bool
	AssertConst(key string, expected any, compare function.Equals[any], msgAndArgs ...any) bool
	Reset()
}

func Constants(t HelperTesting) ConstantsMap {
	return &consts{
		t: t,
		m: sync.Map{},
	}
}

type consts struct {
	t HelperTesting
	m sync.Map
}

func (c *consts) Register(key string, value any) {
	c.t.Helper()

	if _, loaded := c.m.LoadOrStore(key, value); loaded {
		c.t.Fatalf("constant %q is already registered", key)
	}
}

func (c *consts) Get(key string) any {
	c.t.Helper()
	v, ok := c.m.Load(key)
	if !ok {
		c.t.Fatalf("constant %q not found", key)
	}
	return v
}

func (c *consts) Has(key string) bool {
	c.t.Helper()
	_, ok := c.m.Load(key)
	return ok
}

func (c *consts) AssertConst(key string, expected any, compare function.Equals[any], msgAndArgs ...any) bool {
	c.t.Helper()
	v, ok := c.m.Load(key)
	if !ok {
		c.t.Fatalf("constant %q not found", key)
	}
	if !compare(v, expected) {
		fail(c.t, fmt.Sprintf("fail on AssertConst %q: expected %v, got %v", key, expected, v), msgAndArgs...)
		return false
	}
	return true
}

func (c *consts) Reset() {
	c.t.Helper()
	c.m.Clear()
}

func fail(t HelperTesting, defaultMsg string, msgAndArgs ...any) {
	t.Helper()
	// If msgAndArgs contains a format string (first argument), pass it to Fatalf along with additional args
	if len(msgAndArgs) > 0 {
		// Ensure the first argument is a string format
		if format, ok := msgAndArgs[0].(string); ok {
			t.Fatalf(format, msgAndArgs[1:]...)
		} else {
			// If no format string is found, fall back to the default error message
			t.Fatal(defaultMsg)
		}
	} else {
		// No arguments, use the default message
		t.Fatal(defaultMsg)
	}
}
