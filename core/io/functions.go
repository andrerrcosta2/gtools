// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package io

import (
	"io"
	"log"

	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// CloseOrLog closes a resource and logs any error that occurs.
func CloseOrLog(c io.Closer, tags ...string) {
	if err := c.Close(); err != nil {
		label := "resource"

		// Priority 1: Explicitly passed tags
		if len(tags) > 0 {
			label = tags[0]
		} else if n, ok := c.(interface{ Name() string }); ok {
			// Priority 2: Does the object have a Name() method? (like *os.File)
			label = n.Name()
		}

		log.Printf("Error closing %s: %v", label, err)
	}
}

// CloseOrConsume closes a resource and consumes any error that occurs.
func CloseOrConsume(c io.Closer, handler function.BiConsumer[string, error], tags ...string) {
	if err := c.Close(); err != nil {
		label := "resource"

		// Priority 1: Explicitly passed tags
		if len(tags) > 0 {
			label = tags[0]
		} else if n, ok := c.(interface{ Name() string }); ok {
			// Priority 2: Does the object have a Name() method? (like *os.File)
			label = n.Name()
		}
		handler(label, err)
	}
}

// ForceCloseOrLog force-closes a resource and logs any error that occurs.
//
// It returns an error if the resource couldn't be force-closed.
func ForceCloseOrLog(c ForceCloseable, name string) {
	if err := c.ForceClose(); err != nil {
		log.Printf("Error while force-closing resource %q: %+v", name, err)
	}
}

// IsCloseable returns true if the given value implements the Closeable interface,
// false otherwise.
func IsCloseable(c any) bool {
	_, ok := c.(Closeable)
	return ok
}

// AsCloseable returns the given value as a Closeable if it implements the
// interface, or nil if it doesn't. The ok result is true if the given value
// implements the interface, false otherwise.
//
// It's useful when you want to check if a value can be closed.
func AsCloseable(c any) (Closeable, bool) {
	closeable, ok := c.(Closeable)
	return closeable, ok
}

// IsForceCloseable returns true if the given value implements the ForceCloseable interface,
// false otherwise.
func IsForceCloseable(c any) bool {
	_, ok := c.(ForceCloseable)
	return ok
}

// AsForceCloseable returns the given value as a ForceCloseable if it implements the
// interface, or nil if it doesn't. The OK result is true if the given value
// implements the interface, false otherwise.
//
// It's useful when you want to check if a value can be force-closed.
func AsForceCloseable(c any) (ForceCloseable, bool) {
	forceCloseable, ok := c.(ForceCloseable)
	return forceCloseable, ok
}
