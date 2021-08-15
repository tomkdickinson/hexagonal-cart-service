package cart

import (
	"fmt"
	"log"
)

type Service struct {
	repository Repository
}

func ProvideService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) AddItemToCart(cartId string, sku string, quantity int) (*Cart, error) {
	cart, err := s.repository.Find(cartId)

	if err != nil {
		return nil, fmt.Errorf("error finding Cart in repository: %w", err)
	}

	if cart == nil {
		log.Println("No cart in repository, creating new one")
		cart = NewCart(cartId)
	}

	cart.AddItem(sku, quantity)

	if err = s.repository.Save(cart); err != nil {
		return nil, fmt.Errorf("could not send Cart updated message: %w", err)
	}
	return cart, nil
}

func (s *Service) GetCart(cartId string) (*Cart, error) {
	if cart, err := s.repository.Find(cartId); err != nil {
		return nil, fmt.Errorf("error finding cart in repository: %w", err)
	} else if cart == nil {
		return NewCart(cartId), nil
	} else {
		return cart, nil
	}
}
