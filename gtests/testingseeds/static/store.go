// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

func GetProductAsValue(id int, name string, price float64) Product {
	return Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

func GetProductAsPointer(id int, name string, price float64) *Product {
	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}

type Product struct {
	Id    int
	Name  string
	Price float64
}
