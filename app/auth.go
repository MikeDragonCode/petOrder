package app

import (
	"awesomeProject/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var currentUser *user.User

// IsAuthenticated returns true if a user is currently logged in
func IsAuthenticated() bool {
	return currentUser != nil
}

// GetCurrentUser returns the currently logged-in user
func GetCurrentUser() *user.User {
	return currentUser
}

// Login prompts user for email and logs them in if found
func Login(storage *Storage) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите email для входа: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Find user by email
	for _, u := range storage.Users {
		if u.Email == email {
			currentUser = &u
			fmt.Printf("✅ Добро пожаловать, %s!\n", u.Name)
			return
		}
	}

	fmt.Println("❌ Пользователь с таким email не найден.")
}

// Logout logs out the current user
func Logout() {
	if currentUser != nil {
		fmt.Printf("👋 До свидания, %s!\n", currentUser.Name)
		currentUser = nil
	} else {
		fmt.Println("❌ Вы не авторизованы.")
	}
}

// RequireAuth checks if user is authenticated and returns true if so
func RequireAuth() bool {
	if !IsAuthenticated() {
		fmt.Println("🔐 Для выполнения этого действия необходимо войти в систему.")
		return false
	}
	return true
}
