package main
// Go program to create a simple REST API for managing a list of todos using Gin framework
// and a mutex for thread safety.
// The API supports CRUD operations: Create, Read, Update, and Delete.
// The todos are stored in memory in a slice and protected by a mutex to ensure thread safety.
// The API is designed to be simple and easy to understand, making it suitable for beginners
// who are learning Go and web development.
// The code is organized into a main package and a handlers package, where the handlers package

import (
	"github.com/gin-gonic/gin"
	"go-api/src/handlers"
	"sync"
)

var todos = []handlers.Todo{
	{ID: 1, Task: "Learn Go", Done: false},
	{ID: 2, Task: "Build a REST API", Done: false},
}

var mutex sync.Mutex

func main() {
	r := gin.Default()

	r.GET("/todos", handlers.GetTodos(&todos, &mutex))
	r.POST("/todos", handlers.CreateTodo(&todos, &mutex))
	r.PUT("/todos/:id", handlers.UpdateTodo(&todos, &mutex))
	r.DELETE("/todos/:id", handlers.DeleteTodo(&todos, &mutex))

	r.Run(":8080")
}
