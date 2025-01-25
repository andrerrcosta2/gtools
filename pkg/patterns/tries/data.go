// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/iterators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
	"sync"
)

func newDag(root *patternDagNode) *patternDag {
	dag := &patternDag{
		nodes: make(map[*patternDagNode]struct{}),
		edges: make(map[*patternDagNode][]*patternDagNode),
	}
	dag.nodes[root] = struct{}{}
	return dag
}

type patternDag struct {
	mtx   sync.RWMutex
	nodes map[*patternDagNode]struct{}
	edges map[*patternDagNode][]*patternDagNode // from -> [to...]
}

// AddNode adds a new node to the DAG
func (dag *patternDag) AddNode(node *patternDagNode) {
	dag.mtx.Lock()
	defer dag.mtx.Unlock()
	dag.nodes[node] = struct{}{}
}

// HasNode checks if a node exists in the DAG
func (dag *patternDag) HasNode(node *patternDagNode) bool {
	dag.mtx.RLock()
	defer dag.mtx.RUnlock()
	_, exists := dag.nodes[node]
	return exists
}

// AddEdge adds a single edge from one node to another
func (dag *patternDag) AddEdge(from, to *patternDagNode) {
	dag.mtx.Lock()
	defer dag.mtx.Unlock()
	if !dag.HasNode(from) || !dag.HasNode(to) {
		return // Only allow adding edges between existing nodes
	}
	dag.edges[from] = append(dag.edges[from], to) // Single edge (from -> to)
}

// HasEdge checks if an edge exists between two nodes
func (dag *patternDag) HasEdge(from, to *patternDagNode) bool {
	dag.mtx.RLock()
	defer dag.mtx.RUnlock()
	if neighbors, exists := dag.edges[from]; exists {
		for _, neighbor := range neighbors {
			if neighbor == to {
				return true
			}
		}
	}
	return false
}

// Nodes returns a list of all nodes in the DAG
func (dag *patternDag) Nodes() []*patternDagNode {
	dag.mtx.RLock()
	defer dag.mtx.RUnlock()
	keys := make([]*patternDagNode, 0, len(dag.nodes))
	for node := range dag.nodes {
		keys = append(keys, node)
	}
	return keys
}

// Neighbors returns the neighbor of a node (only one child, as it's single-edged)
func (dag *patternDag) Neighbors(node *patternDagNode) []*patternDagNode {
	dag.mtx.RLock()
	defer dag.mtx.RUnlock()
	if neighbors, exists := dag.edges[node]; exists {
		return neighbors
	}
	return []*patternDagNode{}
}

var _ str.Graph[*patternDagNode] = (*patternDag)(nil)

func newDagNode(trieNode nodes.PatternTrie, pattern string) *patternDagNode {
	return &patternDagNode{
		PatternTrie: trieNode,
		pattern:     pattern,
	}
}

type patternDagNode struct {
	nodes.PatternTrie
	pattern string
}

func (n *patternDagNode) Equal(other any) bool {
	return n.pattern == other.(*patternDagNode).pattern
}

func (n *patternDagNode) Less(other any) bool {
	return n.pattern < other.(*patternDagNode).pattern
}

var _ gtools.SortableOf = (*patternDagNode)(nil)

// Entry creates a new patternTrieEntry with the given key and value.
//
// key: The key for the patternTrieEntry.
// value: The value for the patternTrieEntry.
// Returns a pointer to the newly created patternTrieEntry.
func Entry(key string, value symbols.Logical) str.Entry[string, symbols.Logical] {
	return &patternTrieEntry{
		key:   key,
		value: value,
	}
}

type patternTrieEntry struct {
	key   string
	value symbols.Logical
}

func (e *patternTrieEntry) Key() string {
	return e.key
}

func (e *patternTrieEntry) Value() symbols.Logical {
	return e.value
}

func (e *patternTrieEntry) String() string {
	return fmt.Sprintf("%v: %v", e.key, e.value)
}

var _ str.Entry[string, symbols.Logical] = (*patternTrieEntry)(nil)

// Page creates a new PatternPage with the given parameters.
//
// pageNumber: The page number of the PatternPage (0-based index).
// pageSize: The size of the PatternPage.
// totalElements: The total number of elements across all pages.
// elements: The list of elements in this page.
// Returns a pointer to the newly created PatternPage.
func Page(pageNumber, pageSize, totalElements int, elements []grammar.Symbol) *PatternPage {
	return &PatternPage{
		pageNumber:    pageNumber,
		pageSize:      pageSize,
		totalElements: totalElements,
		elements:      elements,
	}
}

type PatternPage struct {
	pageNumber    int
	pageSize      int
	totalElements int
	elements      []grammar.Symbol
}

func (p *PatternPage) PageNumber() int {
	return p.pageNumber
}

func (p *PatternPage) PageSize() int {
	return p.pageSize
}

// TotalElements Returns the total number of elements across all pages
func (p *PatternPage) TotalElements() int {
	return p.totalElements
}

// TotalPages Returns the total number of pages based on the total number of elements and the page size.
func (p *PatternPage) TotalPages() int {
	if p.pageSize == 0 {
		return 0
	}
	return (p.totalElements + p.pageSize - 1) / p.pageSize
}

// HasNext Returns true if there are more pages
func (p *PatternPage) HasNext() bool {
	return data.HasNextPage[grammar.Symbol](p)
}

// IsLast Returns true if this is the last page
func (p *PatternPage) IsLast() bool {
	return data.IsLastPage[grammar.Symbol](p)
}

// Content Returns the list of elements in the page
func (p *PatternPage) Content() []grammar.Symbol {
	return p.elements
}

// Iterator Returns an iterator over the elements in the page
func (p *PatternPage) Iterator() data.Iterator[grammar.Symbol] {
	return iterators.Default(p.Content())
}

var _ data.Page[grammar.Symbol] = (*PatternPage)(nil)
