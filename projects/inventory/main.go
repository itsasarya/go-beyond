package main

import "fmt"

type Product struct {
	Name  string
	Stock int
	Price float32
}

type CartItem struct {
	Product  *Product
	Quantity int
}

type Cart struct {
	Items []CartItem
}

type User struct {
	Name string
	Cart
}

func (p *Product) AddStock(amount int) {
	p.Stock += amount
}

func (p *Product) RemoveStock(amount int) {
	p.Stock -= amount
}

func (p *Product) AddOne() {
	p.Stock++
}

func (p *Product) RemoveOne() {
	p.Stock--
}

func (p *Product) IsAvailable(quantity int) bool {
	return p.Stock >= quantity
}

func (c *Cart) AddToCart(quantity int, p *Product) {
	if p.IsAvailable(quantity) {
		item := CartItem{
			Product:  p,
			Quantity: quantity,
		}

		c.Items = append(c.Items, item)
		p.RemoveStock(quantity)
		fmt.Println(quantity, p.Name, "added to cart")
	} else {
		fmt.Println("Not enough stock")
	}
}

// calculate total price
func (c *Cart) TotalPrice() float32 {
	var total float32

	for _, item := range c.Items {
		total += item.Product.Price * float32(item.Quantity)
	}
	return total
}

func main() {

	laptop := Product{
		Name:  "Laptop",
		Stock: 5,
		Price: 50000.0,
	}

	mouse := Product{
		Name:  "Mouse",
		Stock: 10,
		Price: 1000.0,
	}

	user := User{
		Name: "Aryan",
	}

	user.AddToCart(2, &laptop)
	user.AddToCart(3, &mouse)

	fmt.Println("Total Price:", user.TotalPrice())

	fmt.Println("Laptop stock left:", laptop.Stock)
	fmt.Println("Mouse stock left:", mouse.Stock)
}
