// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package keyers

import "github.com/andrerrcosta2/gtools/core/data/keyer"

func FingerPrint[T any]() keyer.Comparable[T, string] {
	return fingerprinter[T]{}
}

type fingerprinter[T any] struct{}

func (fingerprinter[T]) Key(v T) string {
	return fingerprint[T](v)
}

func Identity[T comparable]() keyer.Comparable[T, T] {
	return identityKeyer[T]{}
}

type identityKeyer[T comparable] struct{}

func (identityKeyer[T]) Key(v T) T {
	return v
}
