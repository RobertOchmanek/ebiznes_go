package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetOrders(c echo.Context) error {

	//Obtain current database connection and fetch products
	db := database.DbManager()
	orders := []model.Order{}
	db.Preload("Products").Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c echo.Context) error {

	//Get product ID from query param
	id := c.Param("id")

	//Obtain current database connection
	db := database.DbManager()

	//Fetch product by ID
	var product model.Product
    db.Where("product_id = ?", id).Find(&product)

	//Fetch order
	var order model.Order
	db.Where("order_id = ?", 1).Find(&order)

	//Update and save DB object
	order.Products = append(order.Products, product)
	db.Save(&order)

	return c.JSON(http.StatusOK, order)
}