package main

type ShoppingCart interface {
	// list of products
	Products() []Product
	// add product to cart
	AddProduct(product Product)
	GetTotalPrice() float64
}

type Cart struct {
	products []Product
}

func (c *Cart) Products() []Product {
	return c.products
}

func (c *Cart) GetTotalPrice() float64 {
	totalPrice := 0.0
	for _, product := range c.products {
		totalPrice += product.GetPrice()
	}
	return totalPrice
}

func (c *Cart) AddProduct(product Product) {
	// before adding product apply coupons on it as decorator
	productWithPercentageDiscount := &PercentageCoupon{
		Product:  product,
		Discount: .1,
	}

	// apply grocery coupon after applying percentage coupon
	productWithGroceryCoupon := &GroceryCoupon{
		Product:  productWithPercentageDiscount,
		Discount: .05,
	}

	c.products = append(c.products, productWithGroceryCoupon)
}
