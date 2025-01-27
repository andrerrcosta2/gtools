// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

type Linked[N any] interface {
	Next() N
}

type DoubleLinked[N any] interface {
	Linked[N]
	// Prev returns the previous linked data.
	//
	// Returns the previous linked data as a value of type N.
	Prev() N
}
