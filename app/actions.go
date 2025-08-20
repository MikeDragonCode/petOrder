package app

import (
	"awesomeProject/order"
	"awesomeProject/user"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowOrders(orders map[int]*order.Order, users []user.User) {
	fmt.Println("📦 Список заказов:")
	for id, o := range orders {
		fmt.Printf("Заказ #%d — клиент: %s\n", id, GetUserNameByID(users, o.UserID))
		for _, p := range o.Products {
			fmt.Printf("  - %s: %.2f₽\n", p.Name, p.Price)
		}
		fmt.Printf("  Общая сумма: %.2f₽\n\n", o.Total())
	}
}

func ShowUsers(users []user.User) {
	fmt.Println("👤 Список пользователей:")
	for _, u := range users {
		fmt.Printf("ID: %d — %s — %s\n", u.ID, u.Name, u.Email)
	}
}
func FindOrderByID(orders map[int]*order.Order, users []user.User) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID заказа: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Неверный формат ID.")
		return
	}

	order, exists := orders[id]
	if !exists {
		fmt.Printf("❌ Заказ с ID %d не найден.\n", id)
		return
	}

	fmt.Printf("🔍 Информация по заказу #%d\n", id)
	fmt.Printf("Клиент: %s\n", GetUserNameByID(users, order.UserID))
	for _, p := range order.Products {
		fmt.Printf("  - %s: %.2f₽\n", p.Name, p.Price)
	}
	fmt.Printf("  Общая сумма: %.2f₽\n", order.Total())
}
