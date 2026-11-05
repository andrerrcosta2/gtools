// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gstrings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

var (
	NE  = errors.New("node already exists")
	NF  = errors.New("node not found")
	NSN = errors.New("the node requested is not a struct node")
	NVN = errors.New("the node requested is not a value node")
)

type Tree[T prim.Hashable] interface {
	NewStructNode(key T) (StructNode[T], error)
	NewValueNode(key T) (Node[T], error)
	String() string
}

type Node[T prim.Hashable] interface {
	SetValue(val string)
}

type StructNode[T prim.Hashable] interface {
	Node[T]
	AddStructChild(key T) error
	GetStructChild(key T) (StructNode[T], error)
	AddValueChild(key T) error
	GetValueChild(key T) (Node[T], error)
}

func TreeBuilder[T prim.Hashable](tabSize uint8, showKeys bool) Tree[T] {
	return &tree[T]{
		nodes:    make(map[T]Node[T]),
		tabSize:  tabSize,
		showKeys: showKeys,
	}
}

type tree[T prim.Hashable] struct {
	nodes    map[T]Node[T]
	tabSize  uint8
	showKeys bool
}

func (t *tree[T]) NewStructNode(key T) (StructNode[T], error) {
	if _, exists := t.nodes[key]; exists {
		return nil, NE
	}
	node := newStructNode(key, t.tabSize, t.showKeys)
	t.nodes[key] = node
	return node, nil
}

func (t *tree[T]) NewValueNode(key T) (Node[T], error) {
	if _, exists := t.nodes[key]; exists {
		return nil, NE
	}

	node := &valueNode[T]{key: key}
	t.nodes[key] = node
	return node, nil
}

func (t *tree[T]) String() string {
	var b strings.Builder
	for _, node := range t.nodes {
		writeNode[T](&b, node, 0, t.tabSize, t.showKeys)
	}
	return b.String()
}

type structNode[T prim.Hashable] struct {
	key      T
	children map[T]Node[T]
	tabSize  uint8
	showKeys bool
}

func newStructNode[T prim.Hashable](key T, tabSize uint8, showKeys bool) *structNode[T] {
	return &structNode[T]{
		key:      key,
		children: make(map[T]Node[T]),
		tabSize:  tabSize,
		showKeys: showKeys,
	}
}

func (n *structNode[T]) AddStructChild(key T) error {
	if _, exists := n.children[key]; exists {
		return NE
	}
	n.children[key] = newStructNode(key, n.tabSize, n.showKeys)
	return nil
}

func (n *structNode[T]) GetStructChild(key T) (StructNode[T], error) {
	if child, ok := n.children[key]; ok {
		if structChild, ok := child.(*structNode[T]); ok {
			return structChild, nil
		}
		return nil, NSN
	}
	return nil, NF
}

func (n *structNode[T]) AddValueChild(key T) error {
	if _, exists := n.children[key]; exists {
		return NE
	}
	n.children[key] = &valueNode[T]{key: key}
	return nil
}

func (n *structNode[T]) GetValueChild(key T) (Node[T], error) {
	if child, ok := n.children[key]; ok {
		if valueChild, ok := child.(*valueNode[T]); ok {
			return valueChild, nil
		}
		return nil, NVN
	}
	return nil, NF
}

func (n *structNode[T]) SetValue(val string) {}

type valueNode[T prim.Hashable] struct {
	key   T
	value string
}

func (v *valueNode[T]) SetValue(val string) {
	v.value = val
}

func writeNode[T prim.Hashable](b *strings.Builder, node Node[T], level uint16, tabSize uint8, showKeys bool) {
	switch n := node.(type) {
	case *structNode[T]:
		if showKeys {
			b.WriteString(fmt.Sprintf("%s%v {\n", writeTab(level, tabSize), n.key))
		}

		for _, child := range n.children {
			writeNode[T](b, child, level+1, tabSize, showKeys)
		}

		if showKeys {
			b.WriteString(fmt.Sprintf("%s}\n", writeTab(level, tabSize)))
		}

	case *valueNode[T]:
		if showKeys {
			b.WriteString(fmt.Sprintf("%s%v: %s\n", writeTab(level, tabSize), n.key, n.value))
		} else {
			b.WriteString(fmt.Sprintf("%s%s\n", writeTab(level, tabSize), n.value))
		}
	}
}

func writeTab(level uint16, size uint8) string {
	return strings.Repeat(strings.Repeat(" ", int(size)), int(level))
}
