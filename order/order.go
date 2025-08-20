package order

import "awesomeProject/product"

type Order struct {
	ID       int
	UserID   int
	Products []product.Product
}

func New(id int, userID int, products []product.Product) *Order {
	return &Order{
		ID:       id,
		UserID:   userID,
		Products: products,
	}
}
func (o Order) Total() float64 {
	total := 0.0
	for _, p := range o.Products {
		total += p.Price
	}
	return total
}

func (o *Order) AddProduct(p product.Product) {
	o.Products = append(o.Products, p)
}
