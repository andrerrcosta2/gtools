// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prims

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/prims/fuzz"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/prims/zero"
)

type Seeds struct{}

func (s *Seeds) Zero() zero.Pack {
	return zero.Pack{}
}

func (s *Seeds) Fuzz() fuzz.Pack {
	return fuzz.Pack{}
}
