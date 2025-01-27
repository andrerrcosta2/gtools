// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package helpers

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"sync"
)

// CreateRowIntoBranchableMatrixIfAbsent It seems for me golang does not consider a component nil checking
// a correct pattern. As it seems a little expensive to sacrifice performance for a simple nil checking inside loops
// that does necessarily require reflection (approximately 21x slower) in such performance critical component,
// the only logical pattern is a nil checking by the caller.
// The reason behind it is simple: interface implementations are not necessarily pointers
// and variables of named types hold zero values. That means components that uses
// generics of interfaces cannot compare its values to nil.
//
// These components may perform better using pointers as map keys instead custom hashes, since most of them
// hashes its addresses. However, I need to study it a little more to check possible trade-offs over different types.
func CreateRowIntoBranchableMatrixIfAbsent[B data.Branchable[B], O any](b B, mtx *sync.RWMutex, branches *map[string][]O, rows *map[string][]string) (string, error) {
	var path []data.Branchable[B]

	if _, hasBranch := b.Branch(); !hasBranch {
		// Handle head node case
		hash := sortables.Unique[B](b)
		mtx.Lock()
		defer mtx.Unlock()
		if _, exists := (*branches)[hash]; !exists {
			(*branches)[hash] = []O{}
			(*rows)[hash] = []string{hash}
		}
		return createNewHeadOnBranchableMatrix(b, rows, branches), nil
	}

	// Backtrack to find an existing branch or head
	for branch, has := b.Branch(); has; branch, has = branch.Branch() {
		row := sortables.Unique[B](branch)

		mtx.RLock()
		exists := (*branches)[row] != nil
		mtx.RUnlock()

		if exists {
			mtx.Lock()
			for i := len(path) - 1; i >= 0; i-- {
				hash := sortables.Unique[B](path[i])
				err := createRowFromBranchOnBranchableMatrix(row, hash, rows)
				if err != nil {
					mtx.Unlock()
					return "", err
				}
				(*branches)[hash] = []O{}
				row = hash
			}
			mtx.Unlock()
			return row, nil
		}
		path = append(path, branch)
	}

	// Handle head when there is a path
	mtx.Lock()
	defer mtx.Unlock()
	from := createNewHeadOnBranchableMatrix(path[len(path)-1].(B), rows, branches)
	path = path[:len(path)-1]

	for i := len(path) - 1; i >= 0; i-- {
		row := sortables.Unique[B](path[i])
		err := createRowFromBranchOnBranchableMatrix(from, row, rows)
		if err != nil {
			return "", err
		}
		(*branches)[row] = []O{}
		from = row
	}

	return from, nil
}

// createNewHeadOnBranchableMatrix creates a new head branch in the matrix.
// It stores the hash of the branchable data in the rows map with the hash as key and a slice with the hash as value.
// It also stores the hash in the branches map with the hash as key and an empty slice as value.
// It returns the hash that was created.
func createNewHeadOnBranchableMatrix[B data.Branchable[B], O any](branchable B, rows *map[string][]string, branches *map[string][]O) string {
	hash := sortables.Unique[B](branchable)
	(*rows)[hash] = []string{hash} // Store the hash in the rows map with the hash as key and a slice with the hash as value
	(*branches)[hash] = []O{}      // Store the hash in the branches map with the hash as key and an empty slice as value
	return hash                    // Return the hash that was created
}

// createRowFromBranchOnBranchableMatrix Links a new branch (`to`) from an existing branch (`from`)
func createRowFromBranchOnBranchableMatrix(from, to string, rows *map[string][]string) error {
	if fromHashes, exists := (*rows)[from]; exists {
		newHashes := append(fromHashes, to) // Extend the branch
		(*rows)[to] = newHashes
		return nil
	}
	return errors.New("source branch does not exist")
}
