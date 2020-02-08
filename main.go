package main

import (
	"jobassistant-server/lambdaapi"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var err error

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "word",
	})
}

func main() {
	port := ":" + os.Getenv("PORT")
	stage := os.Getenv("UP_STAGE")

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong " + stage + " v1 ++ drone" + "v2",
		})
	})
	r.GET("/", hello)
	cardAPI := r.Group("/card")
	{
		cardAPI.POST("/welfare", lambdaapi.Postscore)
		cardAPI.POST("/category", lambdaapi.Category)
		cardAPI.GET("/law/:company", lambdaapi.Lawsearch)
		cardAPI.GET("/qol/:company", lambdaapi.Qollie)
		cardAPI.GET("/salary/:salary", lambdaapi.Salary)
	}

	r.Run(port)
}
