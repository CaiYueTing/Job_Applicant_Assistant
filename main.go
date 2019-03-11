package main

import (
	"net/http"
	"thesis/api"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var err error

// func routerEngine() *gin.Engine {
// 	r := gin.New()

// 	cardAPI := r.Group("/card")
// 	{
// 		cardAPI.POST("/welfare", postscore)
// 		cardAPI.GET("/law/:company", lawsearch)
// 		cardAPI.GET("/salary/:salary", salary)
// 		cardAPI.POST("/category", category)
// 	}

// 	return r
// }

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "word",
	})
}

func main() {

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", hello)
	cardAPI := r.Group("/card")
	{
		cardAPI.POST("/welfare", api.Postscore)
		cardAPI.GET("/law/:company", api.Lawsearch)
		cardAPI.GET("/salary/:salary", api.Salary)
		cardAPI.POST("/category", api.Category)
	}

	// log.Fatal(autotls.Run(r, "welfaredetector.tk", "www.welfaredetector.tk"))

	r.Run(":80")
	// writepoint()
}
