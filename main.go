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

	// Продукты
	p1 := product.NewProduct(1, "iPhone 15", 999.99)
	p2 := product.NewProduct(2, "MacBook Pro", 1999.99)
	products := []product.Product{p1, p2}

	// Пользователи
	user1 := user.New(1, "Миша", "misha@email.com")
	user2 := user.New(2, "Вася", "vasya@email.com")
	newUser := user.New(3, "Анна", "anna@email.com")
	users := []user.User{user1, user2, newUser}

	// Добавляем новый продукт
	newProduct := product.NewProduct(3, "iPad Air", 799.99)
	products = append(products, newProduct)

	// Заказы через Мапу
	ordersMap := map[int]*order.Order{
		1: order.New(1, user1.ID, []product.Product{p1}),
		2: order.New(2, user2.ID, []product.Product{p1, p2}),
		3: order.New(3, newUser.ID, []product.Product{newProduct}),
	}
	delete(ordersMap, 2)
	ordersMap[1].AddProduct(product.NewProduct(4, "Apple Watch", 499.99))
	watch := product.NewProduct(4, "Apple Watch", 499.99)
	products = append(products, watch)
	ordersMap[1].AddProduct(watch)

	// Вывод данных
	fmt.Println("Список товаров:")
	for _, p := range products {
		fmt.Printf("ID: %d — %s — %.2f₽\n", p.ID, p.Name, p.Price)
	}

	fmt.Println("Список пользователей:")
	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Email)
	}

	fmt.Println("Список заказов:")
	for id, o := range ordersMap {
		fmt.Printf("Заказ #%d — клиент %s\n", id, GetUserNameByID(users, o.UserID))

		for _, prod := range o.Products {
			fmt.Printf("  - %s: %.2f₽\n", prod.Name, prod.Price)
		}

		fmt.Printf("  Общая сумма: %.2f₽\n\n", o.Total()) // используем метод Total()
	}
}
