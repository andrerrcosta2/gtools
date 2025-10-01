// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflectrand

import "reflect"

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
