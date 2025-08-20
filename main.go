package main

import (
	"awesomeProject/order"
	"awesomeProject/product"
	"awesomeProject/user"
	"fmt"
)

func GetUserNameByID(users []user.User, id int) string {
	for _, u := range users {
		if u.ID == id {
			return u.Name
		}
	}
	return "Неизвестный пользователь"
}

func main() {
	fmt.Println("Добро пожаловать в Order Management System!")

	// Пользователи
	users := []user.User{
		user.New(1, "Миша", "misha@email.com"),
		user.New(2, "Вася", "vasya@email.com"),
	}

	// Заказы
	orders := []order.Order{
		order.New(1, 1, 25.50),
		order.New(2, 2, 33.99),
	}

	// Продукты
	products := []product.Product{}
	p1 := product.NewProduct(1, "iPhone 15", 999.99)
	p2 := product.NewProduct(2, "MacBook Pro", 1999.99)
	products = append(products, p1, p2)

	fmt.Println("Список товаров:")
	for _, p := range products {
		fmt.Printf("ID: %d — %s — %.2f₽\n", p.ID, p.Name, p.Price)
	}
	fmt.Println("Список пользователей:")
	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Email)
	}
	fmt.Println("Список заказов:")
	for _, o := range orders {
		fmt.Printf("Заказ #%d — клиент %s — сумма: %.2f\n", o.ID, GetUserNameByID(users, o.UserID), o.Amount)
	}
}
