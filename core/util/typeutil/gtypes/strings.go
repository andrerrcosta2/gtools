// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtypes

type StringOf string

// Equal compares the StringOf value with another value.
func (s StringOf) Equal(o interface{}) bool {
	switch other := o.(type) {
	case StringOf:
		return s == other
	case string:
		return string(s) == other
	case []rune:
		return string(s) == string(other)
	case []byte:
		return string(s) == string(other)
	default:
		return false
	}
}

func (s StringOf) Less(o interface{}) bool {
	switch other := o.(type) {
	case StringOf:
		return s < other
	case string:
		return string(s) < other
	case []rune:
		return string(s) < string(other)
	case []byte:
		return string(s) < string(other)
	default:
		return false
	}
}
