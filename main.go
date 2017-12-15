package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	//"github.com/heroku/x/hmetrics/onload"
)

// global variable
var randomID = 1
type Room struct {
	roomID string
}
var RoomsMap  map[string] Room

func main() {
	// env setting
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
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	var userID = ""
	
	router.GET("/", func(c *gin.Context){
		session := sessions.Default(c)
		var v = session.Get("userID")
		if v == nil {
			userID = GenerateID()
			session.Set("userID", userID)
			session.Save()

		} else {
			// get back to his sessions
			fmt.Print(userID)
			userID = v.(string)
			FindWay(userID, c)

		}
		//c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	//router.GET("/about", func(c *gin.Context){
		//c.HTML(http.StatusOK, "<html><p>about</p><html>", nil)
	//})
	router.Run(":" + port)
}

// get to the right way by user ID
func FindWay(userID string, c *gin.Context) int {
	c.HTML(http.StatusOK, "userStatus.tmpl.html", gin.H{"userName": userID})
	return  1;
}

// Generate ID
func GenerateID() string {
	//static varable?
	randomID++;
	return strconv.Itoa(randomID)
}