// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package internal

type PackageInterface interface {
	PackageMethod() *PackageStruct
}

type PackageStruct struct {
	PackageField int
}

func asPackageInterface(p any) (cast PackageInterface, ok bool) {
	cast, ok = p.(PackageInterface)
	return
}
