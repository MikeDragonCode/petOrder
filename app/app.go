package app

import (
	"fmt"
)

func Run() {
	fmt.Println("Добро пожаловать в Order Management System!")
	data := InitData()
	StartMenu(data)
}
