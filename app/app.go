package app

import "fmt"

func Run(data *Storage) {
	fmt.Println("Добро пожаловать в Order Management System!")
	StartMenu(data)
}
