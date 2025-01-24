// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package io

import "io"

type Closeable interface {
	io.Closer
	// IsClosed returns true if the Close() method has been called on this Closeable
	// or if the underlying object has been closed, false otherwise.
	IsClosed() bool
}

type ForceCloseable interface {
	Closeable
	ForceClose() error
}
