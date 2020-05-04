package mongo

import (
	"context"
	"fmt"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/cart"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository struct {
	collection mongo.Collection
}

func (c CartRepository) Save(cart *cart.Cart) error {
	doc := newCartDocument(cart)
	ctx := context.Background()
	_, err := c.collection.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("could not save document to mongo: %w", err)
	}
	return nil
}

func (c CartRepository) Find(id string) (*cart.Cart, error) {
	var doc cartDocument
	ctx := context.Background()
	if err := c.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc); err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error finding cart in mongo collection: %w", err)
	} else {
		return doc.adaptToCart(), nil
	}
}