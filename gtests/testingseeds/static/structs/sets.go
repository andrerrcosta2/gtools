// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/fuzz"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/zero"
)

type Seeds struct{}

func (s *Seeds) Zero() zero.Pack {
	return zero.Pack{}
}

func (s *Seeds) Fuzz() fuzz.Pack {
	return fuzz.Pack{}
}
