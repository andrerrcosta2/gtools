// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

func Triple(left, middle, right int) *TripleSentinel {
	leftSentinel := Sentinel(left)
	middleSentinel := Sentinel(middle)
	rightSentinel := Sentinel(right)
	return &TripleSentinel{
		Left:   &leftSentinel,
		Middle: &middleSentinel,
		Right:  &rightSentinel,
	}
}

type TripleSentinel struct {
	Left   *Sentinel
	Middle *Sentinel
	Right  *Sentinel
}
