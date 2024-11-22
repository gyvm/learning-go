package router

import (
	"github.com/labstack/echo/v4"
	"rest_api/controller"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/loout", uc.LogOut)
	return e
}
