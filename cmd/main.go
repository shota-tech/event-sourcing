package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/shota-tech/event-sourcing/internal"
)

func main() {
	repository := internal.NewWarehouseProductRepository()
	var key string

	for {
		fmt.Println("R: Receive Inventory")
		fmt.Println("S: Ship Inventory")
		fmt.Println("A: Inventory Adjustment")
		fmt.Println("Q: Quantity On Hand")
		fmt.Println("E: Events")
		fmt.Println("Q: Quit")
		fmt.Print("> ")
		fmt.Scan(&key)
		key = strings.ToUpper(key)
		if key == "Q" {
			fmt.Println("Bye!")
			break
		}
		fmt.Println()

		sku := getSKU()
		product := repository.Get(sku)

		switch key {
		case "R":
			quantity := getQuantity()
			product.ReceiveProduct(quantity)
			fmt.Printf("%s Received: %d\n", sku, quantity)
		case "S":
			quantity := getQuantity()
			if err := product.ShipProduct(quantity); err != nil {
				fmt.Printf("error: %s", err.Error())
				break
			}
			fmt.Printf("%s Shipped: %d\n", sku, quantity)
		case "A":
			quantity := getQuantity()
			reason := getAdjustmentReason()
			if err := product.AdjustInventory(quantity, reason); err != nil {
				fmt.Printf("error: %s", err.Error())
				break
			}
			fmt.Printf("%s Adjusted: %d %s\n", sku, quantity, reason)
		case "Q":
			quantity := product.QuantityOnHand()
			fmt.Printf("%s Quantity On Hand: %d\n", sku, quantity)
		case "E":
			fmt.Printf("Events: %s\n", sku)
			for _, event := range product.Events() {
				switch e := event.(type) {
				case internal.ProductShipped:
					fmt.Printf("%s %s Shipped: %d\n", e.DateTime.Format(time.RFC3339), sku, e.Quantity)
				case internal.ProductReceived:
					fmt.Printf("%s %s Received: %d\n", e.DateTime.Format(time.RFC3339), sku, e.Quantity)
				case internal.InventoryAdjusted:
					fmt.Printf("%s %s Adjusted: %d %s\n", e.DateTime.Format(time.RFC3339), sku, e.Quantity, e.Reason)
				}
			}
		}

		repository.Save(product)
		fmt.Println()
	}
}

func getSKU() string {
	fmt.Print("SKU: ")
	var sku string
	fmt.Scan(&sku)
	return sku
}

func getQuantity() int {
	fmt.Print("Quantity: ")
	var quantity int
	fmt.Scan(&quantity)
	return quantity
}

func getAdjustmentReason() string {
	fmt.Print("Reason: ")
	var reason string
	fmt.Scan(&reason)
	return reason
}
