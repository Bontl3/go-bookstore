package routes

import (
	"net/http"

	"github.com/Bontl3/go-bookstore/pkg/controllers"
	"github.com/labstack/echo/v4"
)

// this function will have all your routes
var RegisterBookStoreRoutes = func(router *echo.Echo) {
	router.POST("/book/", echo.WrapHandler(http.HandlerFunc(controllers.CreateBook)))
	router.GET("/book/", echo.WrapHandler(http.HandlerFunc(controllers.GetBook)))
	router.GET("/book/:bookId", echo.WrapHandler(http.HandlerFunc(controllers.GetBookId)))
	router.PUT("/book/:bookId", echo.WrapHandler(http.HandlerFunc(controllers.UpdateBook)))
	router.DELETE("/book/:bookId", echo.WrapHandler(http.HandlerFunc(controllers.DeleteBookId)))
}
