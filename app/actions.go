package app

import (
	"fmt"
)

func ShowOrders() {
	fmt.Println("📦 Список заказов:")
	for id, o := range OrdersMap {
		userName := getUserNameByID(Users, o.UserID)
		fmt.Printf("Заказ #%d — клиент: %s\n", id, userName)
		for _, p := range o.Products {
			fmt.Printf("  - %s: %.2f₽\n", p.Name, p.Price)
		}
		fmt.Printf("  Общая сумма: %.2f₽\n\n", o.Total())
	}
}

func ShowUsers() {
	fmt.Println("👤 Список пользователей:")
	for _, u := range Users {
		fmt.Printf("ID: %d — %s — %s\n", u.ID, u.Name, u.Email)
	}
}
