// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

type Closable interface {
	Close()
	Closed() bool
}
