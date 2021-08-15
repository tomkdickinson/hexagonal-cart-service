package mongo

import (
	"github.com/tomkdickinson/hexagonal-cart-service/internal/domain/cart"
)

type cartDocument struct {
	Id    string         `bson:"_id"`
	Items map[string]int `bson:"items"`
}

func newCartDocument(cart *cart.Cart) *cartDocument {
	return &cartDocument{
		Id:    cart.Id,
		Items: cart.Items(),
	}
}

func (c cartDocument) adaptToCart() *cart.Cart {
	model := cart.NewCart(c.Id)
	for sku, quantity := range c.Items {
		model.AddItem(sku, quantity)
	}
	return model
}
