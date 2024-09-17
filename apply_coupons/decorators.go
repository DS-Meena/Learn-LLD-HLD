package main

// Concrete Decorators
type PercentageCoupon struct {
	discount float64
	Product
}

func (p PercentageCoupon) GetPrice() float64 {
	return p.Product.GetPrice() * (1 - p.discount)
}

type GroceryCoupon struct {
	discount float64
	Product
}

func (p GroceryCoupon) GetPrice() float64 {
	// applies only on grocery products
	if p.Product.GetType() == "Grocery" {
		return p.Product.GetPrice() * (1 - p.discount)
	}
	// otherwise returns original price
	return p.Product.GetPrice()
}
