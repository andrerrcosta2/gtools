// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package _testdata

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func init() {
	SelfReferencedNode.Next = SelfReferencedNode
	for i := 0; i < 5; i++ {
		CyclicTree = &Tree{Value: i, Left: CyclicTree, Right: CyclicTree}
	}
	ReflectedValue.SetInt(100)
}

var AnonymousInterface interface{} = struct {
	Name string
	Age  int
}{"Alice", 30}

type Node struct {
	Value int
	Next  *Node
}

var SelfReferencedNode = &Node{
	Value: 1,
}

type MyInt int

func (m MyInt) String() string { return fmt.Sprintf("MyInt(%d)", m) }

var NamedInt = MyInt(42)

type Person struct {
	Name string
	age  int // Unexported field
}

type Person2 struct {
	Name string
	Age  int
}

var UnexportedField = Person{Name: "Alice", age: 30}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

var CyclicTree *Tree

func FunctionAdd(a, b int) int { return a + b }

var SimpleChannel = make(chan int)

var complexNumber = complex(3, 4)

var X int = 42

var UnsafePointer = unsafe.Pointer(&X)

type MyString = string

var StringAlias = MyString("hello")

var EmptyInterface interface{} = nil

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var StructWithTags = User{Name: "Alice", Age: 30}

var Time = time.Now()
var Duration = 5 * time.Second

var SimpleError = errors.New("something went wrong")

var ContextWithValue = context.WithValue(context.Background(), "key", "value")

type MyType int

func (m MyType) String() string { return fmt.Sprintf("MyType(%d)", m) }

var CustomStringer = MyType(42)

var ReflectedValue = reflect.New(reflect.TypeOf(42)).Elem()
