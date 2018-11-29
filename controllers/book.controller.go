package controllers

import (
	"github.com/labstack/echo"
	"github.com/pskclub/play-with-go/services"
	"net/http"
)

type BookController struct {
}

func (BookController) List(c echo.Context) error {
	bookService := services.BookService{}
	books := bookService.FindAll()
	return c.JSON(http.StatusOK, books)
}

func (BookController) Find(c echo.Context) error {
	id := c.Param("id")
	bookService := services.BookService{}
	book,err := bookService.Find(id)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, book)
}
