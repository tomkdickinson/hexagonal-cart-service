package file

import (
	"encoding/json"
	"fmt"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/cart"
	"io/ioutil"
	"os"
)

type CartRepository struct {
	saveLocation string
}

func ProvideCartRepository() CartRepository {
	saveLocation := os.Getenv("STORAGE_FILE_LOCATION")
	if saveLocation == "" {
		saveLocation = "cart_storage.json"
	}
	c := CartRepository{saveLocation}
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		if err := c.writeDatabase(newDatabase()); err != nil {
			panic(err)
		}
	}
	return c
}


func (c CartRepository) Save(cart *cart.Cart) error {
	if db, err := c.readDatabase(); err != nil {
		return err
	} else {
		db.updateCart(cart)
		return c.writeDatabase(db)
	}
}

func (c CartRepository) Find(id string) (*cart.Cart, error) {
	if db, err := c.readDatabase(); err != nil {
		return nil, err
	} else {
		return db.findCart(id), nil
	}
}


func (c CartRepository) readDatabase() (*Database, error) {
	data, err := ioutil.ReadFile(c.saveLocation)
	if err != nil {
		return nil, fmt.Errorf("could not load cart storage file: %w", err)
	}
	var db Database
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, fmt.Errorf("could not unmarshal cart storage file: %w", err)
	} else {
		return &db, nil
	}
}

func (c CartRepository) writeDatabase(database *Database) error {
	data, err := json.Marshal(database)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(c.saveLocation, data, 0644)
	if err != nil {
		return err
	}
	return nil
}