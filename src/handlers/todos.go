package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	ID  int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func GetTodos(todos *[]Todo, mutex *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		c.JSON(http.StatusOK, *todos)
	}
}

func CreateTodo(todos *[]Todo, mutex *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mutex.Lock()
		newTodo.ID = len(*todos) + 1
		*todos = append(*todos, newTodo)
		mutex.Unlock()
		c.JSON(http.StatusCreated, newTodo)
	}
}

func UpdateTodo(todos *[]Todo, mutex *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var updated Todo
		if err := c.ShouldBindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mutex.Lock()
		defer mutex.Unlock()
		for i, todo := range *todos {
			if todo.ID == id {
				(*todos)[i].Task = updated.Task
				(*todos)[i].Done = updated.Done
				c.JSON(http.StatusOK, (*todos)[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	}
}

func DeleteTodo(todos *[]Todo, mutex *sync.Mutex) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		mutex.Lock()
		defer mutex.Unlock()
		for i, todo := range *todos {
			if todo.ID == id {
				*todos = append((*todos)[:i], (*todos)[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"deleted":id})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	}
}