// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package searchers

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/search/searcher"
)

func DFS[N any, H comparable](cmp comparator.EqCmpKeyer[N, H], start N, expand func(N, func(N) bool) bool) searcher.Graph[N] {
	return &dfs[N, H]{
		start:   start,
		expand:  expand,
		visited: make(map[H]bool),
		store:   make(map[H]N),
		cmp:     cmp,
	}
}

type dfs[N any, H comparable] struct {
	start   N
	expand  func(N, func(N) bool) bool
	visited map[H]bool
	store   map[H]N // store actual elements to retrieve exact reference/value
	cmp     comparator.EqCmpKeyer[N, H]
}

func (s *dfs[N, H]) Has(target N) bool {
	_, ok := s.store[s.cmp.Key(target)]
	return ok
}

func (s *dfs[N, H]) Find(target N) (N, bool) {
	val, ok := s.store[s.cmp.Key(target)]
	return val, ok
}

func (s *dfs[N, H]) Path(from, to N) []N {
	parent := make(map[H]N)
	var path []N

	s.visited = make(map[H]bool)
	s.store = make(map[H]N)

	var visit func(n N) bool

	visit = func(n N) bool {
		hashed := s.cmp.Key(n)
		if s.visited[hashed] {
			return true
		}
		s.visited[hashed] = true
		s.store[hashed] = n

		if s.cmp.Equals(n, to) {
			return false
		}

		return s.expand(n, func(next N) bool {
			h := s.cmp.Key(next)

			if s.visited[h] {
				return true
			}

			parent[h] = n
			return visit(next)
		})
	}

	visit(from)

	// reconstruct path
	if _, ok := parent[s.cmp.Key(to)]; !ok && !s.cmp.Equals(from, to) {
		return nil
	}

	for cur := to; !s.cmp.Equals(cur, from); cur = parent[s.cmp.Key(cur)] {
		path = append(path, cur)
	}
	path = append(path, from)

	// reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (s *dfs[N, H]) Traverse(yield func(N) bool) {
	s.visited = make(map[H]bool)
	s.store = make(map[H]N)

	var visit func(n N) bool
	visit = func(n N) bool {
		hashed := s.cmp.Key(n)
		if s.visited[hashed] {
			return true
		}
		s.visited[hashed] = true
		s.store[hashed] = n // store actual value

		if !yield(n) {
			return false
		}

		return s.expand(n, func(next N) bool {
			return visit(next)
		})
	}

	visit(s.start)
}
