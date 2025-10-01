// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import "reflect"

// HasDepth returns true if the type has a depth and false if it is a direct value
func HasDepth(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct
}

// HasRecursiveRef reports whether type t has recursive references.
func HasRecursiveRef(t reflect.Type) bool {
	visited := map[reflect.Type]bool{}
	return hasRecursive(t, visited)
}

func hasRecursive(t reflect.Type, visited map[reflect.Type]bool) bool {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice || t.Kind() == reflect.Array || t.Kind() == reflect.Chan {
		t = t.Elem()
	}

	if visited[t] {
		return true
	}

	if t.Kind() != reflect.Struct {
		return false
	}

	visited[t] = true
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i).Type
		if hasRecursive(ft, visited) {
			return true
		}
	}
	delete(visited, t)
	return false
}
