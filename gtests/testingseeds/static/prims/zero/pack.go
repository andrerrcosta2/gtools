// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

type (
	Pack struct{}
)

func (c Pack) Refs() []any {
	return referencesSet()
}

func (c Pack) Values() []any {
	return valuesSet()
}
