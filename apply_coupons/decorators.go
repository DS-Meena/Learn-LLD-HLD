package main

// Concrete Decorators
type PercentageCoupon struct {
	Discount float64
	Product  Product
}

func (p PercentageCoupon) GetPrice() float64 {
	return p.Product.GetPrice() * (1 - p.Discount)
}

func (p PercentageCoupon) GetType() string {
	return p.Product.GetType()
}

type GroceryCoupon struct {
	Discount float64
	Product  Product
}

func (p GroceryCoupon) GetPrice() float64 {
	// applies only on grocery products
	if p.Product.GetType() == "Grocery" {
		return p.Product.GetPrice() * (1 - p.Discount)
	}
	// otherwise returns original price
	return p.Product.GetPrice()
}

func (p GroceryCoupon) GetType() string {
	return p.Product.GetType()
}
