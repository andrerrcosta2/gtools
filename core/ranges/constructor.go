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
		return &openRng{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end - 1,
		}
	case ConcurrentClosed:
		return &concOpenRng{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end - 1,
		}
	case Open:
		return &openRng{
			reverse: false,
			current: start,
			start:   start,
			end:     end,
		}
	case ConcurrentOpen:
		return &concOpenRng{
			reverse: false,
			current: start,
			start:   start,
			end:     end,
		}
	case LeftClosed:
		return &openRng{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end,
		}
	case ConcurrentLeftClosed:
		return &concOpenRng{
			reverse: false,
			current: start + 1,
			start:   start + 1,
			end:     end,
		}
	case RightClosed:
		return &openRng{
			reverse: false,
			current: start,
			start:   start,
			end:     end - 1,
		}
	case ConcurrentRightClosed:
		return &concOpenRng{
			reverse: false,
			current: start,
			start:   start,
			end:     end - 1,
		}
	case ReverseClosed:
		return &openRng{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end + 1,
		}
	case ConcurrentReverseClosed:
		return &concOpenRng{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end + 1,
		}
	case ReverseOpen:
		return &openRng{
			reverse: true,
			current: start,
			start:   start,
			end:     end,
		}
	case ConcurrentReverseOpen:
		return &concOpenRng{
			reverse: true,
			current: start,
			start:   start,
			end:     end,
		}
	case ReverseLeftClosed:
		return &openRng{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end,
		}
	case ConcurrentReverseLeftClosed:
		return &concOpenRng{
			reverse: true,
			current: start - 1,
			start:   start - 1,
			end:     end,
		}
	case ReverseRightClosed:
		return &openRng{
			reverse: true,
			current: start,
			start:   start,
			end:     end + 1,
		}
	case ConcurrentReverseRightClosed:
		return &concOpenRng{
			reverse: true,
			current: start,
			start:   start,
			end:     end + 1,
		}
	default:
		return nil
	}
}
