// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/andrerrcosta2/gtools/pkg/generics"
	"github.com/andrerrcosta2/gtools/pkg/tuple"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestNewEntry(t *testing.T) {
	e1 := NewEntry(1, 2)
	if e1.Key != 1 || e1.Value != 2 {
		t.Errorf("NewEntry() = %v, want %v", e1, Entry[int, int]{1, 2})
	}
	e2 := NewEntry("key", "value")
	if e2.Key != "key" || e2.Value != "value" {
		t.Errorf("NewEntry() = %v, want %v", e2, Entry[string, string]{"key", "value"})
	}
}

func TestNewEntrySet(t *testing.T) {
	e1 := NewEntrySet([]*Entry[string, int]{
		NewEntry("1", 2),
		NewEntry("2", 4)},
	)

	if e1.entries["1"].Key != "1" || e1.entries["1"].Value != 2 || e1.entries["2"].Key != "2" || e1.entries["2"].Value != 4 {
		t.Errorf("NewEntrySet() = %v, want %v", e1, EntrySet[string, int]{
			entries: map[string]*Entry[string, int]{
				"1": NewEntry("1", 2),
				"2": NewEntry("2", 4),
			},
		})
	}

	e2 := NewEntrySet([]*Entry[string, string]{NewEntry("key", "value")})
	if e2.entries["key"].Key != "key" || e2.entries["key"].Value != "value" {
		t.Errorf("NewEntrySet() = %v, want %v", e2, EntrySet[string, string]{
			entries: map[string]*Entry[string, string]{"key": NewEntry("key", "value")},
		})
	}
}

func TestAddEntry(t *testing.T) {
	e1 := NewEntrySet([]*Entry[string, int]{
		NewEntry("1", 2),
		NewEntry("2", 4)},
	)
	e1.Add(NewEntry("3", 6), NewEntry("4", 8))
	if e1.entries["3"].Key != "3" || e1.entries["3"].Value != 6 {
		t.Errorf("AddEntry() = %v, want %v", e1, EntrySet[string, int]{
			entries: map[string]*Entry[string, int]{
				"1": NewEntry("1", 2),
				"2": NewEntry("2", 4),
				"3": NewEntry("3", 6),
				"4": NewEntry("4", 8),
			},
		})
	}
}

func TestKeys(t *testing.T) {
	e1 := NewEntrySet([]*Entry[string, int]{
		NewEntry("1", 2),
		NewEntry("2", 4)},
	)
	keys := e1.Keys()
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	if keys[0] != "1" || keys[1] != "2" {
		t.Errorf("Keys() = %v, want %v", keys, []string{"1", "2"})
	}
}

func TestValues(t *testing.T) {
	e1 := NewEntrySet([]*Entry[string, int]{
		NewEntry("1", 2),
		NewEntry("2", 4)},
	)
	values := e1.Values()
	if values[0] != 2 && values[0] != 4 || values[1] != 2 && values[1] != 4 {
		t.Errorf("Values() = %v, want %v", values, []int{2, 4})
	}
}

func TestEntrySet_Len(t *testing.T) {
	e1 := NewEntrySet([]*Entry[string, int]{
		NewEntry("1", 2),
		NewEntry("2", 4)},
	)
	if e1.Len() != 2 {
		t.Errorf("EntrySet.Len() = %v, want %v", e1.Len(), 2)
	}
}

func TestEach(t *testing.T) {
	e1 := NewEntrySet([]*Entry[string, int]{
		NewEntry("1", 2),
		NewEntry("2", 4)},
	)

	Each(e1.entries, func(k string, v *Entry[string, int]) {
		if k != "1" && k != "2" || v.Value != 2 && v.Value != 4 {
			t.Errorf("Each() = %v, want %v", e1, EntrySet[string, int]{
				entries: map[string]*Entry[string, int]{
					"1": NewEntry("1", 2),
					"2": NewEntry("2", 4),
				},
			})
		}
	})
}

func TestMap(t *testing.T) {
	e1 := NewEntrySet([]*Entry[string, int]{
		NewEntry("1", 2),
		NewEntry("2", 4)},
	)
	m := Map(e1.entries, func(k string, v *Entry[string, int]) *Entry[string, int] {
		return NewEntry(k, v.Value*2)
	})

	if m["1"] != 4 || m["2"] != 8 {
		t.Errorf("Map() = %v, want %v", m, map[string]*Entry[string, int]{"1": NewEntry("1", 4), "2": NewEntry("2", 8)})
	}
}

func TestMapEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	e1 := MapEntries(m, func(k string, v int) *Entry[string, int] {
		return NewEntry(k, v)
	})
	if e1.entries["a"].Key != "a" || e1.entries["a"].Value != 1 || e1.entries["b"].Key != "b" || e1.entries["b"].Value != 2 || e1.entries["c"].Key != "c" || e1.entries["c"].Value != 3 {
		t.Errorf("MapEntries() = %v, want %v", e1, EntrySet[string, int]{
			entries: map[string]*Entry[string, int]{
				"a": NewEntry("a", 1),
				"b": NewEntry("b", 2),
				"c": NewEntry("c", 3),
			},
		})
	}
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	v := MapValues(m, func(v int) int {
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

func TestFlatValues(t *testing.T) {
	// Test Case 1: Flattening without transformation (identity function)
	map1 := map[string][]int{
		"numbers": {1, 2, 3},
		"more":    {4, 5, 6},
	}
	expected1 := []int{1, 2, 3, 4, 5, 6}
	result1 := FlatValues(map1, functions.Identity[int])
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
	result2 := FlatValues(map2, func(v string) int { return len(v) })
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
	result3 := FlatValues(map3, func(v string) string { return strings.ToUpper(v) })
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
	result4 := FlatValues(map4, func(v string) string { return string(v[0]) })
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
	result1 := FlatValuesSorted(map1, func(v int) int { return v }, func(i, j int) bool { return i > j })

	if !arrays.Equal(result1, expected1) {
		t.Errorf("Test Case 1 Failed: expected %v, got %v", expected1, result1)
	}

	// Test Case 2: Sort strings by length
	map2 := map[string][]string{
		"fruits":  {"apple", "banana"},
		"colors":  {"carmine", "red"},
		"animals": {"lion", "antelope"},
	}
	expected2 := []string{"antelope", "carmine", "banana", "apple", "lion", "red"}
	result2 := FlatValuesSorted(map2, func(v string) string { return v }, func(i, j string) bool { return len(i) > len(j) })

	if !arrays.Equal(result2, expected2) {
		t.Errorf("Test Case 2 Failed: expected %v, got %v", expected2, result2)
	}

	// Test Case 3: Sort strings alphabetically
	map3 := map[string][]string{
		"words": {"banana", "apple", "cherry"},
	}
	expected3 := []string{"APPLE", "BANANA", "CHERRY"}
	result3 := FlatValuesSorted(map3, func(v string) string { return strings.ToUpper(v) }, func(i, j string) bool { return i < j })

	if !arrays.Equal(result3, expected3) {
		t.Errorf("Test Case 3 Failed: expected %v, got %v", expected3, result3)
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	v := MapKeys(m, func(k string) string {
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

func TestCast(t *testing.T) {
	m := map[string]int{"1": 1, "2": 1, "3": 2, "4": 3, "5": 5, "6": 8, "7": 13}

	pairs := Cast(m, func(k string, v int) generics.BiTypedInterface[string, int] {
		return tuple.NewPair(k, v)
	})

	expected := []*tuple.Pair[string, int]{
		tuple.NewPair("1", 1),
		tuple.NewPair("2", 1),
		tuple.NewPair("3", 2),
		tuple.NewPair("4", 3),
		tuple.NewPair("5", 5),
		tuple.NewPair("6", 8),
		tuple.NewPair("7", 13),
	}

	for i, p := range pairs {
		if pair, ok := p.(*tuple.Pair[string, int]); ok {
			if pair.Second() != m[pair.First()] {
				t.Errorf("Cast() = %v, want %v", pair, expected[i])
			}
		} else {
			t.Errorf("Type assertion failed for pair: %v", p)
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

	result := MapWithKeys(keys, f)

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

	result := MapWithValues(values, f)

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

	result2 := MapWithValues(values2, f)

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

func TestFetch(t *testing.T) {
	entries := []generics.BiTypedInterface[int, string]{
		tuple.NewPair(1, "one"),
		tuple.NewPair(2, "two"),
		tuple.NewPair(3, "three"),
	}
	expected := map[int]string{1: "one", 2: "two", 3: "three"}

	f := func(entry generics.BiTypedInterface[int, string]) (int, string) {
		pair := entry.(*tuple.Pair[int, string])
		return pair.First(), pair.Second()
	}

	result := Fetch(entries, f)

	if len(result) != len(expected) {
		t.Errorf("Expected map length %d, got %d", len(expected), len(result))
	}

	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected value %s for key %d, got %s", v, k, result[k])
		}
	}
}
