package app

import (
	"awesomeProject/order"
	"awesomeProject/product"
	"awesomeProject/user"
	"testing"
)

func TestInitStorage(t *testing.T) {
	storage := InitStorage()

	// Test that storage is initialized with default data
	if len(storage.Users) == 0 {
		t.Error("Storage should have default users")
	}

	if len(storage.Products) == 0 {
		t.Error("Storage should have default products")
	}

	if len(storage.Orders) == 0 {
		t.Error("Storage should have default orders")
	}
}

func TestAddUser(t *testing.T) {
	storage := &Storage{
		Users:    []user.User{},
		Products: []product.Product{},
		Orders:   make(map[int]*order.Order),
	}

	// Test adding first user
	storage.Users = append(storage.Users, user.New(1, "Test User", "test@example.com"))

	if len(storage.Users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(storage.Users))
	}

	if storage.Users[0].Name != "Test User" {
		t.Errorf("Expected user name 'Test User', got '%s'", storage.Users[0].Name)
	}

	if storage.Users[0].Email != "test@example.com" {
		t.Errorf("Expected user email 'test@example.com', got '%s'", storage.Users[0].Email)
	}
}

func TestAddNewOrder(t *testing.T) {
	storage := &Storage{
		Users: []user.User{
			user.New(1, "Test User", "test@example.com"),
		},
		Products: []product.Product{
			product.NewProduct(1, "Test Product", 100.0),
		},
		Orders: make(map[int]*order.Order),
	}

	// Test creating order
	selectedProducts := []product.Product{storage.Products[0]}
	newOrder := order.New(1, 1, selectedProducts)
	storage.Orders[1] = newOrder

	if len(storage.Orders) != 1 {
		t.Errorf("Expected 1 order, got %d", len(storage.Orders))
	}

	if storage.Orders[1].UserID != 1 {
		t.Errorf("Expected user ID 1, got %d", storage.Orders[1].UserID)
	}

	if len(storage.Orders[1].Products) != 1 {
		t.Errorf("Expected 1 product in order, got %d", len(storage.Orders[1].Products))
	}
}

func TestCalculateTotalRevenue(t *testing.T) {
	storage := &Storage{
		Orders: map[int]*order.Order{
			1: order.New(1, 1, []product.Product{
				product.NewProduct(1, "Product 1", 100.0),
				product.NewProduct(2, "Product 2", 200.0),
			}),
			2: order.New(2, 1, []product.Product{
				product.NewProduct(3, "Product 3", 150.0),
			}),
		},
	}

	expectedRevenue := 450.0 // 100 + 200 + 150
	actualRevenue := calculateTotalRevenue(storage.Orders)

	if actualRevenue != expectedRevenue {
		t.Errorf("Expected revenue %.2f, got %.2f", expectedRevenue, actualRevenue)
	}
}

func TestFindMostPopularProduct(t *testing.T) {
	storage := &Storage{
		Products: []product.Product{
			product.NewProduct(1, "Product A", 100.0),
			product.NewProduct(2, "Product B", 200.0),
		},
		Orders: map[int]*order.Order{
			1: order.New(1, 1, []product.Product{
				product.NewProduct(1, "Product A", 100.0),
				product.NewProduct(1, "Product A", 100.0), // Ordered twice
			}),
			2: order.New(2, 1, []product.Product{
				product.NewProduct(2, "Product B", 200.0),
			}),
		},
	}

	mostPopular := findMostPopularProduct(storage.Orders, storage.Products)

	if mostPopular == nil {
		t.Error("Expected to find most popular product")
		return
	}

	if mostPopular.Product.ID != 1 {
		t.Errorf("Expected product ID 1 (most popular), got %d", mostPopular.Product.ID)
	}

	if mostPopular.OrderCount != 2 {
		t.Errorf("Expected order count 2, got %d", mostPopular.OrderCount)
	}
}

func TestFindTopCustomers(t *testing.T) {
	storage := &Storage{
		Users: []user.User{
			user.New(1, "User A", "usera@example.com"),
			user.New(2, "User B", "userb@example.com"),
		},
		Orders: map[int]*order.Order{
			1: order.New(1, 1, []product.Product{
				product.NewProduct(1, "Product 1", 100.0),
			}),
			2: order.New(2, 2, []product.Product{
				product.NewProduct(2, "Product 2", 300.0), // Higher spending
			}),
		},
	}

	topCustomers := findTopCustomers(storage.Orders, storage.Users)

	if len(topCustomers) != 2 {
		t.Errorf("Expected 2 customers, got %d", len(topCustomers))
	}

	// User B should be first (higher spending)
	if topCustomers[0].User.ID != 2 {
		t.Errorf("Expected top customer ID 2, got %d", topCustomers[0].User.ID)
	}

	if topCustomers[0].TotalSpent != 300.0 {
		t.Errorf("Expected top customer spending 300.0, got %.2f", topCustomers[0].TotalSpent)
	}
}
