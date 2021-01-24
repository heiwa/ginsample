package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginsample/model"
	"strconv"
)

func TodoGetAll(todos *model.Todos) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := todos.GetAll()
		ctx.JSON(http.StatusOK, result)
	}
}

func TodoGet(todos *model.Todos) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "id is number")
		}
		todo, err := todos.Get(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
		ctx.JSON(http.StatusOK, *todo)
	}
}

type TodoPostRequest struct {
	Title string `json:title`
}
type TodoPutRequest struct {
	Title string `json:title`
	IsFinish bool `json:isFinish`
}

func TodoPost(todos *model.Todos) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := TodoPostRequest{}
		ctx.Bind(&requestBody)

		todo := model.Todo{
			Id: len(todos.Todos),
			Title: requestBody.Title,
			IsFinish: false,
		}
		todos.Add(todo)

		ctx.Status(http.StatusNoContent)
	}
}

func TodoPut(todos *model.Todos) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "id is number")
		}
		requestBody := TodoPutRequest{}
		ctx.Bind(&requestBody)
		todo, err := todos.Get(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
		todo.Title = requestBody.Title
		todo.IsFinish = requestBody.IsFinish
		todos.Update(*todo)

		ctx.JSON(http.StatusOK, todo)
	}
}
func TodoDelete(todos *model.Todos) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "id is number")
		}
		todo, err := todos.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
		ctx.JSON(http.StatusOK, *todo)
	}
}