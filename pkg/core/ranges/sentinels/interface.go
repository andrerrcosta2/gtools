// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

type Sentinel interface {
	// Next tries to increase the sentinel value and returns true if the value was increased
	Next() bool
	// Inc increases the sentinel value and returns true if the value didn't reach the end
	Inc(i int) bool
	// Prev decreases the sentinel value and returns true if the value didn't reach the start
	Prev() bool
	// Dec tries to decrease the sentinel value and returns true if the value was decreased
	Dec(i int) bool
	// Set sets the sentinel value and returns true if the value is within the range
	Set(i int) bool
	// Less returns true if the value is less than the other
	Less(other int) bool
	// Equal returns true if the value is equal to the other
	Equal(other int) bool
	// Val returns the current value
	Val() int
}

type RangedSentinel interface {
	Sentinel
	// HasNext returns true if there is a next value
	HasNext() bool
	// HasPrev returns true if there is a previous value
	HasPrev() bool
	// FromStart returns the distance from the minimum value and the current value
	FromStart() int
	// ToEnd returns the distance from the maximum value and the current value
	ToEnd() int
}

type CloseableSentinel interface {
	RangedSentinel
	// Close closes the sentinel
	Close()
	// IsClosed returns true if the sentinel is closed
	IsClosed() bool
	// IsOpen returns true if the sentinel is open
	IsOpen() bool
	// Open opens the sentinel
	Open()
}
