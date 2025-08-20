package app

import (
	"awesomeProject/order"
	"awesomeProject/product"
	"awesomeProject/user"
)

var Products = []product.Product{
	product.NewProduct(1, "iPhone 15", 999.99),
	product.NewProduct(2, "MacBook Pro", 1999.99),
	product.NewProduct(3, "iPad Air", 799.99),
	product.NewProduct(4, "Apple Watch", 499.99),
}

var Users = []user.User{
	user.New(1, "Миша", "misha@email.com"),
	user.New(2, "Вася", "vasya@email.com"),
	user.New(3, "Анна", "anna@email.com"),
}

var OrdersMap = map[int]*order.Order{
	1: order.New(1, 1, []product.Product{Products[0]}),
	2: order.New(2, 2, []product.Product{Products[0], Products[1]}),
	3: order.New(3, 3, []product.Product{Products[2]}),
}
