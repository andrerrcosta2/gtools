// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package write

type Opt uint8

const (
	Default                Opt = 0
	IgnoreUnexportedFields Opt = 1 << iota
	AllowNilVsEmpty
)
