package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"thesis/api"
	"time"

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

	mon := time.Now().Month().String()
	date := time.Now().Day()
	myfile, _ := os.Create("server" + mon + strconv.Itoa(date) + ".log")

	gin.DefaultWriter = io.MultiWriter(myfile, os.Stdout)
	// port := ":" + os.Getenv("PORT")
	// stage := os.Getenv("UP_STAGE")

	r := gin.Default()
	r.Use(cors.Default())

	// r.GET("/v1", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong " + stage + " v1 ++ drone" + "v2",
	// 	})
	// })
	// r.GET("/", hello)
	cardAPI := r.Group("/card")
	{
		cardAPI.POST("/welfare", api.Postscore)
		cardAPI.GET("/law/:company", api.Lawsearch)
		cardAPI.GET("/qol/:company", api.Qollie)
		cardAPI.GET("/salary/:salary", api.Salary)
		cardAPI.POST("/category", api.Category)
	}
	// r.Run(port)
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("welfaredetector.tk", "www.welfaredetector.tk"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
	log.Fatal(autotls.Run(r, "welfaredetector.tk", "www.welfaredetector.tk"))

	// r.Run(":80")
	// writepoint()
}
