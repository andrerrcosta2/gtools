// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

type Node interface {
	AddChild(Node)
	GetParent(Node)
}

type GeneralToken struct {
	children []Node
	parent   Node
}
