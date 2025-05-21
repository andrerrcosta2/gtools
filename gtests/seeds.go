// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtests

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/arrays"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/chans"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/funcs"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/maps"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/prims"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/slices"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
)

var Arrays Sets[arrays.ArraySets] = arrays.SetOf()
var Chans Sets[chans.Categories] = chans.SetOf()
var Funcs Sets[funcs.Categories] = funcs.SetOf()
var Interf Sets[interf.Categories] = interf.SetOf()
var Maps Sets[maps.Categories] = maps.SetOf()
var Prims Sets[prims.Categories] = prims.SetOf()
var Slices Sets[slices.Categories] = slices.SetOf()
var Structs Sets[structs.Categories] = structs.SetOf()

type Sets[T any] interface {
	Refs() T
	Values() T
}
