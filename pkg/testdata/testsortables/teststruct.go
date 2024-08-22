// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testsortables

type TestStruct struct {
	Name string
	Age  int
}

func (s TestStruct) Equal(other TestStruct) bool {
	return s.Name == other.Name && s.Age == other.Age
}

func (s TestStruct) Less(other TestStruct) bool {
	return s.Age < other.Age
}
