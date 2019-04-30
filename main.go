package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"thesis/api"

	"github.com/gin-contrib/cors"
	"golang.org/x/crypto/acme/autocert"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "word",
	})
}

func main() {
	myfile, _ := os.Create("server.log")

	gin.DefaultWriter = io.MultiWriter(myfile, os.Stdout)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", hello)
	cardAPI := r.Group("/card")
	{
		cardAPI.POST("/welfare", api.Postscore)
		cardAPI.GET("/law/:company", api.Lawsearch)
		cardAPI.GET("/qol/:company", api.Qollie)
		cardAPI.GET("/salary/:salary", api.Salary)
		cardAPI.POST("/category", api.Category)
	}

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("welfaredetector.tk", "www.welfaredetector.tk"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
	// log.Fatal(autotls.Run(r, "welfaredetector.tk", "www.welfaredetector.tk"))

	// r.Run(":80")
	// writepoint()
}
