// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package casters

func Typed[G any](values ...interface{}) (result []G, allMatches bool) {
	allMatches = true
	for _, value := range values {
		if castValue, ok := value.(G); ok {
			result = append(result, castValue)
		} else {
			allMatches = false
		}
	}
	return
}

func AssertedTyped[G any](values ...interface{}) []G {
	var result []G
	for _, value := range values {
		if castValue, ok := value.(G); ok {
			result = append(result, castValue)
		} else {
			panic("AssertedTyped will panic if the provided value is not of the expected type")
		}
	}
	return result
}
