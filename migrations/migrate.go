package main

import (
	"todo/todo_back/initializers"
	"todo/todo_back/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}
