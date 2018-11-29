package routes

import (
	"github.com/labstack/echo"
	"github.com/pskclub/play-with-go/controllers"
)

func UserRoutes(route *echo.Echo) {
	user := &controllers.UserController{}
	route.GET("/users", user.List)
	route.GET("/users/1", user.Find)
}
