// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build testseed

package testseed

func NewSortableValue(name string, age int) SortableValue {
	return SortableValue{
		Name: name,
		Age:  age,
	}
}

type SortableValue struct {
	Name string
	Age  int
}

func (s SortableValue) Equal(o any) bool {
	if other, ok := o.(SortableValue); ok {
		return s.Name == other.Name && s.Age == other.Age
	}
	return false
}

func (s SortableValue) Less(o any) bool {
	if other, ok := o.(SortableValue); ok {
		return s.Age < other.Age
	}
	return false
}

func NewSortableRef(name string, age int) *SortableRef {
	return &SortableRef{
		Name: name,
		Age:  age,
	}
}

type SortableRef struct {
	Name string
	Age  int
}

func (s *SortableRef) Equal(o any) bool {
	if other, ok := o.(*SortableRef); ok {
		return s.Name == other.Name && s.Age == other.Age
	}
	return false
}

func (s *SortableRef) Less(o any) bool {
	if other, ok := o.(*SortableRef); ok {
		return s.Age < other.Age
	}
	return false
}

func NewSortableNode(value string) *SortableNode {
	return &SortableNode{
		Value: value,
	}
}

type SortableNode struct {
	Value string
}

func (s *SortableNode) Equal(o any) bool {
	if other, ok := o.(*SortableNode); ok {
		return s.Value == other.Value
	}
	return false
}

func (s *SortableNode) Less(o any) bool {
	if other, ok := o.(*SortableNode); ok {
		return s.Value < other.Value
	}
	return false
}

func NewComparableValue(name string, age int) ComparableValue {
	return ComparableValue{
		Name: name,
		Age:  age,
	}
}

type ComparableValue struct {
	Name string
	Age  int
}

func (s ComparableValue) Equal(o any) bool {
	if other, ok := o.(ComparableValue); ok {
		return s.Name == other.Name && s.Age == other.Age
	}
	return false
}

func NewComparableRef(name string, age int) *ComparableRef {
	return &ComparableRef{
		Name: name,
		Age:  age,
	}
}

type ComparableRef struct {
	Name string
	Age  int
}

func (s *ComparableRef) Equal(o any) bool {
	if other, ok := o.(*ComparableRef); ok {
		return s.Name == other.Name && s.Age == other.Age
	}
	return false
}
