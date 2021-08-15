package cart

type Repository interface {
	Save(c *Cart) error
	Find(cartId string) (*Cart, error)
}
