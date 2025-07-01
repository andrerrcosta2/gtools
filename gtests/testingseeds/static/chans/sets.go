// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chans

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/chans/fuzz"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/chans/zero"
)

type Seeds struct{}

func (s *Seeds) Zero() zero.Pack {
	return zero.Pack{}
}

func (s *Seeds) Fuzz() fuzz.Pack {
	return fuzz.Pack{}
}
