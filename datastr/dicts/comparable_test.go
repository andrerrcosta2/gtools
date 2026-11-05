// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
)

func TestComparable_New(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		m := Comparable[string, int]()
		if m == nil {
			t.Fatal("expected non-nil map")
		}
		if m.Len() != 0 {
			t.Errorf("expected length 0, got %d", m.Len())
		}
	})

	t.Run("with entries", func(t *testing.T) {
		entry1 := NewAnyEntry("key1", 100)
		entry2 := NewAnyEntry("key2", 200)

		m := Comparable(entry1, entry2)

		if m.Len() != 2 {
			t.Errorf("expected length 2, got %d", m.Len())
		}

		val, ok := m.Get("key1")
		if !ok || val != 100 {
			t.Errorf("expected (100, true), got (%d, %t)", val, ok)
		}

		val, ok = m.Get("key2")
		if !ok || val != 200 {
			t.Errorf("expected (200, true), got (%d, %t)", val, ok)
		}
	})
}

func TestComparable_PutGet(t *testing.T) {
	m := Comparable[string, string]()

	// Inserir alguns valores
	m.Put("name", "Alice")
	m.Put("city", "New York")
	m.Put("country", "USA")

	// Verificar leitura
	tests := []struct {
		key      string
		expected string
		wantOk   bool
	}{
		{"name", "Alice", true},
		{"city", "New York", true},
		{"country", "USA", true},
	}

	for _, tt := range tests {
		val, ok := m.Get(tt.key)
		if ok != tt.wantOk {
			t.Errorf("Get(%q) ok = %t, want %t", tt.key, ok, tt.wantOk)
		}
		if val != tt.expected {
			t.Errorf("Get(%q) = %q, want %q", tt.key, val, tt.expected)
		}
	}
}

func TestComparable_PutOverwrite(t *testing.T) {
	m := Comparable[string, float64]()

	m.Put("pi", 3.14)
	m.Put("pi", 3.14159)

	val, ok := m.Get("pi")
	if !ok {
		t.Fatal("expected key 'pi' to exist")
	}
	if val != 3.14159 {
		t.Errorf("expected 3.14159, got %f", val)
	}

	if m.Len() != 1 {
		t.Errorf("expected length 1 after overwrite, got %d", m.Len())
	}
}

func TestComparable_GetMissing(t *testing.T) {
	m := Comparable[int, string]()
	m.Put(1, "one")

	// Buscar chave inexistente
	val, ok := m.Get(999)
	if ok {
		t.Error("expected ok=false for missing key")
	}
	if val != "" {
		t.Errorf("expected zero value, got %q", val)
	}
}

func TestComparable_Delete(t *testing.T) {
	m := Comparable[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	m.Delete("b")

	// Verificar que 'b' foi removida
	_, ok := m.Get("b")
	if ok {
		t.Error("expected key 'b' to be deleted")
	}

	// Verificar que outras chaves ainda existem
	if m.Len() != 2 {
		t.Errorf("expected length 2, got %d", m.Len())
	}

	if val, ok := m.Get("a"); !ok || val != 1 {
		t.Error("key 'a' should still exist")
	}
	if val, ok := m.Get("c"); !ok || val != 3 {
		t.Error("key 'c' should still exist")
	}
}

func TestComparable_DeleteMissing(t *testing.T) {
	m := Comparable[string, int]()
	m.Put("x", 10)

	// Não deve causar pânico ao deletar chave inexistente
	m.Delete("nonexistent")

	if m.Len() != 1 {
		t.Errorf("expected length 1, got %d", m.Len())
	}
}

func TestComparable_Contains(t *testing.T) {
	m := Comparable[string, bool]()
	m.Put("active", true)

	tests := []struct {
		key      string
		expected bool
	}{
		{"active", true},
		{"inactive", false},
		{"", false},
	}

	for _, tt := range tests {
		got := m.Contains(tt.key)
		if got != tt.expected {
			t.Errorf("Contains(%q) = %t, want %t", tt.key, got, tt.expected)
		}
	}
}

func TestComparable_Len(t *testing.T) {
	m := Comparable[int, int]()

	// Verificar comprimento inicial
	if m.Len() != 0 {
		t.Errorf("expected initial length 0, got %d", m.Len())
	}

	// Adicionar elementos
	for i := 0; i < 5; i++ {
		m.Put(i, i*10)
	}
	if m.Len() != 5 {
		t.Errorf("expected length 5 after inserts, got %d", m.Len())
	}

	// Remover alguns
	m.Delete(0)
	m.Delete(2)
	if m.Len() != 3 {
		t.Errorf("expected length 3 after deletions, got %d", m.Len())
	}

	// Sobrescrever não deve alterar o comprimento
	m.Put(1, 999)
	if m.Len() != 3 {
		t.Errorf("expected length 3 after overwrite, got %d", m.Len())
	}
}

func TestComparable_Clear(t *testing.T) {
	m := Comparable[string, string]()
	m.Put("key1", "value1")
	m.Put("key2", "value2")
	m.Put("key3", "value3")

	m.Clear()

	if m.Len() != 0 {
		t.Errorf("expected length 0 after clear, got %d", m.Len())
	}

	if m.Contains("key1") {
		t.Error("expected no keys after clear")
	}

	// Verificar que podemos continuar usando o mapa após limpar
	m.Put("newkey", "newvalue")
	if m.Len() != 1 {
		t.Errorf("expected length 1 after re-insertion, got %d", m.Len())
	}
}

func TestComparable_Keys(t *testing.T) {
	t.Run("non-empty map", func(t *testing.T) {
		m := Comparable[string, int]()
		m.Put("z", 1)
		m.Put("a", 2)
		m.Put("m", 3)

		keys := m.Keys()
		if len(keys) != 3 {
			t.Errorf("expected 3 keys, got %d", len(keys))
		}

		// Verificar que todas as chaves estão presentes
		keySet := make(map[string]bool)
		for _, k := range keys {
			keySet[k] = true
		}

		expectedKeys := []string{"z", "a", "m"}
		for _, k := range expectedKeys {
			if !keySet[k] {
				t.Errorf("expected key %q to be present", k)
			}
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m := Comparable[int, string]()
		keys := m.Keys()
		if len(keys) != 0 {
			t.Errorf("expected empty keys slice, got %v", keys)
		}
	})
}

func TestComparable_Values(t *testing.T) {

	m := Comparable[string, int]()
	m.Put("x", 10)
	m.Put("y", 20)
	m.Put("z", 30)

	values := m.Values()
	if len(values) != 3 {
		t.Errorf("expected 3 values, got %d\n%v", len(values), values)

	}

	// Verificar que todos os valores estão presentes
	valueSet := make(map[int]bool)
	for _, v := range values {
		valueSet[v] = true
	}

	expectedValues := []int{10, 20, 30}
	for _, v := range expectedValues {
		if !valueSet[v] {
			t.Errorf("expected value %d to be present", v)
		}
	}
}

func TestComparable_Iterator(t *testing.T) {
	m := Comparable[string, int]()
	m.Put("one", 1)
	m.Put("two", 2)
	m.Put("three", 3)

	iter := m.Iterator()

	visited := make(map[string]int)
	for key, value, hasNext := iter.Next(); hasNext; key, value, hasNext = iter.Next() {
		visited[key] = value
	}

	if len(visited) != 3 {
		t.Errorf("expected to visit 3 entries, got %d", len(visited))
	}

	// Verificar valores corretos
	expected := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range expected {
		got, exists := visited[k]
		if !exists {
			t.Errorf("key %q not visited", k)
		}
		if got != v {
			t.Errorf("value for %q = %d, want %d", k, got, v)
		}
	}
}

func TestComparable_IteratorSorted(t *testing.T) {
	m := Comparable[string, int]()
	m.Put("delta", 4)
	m.Put("alpha", 1)
	m.Put("charlie", 3)
	m.Put("bravo", 2)

	cmp := comparators.Ordered[string]()
	iter := m.Iterator(cmp)

	var keys []string
	for key, _, hasNext := iter.Next(); hasNext; key, _, hasNext = iter.Next() {
		keys = append(keys, key)
	}

	expected := []string{"alpha", "bravo", "charlie", "delta"}
	if len(keys) != len(expected) {
		t.Errorf("expected %d keys, got %d", len(expected), len(keys))
	}

	for i, k := range keys {
		if k != expected[i] {
			t.Errorf("position %d: expected %q, got %q", i, expected[i], k)
		}
	}
}

func TestComparable_IteratorEmpty(t *testing.T) {
	m := Comparable[int, string]()
	iter := m.Iterator()

	if _, _, hasNext := iter.Next(); hasNext {
		t.Error("expected empty iterator to have no next element")
	}
}

func TestComparable_IteratorSnapshot(t *testing.T) {

	m := Comparable[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)

	// Obter iterator - captura snapshot das chaves
	iter := m.Iterator()
	t.Logf("iter: %v", iter)

	// Modificar mapa após obter iterator
	m.Put("c", 3)
	m.Delete("a")

	visited := make(map[string]int)
	for key, value, hasNext := iter.Next(); hasNext; key, value, hasNext = iter.Next() {
		visited[key] = value
	}

	if _, exists := visited["a"]; !exists {
		t.Error("key 'a' should exist on snapshot")
	}
	if _, exists := visited["b"]; !exists {
		t.Error("key 'b' should exist on snapshot")
	}
	if _, exists := visited["c"]; exists {
		t.Error("key 'c' shouldn't exist on snapshot")
	}
}

// Testes com tipos diferentes
func TestComparable_WithCustomTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	m := Comparable[int, Person]()

	p1 := Person{"Alice", 30}
	p2 := Person{"Bob", 25}

	m.Put(1, p1)
	m.Put(2, p2)

	val, ok := m.Get(1)
	if !ok || val != p1 {
		t.Errorf("expected %+v, got %+v", p1, val)
	}
}

func TestComparable_WithComplexKey(t *testing.T) {
	type ComplexKey struct {
		ID   int
		Type string
	}

	m := Comparable[ComplexKey, string]()

	key := ComplexKey{ID: 1, Type: "user"}
	m.Put(key, "data")

	val, ok := m.Get(key)
	if !ok {
		t.Error("expected key to exist")
	}
	if val != "data" {
		t.Errorf("expected 'data', got %q", val)
	}
}
