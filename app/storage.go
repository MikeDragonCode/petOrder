package app

import (
	"awesomeProject/order"
	"awesomeProject/product"
	"awesomeProject/user"
)

type Storage struct {
	Users    []user.User
	Products []product.Product
	Orders   map[int]*order.Order
}

func InitStorage() *Storage {
	p1 := product.NewProduct(1, "iPhone 15", 999.99)
	p2 := product.NewProduct(2, "MacBook Pro", 1999.99)
	p3 := product.NewProduct(3, "iPad Air", 799.99)

	u1 := user.New(1, "Миша", "misha@email.com")
	u2 := user.New(2, "Вася", "vasya@email.com")
	u3 := user.New(3, "Анна", "anna@email.com")

	orders := map[int]*order.Order{
		1: order.New(1, u1.ID, []product.Product{p1}),
		2: order.New(2, u2.ID, []product.Product{p1, p2}),
		3: order.New(3, u3.ID, []product.Product{p3}),
	}

	return &Storage{
		Users:    []user.User{u1, u2, u3},
		Products: []product.Product{p1, p2, p3},
		Orders:   orders,
	}
}
