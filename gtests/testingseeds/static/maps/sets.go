// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/maps/fuzz"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/maps/zero"
)

type Seeds struct{}

func (s *Seeds) Zero() zero.Pack {
	return zero.Pack{}
}

func (s *Seeds) Fuzz() fuzz.Pack {
	return fuzz.Pack{}
}
