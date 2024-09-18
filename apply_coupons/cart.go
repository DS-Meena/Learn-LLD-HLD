package main

type ShoppingCart interface {
	AddProduct(product Product)
	GetTotalPrice() float64
}

type Cart struct {
	Products []Product
}

func (c *Cart) GetTotalPrice() float64 {
	totalPrice := 0.0
	for _, product := range c.Products {
		totalPrice += product.GetPrice()
	}
	return totalPrice
}

func (c *Cart) AddProduct(product Product) {
	c.Products = append(c.Products, product)
}
