package routing

import (
	"github.com/RobertOchmanek/ebiznes_go/api"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	//Create server instance
	e := echo.New()

	//Initialize endpoints and handler methods
	e.GET("/", api.HomePage)
	e.GET("/users", api.FindUsers)
	e.GET("/users/:id", api.FindUser)
	e.POST("/users", api.AddUser)
	e.DELETE("/users/:id", api.RemoveUser)

	return e
}