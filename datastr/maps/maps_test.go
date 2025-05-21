// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/generics"
	"github.com/andrerrcosta2/gtools/datastr/tuple"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/maps"
	"sort"
	"testing"
)

func TestNewComparableEntry(t *testing.T) {
	e1 := NewComparableEntry(1, 2)
	if e1.Key() != 1 || e1.Value() != 2 {
		t.Errorf("NewComparableEntry() = %v, want %v", e1, ComparableEntry[int, int]{1, 2})
	}
	e2 := NewComparableEntry("key", "value")
	if e2.Key() != "key" || e2.Value() != "value" {
		t.Errorf("NewComparableEntry() = %v, want %v", e2, ComparableEntry[string, string]{"key", "value"})
	}
}

func TestNewEntrySet(t *testing.T) {
	ce1 := NewComparableEntry("1", 2)
	ce2 := NewComparableEntry("2", 4)
	e1 := NewEntrySet[string, int](ce1, ce2)

	expKeys1 := []string{"1", "2"}

	if !maps.ContainsAllKeys[string, *ComparableEntry[string, int]](&e1.entries, expKeys1) {
		t.Errorf("NewEntrySet() = %v, want %v", e1, EntrySet[string, int]{
			entries: map[string]*ComparableEntry[string, int]{
				"1": NewComparableEntry("1", 2),
				"2": NewComparableEntry("2", 4),
			},
		})
	}

	if !maps.ContainsAllValues[string, *ComparableEntry[string, int]](&e1.entries, []*ComparableEntry[string, int]{ce1, ce2}) {
		t.Errorf("NewEntrySet() = %v, want %v", e1, EntrySet[string, int]{
			entries: map[string]*ComparableEntry[string, int]{
				"1": NewComparableEntry("1", 2),
				"2": NewComparableEntry("2", 4),
			},
		})
	}
}

func TestAddEntry(t *testing.T) {
	e1 := NewEntrySet(
		NewComparableEntry("1", 2),
		NewComparableEntry("2", 4),
	)

	err := e1.Add(NewComparableEntry("3", 6), NewComparableEntry("4", 8))
	if err != nil {
		t.Errorf("AddEntry() = %v, want nil", err)
	}

	if e1.entries["3"].Key() != "3" || e1.entries["3"].Value() != 6 {
		t.Errorf("AddEntry() = %v, want %v", e1, EntrySet[string, int]{
			entries: map[string]*ComparableEntry[string, int]{
				"1": NewComparableEntry("1", 2),
				"2": NewComparableEntry("2", 4),
				"3": NewComparableEntry("3", 6),
				"4": NewComparableEntry("4", 8),
			},
		})
	}
}

func TestKeys(t *testing.T) {
	e1 := NewEntrySet(
		NewComparableEntry("1", 2),
		NewComparableEntry("2", 4),
	)
	keys := e1.Keys()
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	if keys[0] != "1" || keys[1] != "2" {
		t.Errorf("Keys() = %v, want %v", keys, []string{"1", "2"})
	}
}

func TestValues(t *testing.T) {
	e1 := NewEntrySet(
		NewComparableEntry("1", 2),
		NewComparableEntry("2", 4),
	)
	values := e1.Values()
	if values[0] != 2 && values[0] != 4 || values[1] != 2 && values[1] != 4 {
		t.Errorf("Values() = %v, want %v", values, []int{2, 4})
	}
}

func TestEntrySet_Len(t *testing.T) {
	e1 := NewEntrySet(
		NewComparableEntry("1", 2),
		NewComparableEntry("2", 4),
	)
	if e1.Len() != 2 {
		t.Errorf("Entries.Len() = %v, want %v", e1.Len(), 2)
	}
}

func TestEach(t *testing.T) {
	e1 := NewEntrySet(
		NewComparableEntry("1", 2),
		NewComparableEntry("2", 4),
	)

	maps.Each(e1.entries, func(k string, v *ComparableEntry[string, int]) {
		if k != "1" && k != "2" || v.Value() != 2 && v.Value() != 4 {
			t.Errorf("Each() = %v, want %v", e1, EntrySet[string, int]{
				entries: map[string]*ComparableEntry[string, int]{
					"1": NewComparableEntry("1", 2),
					"2": NewComparableEntry("2", 4),
				},
			})
		}
	})
}

func TestMap(t *testing.T) {
	input := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	transform := func(key int, value string) *ComparableEntry[string, string] {
		return &ComparableEntry[string, string]{
			key:   fmt.Sprintf("Name-%v", key),
			value: fmt.Sprintf("Value-%v", value),
		}
	}

	result := Map(&input, transform)

	// Define the expected output map
	expected := map[string]string{
		"Name-1": "Value-one",
		"Name-2": "Value-two",
		"Name-3": "Value-three",
	}

	// Compare the result with the expected map
	if len(*result) != len(expected) {
		t.Errorf("Expected map of length %d, but got %d", len(expected), len(*result))
	}

	for k, v := range expected {
		if val, exists := (*result)[k]; !exists || val != v {
			t.Errorf("Expected %v: %v, but got %v", k, v, val)
		}
	}
}

func TestMapEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	e1 := MapEntries(m, func(k string, v int) *ComparableEntry[string, int] {
		return NewComparableEntry(k, v)
	})
	if e1.entries["a"].Key() != "a" || e1.entries["a"].Value() != 1 || e1.entries["b"].Key() != "b" ||
		e1.entries["b"].Value() != 2 || e1.entries["c"].Key() != "c" || e1.entries["c"].Value() != 3 {
		t.Errorf("MapEntries() = %v, want %v", e1, EntrySet[string, int]{
			entries: map[string]*ComparableEntry[string, int]{
				"a": NewComparableEntry("a", 1),
				"b": NewComparableEntry("b", 2),
				"c": NewComparableEntry("c", 3),
			},
		})
	}
}

func TestCast(t *testing.T) {
	m := map[string]int{"1": 1, "2": 1, "3": 2, "4": 3, "5": 5, "6": 8, "7": 13}

	pairs := maps.Cast(&m, func(k string, v int) generics.BiTyped[string, int] {
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

func TestFetch(t *testing.T) {
	entries := []*tuple.Pair[int, string]{
		tuple.NewPair(1, "one"),
		tuple.NewPair(2, "two"),
		tuple.NewPair(3, "three"),
	}
	expected := map[int]string{1: "one", 2: "two", 3: "three"}

	f := func(entry *tuple.Pair[int, string]) (int, string) {
		return entry.First(), entry.Second()
	}

	result := *maps.Fetch(entries, f)

	if len(result) != len(expected) {
		t.Errorf("Expected map length %d, got %d", len(expected), len(result))
	}

	for k, v := range expected {
		if result[k] != v {
			t.Errorf("Expected value %s for key %d, got %s", v, k, result[k])
		}
	}
}
