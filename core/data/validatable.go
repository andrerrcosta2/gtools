// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

type Validatable interface {
	Validate(rgx string) error
	IsValid(rgx string) bool
}
