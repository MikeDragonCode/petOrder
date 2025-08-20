package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func startMenu(s *Storage) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n📋 Меню:")
		fmt.Println("1. Показать все заказы")
		fmt.Println("2. Показать всех пользователей")
		fmt.Println("0. Выйти")

		fmt.Print("Выберите действие: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Некорректный ввод. Введите число.")
			continue
		}

		switch choice {
		case 1:
			showAllOrders(s)
		case 2:
			showAllUsers(s)
		case 0:
			fmt.Println("👋 Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
