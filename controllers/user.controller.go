package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type UserController struct {
}

func (UserController) List(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (UserController) Find(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
