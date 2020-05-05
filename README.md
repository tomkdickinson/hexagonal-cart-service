# Hexagonal Cart Service
This is an example repository of how you can write a Go project with Wire and the Hexagonal Pattern.

The service is a very basic ecommerce cart, with the options to view your cart, and add items to it. 


## Starting the service
The easiest way to start the service is to use Docker and Docker Compose.

Just run `docker compose up --build` to build and start the containers.

## Using the service

To add an item to your cart you can make a GET request of format:

`http://localhost:8000/cart/<cart-id>/<item-id>/<quantity>`

To view the items in a cart you can make a GET request of format:

`http://localhost:8000/cart/<cart-id>`

So to add two pairs of socks to a cart with id `my-cart` you would use:

[http://localhost:8000/cart/my-cart/socks/2](http://localhost:8000/cart/my-cart/socks/2)

If you call that a second time, you'll add another 2 pairs of socks to the cart (for a total of 4).

To see those socks persisted in your cart, you can use:

[http://localhost:8000/cart/my-cart](http://localhost:8000/cart/my-cart)

### Switching from Mongo to File storage
To switch from Mongo to persisting on disk, just set the following environment variable:

`CART_STORAGE=file`

This will persist the cart to a file called `cart_storage.json` in the docker container, or in the directory the app was ran if running locally. 

## Building the code

This project uses go modules and go 1.14. 

### Dependency Generation
It also uses [Google Wire](https://github.com/google/wire) as a dependency injection framework. 
Please follow the instructions on how to install it from the repository.

To generate the dependency code, you can run:

`go generate ./...`

Or 

`wire github.com/tomkdickinson/hexagonal-cart-service/cmd/app`

### Building
You can build a native binary with:

`go build github.com/tomkdickinson/hexagonal-cart-service/cmd/app`

Alternatively you can build the service by itself with docker:

`docker build --tag cart-service .`

Running `docker-compose up --build` will also build the container and deploy it locally.