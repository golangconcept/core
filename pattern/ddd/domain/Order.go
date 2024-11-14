package domain

type Order struct {
	ID        string
	Products  []*Product
	Total     Money
	OrderDate string
}

func NewOrder(id, orderDate string) *Order {
	return &Order{
		ID:        id,
		Products:  make([]*Product, 0),
		Total:     NewMoney(0, "USD"),
		OrderDate: orderDate,
	}
}

func (o *Order) AddProduct(product *Product) {
	o.Products = append(o.Products, product)
	o.Total = NewMoney(o.Total.Amount+product.Price.Amount, o.Total.Currency)
}
