// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package matrices

// SetAt sets the value at the given row and column in the given matrix.
//
// It takes the matrix 'm', the row 'row', the column 'col', and the value 'data' as arguments.
// If the row or column is out of bounds, it resizes the matrix to fit the row and column and sets the value at the specified position.
// It returns the modified matrix.
func SetAt[E ~[][]T, T any](m E, row, col int, data T) E {
	// Check if we need to grow the matrix vertically (i.e., add more rows)
	if row >= len(m) {
		// Create a new matrix with the correct number of rows
		newRows := make(E, row+1)
		// Copy all existing rows to the new matrix
		copy(newRows, m)
		// Replace the original matrix with the new one
		m = newRows
	}

	// Check if the row already exists or if it's new
	if m[row] == nil {
		// Create a new row with the correct number of columns if the row didn't exist
		m[row] = make([]T, col+1)
	} else if col >= len(m[row]) {
		// Expand the row if the number of columns is insufficient
		// Create a new row with the correct number of columns
		newCols := make([]T, col+1)
		// Copy all existing columns to the new row
		copy(newCols, m[row])
		// Replace the original row with the new one
		m[row] = newCols
	}

	// Set the data at the specified row and column
	m[row][col] = data

	return m
}

// Grow grows the matrix by adding 'rows' empty rows at the end.
//
// It takes the matrix 'm' and the number of rows 'rows' as arguments.
// It returns the modified matrix.
func Grow[E ~[][]T, T any](m *E, rows int) {
	// Create a new matrix with the correct number of rows
	newRows := make(E, len(*m)+rows)
	// Copy all existing rows to the new matrix
	copy(newRows, *m)
	// Replace the original matrix with the new one
	*m = newRows
}

// GrowIfLessThan grows the matrix by adding the difference in rows between 'rows' and the current number of rows
// if the current number of rows is less than 'rows'.
//
// It takes the matrix 'm' and the number of rows 'rows' as arguments.
// It returns the modified matrix.
func GrowIfLessThan[E ~[][]T, T any](m *E, rows int) {
	if len(*m) < rows {
		// Grow the matrix by adding the difference in rows between 'rows' and the current number of rows
		Grow(m, rows-len(*m))
	}
}
