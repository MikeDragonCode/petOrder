package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

type Order struct {
	ID     int
	UserID int
	Amount float64
}

func NewUser(id int, name string, email string) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func NewOrder(id int, userid int, amount float64) Order {
	return Order{
		ID:     id,
		UserID: userid,
		Amount: amount,
	}
}
func main() {
	fmt.Println("Добро пожаловать в Order Management System!")
	users := []User{}

	user1 := NewUser(1, "Миша", "misha@email.com")
	users = append(users, user1)
	user2 := NewUser(2, "Вася", "vasya@email.com")
	users = append(users, user2)

	orders := []Order{}
	order1 := NewOrder(1, 1, 25.50)
	orders = append(orders, order1)
	order2 := NewOrder(2, 2, 33.99)
	orders = append(orders, order2)

	fmt.Println("Список пользователей:")
	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.Email)
	}
	fmt.Println("Список заказов:")
	for _, order := range orders {
		fmt.Println(order.ID, order.UserID, order.Amount)
	}
}
