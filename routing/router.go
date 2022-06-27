package routing

import (
	"github.com/RobertOchmanek/ebiznes_go/api"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	//Create server instance
	e := echo.New()

	//Initialize endpoints and handler methods
	e.GET("/", api.GetOrders)
	e.PUT("/:id", api.UpdateOrder)

	e.GET("/users", api.FindUsers)
	e.GET("/users/:id", api.FindUser)
	e.POST("/users", api.AddUser)
	e.PUT("/users/:id", api.UpdateUser)
	e.DELETE("/users/:id", api.RemoveUser)

	e.GET("/products", api.FindProducts)
	e.GET("/products/:id", api.FindProduct)
	e.POST("/products", api.AddProduct)
	e.PUT("/products/:id", api.UpdateProduct)
	e.DELETE("/products/:id", api.RemoveProduct)

	return e
}