package main

import "github.com/gin-gonic/gin"

func main ()  {
	router := gin.Default()
	v1 := router.Group("/api/crud"){
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
	
}