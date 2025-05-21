// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package builder

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"strings"
)

func Struct(tab indent.Tab, name string) *StructBuilder {
	return &StructBuilder{tab: tab, name: name}
}

type StructBuilder struct {
	tab    indent.Tab
	name   string
	fields []string
}

func (s *StructBuilder) Field(key, value string) *StructBuilder {
	if key == "" {
		return s
	} else if value == "" {
		return s
	}
	s.fields = append(s.fields, key+": "+sprints.FieldVal(value))
	return s
}

func (s *StructBuilder) Fields(fields ...Tuple) *StructBuilder {
	n := len(s.fields)
	s.fields = append(s.fields, make([]string, len(fields))...) // Pre-allocate

	for i, field := range fields {
		s.fields[n+i] = field.Key + ": " + sprints.FieldVal(field.Value)
	}
	return s
}

func (s *StructBuilder) String() string {
	sb := strings.Builder{}
	sb.Grow(len(s.name) + len(s.fields)*20) // Rough estimate for efficiency

	sb.WriteString(s.tab.Sprint(s.name + "{")) // Open struct

	if len(s.fields) > 0 {
		sb.WriteString("\n\t" + s.tab.Sprint(strings.Join(s.fields, ",\n\t")))
	}

	sb.WriteString(",\n" + s.tab.Sprint("}")) // Close struct
	return sb.String()
}

func NewTuple(key string, value string) Tuple {
	return Tuple{Key: key, Value: value}
}

type Tuple struct {
	Key   string
	Value string
}
