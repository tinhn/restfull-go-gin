package routes

import (
	"net/http"

	demo "../controllers/demo"
	edu "../controllers/edu"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	// Todo group
	tv1Api := router.Group("/v1-demo")

	tv1Api.GET("/todos", demo.GetAllTodos)
	tv1Api.POST("/todo", demo.CreateTodo)
	tv1Api.GET("/todo/:todoId", demo.GetSingleTodo)
	tv1Api.PUT("/todo/:todoId", demo.EditTodo)
	tv1Api.DELETE("/todo/:todoId", demo.DeleteTodo)

	// Student group
	sv1Api := router.Group("/v1-edu")

	sv1Api.GET("/students", edu.GetAllStudents)
	sv1Api.POST("/student", edu.CreateStudent)
	sv1Api.GET("/student/:studentid", edu.GetSingleStudent)
	sv1Api.PUT("/student/:studentid", edu.EditStudent)
	sv1Api.DELETE("/student/:studentid", edu.DeleteStudent)

	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
