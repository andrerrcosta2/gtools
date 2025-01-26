// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

import (
	"cmp"
	"github.com/andrerrcosta2/gtools/core/ranges"
)

func Double(left, right int) *DoubleSentinel {
	return &DoubleSentinel{
		left:  Int(left),
		right: Int(right),
	}
}

type DoubleSentinel struct {
	left  Sentinel
	right Sentinel
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
	return d.left.Val()
}

func (d *DoubleSentinel) Right() int {
	return d.right.Val()
}

func (d *DoubleSentinel) Gap() int {
	less := cmp.Compare(d.Left(), d.Right())
	return less * (d.Left() - d.Right())
}

func DoubleRanged(left int, leftRange ranges.Range, right int, rightRange ranges.Range) *DoubleRangedSentinel {
	return &DoubleRangedSentinel{
		left:  Ranged(left, leftRange),
		right: Ranged(right, rightRange),
	}
}

type DoubleRangedSentinel struct {
	left  RangedSentinel
	right RangedSentinel
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
	left, right = d.left.ToEnd(), d.right.ToEnd()
	less = cmp.Compare(left, right)
	return
}

func (d *DoubleRangedSentinel) DistanceFromStart() (left, right, less int) {
	left, right = d.left.FromStart(), d.right.FromStart()
	less = cmp.Compare(left, right)
	return
}

func (d *DoubleRangedSentinel) Gap() int {
	less := cmp.Compare(d.left.Val(), d.right.Val())
	return less * (d.Left() - d.Right())
}

// DoubleRangedCloseable creates a new DoubleRangedCloseableSentinel
func DoubleRangedCloseable(left ranges.Range, right ranges.Range) *DoubleRangedCloseableSentinel {
	return &DoubleRangedCloseableSentinel{
		left:  RangedCloseable(left),
		right: RangedCloseable(right),
	}
}

type DoubleRangedCloseableSentinel struct {
	left  CloseableSentinel
	right CloseableSentinel
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

func (d *DoubleRangedCloseableSentinel) Set(left, right int) bool {
	if d.left.Set(left) {
		if d.right.Set(right) {
			return true
		}
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
	return d.left.Val()
}

func (d *DoubleRangedCloseableSentinel) Right() int {
	return d.right.Val()
}

func (d *DoubleRangedCloseableSentinel) DistanceToEnd() (left, right, less int) {
	left, right = d.left.ToEnd(), d.right.ToEnd()
	less = cmp.Compare(left, right)
	return
}

func (d *DoubleRangedCloseableSentinel) DistanceFromStart() (left, right, less int) {
	left, right = d.left.FromStart(), d.right.FromStart()
	less = cmp.Compare(left, right)
	return
}

func (d *DoubleRangedCloseableSentinel) Sum() int {
	return d.left.Val() + d.right.Val()
}

func (d *DoubleRangedCloseableSentinel) Gap() int {
	less := cmp.Compare(d.left.Val(), d.right.Val())
	return less * (d.Left() - d.Right())
}

type DoubleRangedCloseableGapSentinel struct {
	left  *DoubleRangedCloseableSentinel
	right *DoubleRangedCloseableSentinel
}
