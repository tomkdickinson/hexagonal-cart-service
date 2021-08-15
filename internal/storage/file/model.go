package file

import (
	"github.com/tomkdickinson/hexagonal-cart-service/internal/domain/cart"
)

type Database struct {
	Carts map[string]Cart
}

func newDatabase() *Database {
	return &Database{Carts: make(map[string]Cart)}
}

type Cart struct {
	Id    string         `json:"id"`
	Items map[string]int `json:"items"`
}

func newCart(c *cart.Cart) Cart {
	return Cart{
		c.Id,
		c.Items(),
	}
}

func (d Database) updateCart(c *cart.Cart) {
	d.Carts[c.Id] = newCart(c)
}

func (d Database) findCart(id string) *cart.Cart {
	if c, ok := d.Carts[id]; !ok {
		return nil
	} else {
		model := cart.NewCart(c.Id)
		for item, quantity := range c.Items {
			model.AddItem(item, quantity)
		}
		return model
	}
}
