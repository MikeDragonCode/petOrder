package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DeleteUser deletes a user by ID and removes all their orders
func DeleteUser(storage *Storage) {
	if !RequireAuth() {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID пользователя для удаления: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Некорректный ID пользователя.")
		return
	}

	// Find user index
	userIndex := -1
	for i, u := range storage.Users {
		if u.ID == userID {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		fmt.Printf("❌ Пользователь с ID %d не найден.\n", userID)
		return
	}

	// Check if user has orders
	userOrders := []int{}
	for orderID, order := range storage.Orders {
		if order.UserID == userID {
			userOrders = append(userOrders, orderID)
		}
	}

	if len(userOrders) > 0 {
		fmt.Printf("⚠️ Пользователь имеет %d заказов. Все заказы будут удалены.\n", len(userOrders))
		fmt.Print("Продолжить? (y/n): ")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(strings.ToLower(confirm))
		if confirm != "y" && confirm != "yes" && confirm != "да" {
			fmt.Println("❌ Удаление отменено.")
			return
		}

		// Delete user's orders
		for _, orderID := range userOrders {
			delete(storage.Orders, orderID)
		}
	}

	// Remove user from slice
	storage.Users = append(storage.Users[:userIndex], storage.Users[userIndex+1:]...)

	fmt.Printf("✅ Пользователь с ID %d успешно удален.\n", userID)
	storage.AutoSave()
}

// DeleteOrder deletes an order by ID
func DeleteOrder(storage *Storage) {
	if !RequireAuth() {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID заказа для удаления: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Некорректный ID заказа.")
		return
	}

	// Check if order exists
	order, exists := storage.Orders[orderID]
	if !exists {
		fmt.Printf("❌ Заказ с ID %d не найден.\n", orderID)
		return
	}

	// Show order details before deletion
	fmt.Printf("🔍 Информация по заказу #%d\n", orderID)
	fmt.Printf("Клиент: %s\n", GetUserNameByID(storage.Users, order.UserID))
	for _, p := range order.Products {
		fmt.Printf("  - %s: %.2f₽\n", p.Name, p.Price)
	}
	fmt.Printf("  Общая сумма: %.2f₽\n", order.Total())

	fmt.Print("Удалить этот заказ? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.TrimSpace(strings.ToLower(confirm))
	if confirm != "y" && confirm != "yes" && confirm != "да" {
		fmt.Println("❌ Удаление отменено.")
		return
	}

	// Delete order
	delete(storage.Orders, orderID)
	fmt.Printf("✅ Заказ #%d успешно удален.\n", orderID)
	storage.AutoSave()
}
