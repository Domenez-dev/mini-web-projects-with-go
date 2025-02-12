package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		todos, err := GetAllTodos()
		if err != nil {
			fmt.Println("Error getting todos:", err)
			ctx.String(http.StatusInternalServerError, "Error retrieving todos")
			return
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	r.POST("/todos", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		status := ctx.PostForm("status")
		_, err := CreateTodo(title, status)
		if err != nil {
			fmt.Println("Error creating todo:", err)
			ctx.String(http.StatusInternalServerError, "Error creating todo")
			return
		}
		ctx.Redirect(http.StatusFound, "/")
	})

	r.GET("/todos/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			ctx.String(http.StatusBadRequest, "Invalid ID")
			return
		}
		err = DeleteTodo(id)
		if err != nil {
			fmt.Println("Error deleting todo:", err)
			ctx.String(http.StatusInternalServerError, "Error deleting todo")
			return
		}
		ctx.Redirect(http.StatusFound, "/")
	})

	r.Run()
}
