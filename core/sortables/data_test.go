// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

type TestSortable struct {
	Name string
	Age  int
}

func (s TestSortable) Equal(o any) bool {
	if other, ok := o.(TestSortable); ok {
		return s.Name == other.Name && s.Age == other.Age
	}
	return false
}

func (s TestSortable) Less(o any) bool {
	if other, ok := o.(TestSortable); ok {
		return s.Age < other.Age
	}
	return false
}

type testInterface interface {
	MethodA() string
	MethodB() int
}

type TestInterfaceImpl1 struct {
	A string
	B int
}

func (s TestInterfaceImpl1) MethodA() string {
	return s.A
}

func (s TestInterfaceImpl1) MethodB() int {
	return s.B
}

type TestInterfaceImpl2 struct {
	A string
	B int
}

func (s TestInterfaceImpl2) MethodA() string {
	return s.A
}

func (s TestInterfaceImpl2) MethodB() int {
	return s.B
}

type TestNonComparable struct {
	id   int
	data []string
}
