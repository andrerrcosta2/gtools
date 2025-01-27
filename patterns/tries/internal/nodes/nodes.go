// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nodes

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/errs"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"strings"
	"sync"
)

type PatternTrie interface {
	str.TrieNode[string, grammar.Symbol, PatternTrie]
	AddChild(key string, child PatternTrie) error
	GetChild(key string) (PatternTrie, bool)
	HasChildren() bool
	Key() string
	RemoveChild(key string)
	SetParent(parent PatternTrie) error
	PrintNode(level int) string
	die() error
}

var _ str.TrieNode[string, grammar.Symbol, PatternTrie] = (PatternTrie)(nil)

// Transition creates a new transitionNode with the given key and parent.
func Transition(key string, parent PatternTrie) PatternTrie {
	return &transitionNode{
		parent:   parent,
		key:      key,
		children: make(map[string]PatternTrie),
	}
}

type transitionNode struct {
	mtx      sync.RWMutex
	inner    sync.RWMutex // Fine-grained lock for inner nodes
	key      string
	parent   PatternTrie            // Parent node
	children map[string]PatternTrie // Map of child nodes (character or Placeholder key)
}

func (p *transitionNode) Value() grammar.Symbol {
	return nil
}

func (p *transitionNode) Children() []PatternTrie {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	children := make([]PatternTrie, 0, len(p.children))
	for _, child := range p.children {
		children = append(children, child)
	}

	return children
}

func (p *transitionNode) SetParent(parent PatternTrie) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.parent = parent
	return nil
}

// Parent returns the parent node
func (p *transitionNode) Parent() PatternTrie {
	p.mtx.RLock()
	p.mtx.RUnlock()
	return p.parent
}

func (p *transitionNode) HasChildren() bool {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return len(p.children) > 0
}

// Key returns the key
func (p *transitionNode) Key() string {
	return p.key
}

// ChildrenKeys returns the keys of the children
func (p *transitionNode) ChildrenKeys() []string {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	keys := make([]string, 0, len(p.children))
	for key := range p.children {
		keys = append(keys, key)
	}

	return keys
}

// ChildrenValues returns the values of the children if the child is a symbol
func (p *transitionNode) ChildrenValues() []grammar.Symbol {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	values := make([]grammar.Symbol, 0, len(p.children))
	for _, child := range p.children {
		if c, ok := child.(*symbolNode); ok {
			values = append(values, c.symbol)
		}
	}
	return values
}

// GetChild returns the children by key
func (p *transitionNode) GetChild(r string) (PatternTrie, bool) {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	child, exists := p.children[r]
	return child, exists
}

// AddChild adds a child
func (p *transitionNode) AddChild(r string, child PatternTrie) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	// Reject double placeholders
	if StartsWithPlaceholder(r) && StartsWithPlaceholder(p.key) {
		return errs.PlaceholderChildOfPlaceholderWithKey(p.key)
	}
	p.children[r] = child
	return nil
}

// RemoveChild removes a child
func (p *transitionNode) RemoveChild(r string) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	delete(p.children, r)
}

func (p *transitionNode) String() string {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.PrintNode(0)
}

// PrintNode for transitionNode returns a string representing the hierarchical structure of the node.
func (p *transitionNode) PrintNode(level int) string {
	var result strings.Builder
	indent := strings.Repeat("  ", level)

	if p.key == def.Root.String() {
		result.WriteString(fmt.Sprintf(" { key: [root-key], type: root, children: ["))
	} else {
		result.WriteString(fmt.Sprintf("\n%s{ key: %s, type: transition, children: [", indent, p.key))
	}

	for _, child := range p.children {
		//result.WriteString(fmt.Sprintf("%s  key: { %s\n", indent, key))
		result.WriteString(child.PrintNode(level + 1))
	}

	if len(p.children) > 0 {
		result.WriteString(fmt.Sprintf("\n%s]}", indent))
	} else {
		result.WriteString("]}")
	}

	return result.String()
}

func (p *transitionNode) die() error {
	p.parent.RemoveChild(p.key)
	p.children = nil
	return nil
}

var _ str.TrieNode[string, grammar.Symbol, PatternTrie] = (*transitionNode)(nil)
var _ PatternTrie = (*transitionNode)(nil)

// Symbol creates a new symbol node with the given key, parent, and symbol.
func Symbol(key string, parent PatternTrie, symbol grammar.Symbol) PatternTrie {
	return &symbolNode{
		key:      key,                          // The key of the node
		parent:   parent,                       // The parent node of this node
		symbol:   symbol,                       // The symbol associated with this node
		children: make(map[string]PatternTrie), // The children of this node
	}
}

type symbolNode struct {
	mtx      sync.RWMutex
	inner    sync.RWMutex // Fine-grained lock for inner nodes
	key      string
	parent   PatternTrie
	symbol   grammar.Symbol
	children map[string]PatternTrie
	port     symbols.Port
}

// Value returns the symbol
func (p *symbolNode) Value() grammar.Symbol {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.symbol
}

func (p *symbolNode) HasChildren() bool {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return len(p.children) > 0
}

// Children return the direct children nodes and its recursive children nodes
func (p *symbolNode) Children() []PatternTrie {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	children := make([]PatternTrie, 0, len(p.children))

	for _, child := range p.children {
		children = append(children, child)
	}

	return children
}

func (p *symbolNode) SetParent(parent PatternTrie) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.parent = parent
	return nil
}

// Parent returns the parent node
func (p *symbolNode) Parent() PatternTrie {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.parent
}

// Key returns the key
// Doesn't need to lock it since the key never changes, and it is created within the own object just once.
func (p *symbolNode) Key() string {
	return p.key // A key shouldn't ever change
}

// ChildrenKeys returns the keys of the children
func (p *symbolNode) ChildrenKeys() []string {
	p.mtx.RLock()
	defer p.mtx.RUnlock()

	keys := make([]string, 0, len(p.children))
	for k := range p.children {
		keys = append(keys, k)
	}

	return keys
}

// ChildrenValues returns the values of the children
// Must return the symbol itself as the first element
func (p *symbolNode) ChildrenValues() []grammar.Symbol {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	values := []grammar.Symbol{p.symbol}
	for _, child := range p.children {
		if c, ok := child.(*symbolNode); ok {
			values = append(values, c.symbol)
		}
	}
	return values
}

func (p *symbolNode) GetChild(s string) (PatternTrie, bool) {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	child, exists := p.children[s]
	return child, exists
}

func (p *symbolNode) AddChild(r string, child PatternTrie) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	// Reject double placeholders
	if StartsWithPlaceholder(r) && StartsWithPlaceholder(p.key) {
		return errs.PlaceholderChildOfPlaceholderWithKey(p.key)
	}
	p.children[r] = child

	return nil
}

func (p *symbolNode) RemoveChild(r string) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	delete(p.children, r)
}

func (p *symbolNode) String() string {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	return p.PrintNode(0)
}

// PrintNode for symbolNode returns a string representing the hierarchical structure of the node.
func (p *symbolNode) PrintNode(level int) string {
	var result strings.Builder
	indent := strings.Repeat("  ", level)
	result.WriteString(fmt.Sprintf("\n%s{ key: %s, type: symbol: (%s), children: [", indent, p.key, p.symbol.String()))

	for _, child := range p.children {
		//result.WriteString(fmt.Sprintf("%s  Child: { %s\n", indent, key))
		result.WriteString(child.PrintNode(level + 1))
	}

	if len(p.children) > 0 {
		result.WriteString(fmt.Sprintf("\n%s]}", indent))
	} else {
		result.WriteString("]}")
	}

	return result.String()
}

func (p *symbolNode) die() error {
	p.parent.RemoveChild(p.key)
	p.children = nil
	p.symbol = nil
	return nil
}

var _ str.TrieNode[string, grammar.Symbol, PatternTrie] = (*symbolNode)(nil)
var _ PatternTrie = (*symbolNode)(nil)

func NewSearch(root PatternTrie, pattern string, output *[]grammar.Symbol) *Search {
	return &Search{
		mtx:     &sync.Mutex{},
		root:    root,
		ref:     root,
		pattern: pattern,
		symbol:  grammar.ByteSymbol{},
		output:  output,
	}
}

type Search struct {
	mtx      *sync.Mutex
	root     PatternTrie       // Root of the trie
	ref      PatternTrie       // Current reference point
	pattern  string            // Current pattern being traversed
	symbol   grammar.Symbol    // Current symbol
	output   *[]grammar.Symbol // Accumulated matches
	maxChars int               // Limit on character repetition
	count    int               // Tracks current repetitions
}

func (s *Search) Clone() *Search {
	return &Search{
		mtx:     s.mtx,
		root:    s.root,
		ref:     s.ref,
		pattern: s.pattern,
		symbol:  s.symbol,
		output:  s.output,
	}
}

func (s *Search) IsEmpty() bool {
	return len(s.pattern) == 0
}

func (s *Search) IsSymbol() bool {
	return IsSymbol(s.ref)
}

func (s *Search) Length() int {
	return len(s.pattern)
}

// Nmlsn returns the next matching literal symbol node in the pattern char by char and if it exists
func (s *Search) Nmlsn() (node *Search, exists bool) {
	length := s.Length()
	iter := s.ref

	for i := 0; i < length; i++ {
		fmt.Printf("next matching literal symbol node: %s\n", string(s.pattern[i]))
		iter, exists = iter.GetChild(string(s.pattern[i]))
		fmt.Printf("exists: %v\n", exists)
		if !exists {
			return nil, false
		}
		if ss, ok := asSymbol(iter); ok {
			fmt.Printf("Found symbol: %s\n", ss.symbol.String())
			var root PatternTrie
			if ss.port == symbols.Breakpoint {
				root = iter
			} else {
				root = s.root
			}
			node = &Search{
				mtx:     s.mtx,
				symbol:  s.symbol.Append(ss.symbol),
				output:  s.output,
				pattern: s.pattern[i+1:],
				root:    root,
				ref:     iter,
			}
			fmt.Printf("Returning node: %+v\n", node)
			return
		}
	}

	return nil, false
}

// Nmpn returns the next matching placeholder node in the pattern char by char and if it exists
func (s *Search) Nmpn() (*Search, bool) {
	length := s.Length()
	iter := s.ref

	// I think this is wrong. i think only symbols should allow placeholders.
	for i := 0; i < length; i++ {
		check, exists := iter.GetChild(string(s.pattern[i]))
		if exists {
			if ss, ok := asSymbol(iter); ok {
				var root PatternTrie
				if ss.port == symbols.Breakpoint {
					root = iter
				} else {
					root = s.root
				}
				return &Search{
					mtx:     s.mtx,
					symbol:  s.symbol.Append(ss.symbol),
					output:  s.output,
					pattern: s.pattern[i+1:],
					root:    root,
					ref:     iter,
				}, true
			} else {
				iter = check
			}
		}
	}

	return nil, false
}

func (s *Search) Pattern() string {
	return s.pattern
}

func (s *Search) Port() symbols.Port {
	if ss, ok := asSymbol(s.ref); ok {
		return ss.port
	}
	return symbols.ChildrenOnly
}

func (s *Search) Ref() PatternTrie {
	return s.ref
}

func (s *Search) Symbol() grammar.Symbol {
	return s.symbol
}

func (s *Search) Write() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	*s.output = append(*s.output, s.symbol)
}
