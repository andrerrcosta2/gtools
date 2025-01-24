// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
)

// This is the simplest dictionary that can be used to build a OfPatterns
// It does not contain any placeholder neither collapses nodes
var linearOpenSymbolsDictionary = &grammar.SymbolDictionary[string, symbols.Logical]{
	"abc": symbols.OpenOf(grammar.ByteSymbol("abc"), false),
	"def": symbols.OpenOf(grammar.ByteSymbol("def"), false),
	"ghi": symbols.OpenOf(grammar.ByteSymbol("ghi"), false),
	"jkl": symbols.OpenOf(grammar.ByteSymbol("jkl"), false),
	"mno": symbols.OpenOf(grammar.ByteSymbol("mno"), false),
	"pqr": symbols.OpenOf(grammar.ByteSymbol("pqr"), false),
	"stu": symbols.OpenOf(grammar.ByteSymbol("stu"), false),
	"vwx": symbols.OpenOf(grammar.ByteSymbol("vwx"), false),
	"yzA": symbols.OpenOf(grammar.ByteSymbol("yzA"), false),
	"BCD": symbols.OpenOf(grammar.ByteSymbol("BCD"), false),
}

var nonLinearOpenSymbolsDictionary = &grammar.SymbolDictionary[string, symbols.Logical]{
	"abc":     symbols.OpenOf(grammar.ByteSymbol("sumbol1"), false),
	"acb":     symbols.OpenOf(grammar.ByteSymbol("sumbol2"), false),
	"bac":     symbols.OpenOf(grammar.ByteSymbol("sumbol3"), false),
	"bca":     symbols.OpenOf(grammar.ByteSymbol("sumbol4"), false),
	"cda":     symbols.OpenOf(grammar.ByteSymbol("sumbol5"), false),
	"cab":     symbols.OpenOf(grammar.ByteSymbol("sumbol6"), false),
	"cba":     symbols.OpenOf(grammar.ByteSymbol("symbol7"), false),
	"abcdefg": symbols.OpenOf(grammar.ByteSymbol("symbol8"), false),
}

var smallRootChildrenDictionary = &grammar.SymbolDictionary[string, symbols.Logical]{
	"a": symbols.OpenOf(grammar.ByteSymbol("abcd"), false),
	"b": symbols.OpenOf(grammar.ByteSymbol("bcde"), false),
	"c": symbols.OpenOf(grammar.ByteSymbol("cdef"), false),
}

var dictionaryWithPlaceholders = &grammar.SymbolDictionary[string, symbols.Logical]{
	"abc":                                   symbols.OpenOf(grammar.ByteSymbol("abc"), false),
	"def[" + def.Placeholder.String() + "]": symbols.OpenOf(symbols.Of("def").Append(def.Placeholder), false),
	"ghi":                                   symbols.OpenOf(grammar.ByteSymbol("ghi"), false),
	"jklm" + def.Placeholder.String() + "m": symbols.OpenOf(symbols.Of("jklM").Append(def.Placeholder), false),
	"pqr":                                   symbols.OpenOf(grammar.ByteSymbol("pqr"), false),
	"stu":                                   symbols.OpenOf(grammar.ByteSymbol("stu"), false),
	"vwx":                                   symbols.OpenOf(grammar.ByteSymbol("vwx"), false),
	"yz" + def.Placeholder.String() + "gp":  symbols.OpenOf(symbols.Of("yz").Append(def.Placeholder).Append(symbols.Of("gp")), false),
}

var invalidDictionary = &grammar.SymbolDictionary[string, symbols.Logical]{
	"abc": symbols.OpenOf(grammar.ByteSymbol("abc").Append(def.Placeholder), false),
	"def": symbols.OpenOf(grammar.ByteSymbol("def"), false),
	"ghi": symbols.OpenOf(grammar.ByteSymbol("ghi"), false),
	"jkl": symbols.OpenOf(grammar.ByteSymbol("jkl"), false),
	"mno": symbols.OpenOf(grammar.ByteSymbol("mno"), false),
	"pqr": symbols.OpenOf(grammar.ByteSymbol("pqr"), false),
	"stu": symbols.OpenOf(grammar.ByteSymbol("stu").Append(def.Root), false),
	"vwx": symbols.OpenOf(grammar.ByteSymbol("vwx"), false),
	"yz":  symbols.OpenOf(grammar.ByteSymbol("yz"), false),
}

var cyclicDictionary = &grammar.SymbolDictionary[string, symbols.Logical]{
	"abcdefg": symbols.OpenOf(grammar.ByteSymbol("abcdefg"), false),
	"a":       symbols.OpenOf(grammar.ByteSymbol("a"), false),
	"ab":      symbols.OpenOf(grammar.ByteSymbol("ab"), false),
	"abc":     symbols.OpenOf(grammar.ByteSymbol("abc"), false),
	"abcd":    symbols.OpenOf(grammar.ByteSymbol("abcd"), false),
	"abcde":   symbols.OpenOf(grammar.ByteSymbol("abcde"), false),
	"abcdef":  symbols.OpenOf(grammar.ByteSymbol("abcdef"), false),
	"b":       symbols.OpenOf(grammar.ByteSymbol("b"), false),
	"bc":      symbols.OpenOf(grammar.ByteSymbol("bc"), false),
	"bcd":     symbols.OpenOf(grammar.ByteSymbol("bcd"), false),
	"bcde":    symbols.OpenOf(grammar.ByteSymbol("bcde"), false),
	"bcdef":   symbols.OpenOf(grammar.ByteSymbol("bcdef"), false),
	"bcdefg":  symbols.OpenOf(grammar.ByteSymbol("bcdefg"), false),
	"c":       symbols.OpenOf(grammar.ByteSymbol("c"), false),
	"cd":      symbols.OpenOf(grammar.ByteSymbol("cd"), false),
	"cde":     symbols.OpenOf(grammar.ByteSymbol("cde"), false),
	"cdef":    symbols.OpenOf(grammar.ByteSymbol("cdef"), false),
	"cdefg":   symbols.OpenOf(grammar.ByteSymbol("cdefg"), false),
	"d":       symbols.OpenOf(grammar.ByteSymbol("d"), false),
	"de":      symbols.OpenOf(grammar.ByteSymbol("de"), false),
	"def":     symbols.OpenOf(grammar.ByteSymbol("def"), false),
	"defg":    symbols.OpenOf(grammar.ByteSymbol("defg"), false),
	"e":       symbols.OpenOf(grammar.ByteSymbol("e"), false),
	"ef":      symbols.OpenOf(grammar.ByteSymbol("ef"), false),
	"efg":     symbols.OpenOf(grammar.ByteSymbol("efg"), false),
	"f":       symbols.OpenOf(grammar.ByteSymbol("f"), false),
	"fg":      symbols.OpenOf(grammar.ByteSymbol("fg"), false),
	"g":       symbols.OpenOf(grammar.ByteSymbol("g"), false),
}

// edgeValidDictionary can be used to test the full capabilities of the trie
// It includes all the resources and all the edge cases a valid dictionary can have.
var edgeValidDictionary = &grammar.SymbolDictionary[string, symbols.Logical]{
	"abcde":    symbols.OpenOf(grammar.ByteSymbol("abcde"), false),
	"abcdef":   symbols.OpenOf(grammar.ByteSymbol("abcdef"), false),
	"abcdefg":  symbols.OpenOf(grammar.ByteSymbol("abcdefg"), false),
	"abcdefgh": symbols.OpenOf(grammar.ByteSymbol("abcdefgh"), false),
	symbols.OfPlaceholder("abcdefg[¨%s]hi", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("one1[%s]", def.Placeholder.String()), false),
	"abcdefghij":   symbols.OpenOf(grammar.ByteSymbol("abcdefghij"), false),
	"abcdefghijk":  symbols.OpenOf(grammar.ByteSymbol("abcdefghijk"), false),
	"abcdefghijkl": symbols.OpenOf(grammar.ByteSymbol("abcdefghijkl"), false),
	symbols.OfPlaceholder("abc[¨%s]defghijklm", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("one2[%s]", def.Placeholder.String()), false),
	"abcdefghijklmn":  symbols.OpenOf(grammar.ByteSymbol("abcdefghijklmn"), false),
	"abcdefghijklmno": symbols.OpenOf(grammar.ByteSymbol("abcdefghijklmno"), false),
	"bcdef":           symbols.OpenOf(grammar.ByteSymbol("bcdef"), false),
	"bcdefg":          symbols.OpenOf(grammar.ByteSymbol("bcdefg"), false),
	"bcdefgh":         symbols.OpenOf(grammar.ByteSymbol("bcdefgh"), false),
	"bcdefghi":        symbols.OpenOf(grammar.ByteSymbol("bcdefghi"), false),
	symbols.OfPlaceholder("bcdef[%s]gh[%s]ij", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("two1[%s]-[%s]", def.Placeholder.String()), false),
	"bcdefghijk":   symbols.OpenOf(grammar.ByteSymbol("bcdefghijk"), false),
	"bcdefghijkl":  symbols.OpenOf(grammar.ByteSymbol("bcdefghijkl"), false),
	"bcdefghijklm": symbols.OpenOf(grammar.ByteSymbol("bcdefghijklm"), false),
	symbols.OfPlaceholder("bcdefgh[%s]i[%s]j", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("two2[%s]-[%s]", def.Placeholder.String()), false),
	"bcdefghijklmno": symbols.OpenOf(grammar.ByteSymbol("bcdefghijklmno"), false),
	"cdefg":          symbols.OpenOf(grammar.ByteSymbol("cdefg"), false),
	"cdefgh":         symbols.OpenOf(grammar.ByteSymbol("cdefgh"), false),
	"cdefghi":        symbols.OpenOf(grammar.ByteSymbol("cdefghi"), false),
	"cdefghij":       symbols.OpenOf(grammar.ByteSymbol("cdefghij"), false),
	"cdefghijk":      symbols.OpenOf(grammar.ByteSymbol("cdefghijk"), false),
	symbols.OfPlaceholder("cdefg[%s]hi[%s]jkl", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("two2[%s]-[%s]", def.Placeholder.String()), false),
	"cdefghijklm":   symbols.OpenOf(grammar.ByteSymbol("cdefghijklm"), false),
	"cdefghijklmn":  symbols.OpenOf(grammar.ByteSymbol("cdefghijklmn"), false),
	"cdefghijklmno": symbols.OpenOf(grammar.ByteSymbol("cdefghijklmno"), false),
	"defgh":         symbols.OpenOf(grammar.ByteSymbol("defgh"), false),
	"defghi":        symbols.OpenOf(grammar.ByteSymbol("defghi"), false),
	"defghij":       symbols.OpenOf(grammar.ByteSymbol("defghij"), false),
	"defghijk":      symbols.OpenOf(grammar.ByteSymbol("defghijk"), false),
	"defghijkl":     symbols.OpenOf(grammar.ByteSymbol("defghijkl"), false),
	"defghijklm":    symbols.OpenOf(grammar.ByteSymbol("defghijklm"), false),
	symbols.OfPlaceholder("defg[%s]h[%s]ijk{%s}lm", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("three1[%s]-[%s]-[%s]", def.Placeholder.String()), false),
	"defghijklmno": symbols.OpenOf(grammar.ByteSymbol("defghijklmno"), false),
	symbols.OfPlaceholder("e[%s]fghi", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("one3[%s]", def.Placeholder.String()), false),
	"efghij":   symbols.OpenOf(grammar.ByteSymbol("efghij"), false),
	"efghijk":  symbols.OpenOf(grammar.ByteSymbol("efghijk"), false),
	"efghijkl": symbols.OpenOf(grammar.ByteSymbol("efghijkl"), false),
	symbols.OfPlaceholder("efgh[%s]ijklm", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("one4[%s]", def.Placeholder.String()), false),
	"efghijklmn":  symbols.OpenOf(grammar.ByteSymbol("efghijklmn"), false),
	"efghijklmno": symbols.OpenOf(grammar.ByteSymbol("efghijklmno"), false),
	"fghij":       symbols.OpenOf(grammar.ByteSymbol("fghij"), false),
	"fghijk":      symbols.OpenOf(grammar.ByteSymbol("fghijk"), false),
	"fghijkl":     symbols.OpenOf(grammar.ByteSymbol("fghijkl"), false),
	"fghijklm":    symbols.OpenOf(grammar.ByteSymbol("fghijklm"), false),
	"fghijklmn":   symbols.OpenOf(grammar.ByteSymbol("fghijklmn"), false),
	"fghijklmno":  symbols.OpenOf(grammar.ByteSymbol("fghijklmno"), false),
	symbols.OfPlaceholder("ghijk[¨%s]", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("one5[%s]", def.Placeholder.String()), false),
	"ghijkl":    symbols.OpenOf(grammar.ByteSymbol("ghijkl"), false),
	"ghijklm":   symbols.OpenOf(grammar.ByteSymbol("ghijklm"), false),
	"ghijklmn":  symbols.OpenOf(grammar.ByteSymbol("ghijklmn"), false),
	"ghijklmno": symbols.OpenOf(grammar.ByteSymbol("ghijklmno"), false),
	"hijk":      symbols.OpenOf(grammar.ByteSymbol("hijk"), false),
	"hijkl":     symbols.OpenOf(grammar.ByteSymbol("hijkl"), false),
	"hijklm":    symbols.OpenOf(grammar.ByteSymbol("hijklm"), false),
	"hijklmn":   symbols.OpenOf(grammar.ByteSymbol("hijklmn"), false),
	symbols.OfPlaceholder("hi|%s|jk[%s]lmno{%s}", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("three2[%s]-[%s]-[%s]", def.Placeholder.String()), false),
	"ijklm":   symbols.OpenOf(grammar.ByteSymbol("ijklm"), false),
	"ijklmn":  symbols.OpenOf(grammar.ByteSymbol("ijklmn"), false),
	"ijklmno": symbols.OpenOf(grammar.ByteSymbol("ijklmno"), false),
	"jklmn":   symbols.OpenOf(grammar.ByteSymbol("jklmn"), false),
	symbols.OfPlaceholder("jkl[%s]mn[%s]o", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("two3[%s]-[%s]", def.Placeholder.String()), false),
	"klmno": symbols.OpenOf(grammar.ByteSymbol("klmno"), false),
	"lmno":  symbols.OpenOf(grammar.ByteSymbol("lmno"), false),
	"mno":   symbols.OpenOf(grammar.ByteSymbol("mno"), false),
	"no":    symbols.OpenOf(grammar.ByteSymbol("no"), false),
	symbols.OfPlaceholder("a{%s}bcdefghijklmno{%s}", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("two4[%s]-[%s]", def.Placeholder.String()), false),
	"b": symbols.OpenOf(grammar.ByteSymbol("b"), false),
	"c": symbols.OpenOf(grammar.ByteSymbol("c"), false),
	"d": symbols.OpenOf(grammar.ByteSymbol("d"), false),
	// two symbols with placeholder after 'e['
	symbols.OfPlaceholder("e[%s]", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("one6[%s]", def.Placeholder.String()), false),
	"o": symbols.OpenOf(grammar.ByteSymbol("o"), false),
	"p": symbols.OpenOf(grammar.ByteSymbol("p"), false),
	"q": symbols.OpenOf(grammar.ByteSymbol("q"), false),
	"r": symbols.OpenOf(grammar.ByteSymbol("r"), false),
	"s": symbols.OpenOf(grammar.ByteSymbol("s"), false),
	"t": symbols.OpenOf(grammar.ByteSymbol("t"), false),
	"u": symbols.OpenOf(grammar.ByteSymbol("u"), false),
	"v": symbols.OpenOf(grammar.ByteSymbol("v"), false),
	symbols.OfPlaceholder("w[%s]", def.Placeholder.String()).String(): symbols.OpenOf(symbols.OfPlaceholder("one7[%s]", def.Placeholder.String()), false),
	"x": symbols.OpenOf(grammar.ByteSymbol("x"), false),
	"y": symbols.OpenOf(grammar.ByteSymbol("y"), false),
	"z": symbols.OpenOf(grammar.ByteSymbol("z"), false),
}

var smallEdgeDictionary = &grammar.SymbolDictionary[string, symbols.Logical]{
	// Simple transition to symbol conversion
	"abc":  symbols.OpenOf(grammar.ByteSymbol("abc"), false),
	"abcd": symbols.OpenOf(grammar.ByteSymbol("abcd"), false), // Extends from "abc", should trigger transition to symbol

	// Placeholder patterns with overlaps
	def.Placeholder.Format(symbols.Of("abc[%s]def"), symbols.Repl).String(): symbols.OpenOf(def.Placeholder.Format(symbols.Of("symbol1[%s]"), symbols.Repl), false),
	def.Placeholder.Format(symbols.Of("abc[%s]d"), symbols.Repl).String():   symbols.OpenOf(def.Placeholder.Format(symbols.Of("symbol2[%s]"), symbols.Repl), false), // Overlaps with "abc[%s]def"

	// More basic symbols
	"a":     symbols.OpenOf(grammar.ByteSymbol("a"), false), // Root child
	"ab":    symbols.OpenOf(grammar.ByteSymbol("ab"), false),
	"abcde": symbols.OpenOf(grammar.ByteSymbol("abcde"), false), // Extends from "abcd"
}

var rootNode = nodes.Transition(def.Root.String(), nil)

var trieBranch = func(chars ...string) nodes.PatternTrie {
	branch := rootNode
	var parent = branch

	for i := 0; i < len(chars); i++ {
		if symbols.Of(chars[i:]).StartsWith(def.Placeholder) {
			if symbols.Of(chars[i:]).Equal(def.Placeholder) {
				_ = parent.AddChild(def.Placeholder.String(), nodes.Symbol(def.Placeholder.String(), parent, def.Placeholder))
				return branch
			}
			_ = parent.AddChild(def.Placeholder.String(), nodes.Transition(def.Placeholder.String(), parent))
			parent, _ = parent.GetChild(def.Placeholder.String())
			continue
		}

		if i == len(chars)-1 {
			_ = parent.AddChild(chars[i], nodes.Symbol(chars[i], parent, symbols.OpenOf(grammar.ByteSymbol(chars[i]), false)))
			return branch
		}

		_ = parent.AddChild(chars[i], nodes.Transition(chars[i], parent))
		parent, _ = parent.GetChild(chars[i])
	}
	return branch
}
