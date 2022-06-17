package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"net/http"
	"github.com/labstack/echo/v4"
)

func FindUsers(c echo.Context) error {

	//Obtain current database connection and fetch users
	db := database.DbManager()
	users := []model.User{}
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func FindUser(c echo.Context) error {

	//Get object ID from query param
	id := c.Param("id")

	//Obtain current database connection and fetch user by ID
	db := database.DbManager()
	var user model.User
    db.Where("id = ?", id).Find(&user)

	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {

	//Bind json from request to object
	newUser := new(model.User)
	c.Bind(newUser)

	//Obtain current database connection and save new user
	db := database.DbManager()
	db.Create(&newUser)

	return c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c echo.Context) error {

	//Get object ID from query param
	id := c.Param("id")

	//Bind json from request to object
	updatedUser := new(model.User)
	c.Bind(updatedUser)

	//Obtain current database connection and update user by ID
	db := database.DbManager()
	var user model.User
    db.Where("id = ?", id).Find(&user)

	//Update and save DB object
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

func RemoveUser(c echo.Context) error {

	//Get object ID from query param
	id := c.Param("id")

	//Obtain current database connection and remove user by ID
	db := database.DbManager()
	var user model.User
    db.Where("id = ?", id).Find(&user)
	db.Delete(&user)

	return c.JSON(http.StatusOK, user)
}