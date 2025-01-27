// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package grammar

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data"
	"testing"
)

func TestNewSymbol(t *testing.T) {
	// Test that NewSymbol works as expected with a string
	if symbol, err := NewSymbol("test"); err != nil {
		t.Errorf("Expected create a string symbol, got %s", err)
	} else {
		if symbol.String() != "test" {
			t.Errorf("Expected 'test', got %s", symbol.String())
		}
	}

	// Test that NewSymbol works as expected with a []byte
	if symbol, err := NewSymbol([]byte("test")); err != nil {
		t.Errorf("Expected create a []byte symbol, got %s", err)
	} else {
		if symbol.String() != "test" {
			t.Errorf("Expected 'test', got %s", symbol.String())
		}
	}

	// Test that NewSymbol works as expected with a ByteSymbol
	if symbol, err := NewSymbol("test"); err != nil {
		t.Errorf("Expected create a []byte symbol, got %s", err)
	} else {
		if symbol.String() != "test" {
			t.Errorf("Expected 'test', got %s", symbol.String())
		}
	}
}

func TestSymbol_NewSymbol_Numbers(t *testing.T) {
	// Test that NewSymbol works as expected with an int
	intSymbol, err := NewSymbol(1234)
	if err != nil {
		t.Fatalf("Expected create a int symbol, got %s", err)
	}

	intResult, err := GetSymbolValue[int](intSymbol)
	if err != nil {
		t.Fatalf("Expected create a int symbol, got %s", err)
	}

	if intResult != 1234 {
		t.Fatalf("Expected '1234', got %d", intResult)
	}

	// Test that NewSymbol works as expected with a float
	floatSymbol, err := NewSymbol(1234.56)
	if err != nil {
		t.Fatalf("Expected create a float symbol, got %s", err)
	}

	floatResult, err := GetSymbolValue[float64](floatSymbol)
	if err != nil {
		t.Fatalf("Expected create a float symbol, got %s", err)
	}

	if floatResult != 1234.56 {
		t.Fatalf("Expected '1234.56', got %f", floatResult)
	}

	// Test that NewSymbol works as expected with a complex
	complexSymbol, err := NewSymbol(1234.56 + 0i)
	if err != nil {
		t.Fatalf("Expected create a complex symbol, got %s", err)
	}

	complexResult, err := GetSymbolValue[complex128](complexSymbol)
	if err != nil {
		t.Fatalf("Expected create a complex symbol, got %s", err)
	}

	if complexResult != (1234.56 + 0i) {
		t.Fatalf("Expected '(1234.56 + 0i)', got %f", complexResult)
	}
}

func TestSymbol_Add(t *testing.T) {
	symbol := ByteSymbol("test").Append(ByteSymbol(".test2"))
	if symbol.String() != "test.test2" {
		t.Errorf("Expected 'test.test2', got %s", symbol.String())
	}

	intSymbol, err := NewSymbol(1234)
	if err != nil {
		t.Fatalf("Expected create a int symbol, got %s", err)
	}

	mergedSymbol := intSymbol.Append(ByteSymbol(".test2"))

	// Try to extract the value from the symbol
	originalSymbol, addedSymbol := mergedSymbol.ExtractSymbol(intSymbol)

	if addedSymbol.String() != ".test2" {
		t.Errorf("Expected '.test2', got %s", addedSymbol.String())
	}

	intResult, err := GetSymbolValue[int](originalSymbol)
	if err != nil {
		t.Fatalf("Expected to get the int symbol value correctly, but got %s", err)
	}
	if intResult != 1234 {
		t.Errorf("Expected '1234', got %d", intResult)
	}
}

func TestSymbol_After(t *testing.T) {
	symbol := ByteSymbol("test").After(1)
	if symbol.String() != "st" {
		t.Errorf("Expected 'st', got %s", symbol.String())
	}
}

func TestSymbol_EdgeCase_After(t *testing.T) {
	symbol := ByteSymbol("test").After(4)
	if symbol.String() != "" {
		t.Errorf("Expected '', got %s", symbol.String())
	}

	symbol = ByteSymbol("test").After(-1)
	if symbol.String() != "test" {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
}

func TestSymbol_Before(t *testing.T) {
	symbol := ByteSymbol("test").Before(3)
	if symbol.String() != "tes" {
		t.Errorf("Expected 'te', got %s", symbol.String())
	}
}

func TestSymbol_EdgeCase_Before(t *testing.T) {
	symbol := ByteSymbol("test").Before(0)
	if symbol.String() != "" {
		t.Errorf("Expected '', got %s", symbol.String())
	}

	symbol = ByteSymbol("test").Before(-1)
	if symbol.String() != "" {
		t.Errorf("Expected '', got %s", symbol.String())
	}

	symbol = ByteSymbol("test").Before(10)
	if symbol.String() != "test" {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
}

func TestSymbol_Between(t *testing.T) {
	symbol := ByteSymbol("test").Between(1, 3)
	if symbol.String() != "es" {
		t.Errorf("Expected 'es', got %s", symbol.String())
	}
}

func TestSymbol_EdgeCase_Between(t *testing.T) {
	symbol := ByteSymbol("test").Between(0, 0)
	if symbol.String() != "" {
		t.Errorf("Expected '', got '%s'", symbol.String())
	}

	symbol = ByteSymbol("test").Between(-1, 0)
	if symbol.String() != "" {
		t.Errorf("Expected '', got '%s'", symbol.String())
	}

	symbol = ByteSymbol("test").Between(0, 10)
	if symbol.String() != "test" {
		t.Errorf("Expected 'test', got '%s'", symbol.String())
	}

	symbol = ByteSymbol("test").Between(0, -1)
	if symbol.String() != "" {
		t.Errorf("Expected '', got '%s'", symbol.String())
	}
}

func TestSymbol_Bytes(t *testing.T) {
	symbol := ByteSymbol("test")
	if string(symbol.Bytes()) != "test" {
		t.Errorf("Expected 'test', got '%s'", symbol.String())
	}
}

func TestSymbol_Len(t *testing.T) {
	symbol := ByteSymbol("test")
	if symbol.Len() != 4 {
		t.Errorf("Expected 'test', got '%s'", symbol.String())
	}
}

func TestSymbol_Equal(t *testing.T) {
	symbol := ByteSymbol("test")
	if !symbol.Equal(ByteSymbol("test")) {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}

	symbol = ByteSymbol("test2")
	if symbol.Equal(ByteSymbol("test")) {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
}

func TestSymbol_Empty(t *testing.T) {
	symbol := ByteSymbol("test")
	if symbol.IsEmpty() {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
	symbol = ByteSymbol("")
	if !symbol.IsEmpty() {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
}

func TestSymbol_Contains(t *testing.T) {
	symbol := ByteSymbol("test-test2")
	if !symbol.Contains(ByteSymbol("test")) {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
	if !symbol.Contains(ByteSymbol("test2")) {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
	if !symbol.Contains(ByteSymbol("test-test2")) {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
	if symbol.Contains(ByteSymbol("test3")) {
		t.Errorf("Expected 'test', got %s", symbol.String())
	}
}

func TestSymbol_CutAll(t *testing.T) {
	symbol := ByteSymbol("test-test2")
	result := symbol.CutAll(ByteSymbol("test"))
	expected := ByteSymbol("-2")
	if !result.Equal(expected) {
		t.Errorf("Expected '-2', got '%s'", result.String())
	}

	symbol = ByteSymbol("test-foo-test-bar-test")
	result = symbol.CutAll(ByteSymbol("test"))
	expected = ByteSymbol("-foo--bar-")
	if !result.Equal(expected) {
		t.Errorf("Expected '-foo--bar-', got '%s'", result.String())
	}

	result = result.CutAll(ByteSymbol("-"))
	expected = ByteSymbol("foobar")
	if !result.Equal(expected) {
		t.Errorf("Expected 'foobar', got '%s'", result.String())
	}
}

func TestSymbol_CutFirst(t *testing.T) {
	symbol := ByteSymbol("test-test2")
	result := symbol.CutFirst(ByteSymbol("test"))
	if !result.Equal(ByteSymbol("-test2")) {
		t.Errorf("Expected '-test2', got %s", symbol.String())
	}
}

func TestSymbol_CutLast(t *testing.T) {
	symbol := ByteSymbol("test-test2")
	result := symbol.CutLast(ByteSymbol("test"))
	if !result.Equal(ByteSymbol("test-2")) {
		t.Errorf("Expected 'test-', got %s", symbol.String())
	}
}

func TestSymbol_IndexOf(t *testing.T) {
	symbol := ByteSymbol("test-test2")
	if idx, found := symbol.IndexOf(ByteSymbol("test")); !found || idx != 0 {
		t.Errorf("Expected idx == '0', got '%d'", idx)
	}
	if idx, found := symbol.IndexOf(ByteSymbol("test2")); !found || idx != 5 {
		t.Errorf("Expected idx == '4', got %d", idx)
	}
	if idx, found := symbol.IndexOf(ByteSymbol("test3")); found || idx != -1 {
		t.Errorf("Expected idx == '-1', got %d", idx)
	}
}

func TestSymbol_Map(t *testing.T) {
	symbol := ByteSymbol("test-test2")
	result := symbol.Map(func(_ int, b byte) byte {
		return b + 1
	})
	if !result.Equal(ByteSymbol("uftu.uftu3")) {
		t.Errorf("Expected 'uftu.uftu3', got %s", result.String())
	}

	if symbol2, err := NewSymbol(12345); err != nil {
		t.Errorf("Expected nil, got %s", err)
	} else {
		result := symbol2.Map(func(i int, b byte) byte {
			if i == symbol2.Len()-1 {
				return (b ^ 0xFF) + 1
			}
			return b ^ 0xFF
		})
		after, err := NewSymbol(-12345)
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
		if !result.Equal(after) {
			val, err := GetSymbolValue[int](result)
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
			t.Errorf("Expected '-12345', got %d", val)
		}
	}
}

func TestSymbol_Remove(t *testing.T) {
	symbol := ByteSymbol("test-test2")
	result := symbol.Remove(0, 5)
	if !result.Equal(ByteSymbol("test2")) {
		t.Errorf("Expected 'test', got %s", result.String())
	}
}

func TestSymbol_EdgeCase_Remove(t *testing.T) {
	symbol := ByteSymbol("test")
	result := symbol.Remove(0, 5)
	if !result.Equal(ByteSymbol("")) {
		t.Errorf("Expected '', got %s", result.String())
	}

	symbol = ByteSymbol("")
	result = symbol.Remove(0, 5)
	if !result.Equal(ByteSymbol("")) {
		t.Errorf("Expected '', got %s", result.String())
	}

	symbol = ByteSymbol("abcdefg")
	result = symbol.Remove(14, -2)
	if !result.Equal(ByteSymbol("")) {
		t.Errorf("Expected '', got %s", result.String())
	}

}

func TestSymbol_ReplaceAll(t *testing.T) {
	// IsEmpty replacement symbol should erase the target
	symbol := ByteSymbol("test-end")
	result, replaced := symbol.ReplaceAll(ByteSymbol("test"), ByteSymbol(""))
	if !replaced || !result.Equal(ByteSymbol("-end")) {
		t.Errorf("Expected 'test-end', got %s", result)
	}

	// Replace target symbol that appears multiple times
	symbol = ByteSymbol("test-test-test")
	result, replaced = symbol.ReplaceAll(ByteSymbol("test"), ByteSymbol("foo"))
	if !replaced || !result.Equal(ByteSymbol("foo-foo-foo")) {
		t.Errorf("Expected 'foo-foo-foo', got %s", result)
	}

	// Replace target symbol that appears once
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceAll(ByteSymbol("test"), ByteSymbol("foo"))
	if !replaced || !result.Equal(ByteSymbol("foo-end")) {
		t.Errorf("Expected 'foo-end', got %s", result)
	}

	// Replace target symbol that appears at the end
	symbol = ByteSymbol("test-end-test")
	result, replaced = symbol.ReplaceAll(ByteSymbol("test"), ByteSymbol("foo"))
	if !replaced || !result.Equal(ByteSymbol("foo-end-foo")) {
		t.Errorf("Expected 'foo-end-foo', got %s", result)
	}

	// Replace target symbol that appears at the beginning
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceAll(ByteSymbol("test"), ByteSymbol("foo"))
	if !replaced || !result.Equal(ByteSymbol("foo-end")) {
		t.Errorf("Expected 'foo-end', got %s", result)
	}

	// Target symbol is not present in the input symbol
	symbol = ByteSymbol("no-match")
	result, replaced = symbol.ReplaceAll(ByteSymbol("test"), ByteSymbol("foo"))
	if replaced || !result.Equal(ByteSymbol("no-match")) {
		t.Errorf("Expected 'no-match', got %s", result)
	}

	// IsEmpty target symbol should return the original symbol
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceAll(ByteSymbol(""), ByteSymbol("foo"))
	if replaced || !result.Equal(ByteSymbol("test-end")) {
		t.Errorf("Expected 'test-end', got %s", result)
	}
}

func TestSymbol_ReplaceEach(t *testing.T) {
	// Replace each occurrence with a different value
	symbol := ByteSymbol("one-two-three-four")
	result, replaced := symbol.ReplaceEach(ByteSymbol("-"), ByteSymbol("A"), ByteSymbol("B"), ByteSymbol("C"))
	if !replaced || !result.Equal(ByteSymbol("oneAtwoBthreeCfour")) {
		t.Errorf("Expected 'oneAtwoBthreeC', got %s", result)
	}

	// Replace each occurrence with fewer replacement values
	symbol = ByteSymbol("a-b-c-d")
	result, replaced = symbol.ReplaceEach(ByteSymbol("-"), ByteSymbol("1"), ByteSymbol("2"))
	if replaced || !result.Equal(ByteSymbol("a-b-c-d")) {
		t.Errorf("Expected 'a1b2c-d', got %s", result)
	}

	// Replace each occurrence with more replacement values
	symbol = ByteSymbol("x-y-z")
	result, replaced = symbol.ReplaceEach(ByteSymbol("-"), ByteSymbol("A"), ByteSymbol("B"), ByteSymbol("C"), ByteSymbol("D"))
	if replaced || !result.Equal(ByteSymbol("x-y-z")) {
		t.Errorf("Expected 'xAyBz', got %s", result)
	}

	// No occurrences of the target symbol
	symbol = ByteSymbol("abc")
	result, replaced = symbol.ReplaceEach(ByteSymbol("-"), ByteSymbol("x"))
	if replaced || !result.Equal(ByteSymbol("abc")) {
		t.Errorf("Expected 'abc', got %s", result)
	}

	// IsEmpty target symbol should return the original symbol
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceEach(ByteSymbol(""), ByteSymbol("foo"))
	if replaced || !result.Equal(ByteSymbol("test-end")) {
		t.Errorf("Expected 'test-end', got %s", result)
	}

	// IsEmpty replacement symbol should erase the target
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceEach(ByteSymbol("test"), ByteSymbol(""))
	if !replaced || !result.Equal(ByteSymbol("-end")) {
		t.Errorf("Expected '', got %s", result)
	}
}

func TestSymbol_ReplaceFirst(t *testing.T) {
	// Replace the first occurrence with a single value
	symbol := ByteSymbol("one-two-three")
	result, replaced := symbol.ReplaceFirst(ByteSymbol("-"), ByteSymbol("A"))
	if !replaced || !result.Equal(ByteSymbol("oneAtwo-three")) {
		t.Errorf("Expected 'oneA-two-three', got %s", result)
	}

	// No occurrences of the target symbol
	symbol = ByteSymbol("abc")
	result, replaced = symbol.ReplaceFirst(ByteSymbol("-"), ByteSymbol("x"))
	if replaced || !result.Equal(ByteSymbol("abc")) {
		t.Errorf("Expected 'abc', got %s", result)
	}

	// Replace first occurrence with an empty replacement symbol
	symbol = ByteSymbol("hello-world")
	result, replaced = symbol.ReplaceFirst(ByteSymbol("-"), ByteSymbol(""))
	if !replaced || !result.Equal(ByteSymbol("helloworld")) {
		t.Errorf("Expected 'helloworld', got %s", result)
	}

	// IsEmpty target symbol should return the original symbol
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceFirst(ByteSymbol(""), ByteSymbol("foo"))
	if replaced || !result.Equal(ByteSymbol("test-end")) {
		t.Errorf("Expected 'test-end', got %s", result)
	}

	// IsEmpty replacement symbol should erase the target
	symbol = ByteSymbol("test-end-test")
	result, replaced = symbol.ReplaceFirst(ByteSymbol("test"), ByteSymbol(""))
	if !replaced || !result.Equal(ByteSymbol("-end-test")) {
		t.Errorf("Expected '', got %s", result)
	}
}

func TestSymbol_ReplaceLast(t *testing.T) {
	// Replace the last occurrence with a single value
	symbol := ByteSymbol("one-two-three")
	result, replaced := symbol.ReplaceLast(ByteSymbol("-"), ByteSymbol("A"))
	if !replaced || !result.Equal(ByteSymbol("one-twoAthree")) {
		t.Errorf("Expected 'one-twoAthree', got %s", result)
	}

	// No occurrences of the target symbol
	symbol = ByteSymbol("abc")
	result, replaced = symbol.ReplaceLast(ByteSymbol("-"), ByteSymbol("x"))
	if replaced || !result.Equal(ByteSymbol("abc")) {
		t.Errorf("Expected 'abc', got %s", result)
	}

	// Replace last occurrence with an empty replacement symbol
	symbol = ByteSymbol("hello-world-planet")
	result, replaced = symbol.ReplaceLast(ByteSymbol("-"), ByteSymbol(""))
	if !replaced || !result.Equal(ByteSymbol("hello-worldplanet")) {
		t.Errorf("Expected 'hello-worldplanet', got %s", result)
	}

	// IsEmpty target symbol should return the original symbol
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceLast(ByteSymbol(""), ByteSymbol("foo"))
	if replaced || !result.Equal(ByteSymbol("test-end")) {
		t.Errorf("Expected 'test-end', got %s", result)
	}

	// IsEmpty replacement symbol should erase the target
	symbol = ByteSymbol("test-end-test")
	result, replaced = symbol.ReplaceLast(ByteSymbol("test"), ByteSymbol(""))
	if !replaced || !result.Equal(ByteSymbol("test-end-")) {
		t.Errorf("Expected 'test-end-', got %s", result)
	}
}

func TestSymbol_ReplaceMany(t *testing.T) {
	// Replace multiple symbols with a single value
	symbol := ByteSymbol("apple-orange-banana")
	result, replaced := symbol.ReplaceMany(ByteSymbol("fruit"), ByteSymbol("apple"), ByteSymbol("banana"))
	if !replaced || !result.Equal(ByteSymbol("fruit-orange-fruit")) {
		t.Errorf("Expected 'fruit-orange-fruit', got %s", result)
	}

	// Replace multiple symbols with the same replacement
	symbol = ByteSymbol("cat-dog-cat-bird")
	result, replaced = symbol.ReplaceMany(ByteSymbol("pet"), ByteSymbol("cat"), ByteSymbol("dog"))
	if !replaced || !result.Equal(ByteSymbol("pet-pet-pet-bird")) {
		t.Errorf("Expected 'pet-pet-pet-bird', got %s", result)
	}

	// No occurrences of symbols to replace
	symbol = ByteSymbol("one-two-three")
	result, replaced = symbol.ReplaceMany(ByteSymbol("num"), ByteSymbol("four"), ByteSymbol("five"))
	if replaced || !result.Equal(ByteSymbol("one-two-three")) {
		t.Errorf("Expected 'one-two-three', got %s", result)
	}

	// Replace many symbols with an empty replacement symbol
	symbol = ByteSymbol("abc-def-ghi")
	result, replaced = symbol.ReplaceMany(ByteSymbol(""), ByteSymbol("abc"), ByteSymbol("ghi"))
	if !replaced || !result.Equal(ByteSymbol("-def-")) {
		t.Errorf("Expected '-def-', got %s", result)
	}

	// IsEmpty replacement symbol and multiple symbols to replace
	symbol = ByteSymbol("test-apple")
	result, replaced = symbol.ReplaceMany(ByteSymbol(""), ByteSymbol("test"), ByteSymbol("apple"))
	if !replaced || !result.Equal(ByteSymbol("-")) {
		t.Errorf("Expected '-', got %s", result)
	}

	// IsEmpty target symbol should return the original symbol
	symbol = ByteSymbol("test-end")
	result, replaced = symbol.ReplaceMany(ByteSymbol("foo-bar"), ByteSymbol(""), ByteSymbol(""))
	if replaced || !result.Equal(ByteSymbol("test-end")) {
		t.Errorf("Expected 'test-end', got %s", result)
	}
}

func TestSymbol_Split(t *testing.T) {
	symbol := ByteSymbol("test-splittable-test")
	splitter := ByteSymbol("test")

	if sym, rest := symbol.ExtractSymbol(splitter); !sym.Equal(splitter) || !rest.Equal(ByteSymbol("-splittable-test")) {
		t.Errorf("Expected sym == 'test', but got '%s'.\n Expected rest == 'splittable-test', but got '%s'\n", sym, rest)
	}
}

func TestSymbol_EdgeCase_SplittingSpecialSymbols(t *testing.T) {
	placeholder := ByteSymbol(`\x01`)

	symbol := ByteSymbol("").Append(placeholder).Append(ByteSymbol("test"))

	if sym, rest := symbol.ExtractSymbol(placeholder); !sym.Equal(placeholder) || !rest.Equal(ByteSymbol("test")) {
		t.Errorf("Expected sym == '\\x01', but got '%s'.\n Expected rest == 'test', but got '%s'\n", sym, rest)
	}
}

func TestSymbol_EdgeCase_SplittingSpecialSymbolsInLoop(t *testing.T) {
	placeholder := ByteSymbol(`\x01`)
	// [\x01,\x01]
	symbol := ByteSymbol("[").Append(placeholder).Append(ByteSymbol(",").Append(placeholder).Append(ByteSymbol("]")))

	if !symbol.Equal(ByteSymbol("[\\x01,\\x01]")) {
		t.Errorf("Expected symbol == '[\\x01,\\x01]', but got '%s'.\n", symbol)
	}

	var symResult Symbol
	var symRest Symbol
	if symResult, symRest = symbol.ExtractSymbol(placeholder); !symResult.Equal(placeholder) || !symRest.Equal(ByteSymbol(",\\x01]")) {
		t.Errorf("Expected sym == '\\x01', but got '%s'.\n Expected rest == ',\\x01]', but got '%s'\n", symResult, symRest)
	}

	fmt.Printf("symResult: %s\n", symResult)
	fmt.Printf("symRest: %s\n", symRest)

	if symResult, symRest = symRest.ExtractSymbol(placeholder); !symResult.Equal(placeholder) || !symRest.Equal(ByteSymbol("]")) {
		t.Errorf("Expected sym == '\\x01', but got '%s'.\n Expected rest == ']', but got '%s'\n", symResult, symRest)
	}
}

func TestSymbol_StartsWith(t *testing.T) {
	symbol := ByteSymbol("test-testing")
	if !symbol.StartsWith(ByteSymbol("test")) {
		t.Errorf("Expected true, got false")
	}
	if symbol.StartsWith(ByteSymbol("testing")) {
		t.Errorf("Expected false, got true")
	}
}

func TestSymbol_Validate(t *testing.T) {
	tests := []struct {
		name      string
		symbol    data.Validatable
		rgx       string
		expectErr bool
	}{
		{
			name:      "Valid symbol matching pattern",
			symbol:    ByteSymbol("abc123"),
			rgx:       "^[a-zA-Z0-9]+$", // Alphanumeric pattern
			expectErr: false,
		},
		{
			name:      "Invalid symbol not matching pattern",
			symbol:    ByteSymbol("abc-123"),
			rgx:       "^[a-zA-Z0-9]+$", // Alphanumeric pattern
			expectErr: true,
		},
		{
			name:      "Invalid regular expression",
			symbol:    ByteSymbol("abc123"),
			rgx:       "[a-zA-Z0-9+", // Invalid regular expression pattern
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.symbol.Validate(tt.rgx)
			if tt.expectErr && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestSymbol_String(t *testing.T) {
	symbol := ByteSymbol("test")
	if symbol.String() != "test" {
		t.Errorf("Expected 'test', got '%s'", symbol.String())
	}
}

func TestSymbol_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		symbol   data.Validatable
		rgx      string
		expected bool
	}{
		{
			name:     "Valid symbol matching pattern",
			symbol:   ByteSymbol("abc123"),
			rgx:      "^[a-zA-Z0-9]+$", // Alphanumeric pattern
			expected: true,
		},
		{
			name:     "Invalid symbol not matching pattern",
			symbol:   ByteSymbol("abc-123"),
			rgx:      "^[a-zA-Z0-9]+$", // Alphanumeric pattern
			expected: false,
		},
		{
			name:     "Invalid regular expression",
			symbol:   ByteSymbol("abc123"),
			rgx:      "[a-zA-Z0-9+", // Invalid regular expression pattern
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.symbol.IsValid(tt.rgx)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
