// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsys

type PathType int

const (
	Relative PathType = iota
	Literal
)

type Path struct {
	Type PathType
	Path string
}

func PathOf(pathType PathType, path string) Path {
	return Path{Type: pathType, Path: path}
}
