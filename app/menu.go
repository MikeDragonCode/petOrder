package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StartMenu(data *Storage) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Show authentication status
		if IsAuthenticated() {
			fmt.Printf("\n🔐 Авторизован как: %s (%s)\n", GetCurrentUser().Name, GetCurrentUser().Email)
		} else {
			fmt.Println("\n🔐 Не авторизован")
		}

		fmt.Println("\n📋 Меню:")
		fmt.Println("1. Показать все заказы")
		fmt.Println("2. Показать всех пользователей")
		fmt.Println("3. Найти заказ по ID")
		fmt.Println("4. Добавить нового пользователя")
		fmt.Println("5. Добавить новый заказ")
		fmt.Println("6. 📊 Показать отчеты")
		fmt.Println("7. 🔐 Войти в систему")
		fmt.Println("8. 🚪 Выйти из системы")
		fmt.Println("9. ❌ Удалить пользователя")
		fmt.Println("10. ❌ Удалить заказ")
		fmt.Println("0. Выйти из программы")

		fmt.Print("Выберите действие: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Введите корректное число.")
			continue
		}

		switch choice {
		case 1:
			ShowOrders(data.Orders, data.Users)
		case 2:
			ShowUsers(data.Users)
		case 3:
			FindOrderByID(data.Orders, data.Users)
		case 4:
			AddUser(data)
		case 5:
			AddNewOrder(data)
		case 6:
			ShowReports(data)
		case 7:
			Login(data)
		case 8:
			Logout()
		case 9:
			DeleteUser(data)
		case 10:
			DeleteOrder(data)
		case 0:
			fmt.Println("👋 Выход из программы.")
			return
		default:
			fmt.Println("❌ Неизвестная команда.")
		}
	}
}
