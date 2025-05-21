// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtypes

func NewTrilean() *Trilean {
	return &Trilean{unknown: true}
}

type Trilean struct {
	unknown bool
	value   bool
}

func (t *Trilean) Set(value bool) {
	t.value = value
	t.unknown = false
}

func (t *Trilean) Unset() {
	t.unknown = true
}

func (t *Trilean) Unknown() bool {
	return t.unknown
}

func (t *Trilean) True() bool {
	return t.value && !t.unknown
}

func (t *Trilean) False() bool {
	return !t.value && !t.unknown
}

func (t *Trilean) String() string {
	if t.unknown {
		return "unknown"
	}
	if t.value {
		return "true"
	}
	return "false"
}
