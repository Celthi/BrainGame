package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/heroku/x/hmetrics/onload"
)

func main() {
	fmt.Print("Hello world \n")
	port := os.Getenv("PORT")
	port = "8090"

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Delims("{[{", "}]}")
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	//router.GET("/about", func(c *gin.Context){
		//c.HTML(http.StatusOK, "<html><p>about</p><html>", nil)
	//})
	router.Run(":" + port)
}