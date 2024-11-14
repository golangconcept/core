package domain

type Product struct {
	ID    string
	Name  string
	Price Money
}

func NewProduct(id, name string, price Money) *Product {
	return &Product{ID: id, Name: name, Price: price}
}
