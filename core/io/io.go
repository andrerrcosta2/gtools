// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package io

import (
	"io"
)

// ErrShortWrite means that a write accepted fewer bytes than requested
// but failed to return an explicit error.
var ErrShortWrite = io.ErrShortWrite

// ErrShortBuffer means that a read required a longer buffer than was provided.
var ErrShortBuffer = io.ErrShortWrite

// EOF is the error returned by Read when no more input is available.
// (Read must return EOF itself, not an error wrapping EOF,
// because callers will test for EOF using ==.)
// Functions should return EOF only to signal a graceful end of input.
// If the EOF occurs unexpectedly in a structured data stream,
// the appropriate error is either [ErrUnexpectedEOF] or some other error
// giving more detail.
var EOF = io.EOF

// ErrUnexpectedEOF means that EOF was encountered in the
// middle of reading a fixed-size block or data structure.
var ErrUnexpectedEOF = io.ErrUnexpectedEOF

// ErrNoProgress is returned by some clients of a [Reader] when
// many calls to Read have failed to return any data or error,
// usually the sign of a broken [Reader] implementation.
var ErrNoProgress = io.ErrNoProgress

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
