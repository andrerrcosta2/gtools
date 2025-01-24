// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

// New creates a new Range instance with the given type and start and end values.
// A range can be:
//
// - Closed, ReverseClosed, ConcurrentClosed, ConcurrentReverseClosed
// - Open, ReverseOpen, ConcurrentOpen, ConcurrentReverseOpen
// - LeftClosed, ReverseLeftClosed, ConcurrentLeftClosed, ConcurrentReverseLeftClosed
// - RightClosed, ReverseRightClosed, ConcurrentRightClosed, ConcurrentReverseRightClosed
func New(typ Type, start int, end int) Range {
	switch typ {
	case Closed:
		return &RangeImpl{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end - 1,
		}
	case ConcurrentClosed:
		return &ConcurrentRangeImpl{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end - 1,
		}
	case Open:
		return &RangeImpl{
			reverse: false,
			current: start,
			start:   start,
			end:     end,
		}
	case ConcurrentOpen:
		return &ConcurrentRangeImpl{
			reverse: false,
			current: start,
			start:   start,
			end:     end,
		}
	case LeftClosed:
		return &RangeImpl{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end,
		}
	case ConcurrentLeftClosed:
		return &ConcurrentRangeImpl{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end,
		}
	case RightClosed:
		return &RangeImpl{
			reverse: false,
			current: start,
			start:   start,
			end:     end - 1,
		}
	case ConcurrentRightClosed:
		return &ConcurrentRangeImpl{
			reverse: false,
			current: start,
			start:   start,
			end:     end - 1,
		}
	case ReverseClosed:
		return &RangeImpl{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end + 1,
		}
	case ConcurrentReverseClosed:
		return &ConcurrentRangeImpl{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end + 1,
		}
	case ReverseOpen:
		return &RangeImpl{
			reverse: true,
			current: start,
			start:   start,
			end:     end,
		}
	case ConcurrentReverseOpen:
		return &ConcurrentRangeImpl{
			reverse: true,
			current: start,
			start:   start,
			end:     end,
		}
	case ReverseLeftClosed:
		return &RangeImpl{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end,
		}
	case ConcurrentReverseLeftClosed:
		return &ConcurrentRangeImpl{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end,
		}
	case ReverseRightClosed:
		return &RangeImpl{
			reverse: true,
			current: start,
			start:   start,
			end:     end + 1,
		}
	case ConcurrentReverseRightClosed:
		return &ConcurrentRangeImpl{
			reverse: true,
			current: start,
			start:   start,
			end:     end + 1,
		}
	default:
		return nil
	}
}
