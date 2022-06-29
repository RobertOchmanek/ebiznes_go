package api

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/model"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {

	//Obtain current database connection and fetch categories
	db := database.DbManager()
	categories := []model.Category{}
	db.Find(&categories)

	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {

	//Get category ID from query param
	id := c.Param("id")

	//Obtain current database connection and fetch category by ID
	db := database.DbManager()
	category := model.Category{}
    db.Where("id = ?", id).Preload("Products").Find(&category)

	return c.JSON(http.StatusOK, category)
}

func AddCategory(c echo.Context) error {

	//Bind json from request to object
	newCategory := new(model.Category)
	c.Bind(newCategory)

	//Obtain current database connection and save new category
	db := database.DbManager()
	db.Create(&newCategory)

	return c.JSON(http.StatusOK, newCategory)
}

func UpdateCategory(c echo.Context) error {

	//Get category ID from query param
	id := c.Param("id")

	//Bind json from request to object
	updatedCategory := new(model.Category)
	c.Bind(updatedCategory)

	//Obtain current database connection and update category by ID
	db := database.DbManager()
	category := model.Category{}
    db.Where("id = ?", id).Find(&category)

	//Update and save DB object
	category.Name = updatedCategory.Name
	db.Save(&category)

	return c.JSON(http.StatusOK, category)
}

func RemoveCategory(c echo.Context) error {

	//Get category ID from query param
	id := c.Param("id")

	//Obtain current database connection and remove category by ID
	db := database.DbManager()
	category := model.Category{}
    db.Where("id = ?", id).Find(&category)
	db.Delete(&category)

	return c.JSON(http.StatusOK, category)
}