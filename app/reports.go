package app

import (
	"awesomeProject/order"
	"awesomeProject/product"
	"awesomeProject/user"
	"fmt"
	"sort"
	"strings"
)

// ShowReports displays various summary reports
func ShowReports(storage *Storage) {
	fmt.Println("📊 Отчеты:")
	fmt.Println(strings.Repeat("=", 40))

	// User count report
	userCount := len(storage.Users)
	fmt.Printf("👥 Общее количество пользователей: %d\n", userCount)

	// Revenue report
	totalRevenue := calculateTotalRevenue(storage.Orders)
	fmt.Printf("💰 Общий доход от всех заказов: %.2f₽\n", totalRevenue)

	// Order count report
	orderCount := len(storage.Orders)
	fmt.Printf("📦 Общее количество заказов: %d\n", orderCount)

	// Most popular product report
	mostPopular := findMostPopularProduct(storage.Orders, storage.Products)
	if mostPopular != nil {
		fmt.Printf("🏆 Самый популярный товар: %s (заказан %d раз)\n",
			mostPopular.Product.Name, mostPopular.OrderCount)
	}

	// Average order value
	if orderCount > 0 {
		avgOrderValue := totalRevenue / float64(orderCount)
		fmt.Printf("📊 Средняя стоимость заказа: %.2f₽\n", avgOrderValue)
	}

	// Top customers by order value
	topCustomers := findTopCustomers(storage.Orders, storage.Users)
	if len(topCustomers) > 0 {
		fmt.Println("\n👑 Топ-3 клиента по сумме заказов:")
		for i, customer := range topCustomers {
			if i >= 3 {
				break
			}
			fmt.Printf("  %d. %s - %.2f₽\n", i+1, customer.User.Name, customer.TotalSpent)
		}
	}

	fmt.Println(strings.Repeat("=", 40))
}

// calculateTotalRevenue calculates total revenue from all orders
func calculateTotalRevenue(orders map[int]*order.Order) float64 {
	total := 0.0
	for _, order := range orders {
		total += order.Total()
	}
	return total
}

// ProductOrderCount represents a product with its order count
type ProductOrderCount struct {
	Product    product.Product
	OrderCount int
}

// findMostPopularProduct finds the product ordered most frequently
func findMostPopularProduct(orders map[int]*order.Order, products []product.Product) *ProductOrderCount {
	// Count orders for each product
	productCounts := make(map[int]int)
	for _, order := range orders {
		for _, product := range order.Products {
			productCounts[product.ID]++
		}
	}

	// Find product with highest count
	var mostPopular *ProductOrderCount
	maxCount := 0

	for productID, count := range productCounts {
		if count > maxCount {
			maxCount = count
			// Find product details
			for _, p := range products {
				if p.ID == productID {
					mostPopular = &ProductOrderCount{
						Product:    p,
						OrderCount: count,
					}
					break
				}
			}
		}
	}

	return mostPopular
}

// CustomerSpending represents a customer with their total spending
type CustomerSpending struct {
	User       user.User
	TotalSpent float64
}

// findTopCustomers finds customers with highest total spending
func findTopCustomers(orders map[int]*order.Order, users []user.User) []CustomerSpending {
	// Calculate total spending for each customer
	customerSpending := make(map[int]float64)
	for _, order := range orders {
		customerSpending[order.UserID] += order.Total()
	}

	// Convert to slice for sorting
	var customers []CustomerSpending
	for userID, totalSpent := range customerSpending {
		// Find user details
		for _, u := range users {
			if u.ID == userID {
				customers = append(customers, CustomerSpending{
					User:       u,
					TotalSpent: totalSpent,
				})
				break
			}
		}
	}

	// Sort by total spending (descending)
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].TotalSpent > customers[j].TotalSpent
	})

	return customers
}
