// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

type auditEvent[T any] struct {
	err      error
	value    T
	hasValue bool
	complete bool
	unsub    bool
	timer    bool
}

type debounceEvent[T any] struct {
	kind    debounceKind
	val     T
	version uint64
}

type debounceKind int

const (
	onNext debounceKind = iota
	onTimeout
)

type timeoutEvent[T any] struct {
	err      error
	value    T
	hasValue bool
	complete bool
	unsub    bool
	timeout  bool
}
