// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package random

type Injectable interface {
	Inject(i any) any
}

type identityInjectable struct{}

func (i identityInjectable) Inject(i2 any) any {
	return i2
}
