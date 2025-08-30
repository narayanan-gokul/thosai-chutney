// The following program has been converted into Go, has not been validated yet
package main

import (
	"fmt"
)

// FoodItem represents a food item with a quantity
type FoodItem struct {
	Quantity int
}

// checkSupply verifies if supplier can meet hub's total demand
func checkSupply(supplierInventory, hubTotals map[string]FoodItem) string {
	for foodID, hubItem := range hubTotals {
		supplierItem, exists := supplierInventory[foodID]
		if !exists || supplierItem.Quantity < hubItem.Quantity {
			return "Unstockable"
		}
		supplierItem.Quantity -= hubItem.Quantity
		supplierInventory[foodID] = supplierItem
	}
	return "Stockable"
}

// fulfillable checks if hub can fulfill consumer's request
func fulfillable(hubTotals, consumer map[string]FoodItem) string {
	for foodID, consumerItem := range consumer {
		hubItem, exists := hubTotals[foodID]
		if !exists || hubItem.Quantity < consumerItem.Quantity {
			return "Unstockable"
		}
		hubItem.Quantity -= consumerItem.Quantity
		hubTotals[foodID] = hubItem
	}
	return "Stockable"
}

// approval returns true if consumer's request is fulfillable
func approval(hubTotals, consumer map[string]FoodItem) bool {
	return fulfillable(hubTotals, consumer) == "Stockable"
}

func main() {
	// Example usage
	supplier := map[string]FoodItem{
		"apple":  {Quantity: 100},
		"banana": {Quantity: 50},
	}
	hub := map[string]FoodItem{
		"apple":  {Quantity: 60},
		"banana": {Quantity: 30},
	}
	consumer := map[string]FoodItem{
		"apple":  {Quantity: 20},
		"banana": {Quantity: 10},
	}

	fmt.Println("Supply Check:", checkSupply(supplier, hub))       // Should print "Stockable"
	fmt.Println("Fulfillable:", fulfillable(hub, consumer))        // Should print "Stockable"
	fmt.Println("Approval:", approval(hub, consumer))              // Should print true
}
