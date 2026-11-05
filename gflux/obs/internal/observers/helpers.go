// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observers

func Next[T any](o *Observer[T], v T) {
	if o != nil && o.Next != nil {
		o.Next(v)
	}
}

func Error[T any](o *Observer[T], err error) {
	if o != nil && o.Error != nil {
		o.Error(err)
	}
}

func Complete[T any](o *Observer[T]) {
	if o != nil && o.Complete != nil {
		o.Complete()
	}
}
