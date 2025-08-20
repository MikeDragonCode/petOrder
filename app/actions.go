package app

import (
	"awesomeProject/order"
	"awesomeProject/product"
	"awesomeProject/user"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Показ всех заказов
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

// Показ всех пользователей
func ShowUsers(users []user.User) {
	fmt.Println("👤 Список пользователей:")
	for _, u := range users {
		fmt.Printf("ID: %d — %s — %s\n", u.ID, u.Name, u.Email)
	}
}

// Поиск заказа по ID
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

// Добавление нового пользователя
func AddUser(data *Storage) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя пользователя: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Введите email пользователя: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Генерация нового ID
	newID := 1
	for _, u := range data.Users {
		if u.ID >= newID {
			newID = u.ID + 1
		}
	}

	newUser := user.New(newID, name, email)
	data.Users = append(data.Users, newUser)

	fmt.Printf("✅ Пользователь %s успешно добавлен с ID %d\n", name, newID)
}

// Добавление нового заказа
func AddNewOrder(orders map[int]*order.Order, users []user.User, products []product.Product) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID пользователя для заказа: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Некорректный ID пользователя.")
		return
	}

	// Проверка пользователя
	var userExists bool
	for _, u := range users {
		if u.ID == userID {
			userExists = true
			break
		}
	}
	if !userExists {
		fmt.Println("❌ Пользователь с таким ID не найден.")
		return
	}

	// Выводим список товаров
	fmt.Println("📦 Доступные товары:")
	for i, p := range products {
		fmt.Printf("%d. %s — %.2f₽\n", i+1, p.Name, p.Price)
	}

	// Выбор товаров
	fmt.Print("Введите номера товаров через запятую (например: 1,3): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	ids := strings.Split(input, ",")

	var selectedProducts []product.Product
	for _, idStr := range ids {
		idStr = strings.TrimSpace(idStr)
		idx, err := strconv.Atoi(idStr)
		if err != nil || idx < 1 || idx > len(products) {
			fmt.Printf("⚠️ Пропущен некорректный номер: %s\n", idStr)
			continue
		}
		selectedProducts = append(selectedProducts, products[idx-1])
	}

	if len(selectedProducts) == 0 {
		fmt.Println("❌ Не выбрано ни одного товара.")
		return
	}

	// Генерация нового ID заказа
	newID := 1
	for id := range orders {
		if id >= newID {
			newID = id + 1
		}
	}

	newOrder := order.New(newID, userID, selectedProducts)
	orders[newID] = newOrder
	fmt.Printf("✅ Заказ #%d успешно создан для пользователя %d\n", newID, userID)
}
