// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package slices

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/slices/fuzz"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/slices/zero"
)

type Seeds struct{}

func (s *Seeds) Zero() zero.Pack {
	return zero.Pack{}
}

func (s *Seeds) Fuzz() fuzz.Pack {
	return fuzz.Pack{}
}
