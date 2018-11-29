package routes

import (
	"github.com/labstack/echo"
	"github.com/pskclub/play-with-go/controllers"
)

func BookRoutes(route *echo.Echo) {
	book := &controllers.BookController{}
	route.GET("/books", book.List)
	route.GET("/books/:id", book.Find)
}
