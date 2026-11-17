
package main

import (
	"fmt"
)

func main() {

	// Create inventory
	inventory := Inventory{}

	fmt.Println("===== INVENTORY MANAGEMENT SYSTEM =====")

	// Add items
	fmt.Println("\nAdding Items:")

	err := inventory.AddItem(Item{
		ID: 1, Name: "Laptop", Quantity: 10, Price: 50000,
	})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Laptop added successfully")
	}

	err = inventory.AddItem(Item{
		ID: 2, Name: "Mouse", Quantity: 20, Price: 500,
	})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mouse added successfully")
	}

	err = inventory.AddItem(Item{
		ID: 3, Name: "Keyboard", Quantity: 15, Price: 1000,
	})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Keyboard added successfully")
	}

	// Display inventory
	DisplayInventory(inventory.Items)

	// Update item
	fmt.Println("\nUpdating Mouse:")

	err = inventory.UpdateItem(2, Item{
		ID: 2, Name: "Mouse", Quantity: 10, Price: 550,
	})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mouse updated successfully")
	}

	// Remove item
	fmt.Println("\nRemoving Keyboard:")

	err = inventory.RemoveItem(3)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Keyboard removed successfully")
	}

	// Display updated inventory
	DisplayInventory(inventory.Items)

	// Calculate summary concurrently
	fmt.Println("\nCalculating inventory summary concurrently...")

	resultChannel := ComputeSummaryAsync(inventory.Items)

	// Receive result from channel
	summary := <-resultChannel

	fmt.Println("\n----- Inventory Summary -----")
	fmt.Printf("Total Items: %d\n", summary.ItemCount)
	fmt.Printf("Total Inventory Value: %.2f\n", summary.TotalValue)

	fmt.Println("\nProgram completed successfully!")
}
