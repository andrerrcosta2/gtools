// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

import (
	"github.com/andrerrcosta2/gtools/core/functions"
	"github.com/andrerrcosta2/gtools/ranges"
)

func Double(left, right int) *DoubleSentinel {
	leftSentinel := Sentinel(left)
	rightSentinel := Sentinel(right)
	return &DoubleSentinel{
		left:  &leftSentinel,
		right: &rightSentinel,
	}
}

type DoubleSentinel struct {
	left  *Sentinel
	right *Sentinel
}

func (d *DoubleSentinel) Walk(left, right int) {
	d.left.Inc(left)
	d.right.Inc(right)
}

func (d *DoubleSentinel) Set(left, right int) {
	d.left.Set(left)
	d.right.Set(right)
}

func (d *DoubleSentinel) Left() int {
	return int(*d.left)
}

func (d *DoubleSentinel) Right() int {
	return int(*d.right)
}

func (d *DoubleSentinel) Gap() int {
	less := functions.Less(d.Left(), d.Right())
	return less * (d.Left() - d.Right())
}

func DoubleRanged(left int, leftRange ranges.Range, right int, rightRange ranges.Range) *DoubleRangedSentinel {
	return &DoubleRangedSentinel{
		left:  Ranged(left, leftRange),
		right: Ranged(right, rightRange),
	}
}

type DoubleRangedSentinel struct {
	left  *RangedSentinel
	right *RangedSentinel
}

func (d *DoubleRangedSentinel) HasNext() (bool, bool) {
	return d.left.HasNext(), d.right.HasNext()
}

func (d *DoubleRangedSentinel) HasAnyNext() bool {
	return d.left.HasNext() || d.right.HasNext()
}

func (d *DoubleRangedSentinel) HasPrev() (bool, bool) {
	return d.left.HasPrev(), d.right.HasPrev()
}

func (d *DoubleRangedSentinel) Walk(left, right int) (leftWalk, rightWalk bool) {
	return d.left.Inc(left), d.right.Inc(right)
}

func (d *DoubleRangedSentinel) Set(left, right int) (leftSet, rightSet bool) {
	return d.left.Set(left), d.right.Set(right)
}

func (d *DoubleRangedSentinel) Left() int {
	return d.left.Val()
}

func (d *DoubleRangedSentinel) Right() int {
	return d.right.Val()
}

func (d *DoubleRangedSentinel) DistanceToEnd() (left, right, less int) {
	left, right = d.left.ToMax(), d.right.ToMax()
	less = functions.Less(left, right)
	return
}

func (d *DoubleRangedSentinel) DistanceFromStart() (left, right, less int) {
	left, right = d.left.FromMin(), d.right.FromMin()
	less = functions.Less(left, right)
	return
}

func (d *DoubleRangedSentinel) Gap() int {
	less := functions.Less(d.left.Val(), d.right.Val())
	return less * (d.Left() - d.Right())
}

func DoubleRangedCloseable(left, minLeft, maxLeft, right, minRight, maxRight int) *DoubleRangedCloseableSentinel {
	return &DoubleRangedCloseableSentinel{
		left: &RangedCloseableSentinel{
			value: left,
			min:   minLeft,
			max:   maxLeft,
		},
		right: &RangedCloseableSentinel{
			value: right,
			min:   minRight,
			max:   maxRight,
		},
	}
}

type DoubleRangedCloseableSentinel struct {
	left  *RangedCloseableSentinel
	right *RangedCloseableSentinel
}

func (d *DoubleRangedCloseableSentinel) HasNext() (bool, bool) {
	return d.left.HasNext(), d.right.HasNext()
}

func (d *DoubleRangedCloseableSentinel) HasLeftNext() bool {
	return d.left.HasNext()
}

func (d *DoubleRangedCloseableSentinel) HasOnlyLeftNext() bool {
	return d.HasLeftNext() && !d.HasRightNext()
}

func (d *DoubleRangedCloseableSentinel) HasRightNext() bool {
	return d.right.HasNext()
}

func (d *DoubleRangedCloseableSentinel) HasOnlyRightNext() bool {
	return d.HasRightNext() && !d.HasLeftNext()
}

func (d *DoubleRangedCloseableSentinel) HasAnyNext() bool {
	return d.left.HasNext() || d.right.HasNext()
}

func (d *DoubleRangedCloseableSentinel) HasAnyDepletedNext() bool {
	return !d.left.HasNext() || !d.right.HasNext()
}

func (d *DoubleRangedCloseableSentinel) HasPrev() (bool, bool) {
	return d.left.HasPrev(), d.right.HasPrev()
}

func (d *DoubleRangedCloseableSentinel) HasLeftPrev() bool {
	return d.left.HasPrev()
}

func (d *DoubleRangedCloseableSentinel) HasRightPrev() bool {
	return d.right.HasPrev()
}

func (d *DoubleRangedCloseableSentinel) HasAnyPrev() bool {
	return d.left.HasPrev() || d.right.HasPrev()
}

func (d *DoubleRangedCloseableSentinel) Walk(left, right int) bool {
	increaseLeft := d.left.Inc(left)
	increaseRight := d.right.Inc(right)
	return !increaseLeft || !increaseRight
}

func (d *DoubleRangedCloseableSentinel) WalkAny(left, right int) {
	d.left.Inc(left)
	d.right.Inc(right)
}

func (d *DoubleRangedCloseableSentinel) Accepts(left, right int) bool {
	return d.left.Accepts(left) && d.right.Accepts(right)
}

func (d *DoubleRangedCloseableSentinel) Set(left, right int) bool {
	if d.left.Accepts(left) && d.right.Accepts(right) {
		d.left.Set(left)
		d.right.Set(right)
		return true
	}
	return false
}

func (d *DoubleRangedCloseableSentinel) CloseLeft() {
	d.left.Close()
}

func (d *DoubleRangedCloseableSentinel) CloseRight() {
	d.right.Close()
}

func (d *DoubleRangedCloseableSentinel) Close(less int) {
	if less < 0 {
		d.left.Close()
	} else {
		d.right.Close()
	}
}

func (d *DoubleRangedCloseableSentinel) OpenLeft() {
	d.left.Open()
}

func (d *DoubleRangedCloseableSentinel) OpenRight() {
	d.right.Open()
}

func (d *DoubleRangedCloseableSentinel) Open(less int) {
	if less < 0 {
		d.left.Open()
	} else {
		d.right.Open()
	}
}

func (d *DoubleRangedCloseableSentinel) IsClosed() (bool, bool) {
	return d.left.IsClosed(), d.right.IsClosed()
}

func (d *DoubleRangedCloseableSentinel) HasAnyClosed() bool {
	return d.left.IsClosed() || d.right.IsClosed()
}

func (d *DoubleRangedCloseableSentinel) IsLeftClosed() bool {
	return d.left.IsClosed()
}

func (d *DoubleRangedCloseableSentinel) IsRightClosed() bool {
	return d.right.IsClosed()
}

func (d *DoubleRangedCloseableSentinel) IsOpen() (bool, bool) {
	return d.left.IsOpen(), d.right.IsOpen()
}

func (d *DoubleRangedCloseableSentinel) HasAnyOpen() bool {
	return d.left.IsOpen() || d.right.IsOpen()
}

func (d *DoubleRangedCloseableSentinel) IsHalfOpen() bool {
	return (d.left.IsOpen() && d.right.IsClosed()) ||
		(d.left.IsClosed() && d.right.IsOpen())
}

func (d *DoubleRangedCloseableSentinel) IsLeftOpen() bool {
	return d.left.IsOpen()
}

func (d *DoubleRangedCloseableSentinel) IsRightOpen() bool {
	return d.right.IsOpen()
}

func (d *DoubleRangedCloseableSentinel) Next(less int) {
	if less < 0 {
		d.left.Next()
	} else {
		d.right.Next()
	}
}

func (d *DoubleRangedCloseableSentinel) Prev(less int) {
	if less < 0 {
		d.left.Next()
	} else {
		d.right.Next()
	}
}

func (d *DoubleRangedCloseableSentinel) Left() int {
	return d.left.value
}

func (d *DoubleRangedCloseableSentinel) Right() int {
	return d.right.value
}

func (d *DoubleRangedCloseableSentinel) DistanceToEnd() (left, right, less int) {
	left, right = d.left.DistanceToEnd(), d.right.DistanceToEnd()
	less = functions.Less(left, right)
	return
}

func (d *DoubleRangedCloseableSentinel) DistanceFromStart() (left, right, less int) {
	left, right = d.left.DistanceFromStart(), d.right.DistanceFromStart()
	less = functions.Less(left, right)
	return
}

func (d *DoubleRangedCloseableSentinel) Sum() int {
	return d.left.value + d.right.value
}

func (d *DoubleRangedCloseableSentinel) Gap() int {
	less := functions.Less(d.left.value, d.right.value)
	return less * (d.Left() - d.Right())
}

type DoubleRangedCloseableGapSentinel struct {
	left  *DoubleRangedCloseableSentinel
	right *DoubleRangedCloseableSentinel
}
