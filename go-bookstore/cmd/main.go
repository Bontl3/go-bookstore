package main

// we are going to create a servere which ill define our local host
// we tell golang where our routers are defined

import (
	"log" // to be able to log out if there is any error
	"net/http"

	"github.com/Bontl3/go-bookstore/pkg/routes" // importing the routes
	"github.com/labstack/echo/v4"
)

func main() {

	// create a new echo instance
	r := echo.New()

	// parsing the new echo route instance
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	// Listen and serve helps us to create a server and we parse to it the  port where we want start the server
	// if there is an error it will log it out
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
