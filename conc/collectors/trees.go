// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package collectors provides implementations of the Collector interface for organizing tree structures in parallel.
// To achieve that, most of these collectors increases spatial complexity to support scalability or latency-sensitive flows.
//
// Name Aspects of Parallel Tree Collection:
//
// 1. Parallelization Overhead:
// Designed for scalability, this package organizes tree rows in parallel, beneficial for large datasets.
// However, parallel processing adds synchronization overhead, which may slow down small datasets where
// sequential methods (e.g., depth-first or breadth-first traversal) can perform faster.
//
// 2. Real-Time Organization:
// Unlike traditional methods that organize nodes after traversal, this package organizes nodes as they are processed.
// This real-time approach saves post-processing time, especially on large datasets.
//
// 3. Scalability vs. Small Data Efficiency:
// While parallelism boosts performance on large datasets, it may add unnecessary complexity for small datasets.
// For smaller data, sequential traversal is often simpler and faster due to lower coordination costs.
//
// 4. Trade-Off Strategies:
// Future improvements could include adaptive parallelization, hybrid structures, batched processing, and
// dynamic depth/branch management to balance performance and scalability.
package collectors

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/collectors/internal/datastr"
	"github.com/andrerrcosta2/gtools/conc/collectors/internal/helpers"
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/gflux/core/tasks"
)

//type BranchableCollector[B data.Branchable, H any, C ~[][]B] Collector[B, C]

// Branchable returns a Collector which can collect any data that implements data.Branchable.
// It returns a matrix of unsorted branchable values organized by rows.
func Branchable[B data.Branchable[B], C ~[][]B]() Collector[B, C] {
	return &branchableCollector[B, C]{
		errors:   make([]error, 0),
		branches: make(map[string][]B),
		rows:     make(map[string][]string),
	}
}

// This collector organizes any data that implements data.Branchable into an unsorted matrix.
// The branch must be naturally comparable. Relying on pointers as keys is unsafe, as this collector
// dereferences pointers before comparing them, which may lead to incorrect collections.
//
// Due to the current philosophy of golang, this collector, while working with interfaces,
// doesn't prevent processing of nil stream data inside the collector without a reflection overhead.
// For that reason this collector may panic when a nil value is retrieved from the stream as data.Branchable.
// That was a decision made to prioritize performance over safety.
type branchableCollector[B data.Branchable[B], C ~[][]B] struct {
	mtx      sync.RWMutex
	branches map[string][]B
	rows     map[string][]string // hashes of branches in a row
	errors   []error
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a matrix of rows.
func (c *branchableCollector[B, C]) Collect(stream gtools.Stream[B]) C {
	for branchable := range stream {
		c.collect(branchable)
	}
	return c.get()
}

func (c *branchableCollector[B, C]) collect(b B) {
	branch, err := helpers.CreateRowIntoBranchableMatrixIfAbsent(b, &c.mtx, &c.branches, &c.rows)
	if err != nil {
		c.errors = append(c.errors, err)
		return
	}
	c.appendBranchable(branch, b)
}

func (c *branchableCollector[B, C]) appendBranchable(hash string, b B) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	branch := c.branches[hash]
	branch = append(branch, b)
}

func (c *branchableCollector[B, C]) get() (branches [][]B) {
	branches = make([][]B, 0, len(c.rows))
	for _, hashes := range c.rows {
		branches = append(branches, c.flatRow(hashes))
	}
	return
}

func (c *branchableCollector[B, C]) flatRow(hashes []string) []B {
	flatRow := make([]B, 0)
	for _, hash := range hashes {
		if branch, exists := c.branches[hash]; exists {
			flatRow = append(flatRow, branch...) // Append the dereferenced branch
		}
	}
	return flatRow
}

func (c *branchableCollector[B, C]) Errors() []error {
	return c.errors
}

// Branch returns a Collector that collects the bch of any data that implements data.Branchable.
// It returns a matrix organized by the hierarchy of the bch in the order of the data.
// That means this collector do not collect the data itself, only the bch.
func Branch[B data.Branchable[B], C ~[][]B]() Collector[B, C] {
	return &branchCollector[B, C]{
		matrix:   make([][]B, 0),
		branches: make(map[string]int),
	}
}

type branchCollector[B data.Branchable[B], C ~[][]B] struct {
	mtx      sync.RWMutex
	matrix   [][]B
	branches map[string]int
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a matrix of rows.
// It does not collect nil values or values with nil rows.
func (c *branchCollector[B, C]) Collect(stream gtools.Stream[B]) C {
	// Iterate over the stream
	for branchable := range stream {
		// Collect the value to generate the matrix
		c.collect(branchable)
	}
	return c.matrix
}

func (c *branchCollector[B, C]) collect(b B) {
	// This is a significant note. a branchable data that has a branch
	// that aren't branchable can represent a key but not a tree-node-like structure.
	// That means a data with no branch can be assumed to be a root branch on recursive data.
	// (The problem goes further when it is not, but a different type of recursive data
	// e.g [D data.Branchable[B], B data.Branchable[B], C ~[][]B]
	// this is a case where I will need to create too many different collectors
	// for different types of data. I will stick with this for now.)
	if branch, has := b.Branch(); has {
		hash := sortables.Unique[B](branch)
		c.mtx.RLock()
		if _, exists := c.branches[hash]; exists {
			c.mtx.RUnlock()
			return
		}
		c.mtx.RUnlock()

		c.mtx.Lock()
		defer c.mtx.Unlock()

		idx := len(c.matrix)
		c.branches[hash] = idx
		c.matrix = append(c.matrix, []B{}) // Initialize a new slice for this branch

		// The data should be inserted in reverse order
		for itr := branch; has; itr, has = itr.Branch() {
			hash = sortables.Unique[B](itr)
			// Prepend the new element (last inserted should be first)
			c.matrix[idx] = append([]B{itr}, c.matrix[idx]...)
		}
	}
}

func (c *branchCollector[B, C]) createHead(b B) {

}

// Errors returns the list of err that occurred during the collection process.
// It returns all err collected during the collection process.
func (c *branchCollector[B, C]) Errors() []error {
	return []error{}
}

// PipedBranchable returns a Collector that can collect any data the implements data.Branchable.
// It returns a matrix of unsorted branchable values organized by rows after being piped.
func PipedBranchable[B data.Branchable[B], O any, C ~[][]O]() Piped[B, O, C] {
	return &pipedBranchableCollector[B, O, C]{
		branches: make(map[string][]O),
		rows:     make(map[string][]string),
		errors:   make([]error, 0),
	}
}

type pipedBranchableCollector[B data.Branchable[B], O any, C ~[][]O] struct {
	mtx      sync.RWMutex
	branches map[string][]O
	rows     map[string][]string // hashes of branches in a row
	pipe     functions.Function[B, O]
	errors   []error
}

func (c *pipedBranchableCollector[B, O, C]) Collect(stream gtools.Stream[B]) C {
	// Iterate over the stream
	for branchable := range stream {
		// Collect the value to generate the matrix
		c.collect(branchable)

	}
	return c.get()
}

func (c *pipedBranchableCollector[B, O, C]) collect(b B) {
	branch, err := helpers.CreateRowIntoBranchableMatrixIfAbsent(b, &c.mtx, &c.branches, &c.rows)
	if err != nil {
		c.errors = append(c.errors, err)
		return
	}
	c.appendBranchable(branch, b)
}

func (c *pipedBranchableCollector[B, O, C]) get() (branches [][]O) {
	branches = make([][]O, 0, len(c.rows))
	for _, hashes := range c.rows {
		branches = append(branches, c.flatRow(hashes))
	}
	return
}

func (c *pipedBranchableCollector[B, O, C]) appendBranchable(hash string, b B) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	branch := c.branches[hash]
	branch = append(branch, c.pipe(b))
}

func (c *pipedBranchableCollector[B, O, C]) flatRow(hashes []string) []O {
	flatRow := make([]O, 0)
	for _, hash := range hashes {
		if branch, exists := c.branches[hash]; exists {
			flatRow = append(flatRow, branch...) // Append the dereferenced branch
		}
	}
	return flatRow
}

// Errors returns the list of err that occurred during the collection process.
// It returns all err collected during the collection process.
func (c *pipedBranchableCollector[B, O, C]) Errors() []error { return c.errors }

// Pipe returns a Piped that collects all values from the given Stream,
// applies the given flt to each value and collects the results into a slice.
func (c *pipedBranchableCollector[B, O, C]) Pipe(pipe functions.Function[B, O]) Piped[B, O, C] {
	// Set the flt to be applied to each collected value
	c.pipe = pipe
	// Return itself as a Piped
	return c
}

type PipedBranchCollector[B data.Branchable[B], H any, O any, C ~[][]O] Piped[H, O, C]

type SerialBranchableCollector[B data.SerializableBranchable[B, S], H any, S prim.Serializable, C ~[]B] Collector[B, C]

type SequentialBranchableCollector[B data.SerializableBranchable[B, int], H any, C ~[]B] SerialBranchableCollector[B, H, int, C]

// SequentialBranchable returns a Collector that can collect any data that implements
// data.SerializableBranchable[H, int].
// It returns a matrix of unsorted branchable values organized by rows.
// A recursive branchable data is a branchable data which branch is also a branchable data.
// If the flag is set to true and a non-branchable data is retrieved as branch in the stream,
// the data will be ignored.
func SequentialBranchable[B data.SerializableBranchable[B, int], H any, C ~[][]B]() Collector[B, C] {
	return &sequentialBranchableCollector[B, H, C]{
		errors:   make([]error, 0),
		branches: make(map[string][]B),
		rows:     make(map[string][]string),
	}
}

type sequentialBranchableCollector[B data.SerializableBranchable[B, int], H any, C ~[][]B] struct {
	mtx      sync.RWMutex
	branches map[string][]B
	rows     map[string][]string // hashes of branches in a row
	errors   []error
}

func (c *sequentialBranchableCollector[B, H, C]) Collect(stream gtools.Stream[B]) C {
	for branchable := range stream {
		c.collect(branchable)
	}
	return c.get()
}

func (c *sequentialBranchableCollector[B, H, C]) collect(b B) {
	branch, err := helpers.CreateRowIntoBranchableMatrixIfAbsent(b, &c.mtx, &c.branches, &c.rows)
	if err != nil {
		c.errors = append(c.errors, err)
		return
	}
	c.appendBranchable(branch, b)
}

func (c *sequentialBranchableCollector[B, H, C]) get() (branches [][]B) {
	branches = make([][]B, 0, len(c.rows))
	for _, hashes := range c.rows {
		branches = append(branches, c.flatRow(hashes))
	}
	return
}

func (c *sequentialBranchableCollector[B, H, C]) appendBranchable(hash string, b B) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	branch := c.branches[hash]
	index := b.Serial()

	// Ensure the slice is large enough to hold the index
	if len(branch) <= index {
		// Grow the slice up to the required index + 1
		newBranch := make([]B, index+1)
		copy(newBranch, branch)
		branch = newBranch
	}

	// Assign the value directly
	branch[index] = b
}

func (c *sequentialBranchableCollector[B, H, C]) flatRow(hashes []string) []B {
	row := make([]B, 0, len(hashes))
	for _, hash := range hashes {
		if branch, exists := c.branches[hash]; exists {
			row = append(row, branch...)
		}
	}
	return row
}

// Errors returns the list of err that occurred during the collection process.
// It returns all err collected during the collection process.
func (c *sequentialBranchableCollector[B, H, C]) Errors() []error {
	return c.errors
}

type PipedSerializedBranchableCollector[B data.SerializableBranchable[B, S], S prim.Serializable, O any, C ~[][]O] Piped[B, O, C]

// PipedSequentialBranchableCollector is a Collector for sequential branchable nodes
type PipedSequentialBranchableCollector[B data.SerializableBranchable[B, int], O any, C ~[][]O] PipedSerializedBranchableCollector[B, int, O, C]

// BranchFlattener returns a Flattener that can collect any data that implements data.Branchable.
// It returns a slice of unsorted branchable values organized by rows after being flattened.
//
// This is exhaustively complex (many abstract patterns hidden inside a single flow) and IT'S NOT FINISHED!
// I already found its patterns, but as I'm not using it anymore on my pattern-trie, I will finish it on the next mr.
// If you are curious you can see there is an async value which assumes two states:
// 1. The value is not ready, it is a supplier
// 2. The value is ready
// Using correct components for the correct synchronization pattern will decrease the complexity
// inside the "flat" method. The map must be async and pattern-biased, I don't remember exactly what but its registered somewhere.
//
// Is worth to mention that these are my solutions for in-memory collectors. Collectors of serializable
// data for sure are way more complex and may require too many different compositions.
func BranchFlattener[B data.Branchable[B], O any, C ~[]O]() Flattener[B, O, C] {
	return &branchFlattenerCollector[B, O, C]{
		bch: make(map[string]*datastr.Pair[int, *tasks.StepBranch]),
	}
}

type branchFlattenerCollector[B data.Branchable[B], O any, C ~[]O] struct {
	mtx sync.RWMutex
	out C
	bch map[string]*datastr.Pair[int, *tasks.StepBranch]
	err []error
	flt functions.BiFunction[O, B, O]
}

func (c *branchFlattenerCollector[B, O, C]) Collect(stream gtools.Stream[B]) C {
	for node := range stream {
		c.collect(node)
	}
	return c.out
}

func (c *branchFlattenerCollector[B, O, C]) collect(b B) {
	if branch, has := b.Branch(); has {
		hash := sortables.Unique[B](branch)
		c.mtx.Lock()
		defer c.mtx.Unlock()
		if _, exists := c.bch[hash]; exists {
			return
		}
		c.flat(hash, branch)
	}
}

func (c *branchFlattenerCollector[B, O, C]) flat(hash string, b B) {
	idx := len(c.out)
	c.bch[hash] = datastr.PairOf[int, *tasks.StepBranch](idx, nil)
	var ops []functions.Runnable

	for branch, exists := b.Branch(); exists; branch, exists = branch.Branch() {
		// the scheduled chain must be the starting point
		chain := sortables.Unique[B](branch)
		// if exists should be used as flat object
		if flat, has := c.bch[chain]; has {
			if len(ops) == 0 {
				// there is no gap between bch
				c.out[idx] = c.flt(c.out[flat.First()], b)
			} else {
				// If intermediate bch weren't out they must be scheduled.
				// that happens because bch are accessed by its memory addresses
				// not by channel access, so we may be facing a race condition when we
				// are doing an early processing.
				//c.runner.Append(chain, ops...) The runner is dead. Maybe not.
			}
			return
		} else {
			ops = append([]functions.Runnable{
				func() { c.out[idx] = c.flt(c.out[idx], branch) },
			}, ops...)
		}
	}
	// if no branch is found it is a root branch
	var zero O
	c.out = append(c.out, c.flt(zero, b))
}

func (c *branchFlattenerCollector[B, O, C]) Errors() []error {
	return c.err
}

func (c *branchFlattenerCollector[B, O, C]) Pipe(pipe functions.BiFunction[O, B, O]) Flattener[B, O, C] {
	c.flt = pipe
	return c
}
