package app

import (
	"awesomeProject/user"
	"fmt"
)

func showAllOrders(s *Storage) {
	for id, o := range s.Orders {
		username := getUserNameByID(s.Users, o.UserID)
		fmt.Printf("Заказ #%d — клиент %s\n", id, username)
		for _, p := range o.Products {
			fmt.Printf("  - %s: %.2f₽\n", p.Name, p.Price)
		}
		fmt.Printf("  Общая сумма: %.2f₽\n\n", o.Total())
	}
}

func showAllUsers(s *Storage) {
	for _, u := range s.Users {
		fmt.Printf("ID: %d — %s — %s\n", u.ID, u.Name, u.Email)
	}
}

func getUserNameByID(users []user.User, id int) string {
	for _, u := range users {
		if u.ID == id {
			return u.Name
		}
	}
	return "Неизвестный пользователь"
}
