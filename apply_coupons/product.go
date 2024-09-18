package main

// Product interface
type Product interface {
	GetPrice() float64
	GetType() string
}

// Concrete Products
type GroceryProduct struct {
	Name  string
	price float64
}

func (p *GroceryProduct) GetPrice() float64 {
	return p.price
}

func (p *GroceryProduct) GetType() string {
	return "Grocery"
}

type ElectronicProduct struct {
	Name  string
	price float64
}

func (p *ElectronicProduct) GetPrice() float64 {
	return p.price
}

func (p *ElectronicProduct) GetType() string {
	return "Electronic"
}
