// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

//// XorRanged is a sentinel with a RightClosed range
//func XorRanged(min, max, subset int) (*XorRangedSentinel, error) {
//	if min > max {
//		return nil, errors.New("min > max")
//	}
//
//	if subset > max-min+1 {
//		return nil, errors.New("subset > max-min+1")
//	}
//
//	return &XorRangedSentinel{
//		subset:   createSubset(min, subset),
//		subsize:  subset,
//		leftPos:  Ranged(subset-2, 0, subset-1),
//		rightPos: Ranged(subset-1, 1, subset),
//		min:      min,
//		max:      max,
//	}, nil
//}
//
//func createSubset(min, subset int) []int {
//	sub := make([]int, subset)
//	for i := min; i < subset; i++ {
//		sub[i] = i
//	}
//	return sub
//}
//
//type XorRangedSentinel struct {
//	subset   []int
//	subsize  int
//	min      int
//	max      int
//	leftPos  *RangedSentinel
//	rightPos *RangedSentinel
//}
//
//func (d *XorRangedSentinel) Current() []int {
//	return d.subset
//}
//
//// Next [[2 3 4 5] [2 3 4 6] [2 3 4 7] [2 3 4 8] [2 3 5 6] [2 3 5 7] [2 3 5 8] [2 3 6 7]
//// [2 3 6 8] [2 3 7 8] [2 4 5 6] [2 4 5 7] [2 4 5 8] [2 4 6 7] [2 4 6 8] [2 4 7 8] [2 5 6 7]
//// [2 5 6 8] [2 5 7 8] [2 6 7 8] [3 4 5 6] [3 4 5 7] [3 4 5 8] [3 4 6 7] [3 4 6 8] [3 4 7 8]
//// [3 5 6 7] [3 5 6 8] [3 5 7 8] [3 6 7 8] [4 5 6 7] [4 5 6 8] [4 5 7 8] [4 6 7 8] [5 6 7 8]]
//func (d *XorRangedSentinel) Next() ([]int, bool) {
//	if d.subset[d.rightPos.value] < d.maxRightValue() {
//		d.subset[d.rightPos.value]++
//		return d.subset, true
//	}
//
//	for d.subset[d.leftPos.value] >= d.maxLeftValue() && d.leftPos.Prev() {
//		// Continue moving left until we find a position to increment
//	}
//
//	// Check if the left position can increment and ride the subset
//	if d.subset[d.leftPos.value] < d.maxLeftValue() {
//		if d.rideSubset() {
//			return d.subset, true
//		}
//	}
//
//	return []int{}, false
//}
//
//func (d *XorRangedSentinel) rideSubset() bool {
//	for {
//		// Check if the new starting point + the size of the subset would exceed the max
//		if d.subset[d.leftPos.value]+d.subsize-d.leftPos.value > d.max {
//			// If so, try to move the left position back and retry
//			if !d.leftPos.Prev() {
//
//				return false // No more positions to decrement, no more combinations possible
//			}
//			continue // Retry with the adjusted left position
//		}
//
//		start := d.subset[d.leftPos.value] + 1
//
//		// Update the subset with the new sequence of values
//		for i := d.leftPos.value; i < d.subsize; i++ {
//			d.subset[i] = start
//			start++
//		}
//
//		// Adjust the right position to follow the left
//		d.rightPos.Set(d.leftPos.value + 1)
//		return true
//	}
//}
//
//func (d *XorRangedSentinel) maxRightValue() int {
//	out := d.max - d.rightPos.ToMax()
//	return out
//}
//
//func (d *XorRangedSentinel) maxLeftValue() int {
//	out := d.max - d.leftPos.ToMax()
//	return out
//}

// Ride [[2 3 4 5] [2 3 4 6] [2 3 4 7] [2 3 4 8] [2 3 5 6] [2 3 5 7] [2 3 5 8] [2 3 6 7]
// [2 3 6 8] [2 3 7 8] [2 4 5 6] [2 4 5 7] [2 4 5 8] [2 4 6 7] [2 4 6 8] [2 4 7 8] [2 5 6 7]
// [2 5 6 8] [2 5 7 8] [2 6 7 8] [3 4 5 6] [3 4 5 7] [3 4 5 8] [3 4 6 7] [3 4 6 8] [3 4 7 8]
// [3 5 6 7] [3 5 6 8] [3 5 7 8] [3 6 7 8] [4 5 6 7] [4 5 6 8] [4 5 7 8] [4 6 7 8] [5 6 7 8]]
//func (d *XorRangedSentinel) Ride(subset int, ride int) ([][]int, bool) {
//	// Conditions to right.Next():
//	// 1. If d.Right + 1 is less than d.to
//	// 2. If d.Right + 1 must restart
//	subsets := make([][]int, ride)
//	combination := make([]int, subset)
//	rightRides := subset
//	removing := 1
//
//	for i := 0; i < ride; i++ { // whole combination
//		if i == 0 { // the first ride must create a snapshot
//			// Initialize combination with the current left sentinel
//			combination[0] = d.Left.Val()
//			for j := 1; j < rightRides; j++ { // each value of combination
//				if d.Walk(0, 1) {
//					combination[j] = d.Right.Val()
//				} else {
//					if removing == 1 {
//
//					}
//					removing++
//				}
//			}
//		} else {
//			combination = arrays.RemoveLastN(combination, removing)
//			for j := len(combination) - removing; j < len(combination); j++ {
//				if d.Walk(0, 1) {
//					combination[j] = d.Right.Val()
//				} else {
//					removing++
//					newRemoved := combination[len(combination)-removing]
//					d.Right.Set(newRemoved)
//					if d.Walk(0, 1) {
//						combination = arrays.RemoveLastN(combination, removing)
//						combination[j] = d.Right.Val()
//						removing++
//					} else {
//						return subsets, false // There's no more
//					}
//				}
//			}
//		}
//
//		// Addf the combination to the result
//		subsets[i] = combination
//	}
//	return subsets, d.Left.Less(d.to)
//}
