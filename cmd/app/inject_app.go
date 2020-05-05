//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/api"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/cart"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/storage/file"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/storage/mongo"
)

func handlersWithMongoStorage() (*api.Handlers, error) {
	panic(
		wire.Build(
			mongo.ProvideCollection,
			mongo.ProvideRepository,
			wire.Bind(new(cart.Repository), new(mongo.Repository)),
			cart.ProvideService,
			api.ProvideHandlers,
		))
}

func handlersWithFileStorage() (*api.Handlers, error) {
	panic(
		wire.Build(
			file.ProvideRepository,
			wire.Bind(new(cart.Repository), new(file.Repository)),
			cart.ProvideService,
			api.ProvideHandlers,
		))
}
