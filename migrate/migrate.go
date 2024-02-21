package main

import (
	"github.com/nikkbh/echo-crud-api/initializers"
	"github.com/nikkbh/echo-crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}
