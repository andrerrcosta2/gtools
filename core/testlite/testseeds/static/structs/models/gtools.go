// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

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

func (s ComparableValue) Equal(other ComparableValue) bool {
	return s.Name == other.Name && s.Age == other.Age
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

func (s *ComparableRef) Equal(other *ComparableRef) bool {

	return s.Name == other.Name && s.Age == other.Age

}

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

func (s SortableValue) Equal(o SortableValue) bool {
	return s.Name == o.Name && s.Age == o.Age
}

func (s SortableValue) Less(o SortableValue) bool {
	return s.Age < o.Age
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

func (s *SortableRef) Equal(other *SortableRef) bool {
	return s.Name == other.Name && s.Age == other.Age
}

func (s *SortableRef) Less(o *SortableRef) bool {
	return s.Age < o.Age
}

func NewSortableNode(value string) *SortableNode {
	return &SortableNode{
		Value: value,
	}
}

type SortableNode struct {
	Value string
}

func (s *SortableNode) Equal(o *SortableNode) bool {
	return s.Value == o.Value
}

func (s *SortableNode) Less(o *SortableNode) bool {
	return s.Value < o.Value
}

type NaturallyComparablePerson struct {
	Age  int
	Name string
}

func (s *NaturallyComparablePerson) Equal(o *NaturallyComparablePerson) bool {
	return *s == *o
}

type ImplementablyComparablePerson struct {
	Age    int
	Name   string
	Phones []string
}

func (s *ImplementablyComparablePerson) Equal(other *ImplementablyComparablePerson) bool {

	if s.Age != other.Age || s.Name != other.Name {
		return false
	}

	if len(s.Phones) != len(other.Phones) {
		return false
	}
	for i := range s.Phones {
		if s.Phones[i] != other.Phones[i] {
			return false
		}
	}

	return true
}
