package main

import (
	"github.com/gin-gonic/gin"
)

func homepage(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "hello" + name,
	})
}

func lawcount(c *gin.Context) {
	count := c.Param("company")
	c.JSON(200, gin.H{
		"message": len(count),
	})
}

func main() {
	r := gin.Default()
	r.GET("/welfare/:name", homepage)
	r.GET("/law/:company", lawcount)
	r.Run()
}
