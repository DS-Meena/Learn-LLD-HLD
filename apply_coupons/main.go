package main

import "fmt"

func main() {

	// create products
	item1 := &GroceryProduct{"Milk", 199}
	item2 := &ElectronicProduct{"Laptop", 79999}
	item3 := &GroceryProduct{"Bread", 99}
	item4 := &ElectronicProduct{"Smartphone", 49999}
	products := []Product{item1, item2, item3, item4}

	shoppingCart := Cart{}

	// iterate the products and apply coupons
	for _, product := range products {
		productWithPercentageDiscount := &PercentageCoupon{Product: product, Discount: .1}
		productWithGroceryCoupon := &GroceryCoupon{Product: productWithPercentageDiscount, Discount: .05}

		// can add more coupons here with different conditions
		// based on product type and index (pass it also)

		shoppingCart.AddProduct(productWithGroceryCoupon)
	}

	// print the price of shopping cart
	fmt.Printf("Total price after applying coupons is ₹%+v \n", shoppingCart.GetTotalPrice())
}
