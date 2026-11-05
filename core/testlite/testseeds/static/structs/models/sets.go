// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

type Seeds[ZC, FC any, ZV, FV any, ZR, FR any] interface {
	Zero() Categories[ZC, ZV, ZR]
	Fuzz() Categories[FC, FV, FR]
}

type Categories[C any, V any, R any] interface {
	Categories() C
	Values() V
	Refs() R
}
