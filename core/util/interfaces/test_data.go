// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package interfaces

type testInterfaceA interface {
	MethodA() string
	MethodB() int
}

type testInterfaceB interface {
	MethodC() bool
}

type interfaceAImpl1 struct {
	A string
	B int
}

func (s interfaceAImpl1) MethodA() string {
	return s.A
}

func (s interfaceAImpl1) MethodB() int {
	return s.B
}

type interfaceAImpl2 struct {
	A string
	B int
}

func (s interfaceAImpl2) MethodA() string {
	return s.A
}

func (s interfaceAImpl2) MethodB() int {
	return s.B
}

type interfaceBImpl1 struct{}

func (s interfaceBImpl1) MethodC() bool {
	return true
}

type interfaceBImpl2 struct{}

func (s interfaceBImpl2) MethodC() bool {
	return false
}
