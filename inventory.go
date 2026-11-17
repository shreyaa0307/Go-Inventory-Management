
package main

import (
	"errors"
	"fmt"
	"strings"
)

// Item structure
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Inventory stores multiple items using a slice
type Inventory struct {
	Items []Item
}

// Summary structure
type Summary struct {
	ItemCount int
	TotalValue float64
}

// Validate item details
func validateItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("item ID must be positive")
	}

	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	if item.Price < 0 {
		return errors.New("price cannot be negative")
	}

	return nil
}

// Add a new item
func (inv *Inventory) AddItem(item Item) error {
	if err := validateItem(item); err != nil {
		return err
	}

	for _, existing := range inv.Items {
		if existing.ID == item.ID {
			return errors.New("item already exists")
		}
	}

	inv.Items = append(inv.Items, item)
	return nil
}

// Update an existing item
func (inv *Inventory) UpdateItem(id int, updated Item) error {
	if id <= 0 {
		return errors.New("invalid item ID")
	}

	if updated.ID != id {
		return errors.New("updated item ID must match")
	}

	if err := validateItem(updated); err != nil {
		return err
	}

	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items[i] = updated
			return nil
		}
	}

	return errors.New("item not found")
}

// Remove an item
func (inv *Inventory) RemoveItem(id int) error {
	if id <= 0 {
		return errors.New("invalid item ID")
	}

	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items = append(
				inv.Items[:i],
				inv.Items[i+1:]...,
			)
			return nil
		}
	}

	return errors.New("item not found")
}

// Calculate summary concurrently using goroutine and channel
func ComputeSummaryAsync(items []Item) <-chan Summary {
	result := make(chan Summary, 1)

	// Start a goroutine
	go func() {
		defer close(result)

		var summary Summary

		for _, item := range items {
			summary.ItemCount++
			summary.TotalValue += float64(item.Quantity) * item.Price
		}

		// Send result through channel
		result <- summary
	}()

	return result
}

// Display inventory items
func DisplayInventory(items []Item) {
	fmt.Println("\n----- Inventory Items -----")

	for _, item := range items {
		fmt.Printf(
			"ID: %d | Name: %s | Quantity: %d | Price: %.2f\n",
			item.ID,
			item.Name,
			item.Quantity,
			item.Price,
		)
	}
}
