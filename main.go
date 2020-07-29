package main

import (
	"log"
	//"fmt"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Calculation struct {
	Equation string `json:"equation"`
}

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	r := gin.Default() //router

	r.Use(static.Serve("/", static.LocalFile("./public/index.html", true)))
	r.NoRoute(serveIndex)

	r.GET("/", info)

	//routes start with get, and get upgraded to ws.
	r.GET("/subscribe", subscribe)
	r.GET("/publish", publish)

	r.Run(":8000")
}

func info(c *gin.Context) {
	c.String(200, "Web API for wscalc.com\nWritten by Tyler Beverley. Copywrite (C) 2020, all rights reserverd")
}

func serveIndex(c *gin.Context) {
	c.File("./public.index.html")
}

//subscribe to calculations made by other users.
func subscribe(c *gin.Context) {

	//upgrade the connection to use ws
	ws, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("error upgrading to websockets: %s\n.", err.Error())
		return
	}

	defer ws.Close()

	//for testing.
	var calc Calculation = Calculation{
		Equation: "1+2+3=6",
	}

	err = ws.WriteJSON(&calc)
	if err != nil {
		log.Printf("Error writing client bound json: %s\n", err.Error())
	}
}

//publish a calculation for other users to view.
func publish(c *gin.Context) {

	//upgrade the connection to use ws
	ws, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading to websockets: %s\n.", err.Error())
	}

	defer ws.Close()

	var calc Calculation
	err = ws.ReadJSON(&calc)
	if err != nil {
		log.Printf("Error reading server bound json: %s\n", err.Error())
		return
	}

}
