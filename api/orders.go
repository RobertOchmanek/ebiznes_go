package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/RobertOchmanek/ebiznes_go/model/rest"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetOrders(c echo.Context) error {

	//Obtain current database connection and fetch orders
	db := database.DbManager()
	orders := []model.Order{}
	
	db.Preload("Payment").Preload("Products").Find(&orders)

	return c.JSON(http.StatusOK, orders)
}

func GetOrder(c echo.Context) error {

	//Get order ID from query param
	id := c.Param("id")

	//Obtain current database connection and fetch order by ID
	db := database.DbManager()
	order := model.Order{}
	//Preload all category's products and payment and include in response
    db.Where("id = ?", id).Preload("Payment").Preload("Products").Find(&order)

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {

	//Bind json from request to object
	restOrder := new(rest.RestOrder)
	c.Bind(restOrder)

	//Obtain current database connection
	db := database.DbManager()

	//Fetch products from request by IDs
	products := []model.Product{}
	for _, id := range restOrder.ProductsIds {
		product := model.Product{}
		db.Where("id = ?", id).Find(&product)
		products = append(products, product)
	}

	//Convert payment from REST DTO to model object
	payment := model.Payment{}
	payment.Accepted = restOrder.Payment.Accepted
	payment.PaymentType = restOrder.Payment.PaymentType

	//Save new order, order ID is added by GORM
	newOrder := model.Order{}
	newOrder.Products = products
	newOrder.Payment = payment
	db.Create(&newOrder)

	user := model.User{}
	db.Where("id = ?", restOrder.UserId).Find(&user)
	user.Orders = append(user.Orders, newOrder)
	db.Save(&user)

	return c.JSON(http.StatusOK, newOrder)
}