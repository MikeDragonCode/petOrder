package app

import (
	"awesomeProject/order"
	"awesomeProject/product"
	"awesomeProject/user"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const dataFileName = "data.json"

// SaveToFile saves the storage data to a JSON file
func (s *Storage) SaveToFile() error {
	// Create data directory if it doesn't exist
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	filePath := filepath.Join(dataDir, dataFileName)

	// Convert orders map to slice for JSON serialization
	ordersSlice := make([]*order.Order, 0, len(s.Orders))
	for _, order := range s.Orders {
		ordersSlice = append(ordersSlice, order)
	}

	// Create a temporary struct for serialization
	dataToSave := struct {
		Users    []user.User       `json:"users"`
		Products []product.Product `json:"products"`
		Orders   []*order.Order    `json:"orders"`
	}{
		Users:    s.Users,
		Products: s.Products,
		Orders:   ordersSlice,
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(dataToSave, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("💾 Данные сохранены в файл %s\n", filePath)
	return nil
}

// LoadFromFile loads storage data from a JSON file
func (s *Storage) LoadFromFile() error {
	filePath := filepath.Join("data", dataFileName)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("📁 Файл %s не найден. Используются данные по умолчанию.\n", filePath)
		return nil
	}

	// Read file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Create temporary struct for deserialization
	var dataFromFile struct {
		Users    []user.User       `json:"users"`
		Products []product.Product `json:"products"`
		Orders   []*order.Order    `json:"orders"`
	}

	// Unmarshal from JSON
	if err := json.Unmarshal(jsonData, &dataFromFile); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Convert orders slice back to map
	ordersMap := make(map[int]*order.Order)
	for _, order := range dataFromFile.Orders {
		ordersMap[order.ID] = order
	}

	// Update storage
	s.Users = dataFromFile.Users
	s.Products = dataFromFile.Products
	s.Orders = ordersMap

	fmt.Printf("📂 Данные загружены из файла %s\n", filePath)
	return nil
}

// AutoSave saves data automatically (can be called after important operations)
func (s *Storage) AutoSave() {
	if err := s.SaveToFile(); err != nil {
		fmt.Printf("⚠️ Ошибка автоматического сохранения: %v\n", err)
	}
}
