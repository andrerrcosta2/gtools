// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"math/rand"
)

var EmptyZeroInst = new(Empty)

func EmptyAsValue() Empty {
	return Empty{}
}

func EmptyAsRef() *Empty {
	return new(Empty)
}

type Empty struct{}

func FunctionAsRandValue[T any]() Function[T] {
	return Function[T]{
		fn: func() T {
			return random.SingleOf[T]()
		},
	}
}

func FunctionAsRandRef[T any]() *Function[T] {
	var fn = FunctionAsRandValue[T]()
	return &fn
}

type Function[T any] struct {
	fn func() T
}

func (f *Function[T]) String() string {
	var zero T
	return fmt.Sprintf("func() %T", zero)
}

func (f *Function[T]) Call() {
	f.fn()
}

var OneDataZeroInst = new(OneData)

func OneDataAsRef(name string) *OneData {
	return &OneData{
		Name: name,
	}
}

func OneDataAsValue(name string) OneData {
	return OneData{
		Name: name,
	}
}

func OneDataAsRandRef() *OneData {
	return &OneData{
		Name: random.SingleOf[string](),
	}
}

func OneDataAsRandValue() OneData {
	return OneData{
		Name: random.SingleOf[string](),
	}
}

type OneData struct {
	Name string
}

func (o *OneData) String() string {
	return o.Name
}

var PublicZeroInst = new(Public)

func PublicAsRef() *Public {
	return &Public{}
}

func PublicAsValue() Public {
	return Public{}
}

func PublicAsRandRef() *Public {
	return &Public{}
}

func PublicAsRandValue() Public {
	return Public{}
}

type Public struct{}

func (p *Public) PublicMethod() int {
	return rand.Intn(1000)
}

var SimpleZeroInst = new(Simple)

func SimpleAsRef() *Simple {
	return &Simple{}
}

func SimpleAsValue() Simple {
	return Simple{}
}

func SimpleAsRandRef() *Simple {
	return &Simple{}
}

func SimpleAsRandValue() Simple {
	return Simple{}
}

type Simple struct{}

func (s *Simple) Simple() int {
	return rand.Intn(1000)
}
