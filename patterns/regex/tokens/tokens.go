// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tokens

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

const PH = '\x02'
const PHStr = `\x02`
const GROUP = '\x03'

// General general tokens
type General symbols.Logical

func gen(port symbols.Port, s string, tags ...string) General {
	return symbols.LogicalOf(grammar.ByteSymbol(s), port, tags...)
}

var (
	// AnyChar : .
	// Matches any single character except newline
	AnyChar = gen(symbols.Open, ".") // General{ByteSymbol: []byte{'.'}}

	// NewLineChar : \n r
	// Matches newline character
	NewLineChar = gen(symbols.RecursiveOnly, "\\n") // General{ByteSymbol: []byte{'\n'}}

	// TabChar : \t
	// Matches tab character
	TabChar = gen(symbols.RecursiveOnly, "\\t") // General{ByteSymbol: []byte{'\t'}}

	// CarriageReturn : \r
	// Matches carriage return character
	CarriageReturn = gen(symbols.RecursiveOnly, "\\r") // General{ByteSymbol: []byte{'\r'}}

	// StartOfLine : ^
	// Matches the start of a line
	// let text = "Is this his";
	// let pattern = /^Is/g;
	StartOfLine = gen(symbols.RecursiveOnly, "^") // General{ByteSymbol: []byte{'^'}}

	// EndOfLine : $
	// Matches the end of a line
	EndOfLine = gen(symbols.Open, "$") // General{ByteSymbol: []byte{'$'}}

	// LiteralBackslash : \\
	// Represents a literal backslash character
	LiteralBackslash = gen(symbols.RecursiveOnly, "\\") // General{ByteSymbol: []byte{'\\'}}

	// Or : |
	// Matches either the expression before or after it
	Or = gen(symbols.RecursiveOnly, "|") // General{ByteSymbol: []byte{'|'}}

	// ExactCount : {ph}
	// Matches exactly {ph} occurrences of the preceding element
	// Use 'ph' for the count value
	ExactCount = gen(symbols.RecursiveOnly, fmt.Sprintf("{%s}", PHStr)) // General{ByteSymbol: []byte{'{', PH, '}'}}

	// RangeCount : {ph,ph}
	// Matches between {ph} and {ph} occurrences of the preceding element
	// Use 'ph' for the min and max count values
	RangeCount = gen(symbols.RecursiveOnly, fmt.Sprintf("{%s,%s}", PHStr, PHStr)) // General{ByteSymbol: []byte{'{', PH, ',', PH, '}'}}
)

// Quantifier Quantifiers
type Quantifier symbols.Logical

func qnt(port symbols.Port, s string, tags ...string) Quantifier {
	return symbols.LogicalOf(grammar.ByteSymbol(s), port, tags...)
}

var (
	// ZeroOrMore : *
	// Matches 0 or more of the previous element
	ZeroOrMore = qnt(symbols.RecursiveOnlyQt, "*") // Quantifier{ByteSymbol: []byte{'*'}}

	// OneOrMore : +
	// Matches 1 or more of the previous element
	OneOrMore = qnt(symbols.RecursiveOnlyQt, "+") // Quantifier{ByteSymbol: []byte{'+'}}

	// ZeroOrOne : ?
	// Matches 0 or 1 of the previous element
	ZeroOrOne = qnt(symbols.RecursiveOnly, "?") // Quantifier{ByteSymbol: []byte{'?'}}
)

// CharClass Character Classes
type CharClass symbols.Logical

func char(port symbols.Port, s string, tags ...string) CharClass {
	return symbols.LogicalOf(grammar.ByteSymbol(s), port, tags...)
}

var (
	// AnyWhitespaceChar : \s
	// Matches any whitespace character
	AnyWhitespaceChar = char(symbols.RecursiveOnly, "\\s") // CharClass{ByteSymbol: []byte{'\\', 's'}}

	// AnyNonWhitespaceChar : \S
	// Matches any non-whitespace character
	AnyNonWhitespaceChar = char(symbols.RecursiveOnly, "\\S") // CharClass{ByteSymbol: []byte{'\\', 'S'}}

	// AnyWordChar : \w
	// Matches any word character (alphanumeric and underscore)
	AnyWordChar = char(symbols.RecursiveOnly, "\\w") // CharClass{ByteSymbol: []byte{'\\', 'w'}}

	// AnyNonWordChar : \W
	// Matches any non-word character
	AnyNonWordChar = char(symbols.RecursiveOnly, "\\W") // CharClass{ByteSymbol: []byte{'\\', 'W'}}

	// AnyDigit : \d
	// Matches any digit
	AnyDigit = char(symbols.RecursiveOnly, "\\d") // CharClass{ByteSymbol: []byte{'\\', 'd'}}

	// AnyNonDigit : \D
	// Matches any non-digit character
	AnyNonDigit = char(symbols.RecursiveOnly, "\\D") // CharClass{ByteSymbol: []byte{'\\', 'D'}}

	// WordBoundaryEnds : {WORD}\b
	// Matches a word boundary that ends with a word boundary
	WordBoundaryEnds = char(symbols.RecursiveOnly, fmt.Sprintf("%s\\b", PHStr)) // CharClass{ByteSymbol: []byte{WORD, '\\', 'b'}}

	// WordBoundaryStarts : \b{WORD}
	// Matches a word that starts with a word boundary
	WordBoundaryStarts = char(symbols.RecursiveOnly, fmt.Sprintf("\\b%s", PHStr)) // CharClass{ByteSymbol: []byte{'\\', 'b', WORD}}

	// WordBoundaryNotEnds : {WORD}\B
	// Return the first position where it is present, NOT in the end of a word
	WordBoundaryNotEnds = char(symbols.RecursiveOnly, fmt.Sprintf("%s\\B", PHStr)) // CharClass{ByteSymbol: []byte{WORD, '\\', 'B'}}

	// WordBoundaryNotStarts : \B{WORD}
	// Return the first position where it is present, NOT in the start of a word
	WordBoundaryNotStarts = char(symbols.RecursiveOnly, fmt.Sprintf("\\B%s", PHStr)) // CharClass{ByteSymbol: []byte{'\\', 'B', WORD}}
)

// SpecialChar Special Characters
type SpecialChar symbols.Logical

func spc(port symbols.Port, s string, tags ...string) SpecialChar {
	return symbols.LogicalOf(grammar.ByteSymbol(s), port, tags...)
}

var (
	// Backspace : \b ([\b]
	// Matches a backspace character when it is inside a []
	Backspace = spc(symbols.RecursiveOnly, "[\\b]") // SpecialChar{ByteSymbol: []byte{'\\', 'b'}}

	// ControlChar : \c
	// Placeholder for control characters
	ControlChar = spc(symbols.RecursiveOnly, "\\c") // SpecialChar{ByteSymbol: []byte{'\\', 'c'}}

	// NullChar : \0
	// Matches a null character
	NullChar = spc(symbols.RecursiveOnly, "\\0") // SpecialChar{ByteSymbol: []byte{'\\', '0'}}

	// VerticalTab : \v
	// Matches a vertical tab character
	VerticalTab = spc(symbols.RecursiveOnly, "\\v") // SpecialChar{ByteSymbol: []byte{'\\', 'v'}}

	// OctalChar : \0
	// Matches octal character; requires dynamic value
	OctalChar = spc(symbols.RecursiveOnly, "\\O") // SpecialChar{ByteSymbol: []byte{'\\', 'O'}}

	// HexChar : \x
	// Matches hex character; requires dynamic value
	HexChar = spc(symbols.RecursiveOnly, "\\x") // SpecialChar{ByteSymbol: []byte{'\\', 'x'}}

	// UnicodeChar : \u
	// Matches unicode character; requires dynamic value
	UnicodeChar = spc(symbols.RecursiveOnly, "\\u") // SpecialChar{ByteSymbol: []byte{'\\', 'u'}}

	// LineBreak : \r
	// Matches a line break character
	LineBreak = spc(symbols.RecursiveOnly, "\\r") // SpecialChar{ByteSymbol: []byte{'\\', 'r'}}
)

// Group Grouping
type Group symbols.Logical

func group(port symbols.Port, s string, tags ...string) Group {
	return symbols.LogicalOf(grammar.ByteSymbol(s), port, tags...)
}

var (
	// CapturingGroup : (ph)
	// Use 'ph' with placeholders for content inside
	CapturingGroup = group(symbols.RecursiveOnly, fmt.Sprintf("(%s)", PHStr)) // Group{ByteSymbol: []byte{'(', PH, ')')

	// NonCapturingGroup : (?:ph)
	// Use 'ph' with placeholders for content inside
	NonCapturingGroup = group(symbols.RecursiveOnly, fmt.Sprintf("(?:%s)", PHStr)) // Group{ByteSymbol: []byte{'(', '?', ':', PH, ')'}}

	// NamedCapturingGroup : (?P<ph>ph)
	// 'ph' represents the placeholder for the group name and content
	NamedCapturingGroup = group(symbols.RecursiveOnly, fmt.Sprintf("(?P<%s>%s)", PHStr, PHStr)) // Group{ByteSymbol: []byte{"'(', '?', 'P', '<', PH, '>', PH, ')')

	// CharacterSet : [ph]
	// Use 'ph' with placeholders for content inside
	CharacterSet = group(symbols.RecursiveOnly, fmt.Sprintf("[%s]", PHStr)) // Group{ByteSymbol: []byte{'[', PH, ']') // Group{ByteSymbol: []byte{'[', PH, ']'}}

	// NegatedCharacterSet : [^ph]
	// Use 'ph' with placeholders for content inside
	NegatedCharacterSet = group(symbols.RecursiveOnly, fmt.Sprintf("[^%s]", PHStr)) // Group{ByteSymbol: []byte{'[', '^', PH, ']') // Group{ByteSymbol: []byte{'[', '^', PH, ']'}}

	// CharacterRange : {ph,ph}
	// Use 'ph' with placeholders for content inside
	CharacterRange = group(symbols.RecursiveOnly, fmt.Sprintf("{%s,%s}", PHStr, PHStr)) // Group{ByteSymbol: []byte{'{', PH, ',', PH, '}') // Group{ByteSymbol: []byte{'{', PH, ',', PH, '}'}}

	// CharacterClass : \p{ph}
	// Use 'ph' with placeholders for content inside
	CharacterClass = group(symbols.RecursiveOnly, fmt.Sprintf("\\p{%s}", PHStr)) // Group{ByteSymbol: []byte{'\\', 'p', '{', PH, '}') // Group{ByteSymbol: []byte{'\\', 'p', '{', PH, '}'}}

	// Alternation : (ph|ph)
	// Use 'ph' with placeholders for content inside
	Alternation = group(symbols.RecursiveOnly, fmt.Sprintf("(%s|%s)", PHStr, PHStr)) // Group{ByteSymbol: []byte{"'(', PH, '|', PH, ')') // Group{ByteSymbol: []byte{'(', PH, '|', PH, ')'}}
)

// StringReplacement String Replacements
type StringReplacement symbols.Logical

func str(port symbols.Port, s string, tags ...string) StringReplacement {
	return symbols.LogicalOf(grammar.ByteSymbol(s), port, tags...)
}

var (
	// InsertBeforeMatchedString : $`
	// Insert before the matched string
	InsertBeforeMatchedString = str(symbols.RecursiveOnly, "$`") // StringReplacement{ByteSymbol: []byte("$`")}

	// InsertAfterMatchedString : $'
	// Insert after the matched string
	InsertAfterMatchedString = str(symbols.RecursiveOnly, "$'") // StringReplacement{ByteSymbol: []byte("$'")}

	// InsertLastMatched : $+
	// Insert last matched string
	InsertLastMatched = str(symbols.RecursiveOnly, "$+") // StringReplacement{ByteSymbol: []byte("$+")}

	// InsertEntireMatch : $&
	// Insert the entire match
	InsertEntireMatch = str(symbols.RecursiveOnly, "$&") // StringReplacement{ByteSymbol: []byte("$&")}

	// InsertNthCaptureGroup : $n
	// Insert the nth capture group
	InsertNthCaptureGroup = str(symbols.RecursiveOnly, "$n") // StringReplacement{ByteSymbol: []byte("$n")}
)

// Assertion Assertions
type Assertion symbols.Logical

func ass(port symbols.Port, s string, tags ...string) Assertion {
	return symbols.LogicalOf(grammar.ByteSymbol(s), port, tags...)
}

var (
	// PositiveLookahead : (?=content)
	// Use placeholder for content inside the lookahead
	PositiveLookahead = ass(symbols.RecursiveOnly, fmt.Sprintf("(?=%s)", PHStr)) // Assertion{ByteSymbol: []byte{'(', '?', '=', PH, ')'}}

	// NegativeLookahead : (?!content)
	// Use placeholder for content inside the lookahead
	NegativeLookahead = ass(symbols.RecursiveOnly, fmt.Sprintf("(?!%s)", PHStr)) // Assertion{ByteSymbol: []byte{'(', '?', '!', PH, ')') // Assertion{ByteSymbol: []byte{'(', '?', '!', PH, ')'}}

	// PositiveLookbehind : (?<=content)
	// Use placeholder for content inside the lookbehind
	PositiveLookbehind = ass(symbols.RecursiveOnly, fmt.Sprintf("(?<=%s)", PHStr)) // Assertion{ByteSymbol: []byte{'(', '?', '<', '=', PH, ')') // Assertion{ByteSymbol: []byte{'(', '?', '<', '=', PH, ')'}}

	// NegativeLookbehind : (?<!content)
	// Use placeholder for content inside the lookbehind
	NegativeLookbehind = ass(symbols.RecursiveOnly, fmt.Sprintf("(?<!%s)", PHStr)) // Assertion{ByteSymbol: []byte{'(', '?', '<', '!', PH, ')') // Assertion{ByteSymbol: []byte{'(', '?', '<', '!', PH, ')'}}
)
