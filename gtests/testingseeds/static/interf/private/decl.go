// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package private

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf/private/internal"
)

func AsInternalPackageInterface(p any) (cast internal.PackageInterface, ok bool) {
	cast, ok = p.(internal.PackageInterface)
	return
}

type PackageInterface interface {
	PackageMethod() *PackageStruct
}

type PackageStruct = internal.PackageStruct
