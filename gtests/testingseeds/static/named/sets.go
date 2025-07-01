// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package named

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/named/fuzz"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/named/zero"
)

type Seeds struct{}

func (s *Seeds) Zero() zero.Pack {
	return zero.Pack{}
}

func (s *Seeds) Fuzz() fuzz.Pack {
	return fuzz.Pack{}
}
