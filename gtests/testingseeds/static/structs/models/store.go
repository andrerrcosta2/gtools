// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

type Product struct {
	Id    int
	Name  string
	Price float64
}

func ProductAsValue(id int, name string, price float64) Product {
	return Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

func ProductAsPointer(id int, name string, price float64) *Product {
	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

func ProductAsRandValue() Product {
	return Product{
		Id:    random.SingleOf[int](),
		Name:  random.Alphanumeric(1, 2, 30).At(0),
		Price: random.SingleOf[float64](),
	}
}

func ProductAsRandRef() *Product {
	return &Product{
		Id:    random.SingleOf[int](),
		Name:  random.Alphanumeric(1, 2, 30).At(0),
		Price: random.SingleOf[float64](),
	}
}
