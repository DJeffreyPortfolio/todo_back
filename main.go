package main

import (
	"todo/todo_back/controllers"
	"todo/todo_back/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	//config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost.com:5173"}

	r.Use(cors.Default())

	// Route to create todo
	r.POST("/todos", controllers.CreateTodo)
	// Route to update todo
	r.PUT("/todos/:id", controllers.UpdateTodo)

	// Route to gather all todos
	r.GET("/todos", controllers.GetAllTodos)
	// Route to call out one todo by id
	r.GET("/todos/:id", controllers.GetTodoByID)

	// Route to delete todo by id.
	r.DELETE("todos/:id", controllers.DeleteTodo)

	r.Run()
}
