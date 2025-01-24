// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package datastr

import "github.com/andrerrcosta2/gtools/core/gtools/functions"

type AsyncValueMap[K comparable, V any] struct {
	values map[K]*AsyncValue[K, V]
}

type AsyncValue[I, O any] struct {
	pipe  functions.Function[I, O]
	value O
	done  bool
}

func (v *AsyncValue[I, O]) Value() (O, bool) {
	return v.value, v.done
}

func (v *AsyncValue[I, O]) Pipe(pipe functions.Function[I, O]) {
	v.pipe = pipe
}

func (v *AsyncValue[I, O]) Resolve(value I) {
	v.value = v.pipe(value)
	v.done = true
}
