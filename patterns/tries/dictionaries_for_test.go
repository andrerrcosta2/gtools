// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
)

type K = string
type S = symbols.Logical

// This is the simplest dictionary that can be used to build a OfPatterns
// It does not contain any placeholder neither collapses nodes
var linearOpenSymbolsDictionary = &grammar.SymbolDictionary[K, []S, S]{
	"abc": []S{symbols.OpenOf(grammar.ByteSymbol("abc"))},
	"def": []S{symbols.OpenOf(grammar.ByteSymbol("def"))},
	"ghi": []S{symbols.OpenOf(grammar.ByteSymbol("ghi"))},
	"jkl": []S{symbols.OpenOf(grammar.ByteSymbol("jkl"))},
	"mno": []S{symbols.OpenOf(grammar.ByteSymbol("mno"))},
	"pqr": []S{symbols.OpenOf(grammar.ByteSymbol("pqr"))},
	"stu": []S{symbols.OpenOf(grammar.ByteSymbol("stu"))},
	"vwx": []S{symbols.OpenOf(grammar.ByteSymbol("vwx"))},
	"yzA": []S{symbols.OpenOf(grammar.ByteSymbol("yzA"))},
	"BCD": []S{symbols.OpenOf(grammar.ByteSymbol("BCD"))},
}

var nonLinearOpenSymbolsDictionary = &grammar.SymbolDictionary[K, []S, S]{
	"abc":     []S{symbols.OpenOf(grammar.ByteSymbol("sumbol1"))},
	"acb":     []S{symbols.OpenOf(grammar.ByteSymbol("sumbol2"))},
	"bac":     []S{symbols.OpenOf(grammar.ByteSymbol("sumbol3"))},
	"bca":     []S{symbols.OpenOf(grammar.ByteSymbol("sumbol4"))},
	"cda":     []S{symbols.OpenOf(grammar.ByteSymbol("sumbol5"))},
	"cab":     []S{symbols.OpenOf(grammar.ByteSymbol("sumbol6"))},
	"cba":     []S{symbols.OpenOf(grammar.ByteSymbol("symbol7"))},
	"abcdefg": []S{symbols.OpenOf(grammar.ByteSymbol("symbol8"))},
}

var smallRootChildrenDictionary = &grammar.SymbolDictionary[K, []S, S]{
	"a": []S{symbols.OpenOf(grammar.ByteSymbol("abcd"))},
	"b": []S{symbols.OpenOf(grammar.ByteSymbol("bcde"))},
	"c": []S{symbols.OpenOf(grammar.ByteSymbol("cdef"))},
}

var dictionaryWithPlaceholders = &grammar.SymbolDictionary[K, []S, S]{
	"abc":                                   []S{symbols.OpenOf(grammar.ByteSymbol("abc"))},
	"def[" + def.Placeholder.String() + "]": []S{symbols.OpenOf(symbols.Of("def").Append(def.Placeholder))},
	"ghi":                                   []S{symbols.OpenOf(grammar.ByteSymbol("ghi"))},
	"jklm" + def.Placeholder.String() + "m": []S{symbols.OpenOf(symbols.Of("jklM").Append(def.Placeholder))},
	"pqr":                                   []S{symbols.OpenOf(grammar.ByteSymbol("pqr"))},
	"stu":                                   []S{symbols.OpenOf(grammar.ByteSymbol("stu"))},
	"vwx":                                   []S{symbols.OpenOf(grammar.ByteSymbol("vwx"))},
	"yz" + def.Placeholder.String() + "gp":  []S{symbols.OpenOf(symbols.Of("yz").Append(def.Placeholder).Append(symbols.Of("gp")))},
}

var invalidDictionary = &grammar.SymbolDictionary[K, []S, S]{
	"abc": []S{symbols.OpenOf(grammar.ByteSymbol("abc").Append(def.Placeholder))},
	"def": []S{symbols.OpenOf(grammar.ByteSymbol("def"))},
	"ghi": []S{symbols.OpenOf(grammar.ByteSymbol("ghi"))},
	"jkl": []S{symbols.OpenOf(grammar.ByteSymbol("jkl"))},
	"mno": []S{symbols.OpenOf(grammar.ByteSymbol("mno"))},
	"pqr": []S{symbols.OpenOf(grammar.ByteSymbol("pqr"))},
	"stu": []S{symbols.OpenOf(grammar.ByteSymbol("stu").Append(def.Root))},
	"vwx": []S{symbols.OpenOf(grammar.ByteSymbol("vwx"))},
	"yz":  []S{symbols.OpenOf(grammar.ByteSymbol("yz"))},
}

var cyclicDictionary = &grammar.SymbolDictionary[K, []S, S]{
	"abcdefg": []S{symbols.OpenOf(grammar.ByteSymbol("abcdefg"))},
	"a":       []S{symbols.OpenOf(grammar.ByteSymbol("a"))},
	"ab":      []S{symbols.OpenOf(grammar.ByteSymbol("ab"))},
	"abc":     []S{symbols.OpenOf(grammar.ByteSymbol("abc"))},
	"abcd":    []S{symbols.OpenOf(grammar.ByteSymbol("abcd"))},
	"abcde":   []S{symbols.OpenOf(grammar.ByteSymbol("abcde"))},
	"abcdef":  []S{symbols.OpenOf(grammar.ByteSymbol("abcdef"))},
	"b":       []S{symbols.OpenOf(grammar.ByteSymbol("b"))},
	"bc":      []S{symbols.OpenOf(grammar.ByteSymbol("bc"))},
	"bcd":     []S{symbols.OpenOf(grammar.ByteSymbol("bcd"))},
	"bcde":    []S{symbols.OpenOf(grammar.ByteSymbol("bcde"))},
	"bcdef":   []S{symbols.OpenOf(grammar.ByteSymbol("bcdef"))},
	"bcdefg":  []S{symbols.OpenOf(grammar.ByteSymbol("bcdefg"))},
	"c":       []S{symbols.OpenOf(grammar.ByteSymbol("c"))},
	"cd":      []S{symbols.OpenOf(grammar.ByteSymbol("cd"))},
	"cde":     []S{symbols.OpenOf(grammar.ByteSymbol("cde"))},
	"cdef":    []S{symbols.OpenOf(grammar.ByteSymbol("cdef"))},
	"cdefg":   []S{symbols.OpenOf(grammar.ByteSymbol("cdefg"))},
	"d":       []S{symbols.OpenOf(grammar.ByteSymbol("d"))},
	"de":      []S{symbols.OpenOf(grammar.ByteSymbol("de"))},
	"def":     []S{symbols.OpenOf(grammar.ByteSymbol("def"))},
	"defg":    []S{symbols.OpenOf(grammar.ByteSymbol("defg"))},
	"e":       []S{symbols.OpenOf(grammar.ByteSymbol("e"))},
	"ef":      []S{symbols.OpenOf(grammar.ByteSymbol("ef"))},
	"efg":     []S{symbols.OpenOf(grammar.ByteSymbol("efg"))},
	"f":       []S{symbols.OpenOf(grammar.ByteSymbol("f"))},
	"fg":      []S{symbols.OpenOf(grammar.ByteSymbol("fg"))},
	"g":       []S{symbols.OpenOf(grammar.ByteSymbol("g"))},
}

// edgeValidDictionary can be used to test the full capabilities of the trie
// It includes all the resources and all the edge cases a valid dictionary can have.
var edgeValidDictionary = &grammar.SymbolDictionary[K, []S, S]{
	"abcde":    []S{symbols.OpenOf(grammar.ByteSymbol("abcde"))},
	"abcdef":   []S{symbols.OpenOf(grammar.ByteSymbol("abcdef"))},
	"abcdefg":  []S{symbols.OpenOf(grammar.ByteSymbol("abcdefg"))},
	"abcdefgh": []S{symbols.OpenOf(grammar.ByteSymbol("abcdefgh"))},
	symbols.OfPlaceholder("abcdefg[¨%s]hi", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("one1[%s]", def.Placeholder.String()))},
	"abcdefghij":   []S{symbols.OpenOf(grammar.ByteSymbol("abcdefghij"))},
	"abcdefghijk":  []S{symbols.OpenOf(grammar.ByteSymbol("abcdefghijk"))},
	"abcdefghijkl": []S{symbols.OpenOf(grammar.ByteSymbol("abcdefghijkl"))},
	symbols.OfPlaceholder("abc[¨%s]defghijklm", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("one2[%s]", def.Placeholder.String()))},
	"abcdefghijklmn":  []S{symbols.OpenOf(grammar.ByteSymbol("abcdefghijklmn"))},
	"abcdefghijklmno": []S{symbols.OpenOf(grammar.ByteSymbol("abcdefghijklmno"))},
	"bcdef":           []S{symbols.OpenOf(grammar.ByteSymbol("bcdef"))},
	"bcdefg":          []S{symbols.OpenOf(grammar.ByteSymbol("bcdefg"))},
	"bcdefgh":         []S{symbols.OpenOf(grammar.ByteSymbol("bcdefgh"))},
	"bcdefghi":        []S{symbols.OpenOf(grammar.ByteSymbol("bcdefghi"))},
	symbols.OfPlaceholder("bcdef[%s]gh[%s]ij", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("two1[%s]-[%s]", def.Placeholder.String()))},
	"bcdefghijk":   []S{symbols.OpenOf(grammar.ByteSymbol("bcdefghijk"))},
	"bcdefghijkl":  []S{symbols.OpenOf(grammar.ByteSymbol("bcdefghijkl"))},
	"bcdefghijklm": []S{symbols.OpenOf(grammar.ByteSymbol("bcdefghijklm"))},
	symbols.OfPlaceholder("bcdefgh[%s]i[%s]j", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("two2[%s]-[%s]", def.Placeholder.String()))},
	"bcdefghijklmno": []S{symbols.OpenOf(grammar.ByteSymbol("bcdefghijklmno"))},
	"cdefg":          []S{symbols.OpenOf(grammar.ByteSymbol("cdefg"))},
	"cdefgh":         []S{symbols.OpenOf(grammar.ByteSymbol("cdefgh"))},
	"cdefghi":        []S{symbols.OpenOf(grammar.ByteSymbol("cdefghi"))},
	"cdefghij":       []S{symbols.OpenOf(grammar.ByteSymbol("cdefghij"))},
	"cdefghijk":      []S{symbols.OpenOf(grammar.ByteSymbol("cdefghijk"))},
	symbols.OfPlaceholder("cdefg[%s]hi[%s]jkl", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("two2[%s]-[%s]", def.Placeholder.String()))},
	"cdefghijklm":   []S{symbols.OpenOf(grammar.ByteSymbol("cdefghijklm"))},
	"cdefghijklmn":  []S{symbols.OpenOf(grammar.ByteSymbol("cdefghijklmn"))},
	"cdefghijklmno": []S{symbols.OpenOf(grammar.ByteSymbol("cdefghijklmno"))},
	"defgh":         []S{symbols.OpenOf(grammar.ByteSymbol("defgh"))},
	"defghi":        []S{symbols.OpenOf(grammar.ByteSymbol("defghi"))},
	"defghij":       []S{symbols.OpenOf(grammar.ByteSymbol("defghij"))},
	"defghijk":      []S{symbols.OpenOf(grammar.ByteSymbol("defghijk"))},
	"defghijkl":     []S{symbols.OpenOf(grammar.ByteSymbol("defghijkl"))},
	"defghijklm":    []S{symbols.OpenOf(grammar.ByteSymbol("defghijklm"))},
	symbols.OfPlaceholder("defg[%s]h[%s]ijk{%s}lm", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("three1[%s]-[%s]-[%s]", def.Placeholder.String()))},
	"defghijklmno": []S{symbols.OpenOf(grammar.ByteSymbol("defghijklmno"))},
	symbols.OfPlaceholder("e[%s]fghi", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("one3[%s]", def.Placeholder.String()))},
	"efghij":   []S{symbols.OpenOf(grammar.ByteSymbol("efghij"))},
	"efghijk":  []S{symbols.OpenOf(grammar.ByteSymbol("efghijk"))},
	"efghijkl": []S{symbols.OpenOf(grammar.ByteSymbol("efghijkl"))},
	symbols.OfPlaceholder("efgh[%s]ijklm", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("one4[%s]", def.Placeholder.String()))},
	"efghijklmn":  []S{symbols.OpenOf(grammar.ByteSymbol("efghijklmn"))},
	"efghijklmno": []S{symbols.OpenOf(grammar.ByteSymbol("efghijklmno"))},
	"fghij":       []S{symbols.OpenOf(grammar.ByteSymbol("fghij"))},
	"fghijk":      []S{symbols.OpenOf(grammar.ByteSymbol("fghijk"))},
	"fghijkl":     []S{symbols.OpenOf(grammar.ByteSymbol("fghijkl"))},
	"fghijklm":    []S{symbols.OpenOf(grammar.ByteSymbol("fghijklm"))},
	"fghijklmn":   []S{symbols.OpenOf(grammar.ByteSymbol("fghijklmn"))},
	"fghijklmno":  []S{symbols.OpenOf(grammar.ByteSymbol("fghijklmno"))},
	symbols.OfPlaceholder("ghijk[¨%s]", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("one5[%s]", def.Placeholder.String()))},
	"ghijkl":    []S{symbols.OpenOf(grammar.ByteSymbol("ghijkl"))},
	"ghijklm":   []S{symbols.OpenOf(grammar.ByteSymbol("ghijklm"))},
	"ghijklmn":  []S{symbols.OpenOf(grammar.ByteSymbol("ghijklmn"))},
	"ghijklmno": []S{symbols.OpenOf(grammar.ByteSymbol("ghijklmno"))},
	"hijk":      []S{symbols.OpenOf(grammar.ByteSymbol("hijk"))},
	"hijkl":     []S{symbols.OpenOf(grammar.ByteSymbol("hijkl"))},
	"hijklm":    []S{symbols.OpenOf(grammar.ByteSymbol("hijklm"))},
	"hijklmn":   []S{symbols.OpenOf(grammar.ByteSymbol("hijklmn"))},
	symbols.OfPlaceholder("hi|%s|jk[%s]lmno{%s}", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("three2[%s]-[%s]-[%s]", def.Placeholder.String()))},
	"ijklm":   []S{symbols.OpenOf(grammar.ByteSymbol("ijklm"))},
	"ijklmn":  []S{symbols.OpenOf(grammar.ByteSymbol("ijklmn"))},
	"ijklmno": []S{symbols.OpenOf(grammar.ByteSymbol("ijklmno"))},
	"jklmn":   []S{symbols.OpenOf(grammar.ByteSymbol("jklmn"))},
	symbols.OfPlaceholder("jkl[%s]mn[%s]o", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("two3[%s]-[%s]", def.Placeholder.String()))},
	"klmno": []S{symbols.OpenOf(grammar.ByteSymbol("klmno"))},
	"lmno":  []S{symbols.OpenOf(grammar.ByteSymbol("lmno"))},
	"mno":   []S{symbols.OpenOf(grammar.ByteSymbol("mno"))},
	"no":    []S{symbols.OpenOf(grammar.ByteSymbol("no"))},
	symbols.OfPlaceholder("a{%s}bcdefghijklmno{%s}", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("two4[%s]-[%s]", def.Placeholder.String()))},
	"b": []S{symbols.OpenOf(grammar.ByteSymbol("b"))},
	"c": []S{symbols.OpenOf(grammar.ByteSymbol("c"))},
	"d": []S{symbols.OpenOf(grammar.ByteSymbol("d"))},
	// two symbols with placeholder after 'e['
	symbols.OfPlaceholder("e[%s]", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("one6[%s]", def.Placeholder.String()))},
	"o": []S{symbols.OpenOf(grammar.ByteSymbol("o"))},
	"p": []S{symbols.OpenOf(grammar.ByteSymbol("p"))},
	"q": []S{symbols.OpenOf(grammar.ByteSymbol("q"))},
	"r": []S{symbols.OpenOf(grammar.ByteSymbol("r"))},
	"s": []S{symbols.OpenOf(grammar.ByteSymbol("s"))},
	"t": []S{symbols.OpenOf(grammar.ByteSymbol("t"))},
	"u": []S{symbols.OpenOf(grammar.ByteSymbol("u"))},
	"v": []S{symbols.OpenOf(grammar.ByteSymbol("v"))},
	symbols.OfPlaceholder("w[%s]", def.Placeholder.String()).String(): []S{symbols.OpenOf(symbols.OfPlaceholder("one7[%s]", def.Placeholder.String()))},
	"x": []S{symbols.OpenOf(grammar.ByteSymbol("x"))},
	"y": []S{symbols.OpenOf(grammar.ByteSymbol("y"))},
	"z": []S{symbols.OpenOf(grammar.ByteSymbol("z"))},
}

var smallEdgeDictionary = &grammar.SymbolDictionary[K, []S, S]{
	// Simple transition to symbol conversion
	"abc":  []S{symbols.OpenOf(grammar.ByteSymbol("abc"))},
	"abcd": []S{symbols.OpenOf(grammar.ByteSymbol("abcd"))}, // Extends from "abc", should trigger transition to symbol

	// Placeholder patterns with overlaps
	def.Placeholder.Format(symbols.Of("abc[%s]def"), symbols.Repl).String(): []S{symbols.OpenOf(def.Placeholder.Format(symbols.Of("symbol1[%s]"), symbols.Repl))},
	def.Placeholder.Format(symbols.Of("abc[%s]d"), symbols.Repl).String():   []S{symbols.OpenOf(def.Placeholder.Format(symbols.Of("symbol2[%s]"), symbols.Repl))}, // Overlaps with "abc[%s]def"

	// More basic symbols
	"a":     []S{symbols.OpenOf(grammar.ByteSymbol("a"))}, // Root child
	"ab":    []S{symbols.OpenOf(grammar.ByteSymbol("ab"))},
	"abcde": []S{symbols.OpenOf(grammar.ByteSymbol("abcde"))}, // Extends from "abcd"
}

var rootNode = nodes.Transition(def.Root.String(), nil)

// I THINK THIS IS DEPRECATED
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
			_ = parent.AddChild(chars[i], nodes.Symbol(chars[i], parent, symbols.OpenOf(grammar.ByteSymbol(chars[i]))))
			return branch
		}

		_ = parent.AddChild(chars[i], nodes.Transition(chars[i], parent))
		parent, _ = parent.GetChild(chars[i])
	}
	return branch
}
