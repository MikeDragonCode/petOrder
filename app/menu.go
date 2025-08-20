package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StartMenu(data *Data) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n📋 Меню:")
		fmt.Println("1. Показать все заказы")
		fmt.Println("2. Показать всех пользователей")
		fmt.Println("3. Найти заказ по ID")
		fmt.Println("0. Выйти")

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
		case 0:
			fmt.Println("👋 Выход из программы.")
			return
		default:
			fmt.Println("❌ Неизвестная команда.")
		}
	}
}
