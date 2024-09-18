package main

import "fmt"

func main() {

	// create products
	item1 := &GroceryProduct{"Milk", 199}
	item2 := &ElectronicProduct{"Laptop", 79999}
	item3 := &GroceryProduct{"Bread", 99}
	item4 := &ElectronicProduct{"Smartphone", 49999}

	shoppingCart := Cart{}

	// apply coupon and add to cart
	shoppingCart.AddProduct(item1)
	shoppingCart.AddProduct(item2)
	shoppingCart.AddProduct(item3)
	shoppingCart.AddProduct(item4)

	// // coupons
	// productWithPercentageDiscount := &PercentageCoupon{Discount: 10.0}
	// productWithGroceryCoupon := &GroceryCoupon{Discount: 5.0}
	// shoppingCart.AddCoupon(productWithPercentageDiscount)
	// shoppingCart.AddCoupon(productWithGroceryCoupon)

	// print the price of shopping cart
	fmt.Printf("Total price after applying coupons is %+v", shoppingCart.GetTotalPrice())
}
