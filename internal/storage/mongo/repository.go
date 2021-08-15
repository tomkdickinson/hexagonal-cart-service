package mongo

import (
	"context"
	"fmt"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/domain/cart"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func ProvideRepository(collection *mongo.Collection) Repository {
	return Repository{collection}
}

func (c Repository) Save(cart *cart.Cart) error {
	doc := newCartDocument(cart)
	ctx := context.Background()
	opts := options.Update().SetUpsert(true)
	_, err := c.collection.UpdateOne(ctx, bson.M{"_id": cart.Id}, bson.M{"$set": doc}, opts)
	if err != nil {
		return fmt.Errorf("could not save document to mongo: %w", err)
	}
	return nil
}

func (c Repository) Find(id string) (*cart.Cart, error) {
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
