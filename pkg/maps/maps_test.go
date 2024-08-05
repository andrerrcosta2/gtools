// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/pkg/generics"
	"github.com/andrerrcosta2/gtools/pkg/tuple"
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
	if v[0] != 2 || v[1] != 4 || v[2] != 6 {
		t.Errorf("MapValues() = %v, want %v", v, []int{2, 4, 6})
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	v := MapKeys(m, func(k string) string {
		return k + "1"
	})
	if v[0] != "a1" || v[1] != "b1" || v[2] != "c1" {
		t.Errorf("MapKeys() = %v, want %v", v, []string{"a1", "b1", "c1"})
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
