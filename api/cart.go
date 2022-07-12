package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetCart(c echo.Context) error {

	//Get user ID from query param
	userId := c.Param("userId")

	//Obtain current database connection and fetch cart by user ID
	db := database.DbManager()
	cart := model.Cart{}
	//Preload all cart items and include in response
    db.Where("user_id = ?", userId).Preload("CartItems").Find(&cart)

	return c.JSON(http.StatusOK, cart)
}

func UpdateCart(c echo.Context) error {

	//Bind json from request to object
	updatedCart := new(model.Cart)
	c.Bind(updatedCart)

	//Obtain current database connection
	db := database.DbManager()

	db.Save(&updatedCart)
	return c.JSON(http.StatusOK, updatedCart)
}