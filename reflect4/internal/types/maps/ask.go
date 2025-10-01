// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import "reflect"

// HasKey reports whether a map contains a key
func HasKey(m reflect.Value, key reflect.Value) bool {
	return m.MapIndex(key).IsValid()
}
