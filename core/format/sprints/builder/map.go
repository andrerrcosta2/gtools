// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package builder

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"strings"
)

func Map(tab indent.Tab, name string) *MapBuilder {
	return &MapBuilder{
		tab:    tab,
		name:   name,
		fields: make(map[string]string),
	}
}

type MapBuilder struct {
	tab    indent.Tab
	name   string
	fields map[string]string
}

func (b *MapBuilder) Entry(key, value string) *MapBuilder {
	b.fields[key] = value
	return b
}

func (b *MapBuilder) String() string {
	sb := strings.Builder{}
	sb.WriteString(sprints.TypeValue(b.tab, b.name, "{"))

	for key, value := range b.fields {
		sb.WriteString(sprints.TypeValue(b.tab.Inc(), key, ": "))
		sb.WriteString(sprints.TypeValue(b.tab.Inc(), value, ","))
	}

	sb.WriteString("}")
	return sb.String()
}
