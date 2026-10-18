package main

import (
	"fmt"
	"strings"
)

var productsPrice = map[string]float64{
	"Apple":     20.00,
	"Banana":    40.00,
	"Cherry":    80.00,
	"Orange":    80.00,
	"Pineapple": 120.00,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	// Sale items: e.g. "Banana_SALE"
	if strings.HasSuffix(itemCode, "_SALE") {
		originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
		basePrice, found := productsPrice[originalItemCode]
		if !found {
			fmt.Printf("-item: %s (Product not found)\n", itemCode)
			return 0.0, false
		}
		salePrice := basePrice * 0.78
		fmt.Printf("-item %s (Sale! original: $%.2f, Sale price: $%.2f)\n", originalItemCode, basePrice, salePrice)
		return salePrice, true
	}

	// Regular items
	basePrice, found := productsPrice[itemCode]
	if !found {
		fmt.Printf("-item: %s (Product not found)\n", itemCode)
		return 0.0, false
	}
	fmt.Printf("-item %s: $%.2f\n", itemCode, basePrice)
	return basePrice, true
}

func main() {
	fmt.Println("------------------- Simple Order Processor ----------------------")

	orderItems := []string{
		"Apple", "Banana_SALE", "Cherry", "Orange",
	}

	var subtotal float64
	fmt.Println("------------------- Processing Order Items ----------------------")

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Subtotal Price: %.2f\n", subtotal)
}
