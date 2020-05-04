package cart

type Cart struct {
	Id    string
	items map[string]int
}

func NewCart(id string) *Cart {
	return &Cart{
		Id:    id,
		items: make(map[string]int),
	}
}

func (c *Cart) AddItem(sku string, quantityToAdd int) {
	if quantityInCart, ok := c.items[sku]; !ok {
		c.items[sku] = quantityToAdd
	} else {
		c.items[sku] = quantityInCart + quantityToAdd
	}
}

func (c Cart) Items() map[string]int {
	return c.items
}
