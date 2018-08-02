package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
	"strconv"
	"github.com/BrainGame/Games"
	"github.com/BrainGame/PlayerInfo"
	"github.com/BrainGame/RoomInfo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/rand"
)

// user count/user ID
var clientIDSeq = 0

// Room ID -> Room
var RoomsMap = make(map[string]*RoomInfo.Room)

func main() {
	/* const (
        Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
        Ltime                         // the time in the local time zone: 01:23:23
        Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
        Llongfile                     // full file name and line number: /a/b/c/d.go:23
        Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
        LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
        LstdFlags     = Ldate | Ltime // initial values for the standard logger
	) */
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	file, err := os.OpenFile("./ErrorLog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	log.SetOutput(file)
// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// env setting
	fmt.Println("Let the game begin!")
	port := os.Getenv("PORT") // try to get the port from env
	//port = "8090" // server port define for test only

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Delims("{[{", "}]}") // This is used for Vue at the beginning, but deprecated since shift to React
	router.Use(gin.Logger())
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("UserSession", store))
	var clientID = ""

	router.GET("/", func(c *gin.Context) {
		roomID := c.DefaultQuery("roomID", "") // shortcut for c.Request.URL.Query().Get("lastname")
		// The index page
		if "" == roomID {
			// visit first time
			//roomID = strconv.Itoa(rand.Intn(1000))
			//RoomsMap[roomID] = &RoomInfo.Room{roomID, make(map[string]*PlayerInfo.Player), "starting"}
			c.JSON(http.StatusOK, gin.H{"Title": "Oyster Game", "status": "There is no room created yet, you could create room."})


		} else {
			// he knows which room to seat
			// but the room is not created yet
			if nil == RoomsMap[roomID] {
				log.Println("Room %s is not created yet!", roomID)
				c.String(http.StatusOK, "Please create room")
				return
				// create room
				// RoomsMap[roomID] = &RoomInfo.Room{roomID, make(map[string]*PlayerInfo.Player), "starting"}
			} 	
	
			log.Println("Handle the client.")
			session := sessions.Default(c)
			var v = session.Get("clientID")
			// the client does not login yet
			if v == nil {
				clientID = GenerateID()

				PlayerID := strconv.Itoa(rand.Intn(1000))

				// create player
				player := &PlayerInfo.Player{Games.LRoles{"Wolf", "Kill the smart"}, Games.Life{1}, roomID, PlayerID}
				RoomsMap[roomID].Seats[clientID] = player
				session.Set("clientID", clientID)
				session.Save()
				c.String(http.StatusOK, "First time log, you Player ID is %s", PlayerID)
				
			} else {
				log.Println("user has been logged.")
				// get back to his sessions

				clientID = v.(string)
				// get the room he seats
				room := RoomsMap[roomID]
				if nil == room {
					log.Println("Room %s has vanished", roomID)
					c.String(http.StatusOK, "The room %s is invalid, please create a new room", roomID)
					return
				}
				playerNames := ""
				for _, player := range room.Seats {
					playerNames += player.Name + " "
				}
				c.JSON(http.StatusOK, gin.H{"users": playerNames, "status": "list players", "roomID": roomID})


		}
	}
	
		//c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	//c.HTML(http.StatusOK, "<html><p>about</p><html>", nil)
	//})

	router.POST("/Room", func(c *gin.Context) {
		var roomConfig RoomInfo.RoomConfig
		if err := c.ShouldBindJSON(&roomConfig); err == nil {
			roomID := roomConfig.RoomID

			fmt.Println(roomConfig.RoomName)
			if (nil != RoomsMap) {
				c.String(200, "Room %s has been created, please use another room ID or join the room", roomID)
				return 
			}
			RoomsMap[roomID] = &RoomInfo.Room{roomID, make(map[string]*PlayerInfo.Player), "starting"}

			c.JSON(200, gin.H{"status": "Room created"})
		}
	})

	router.Run(":" + port)
}

// get to the right way by user ID
func FindWay(userID string, c *gin.Context) int {
	
	if (nil == c) {
		return 2
	}
	c.String(http.StatusOK, "userName %s", userID)
	return 1
}

// Generate ID
func GenerateID() string {
	//static varable?
	clientIDSeq++
	return strconv.Itoa(clientIDSeq)
}
