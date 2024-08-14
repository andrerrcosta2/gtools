// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsutil

type PathType int

const (
	Relative PathType = iota
	Root
	Literal
)

type Path struct {
	Type PathType
	Path string
}

func NewPath(pathType PathType, path string) Path {
	return Path{Type: pathType, Path: path}
}
