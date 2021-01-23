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