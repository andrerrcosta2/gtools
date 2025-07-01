// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package funcs

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/funcs/consumer"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/funcs/supplier"
)

type Seeds struct{}

func (s *Seeds) Consumers() consumer.Pack {
	return consumer.Pack{}
}

func (s *Seeds) Suppliers() supplier.Pack {
	return supplier.Pack{}
}
