package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nikkbh/echo-crud-api/initializers"
	"github.com/nikkbh/echo-crud-api/models"
	"github.com/nikkbh/echo-crud-api/utils"
	"gorm.io/gorm"
)

func CreateTodo(c echo.Context) (err error) {
	var requestBody struct {
		models.Todo
		Task string `json:"task" validate:"required"`
	}
	c.Bind(&requestBody)
	if err = c.Validate(requestBody); err != nil {
		return utils.BadRequestError()
	}
	todo := models.Todo{Task: requestBody.Task, Status: "pending"}
	result := initializers.DB.Create(&todo)
	if result.Error != nil {
		return utils.DBCreateError()
	}
	return c.JSON(http.StatusCreated, todo.ID)
}

func GetTodos(c echo.Context) (err error) {
	todos := []models.Todo{}
	status := c.QueryParam("status")
	var dbResult *gorm.DB
	if status != "" {
		dbResult = initializers.DB.Where("status = ?", status).Find(&todos)
	} else {
		dbResult = initializers.DB.Find(&todos)
	}
	if dbResult.Error != nil {
		return utils.DBFetchError()
	}
	return c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	// Find the TODO to be updated by id
	initializers.DB.Find(&todo, id)
	dbResult := initializers.DB.Model(&todo).Updates(models.Todo{Status: "done"})
	if dbResult.Error != nil {
		return utils.DBUpdateError()
	}
	return c.NoContent(http.StatusAccepted)
}

func DeleteTodo(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	dbResult := initializers.DB.Delete(&models.Todo{}, id)
	if dbResult.Error != nil {
		return utils.DBDeleteError()
	}
	return c.NoContent(http.StatusAccepted)
}
