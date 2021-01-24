package main

import (
	"github.com/gin-gonic/gin"
	"ginsample/model"
	"ginsample/handler"
)

func main() {
	todos := model.NewTodos()
	r := gin.Default()
	r.GET("/todo", handler.TodoGetAll(todos))
	r.GET("/todo/:id", handler.TodoGet(todos))
	r.POST("/todo", handler.TodoPost(todos))
	r.PUT("/todo/:id", handler.TodoPut(todos))
	r.DELETE("/todo/:id", handler.TodoDelete(todos))
	r.Run()
}
