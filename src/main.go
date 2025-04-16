package main

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
