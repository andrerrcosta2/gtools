// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/arrays"
	"reflect"
	"sort"
	"strings"
	"testing"
)

// I created the tests on the datastr maps package
// I don't know why I did that. But now the tests uses
// data structure domain objects and I won't bring them here now

// Test MapValues
func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	v := MapValues(&m, func(v int) int {
		return v * 2
	})

	// Check that the length of the result slice is correct
	if len(v) != 3 {
		t.Errorf("MapValues() length = %v, want %v", len(v), 3)
	}

	// Create a map to track the expected values
	expected := map[int]bool{2: false, 4: false, 6: false}

	// Check that all expected values are in the result slice
	for _, value := range v {
		if _, ok := expected[value]; !ok {
			t.Errorf("MapValues() = %v, unexpected value found", value)
		}
		expected[value] = true
	}

	// Check that all expected values were found
	for value, found := range expected {
		if !found {
			t.Errorf("MapValues() missing expected value %v", value)
		}
	}
}

// Test FlatValues
func TestFlatValues(t *testing.T) {
	// Test Case 1: Flattening without transformation (identity function)
	map1 := map[string][]int{
		"numbers": {1, 2, 3},
		"more":    {4, 5, 6},
	}
	expected1 := []int{1, 2, 3, 4, 5, 6}
	result1 := FlatValues(&map1, functions.Identity[int])
	sort.Ints(result1)
	sort.Ints(expected1)

	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test Case 1 Failed: expected %v, got %v", expected1, result1)
	}

	// Test Case 2: Flattening with transformation (string length)
	map2 := map[string][]string{
		"fruits":  {"apple", "banana"},
		"colors":  {"red", "blue"},
		"animals": {"cat", "dog"},
	}
	expected2 := []int{5, 6, 3, 4, 3, 3}
	result2 := FlatValues(&map2, func(v string) int { return len(v) })
	sort.Ints(result2)
	sort.Ints(expected2)

	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test Case 2 Failed: expected %v, got %v", expected2, result2)
	}

	// Test Case 3: Flattening with transformation to uppercase
	map3 := map[string][]string{
		"words": {"hello", "world"},
	}
	expected3 := []string{"HELLO", "WORLD"}
	result3 := FlatValues(&map3, func(v string) string { return strings.ToUpper(v) })
	sort.Strings(result3)
	sort.Strings(expected3)

	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test Case 3 Failed: expected %v, got %v", expected3, result3)
	}

	// Test Case 4: Flattening with transformation to get first characters
	map4 := map[string][]string{
		"names": {"Alice", "Bob", "Charlie"},
	}
	expected4 := []string{"A", "B", "C"}
	result4 := FlatValues(&map4, func(v string) string { return string(v[0]) })
	sort.Strings(result4)
	sort.Strings(expected4)

	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Test Case 4 Failed: expected %v, got %v", expected4, result4)
	}
}

func TestFlatValuesSorted(t *testing.T) {
	// Test Case 1: Sort integers in descending order
	map1 := map[string][]int{
		"numbers": {1, 2, 3},
		"more":    {4, 5, 6},
	}
	expected1 := []int{6, 5, 4, 3, 2, 1}
	result1 := FlatValuesSorted(&map1, func(v int) int { return v }, func(i, j int) bool { return i > j })

	if !arrays.Equals[int](result1, expected1) {
		t.Errorf("Test Case 1 Failed: expected %v, got %v", expected1, result1)
	}

	// Test Case 2: Sort strings by length
	map2 := map[string][]string{
		"fruits":  {"apple", "banana"},
		"colors":  {"carmine", "red"},
		"animals": {"lion", "antelope"},
	}
	expected2 := []string{"antelope", "carmine", "banana", "apple", "lion", "red"}
	result2 := FlatValuesSorted(&map2, func(v string) string { return v }, func(i, j string) bool { return len(i) > len(j) })

	if !arrays.Equals(result2, expected2) {
		t.Errorf("Test Case 2 Failed: expected %v, got %v", expected2, result2)
	}

	// Test Case 3: Sort strings alphabetically
	map3 := map[string][]string{
		"words": {"banana", "apple", "cherry"},
	}
	expected3 := []string{"APPLE", "BANANA", "CHERRY"}
	result3 := FlatValuesSorted(&map3, func(v string) string { return strings.ToUpper(v) }, func(i, j string) bool { return i < j })

	if !arrays.Equals(result3, expected3) {
		t.Errorf("Test Case 3 Failed: expected %v, got %v", expected3, result3)
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	v := MapKeys(&m, func(k string) string {
		return k + "1"
	})

	sort.Strings(v)

	expected := []string{"a1", "b1", "c1"}
	for i := range expected {
		if v[i] != expected[i] {
			t.Errorf("MapKeys() = %v, want %v", v, expected)
			break
		}
	}
}

func TestMapWithKeys(t *testing.T) {
	keys := []int{1, 2, 3}
	expected := map[string]int{"one": 1, "two": 2, "three": 3}

	f := func(k int) (string, int) {
		switch k {
		case 1:
			return "one", 1
		case 2:
			return "two", 2
		case 3:
			return "three", 3
		}
		return "", 0
	}

	result := *MapWithKeys[int, string, int](keys, f)

	if len(result) != len(expected) {
		t.Errorf("Expected map length %d, got %d", len(expected), len(result))
	}

	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected value %d for key %s, got %d", v, k, result[k])
		}
	}
}

func TestMapWithValues(t *testing.T) {
	values := []string{"apple", "banana", "strawberry"}
	expected := map[int]string{5: "apple", 6: "banana", 10: "strawberry"}

	f := func(v string) (int, string) {
		return len(v), v
	}

	result := *MapWithValues(values, f)

	if len(result) != len(expected) {
		t.Errorf("Expected map length %d, got %d", len(expected), len(result))
	}

	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected value (%s) for key (%d), but got %s", v, k, result[k])
		}
	}

	values2 := []string{"apple", "banana", "cherry"}
	possible := []map[int]string{
		{5: "apple", 6: "banana"}, // If "banana" overwrites "cherry"
		{5: "apple", 6: "cherry"}, // If "cherry" overwrites "banana"
	}

	result2 := *MapWithValues(values2, f)

	matchLength := false
	matchContent := false

	for _, exp := range possible {
		if len(result2) == len(exp) {
			matchLength = true
			for k, v := range exp {
				if result[k] != v {
					matchContent = false
					continue
				} else {
					matchContent = true
					break
				}
			}
		}
	}

	if !matchLength || !matchContent {
		t.Errorf("Result %v does not match any of the expected outcomes %v", result2, possible)
	}
}

func TestContainsKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		key      string
		expected bool
	}{
		{"a", true},
		{"b", true},
		{"d", false},
	}

	for _, test := range tests {
		result := ContainsKey(&m, test.key)
		if result != test.expected {
			t.Errorf("ContainsKey(%v, %v) = %v; expected %v", m, test.key, result, test.expected)
		}
	}
}

// Test ContainsValue
func TestContainsValue(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{2, true},
		{4, false},
	}

	for _, test := range tests {
		result := ContainsValue(&m, test.value)
		if result != test.expected {
			t.Errorf("ContainsValue(%v, %v) = %v; expected %v", m, test.value, result, test.expected)
		}
	}
}

// Test ContainsAllKeys
func TestContainsAllKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		keys     []string
		expected bool
	}{
		{[]string{"a", "b"}, true},
		{[]string{"a", "b", "c"}, true},
		{[]string{"a", "b", "d"}, false},
	}

	for _, test := range tests {
		result := ContainsAllKeys(&m, test.keys)
		if result != test.expected {
			t.Errorf("ContainsAllKeys(%v, %v) = %v; expected %v", m, test.keys, result, test.expected)
		}
	}
}

// Test ContainsAllValues
func TestContainsAllValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		values   []int
		expected bool
	}{
		{[]int{1, 2}, true},
		{[]int{1, 2, 3}, true},
		{[]int{1, 2, 4}, false},
	}

	for _, test := range tests {
		result := ContainsAllValues(&m, test.values)
		if result != test.expected {
			t.Errorf("ContainsAllValues(%v, %v) = %v; expected %v", m, test.values, result, test.expected)
		}
	}
}

// Test AreSameKeys
func TestAreSameKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		keys     []string
		expected bool
	}{
		{[]string{"a", "b", "c"}, true},
		{[]string{"a", "b"}, false},
		{[]string{"a", "b", "c", "d"}, false},
		{[]string{"c", "b", "a"}, true}, // Order doesn't matter
	}

	for _, test := range tests {
		result := AreSameKeys(&m, test.keys)
		if result != test.expected {
			t.Errorf("AreSameKeys(%v, %v) = %v; expected %v", m, test.keys, result, test.expected)
		}
	}
}

// Test AreSameValues
func TestAreSameValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		values   []int
		expected bool
	}{
		{[]int{1, 2, 3}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 4}, false},
		{[]int{3, 2, 1}, true}, // Order doesn't matter
	}

	for _, test := range tests {
		result := AreSameValues(&m, test.values)
		if result != test.expected {
			t.Errorf("AreSameValues(%v, %v) = %v; expected %v", m, test.values, result, test.expected)
		}
	}
}
