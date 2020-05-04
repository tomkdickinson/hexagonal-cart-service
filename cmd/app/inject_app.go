//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/api"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/cart"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/storage/file"
)

//func handlersWithMongoStorage() (*api.Handlers, error) {
//	panic(
//		wire.Build(
//			api.ProvideHandlers,
//		))
//}

func handlersWithFileStorage() (*api.Handlers, error) {
	panic(
		wire.Build(
			file.ProvideCartRepository,
			wire.Bind(new(cart.Repository), new(file.CartRepository)),
			cart.ProvideService,
			api.ProvideHandlers,
		))
}
