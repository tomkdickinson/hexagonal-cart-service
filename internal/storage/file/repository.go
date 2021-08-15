package file

import (
	"encoding/json"
	"fmt"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/domain/cart"
	"io/ioutil"
	"log"
	"os"
)

const saveLocation = "cart_storage.json"

type Repository struct {
}

func ProvideRepository() Repository {
	c := Repository{}
	if _, err := os.Stat(saveLocation); os.IsNotExist(err) {
		log.Println("Creating file database")
		if err := c.writeDatabase(newDatabase()); err != nil {
			panic(err)
		}
	}
	return c
}

func (c Repository) Save(cart *cart.Cart) error {
	if db, err := c.readDatabase(); err != nil {
		return err
	} else {
		db.updateCart(cart)
		return c.writeDatabase(db)
	}
}

func (c Repository) Find(id string) (*cart.Cart, error) {
	if db, err := c.readDatabase(); err != nil {
		return nil, err
	} else {
		return db.findCart(id), nil
	}
}

func (c Repository) readDatabase() (*Database, error) {
	data, err := ioutil.ReadFile(saveLocation)
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

func (c Repository) writeDatabase(database *Database) error {
	data, err := json.Marshal(database)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(saveLocation, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
