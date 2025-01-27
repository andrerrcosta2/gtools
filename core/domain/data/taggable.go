// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

// AsTaggable returns the given error as a data.Taggable[string] if it
// implements the data.Taggable[string] interface, or nil if it does not.
// The boolean return value is true if the error implements the
// data.Taggable[string] interface, and false otherwise.
func AsTaggable[T prim.Hashable](data any) (Taggable[T], bool) {
	if taggable, ok := data.(Taggable[T]); ok {
		return taggable, true
	}
	return nil, false
}

func IsTaggable[T prim.Hashable](data any) bool {
	_, ok := data.(Taggable[T])
	return ok
}

// HasAnyTag checks if the given Taggable has any of the given tags.
// It takes a Taggable and a variable number of tags as arguments.
// It returns true if the Taggable has any of the given tags, and false otherwise.
func HasAnyTag[T prim.Hashable](taggable Taggable[T], tags ...T) bool {
	// Iterate over the given tags
	for _, tag := range tags {
		// Iterate over the current tags of the Taggable
		for _, t := range taggable.Tags() {
			// If the Taggable has the current tag, return true
			if t == tag {
				return true
			}
		}
	}
	// If the Taggable does not have any of the given tags, return false
	return false
}

// HasAllTags checks if the given Taggable has all the given tags.
// It takes a Taggable and a variable number of tags as arguments.
// It returns true if the Taggable has all the given tags, and false otherwise.
func HasAllTags[T prim.Hashable](taggable Taggable[T], tags ...T) bool {
	// If the Taggable does not have any of the given tags, return false
	if !HasAnyTag(taggable, tags...) {
		return false
	}
	// If the Taggable has all the given tags, return true
	return true
}
