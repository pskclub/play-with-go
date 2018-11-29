package routes

import (
	"github.com/labstack/echo"
)

func init() {
	e := echo.New()

	UserRoutes(e)
	BookRoutes(e)

	e.Logger.Fatal(e.Start(":9000"))
}
