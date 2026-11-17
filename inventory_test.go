
package main

import "testing"

// Test 1: Add item successfully
func TestAddItem(t *testing.T) {
	inv := Inventory{}

	item := Item{
		ID: 1, Name: "Laptop", Quantity: 10, Price: 50000,
	}

	err := inv.AddItem(item)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(inv.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(inv.Items))
	}
}

// Test 2: Prevent duplicate items
func TestDuplicateItem(t *testing.T) {
	inv := Inventory{}

	item := Item{
		ID: 1, Name: "Laptop", Quantity: 10, Price: 50000,
	}

	inv.AddItem(item)

	err := inv.AddItem(item)

	if err == nil {
		t.Error("Expected duplicate item error")
	}
}

// Test 3: Update existing item
func TestUpdateItem(t *testing.T) {
	inv := Inventory{}

	inv.AddItem(Item{
		ID: 1, Name: "Laptop", Quantity: 10, Price: 50000,
	})

	updated := Item{
		ID: 1, Name: "Laptop", Quantity: 20, Price: 45000,
	}

	err := inv.UpdateItem(1, updated)

	if err != nil {
		t.Errorf("Update failed: %v", err)
	}

	if inv.Items[0].Quantity != 20 {
		t.Errorf("Expected quantity 20, got %d",
			inv.Items[0].Quantity)
	}
}

// Test 4: Remove an item
func TestRemoveItem(t *testing.T) {
	inv := Inventory{}

	inv.AddItem(Item{
		ID: 1, Name: "Laptop", Quantity: 10, Price: 50000,
	})

	err := inv.RemoveItem(1)

	if err != nil {
		t.Errorf("Remove failed: %v", err)
	}

	if len(inv.Items) != 0 {
		t.Errorf("Expected empty inventory")
	}
}

// Test 5: Reject invalid item
func TestInvalidItem(t *testing.T) {
	inv := Inventory{}

	item := Item{
		ID: 0, Name: "", Quantity: -5, Price: -100,
	}

	err := inv.AddItem(item)

	if err == nil {
		t.Error("Expected validation error")
	}
}

// Test 6: Calculate inventory summary
func TestComputeSummary(t *testing.T) {
	items := []Item{
		{ID: 1, Name: "Laptop", Quantity: 2, Price: 1000},
		{ID: 2, Name: "Mouse", Quantity: 3, Price: 500},
	}

	result := <-ComputeSummaryAsync(items)

	if result.ItemCount != 2 {
		t.Errorf("Expected 2 items, got %d", result.ItemCount)
	}

	if result.TotalValue != 3500 {
		t.Errorf("Expected total value 3500, got %.2f",
			result.TotalValue)
	}
}
