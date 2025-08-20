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

func PrintOrderByID(ordersMap map[int]*order.Order, users []user.User, id int) {
	orderPtr, ok := ordersMap[id]
	if !ok {
		fmt.Printf("Заказ с ID %d не найден.\n", id)
		return
	}

	fmt.Printf("Информация по заказу #%d:\n", id)
	fmt.Printf("Клиент: %s\n", GetUserNameByID(users, orderPtr.UserID))
	for _, p := range orderPtr.Products {
		fmt.Printf("  - %s: %.2f₽\n", p.Name, p.Price)
	}
	fmt.Printf("  Общая сумма: %.2f₽\n", orderPtr.Total())
}

func UpdateEmailByID(users *[]user.User, id int, newEmail string) bool {
	for i, u := range *users {
		if u.ID == id {
			(*users)[i].Email = newEmail
			return true
		}
	}
	return false
}

func DeleteUserByID(users *[]user.User, id int) bool {
	for i, u := range *users {
		if u.ID == id {
			// Удаление элемента из слайса: всё до i + всё после i
			*users = append((*users)[:i], (*users)[i+1:]...)
			return true
		}
	}
	return false
}

func FindOrdersByProductName(orders map[int]*order.Order, productName string) []int {
	var result []int
	for id, ord := range orders {
		for _, p := range ord.Products {
			if p.Name == productName {
				result = append(result, id)
				break
			}
		}
	}
	return result
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
	fmt.Println("Удаляем пользователя с ID: 2")
	if DeleteUserByID(&users, 2) {
		fmt.Println("Пользователь удален успешно")
	} else {
		fmt.Println("Пользователь не найден.")

	}

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

	//  Поиск товара
	fmt.Println("Поиск заказа с ID 1:")
	PrintOrderByID(ordersMap, users, 1)

	fmt.Println("Поиск заказа с ID 99:")
	PrintOrderByID(ordersMap, users, 99)

	fmt.Println("\nПоиск заказов, содержащих продукт 'iPhone 15':")
	matchingOrders := FindOrdersByProductName(ordersMap, "iPhone 15")
	if len(matchingOrders) == 0 {
		fmt.Println("Таких заказов не найдено.")
	} else {
		fmt.Println("Найденные ID заказов:", matchingOrders)
	}

	// Меняем почту юзеру

	fmt.Println("Смена почты пользователю ID 1:")
	if UpdateEmailByID(&users, 1, "newmail@checkmail.com") {
		fmt.Printf("Почта изменена успешна!")
	} else {
		fmt.Println("Пользователь не найден!!!")
	}

	// Вывод данных
	fmt.Println("Список товаров:")
	for _, p := range products {
		fmt.Printf("ID: %d — %s — %.2f₽\n", p.ID, p.Name, p.Price)
	}

	fmt.Println("\nСписок пользователей:")
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
