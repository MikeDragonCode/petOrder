package app

import "fmt"

func Run(data *Storage) {
	fmt.Println("Добро пожаловать в Order Management System!")

	// Load data from file on startup
	if err := data.LoadFromFile(); err != nil {
		fmt.Printf("⚠️ Ошибка загрузки данных: %v\n", err)
	}

	StartMenu(data)

	// Save data to file before exit
	fmt.Println("💾 Сохранение данных...")
	if err := data.SaveToFile(); err != nil {
		fmt.Printf("⚠️ Ошибка сохранения данных: %v\n", err)
	}
}
