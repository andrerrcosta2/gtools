// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"errors"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/errs"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/validate"
	"testing"
)

func TestFilterPatternAgainstSymbol_EmptySymbolAndPattern(t *testing.T) {
	// Test with empty symbol and pattern
	err := validate.PatternAgainstSymbol("", symbols.Empty)
	if !errors.Is(err, errs.EmptyEntry) {
		t.Fatalf("Expected error for empty symbol and pattern, got %v", err)
	}
}

func TestFilterPatternAgainstSymbol_RootSymbolInPattern(t *testing.T) {
	// Test when pattern contains the Root symbol
	err := validate.PatternAgainstSymbol(string(def.Root), symbols.String("validSymbol"))
	if err == nil {
		t.Fatalf("Expected error for pattern containing the Root symbol, but got none")
	}
}

func TestFilterPatternAgainstSymbol_PlaceholderInSymbolAndPattern(t *testing.T) {
	// Test with placeholders in both symbol and pattern
	err := validate.PatternAgainstSymbol(def.Placeholder.String()+"a", def.Placeholder.Append(symbols.String("b")))
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
}

func TestFilterPatternAgainstSymbol_PlaceholderMissingInPattern(t *testing.T) {
	// Test when Placeholder exists in the symbol but not in the pattern
	err := validate.PatternAgainstSymbol("ab", def.Placeholder.Append(symbols.String("c")))
	if err == nil {
		t.Fatalf("Expected error for missing Placeholder in the pattern, but got none")
	}
}

func TestFilterPatternAgainstSymbol_ConsecutivePlaceholders(t *testing.T) {
	// Test when placeholders appear consecutively in the symbol and pattern
	err := validate.PatternAgainstSymbol(def.Placeholder.Append(def.Placeholder).String(), def.Placeholder.Append(def.Placeholder))
	if err == nil {
		t.Fatalf("Expected error for consecutive placeholders, but got none")
	}
}

func TestFilterPatternAgainstSymbol_ValidPattern(t *testing.T) {
	// Test with a valid pattern and symbol
	err := validate.PatternAgainstSymbol("a"+string(def.Placeholder)+"b", def.Placeholder.Format(symbols.String("x%sy"), symbols.String("%s")))
	if err != nil {
		t.Fatalf("Expected nil for valid pattern, got %v", err)
	}
}
