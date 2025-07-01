// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

type NaturallyComparablePerson struct {
	Age  int
	Name string
}

func (s *NaturallyComparablePerson) Equal(o any) bool {
	if other, ok := o.(*NaturallyComparablePerson); ok {
		return *s == *other
	}
	return false
}

type ImplementablyComparablePerson struct {
	Age    int
	Name   string
	Phones []string
}

func (s *ImplementablyComparablePerson) Equal(o any) bool {
	other, ok := o.(*ImplementablyComparablePerson)
	if !ok {
		return false
	}

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
