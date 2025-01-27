// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tokens

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

// Regexp is a dictionary of logical regular expression symbols.
var Regexp = &grammar.SymbolDictionary[string, []symbols.Logical, symbols.Logical]{
	// A few symbols must be declared as root as key and children only as logical port

	// Same Keys
	PHStr: []symbols.Logical{
		// General
		AnyChar, StartOfLine, EndOfLine,
		// Quantifiers
		ZeroOrMore, OneOrMore, ZeroOrOne,
	},

	LiteralBackslash.String(): []symbols.Logical{
		// General
		LiteralBackslash,
	},

	// General
	NewLineChar.String():    []symbols.Logical{NewLineChar},
	TabChar.String():        []symbols.Logical{TabChar},
	CarriageReturn.String(): []symbols.Logical{CarriageReturn},
	Or.String():             []symbols.Logical{Or},

	// Quantifiers
	ExactCount.String(): []symbols.Logical{ExactCount},
	RangeCount.String(): []symbols.Logical{RangeCount},

	// Character Classes
	AnyWhitespaceChar.String():     []symbols.Logical{AnyWhitespaceChar},
	AnyNonWhitespaceChar.String():  []symbols.Logical{AnyNonWhitespaceChar},
	AnyWordChar.String():           []symbols.Logical{AnyWordChar},
	AnyNonWordChar.String():        []symbols.Logical{AnyNonWordChar},
	AnyDigit.String():              []symbols.Logical{AnyDigit},
	AnyNonDigit.String():           []symbols.Logical{AnyNonDigit},
	WordBoundaryEnds.String():      []symbols.Logical{WordBoundaryEnds},
	WordBoundaryStarts.String():    []symbols.Logical{WordBoundaryStarts},
	WordBoundaryNotEnds.String():   []symbols.Logical{WordBoundaryNotEnds},
	WordBoundaryNotStarts.String(): []symbols.Logical{WordBoundaryNotStarts},
	ControlChar.String():           []symbols.Logical{ControlChar},
	NullChar.String():              []symbols.Logical{NullChar},
	VerticalTab.String():           []symbols.Logical{VerticalTab},
	HexChar.String():               []symbols.Logical{HexChar},
	CharacterClass.String():        []symbols.Logical{CharacterClass},
	Alternation.String():           []symbols.Logical{Alternation},

	// Special
	InsertBeforeMatchedString.String(): []symbols.Logical{InsertBeforeMatchedString},
	InsertAfterMatchedString.String():  []symbols.Logical{InsertAfterMatchedString},
	InsertLastMatched.String():         []symbols.Logical{InsertLastMatched},
	InsertEntireMatch.String():         []symbols.Logical{InsertEntireMatch},
	InsertNthCaptureGroup.String():     []symbols.Logical{InsertNthCaptureGroup},
	Backspace.String():                 []symbols.Logical{Backspace},
	OctalChar.String():                 []symbols.Logical{OctalChar},
	UnicodeChar.String():               []symbols.Logical{UnicodeChar},
	LineBreak.String():                 []symbols.Logical{LineBreak},

	// Group
	CapturingGroup.String():      []symbols.Logical{CapturingGroup},
	NonCapturingGroup.String():   []symbols.Logical{NonCapturingGroup},
	NamedCapturingGroup.String(): []symbols.Logical{NamedCapturingGroup},
	NegatedCharacterSet.String(): []symbols.Logical{NegatedCharacterSet},
	CharacterRange.String():      []symbols.Logical{CharacterRange},
	CharacterSet.String():        []symbols.Logical{CharacterSet},

	// Assertions
	PositiveLookahead.String():  []symbols.Logical{PositiveLookahead},
	NegativeLookahead.String():  []symbols.Logical{NegativeLookahead},
	PositiveLookbehind.String(): []symbols.Logical{PositiveLookbehind},
	NegativeLookbehind.String(): []symbols.Logical{NegativeLookbehind},

	// Quanfiers
	ZeroOrMore.String(): []symbols.Logical{ZeroOrMore},
	OneOrMore.String():  []symbols.Logical{OneOrMore},
	ZeroOrOne.String():  []symbols.Logical{ZeroOrOne},
}
