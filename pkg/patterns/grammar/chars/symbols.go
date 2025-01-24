// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chars

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

type Digit symbols.Logical

func digit(s byte, tags ...string) Digit {
	return symbols.LogicalOf(grammar.ByteSymbol([]byte{s}), symbols.Open, false, tags...)
}

var (
	Zero  = digit('0', "WORD", "DIGIT", "ALPHANUMERIC")
	One   = digit('1', "WORD", "DIGIT", "ALPHANUMERIC")
	Two   = digit('2', "WORD", "DIGIT", "ALPHANUMERIC")
	Three = digit('3', "WORD", "DIGIT", "ALPHANUMERIC")
	Four  = digit('4', "WORD", "DIGIT", "ALPHANUMERIC")
	Five  = digit('5', "WORD", "DIGIT", "ALPHANUMERIC")
	Six   = digit('6', "WORD", "DIGIT", "ALPHANUMERIC")
	Seven = digit('7', "WORD", "DIGIT", "ALPHANUMERIC")
	Eight = digit('8', "WORD", "DIGIT", "ALPHANUMERIC")
	Nine  = digit('9', "WORD", "DIGIT", "ALPHANUMERIC")
)

type Alphabet symbols.Logical

func letter(s byte, tags ...string) Alphabet {
	return symbols.LogicalOf(grammar.ByteSymbol([]byte{s}), symbols.Open, false, tags...)
}

var (
	ALower = letter('a', "WORD", "LETTER", "ALPHANUMERIC")
	BLower = letter('b', "WORD", "LETTER", "ALPHANUMERIC")
	CLower = letter('c', "WORD", "LETTER", "ALPHANUMERIC")
	DLower = letter('d', "WORD", "LETTER", "ALPHANUMERIC")
	ELower = letter('e', "WORD", "LETTER", "ALPHANUMERIC")
	FLower = letter('f', "WORD", "LETTER", "ALPHANUMERIC")
	GLower = letter('g', "WORD", "LETTER", "ALPHANUMERIC")
	HLower = letter('h', "WORD", "LETTER", "ALPHANUMERIC")
	ILower = letter('i', "WORD", "LETTER", "ALPHANUMERIC")
	JLower = letter('j', "WORD", "LETTER", "ALPHANUMERIC")
	KLower = letter('k', "WORD", "LETTER", "ALPHANUMERIC")
	LLower = letter('l', "WORD", "LETTER", "ALPHANUMERIC")
	MLower = letter('m', "WORD", "LETTER", "ALPHANUMERIC")
	NLower = letter('n', "WORD", "LETTER", "ALPHANUMERIC")
	OLower = letter('o', "WORD", "LETTER", "ALPHANUMERIC")
	PLower = letter('p', "WORD", "LETTER", "ALPHANUMERIC")
	QLower = letter('q', "WORD", "LETTER", "ALPHANUMERIC")
	RLower = letter('r', "WORD", "LETTER", "ALPHANUMERIC")
	SLower = letter('s', "WORD", "LETTER", "ALPHANUMERIC")
	TLower = letter('t', "WORD", "LETTER", "ALPHANUMERIC")
	ULower = letter('u', "WORD", "LETTER", "ALPHANUMERIC")
	VLower = letter('v', "WORD", "LETTER", "ALPHANUMERIC")
	WLower = letter('w', "WORD", "LETTER", "ALPHANUMERIC")
	XLower = letter('x', "WORD", "LETTER", "ALPHANUMERIC")
	YLower = letter('y', "WORD", "LETTER", "ALPHANUMERIC")
	ZLower = letter('z', "WORD", "LETTER", "ALPHANUMERIC")
	AUpper = letter('A', "WORD", "LETTER", "ALPHANUMERIC")
	BUpper = letter('B', "WORD", "LETTER", "ALPHANUMERIC")
	CUpper = letter('C', "WORD", "LETTER", "ALPHANUMERIC")
	DUpper = letter('D', "WORD", "LETTER", "ALPHANUMERIC")
	EUpper = letter('E', "WORD", "LETTER", "ALPHANUMERIC")
	FUpper = letter('F', "WORD", "LETTER", "ALPHANUMERIC")
	GUpper = letter('G', "WORD", "LETTER", "ALPHANUMERIC")
	HUpper = letter('H', "WORD", "LETTER", "ALPHANUMERIC")
	IUpper = letter('I', "WORD", "LETTER", "ALPHANUMERIC")
	JUpper = letter('J', "WORD", "LETTER", "ALPHANUMERIC")
	KUpper = letter('K', "WORD", "LETTER", "ALPHANUMERIC")
	LUpper = letter('L', "WORD", "LETTER", "ALPHANUMERIC")
	MUpper = letter('M', "WORD", "LETTER", "ALPHANUMERIC")
	NUpper = letter('N', "WORD", "LETTER", "ALPHANUMERIC")
	OUpper = letter('O', "WORD", "LETTER", "ALPHANUMERIC")
	PUpper = letter('P', "WORD", "LETTER", "ALPHANUMERIC")
	QUpper = letter('Q', "WORD", "LETTER", "ALPHANUMERIC")
	RUpper = letter('R', "WORD", "LETTER", "ALPHANUMERIC")
	SUpper = letter('S', "WORD", "LETTER", "ALPHANUMERIC")
	TUpper = letter('T', "WORD", "LETTER", "ALPHANUMERIC")
	UUpper = letter('U', "WORD", "LETTER", "ALPHANUMERIC")
	VUpper = letter('V', "WORD", "LETTER", "ALPHANUMERIC")
	WUpper = letter('W', "WORD", "LETTER", "ALPHANUMERIC")
	XUpper = letter('X', "WORD", "LETTER", "ALPHANUMERIC")
	YUpper = letter('Y', "WORD", "LETTER", "ALPHANUMERIC")
	ZUpper = letter('Z', "WORD", "LETTER", "ALPHANUMERIC")
)

type Punctuation symbols.Logical

func punctuation(s byte, tags ...string) Punctuation {
	return symbols.LogicalOf(grammar.ByteSymbol([]byte{s}), symbols.Open, false, tags...)
}

var (
	Comma       = punctuation(',', "WORD", "PUNCTUATION")
	Dot         = punctuation('.', "WORD", "PUNCTUATION")
	Colon       = punctuation(':', "WORD", "PUNCTUATION")
	Semicolon   = punctuation(';', "WORD", "PUNCTUATION")
	Question    = punctuation('?', "WORD", "PUNCTUATION")
	Exclamation = punctuation('!', "WORD", "PUNCTUATION")
	AtSign      = punctuation('@', "WORD", "PUNCTUATION")
	Underscore  = punctuation('_', "WORD", "PUNCTUATION")
	Hash        = punctuation('#', "WORD", "PUNCTUATION")
	Dollar      = punctuation('$', "WORD", "PUNCTUATION")
	Percent     = punctuation('%', "WORD", "PUNCTUATION")
	Caret       = punctuation('^', "WORD", "PUNCTUATION")
	Ampersand   = punctuation('&', "WORD", "PUNCTUATION")
	Star        = punctuation('*', "WORD", "PUNCTUATION")
	Parenthesis = punctuation('(', "WORD", "PUNCTUATION")
	Brace       = punctuation('{', "WORD", "PUNCTUATION")
	Bracket     = punctuation('[', "WORD", "PUNCTUATION")
)
