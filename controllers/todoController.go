package controllers

import (
	"net/http"
	"todo/todo_back/initializers"
	"todo/todo_back/models"

	"github.com/gin-gonic/gin"
)

type CreateTodoInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Create todos.
func CreateTodo(c *gin.Context) {
	var input CreateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Title: input.Title, Body: input.Body}

	result := initializers.DB.Create(&todo)

	// Check for errors
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

// Gather all Todos into one place.
func GetAllTodos(c *gin.Context) {
	// Get all records from db.
	var todos []models.Todo

	if err := initializers.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos not found!"})
		return
	}

	// Return all todos
	c.JSON(http.StatusOK, gin.H{"data": todos})

}

// Call out a single todo.
func GetTodoByID(c *gin.Context) {
	var todo models.Todo
	// get id from url
	id := c.Param("id")

	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found!"})
		return
	}

	// Return single todo
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// Update that single todo.
func UpdateTodo(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// Create stuct to bind data from request to.
	var todoBody struct {
		Title string
		Body  string
	}
	c.ShouldBind(&todoBody)

	var todo models.Todo
	initializers.DB.First(&todo, id)

	// Update todo
	initializers.DB.Model(&todo).Updates(models.Todo{Title: todoBody.Title, Body: todoBody.Body})

	// Return single todo
	c.JSON(200, gin.H{
		"todo": todo,
	})
}

// Update that single todo.
func DeleteTodo(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// Delete post by id
	initializers.DB.Delete(&models.Todo{}, id)

	// Return single todo
	c.JSON(200, gin.H{
		"message": "Deleted Successfully",
	})
}
