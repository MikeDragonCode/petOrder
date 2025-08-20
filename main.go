package main

import (
	"awesomeProject/app"
)

func main() {
	storage := app.InitStorage()
	app.Run(storage)
}
