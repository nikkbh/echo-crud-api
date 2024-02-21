package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nikkbh/echo-crud-api/controllers"
	"github.com/nikkbh/echo-crud-api/initializers"
	"github.com/nikkbh/echo-crud-api/utils"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{validator: validator.New()}
	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.POST("/todo", controllers.CreateTodo)
	e.GET("/todos", controllers.GetTodos)
	e.DELETE("/todo/:id", controllers.DeleteTodo)
	e.PUT("/todo/:id", controllers.UpdateTodo)

	// Error Handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if apiError, ok := err.(utils.APIError); ok {
			c.JSON(apiError.Status, map[string]any{apiError.Code: apiError.Msg})
			return
		} else if c.Param("id") == "" {
			c.JSON(http.StatusMethodNotAllowed, map[string]any{"METHOD_NOT_ALLOWED": "Path parameter missing"})
		} else {
			c.JSON(http.StatusInternalServerError, map[string]any{"INTERNAL_SERVER_ERROR": "Internal Server Error"})
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
