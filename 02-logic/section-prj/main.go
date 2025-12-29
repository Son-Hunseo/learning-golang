package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]
	if !found { // 해당하는 상품이 없다면
		if strings.HasSuffix(itemCode, "_SALE") { // 해당하는 상품 이름에 _SALE이 붙어있는지 확인
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90 // 10퍼센트 세일
				fmt.Printf(" - Item %s (Sale! Original: $%.2f, Sale Price: $%.2f)\n",
					originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}

		fmt.Printf(" - Item %s (Product not found)\n", itemCode)
		return 0.0, false
	}

	return basePrice, true
}

func main() {
	fmt.Println("---------- Simple Sales Order Processor ----------")

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	var subtotal float64
	fmt.Println("------ Processing Order Items:")

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Subtotal Price: %.2f\n", subtotal)
}
