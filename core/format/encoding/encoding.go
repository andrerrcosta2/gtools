// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package encoding

type Type uint8

const (
	JSON Type = iota
	Bin
	Gob
	Hex
	Base64
)
