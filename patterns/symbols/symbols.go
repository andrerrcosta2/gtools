// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package symbols

import (
	"github.com/andrerrcosta2/gtools/core/domain/data"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
)

type Port uint8

const (
	RecursiveOnly    Port = iota // Recursion only
	Open                         // Open to children and to recursion
	ChildrenOnly                 // Children only
	RecursiveToSelf              // Recursive to self only
	Breakpoint                   // turns into root
	ChildrenNegative             // Children negative
	ChildrenExact                // Children
	RecursiveOnlyQt              // exact
	OpenQt
	ChildrenOnlyQt
	Closed // Dead end
)

func OpenOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   Open,
		tags:   tags,
	}
}

func OpenQtOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   OpenQt,
		tags:   tags,
	}
}

func ClosedOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   Closed,
		tags:   tags,
	}
}

func RcvOnlyOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   RecursiveOnly,
		tags:   tags,
	}
}

func RecOnlyQtOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   RecursiveOnlyQt,
		tags:   tags,
	}
}

func ChdOnlyOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   ChildrenOnly,
		tags:   tags,
	}
}

func ChdOnlyQtOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   ChildrenOnlyQt,
		tags:   tags,
	}
}

func ChdExactOf(symbol grammar.Symbol, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   ChildrenExact,
		tags:   tags,
	}
}

func LogicalOf(symbol grammar.Symbol, port Port, tags ...string) Logical {
	return logical{
		Symbol: symbol,
		port:   port,
		tags:   tags,
	}
}

type Logical interface {
	grammar.Symbol
	data.Taggable[string]
	Port() Port
}

type logical struct {
	grammar.Symbol
	tags []string
	port Port
}

func (s logical) Port() Port {
	return s.port
}

func (s logical) Tag(tags ...string) {
	s.tags = append(s.tags, tags...)
}

func (s logical) Tags() []string {
	return s.tags
}
