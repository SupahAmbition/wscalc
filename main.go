package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"wscalc/calculations"
)

//used for upgrading a get request to ws
var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := setupRouter()
	r.Run(":8000")
}

func setupRouter() *gin.Engine {

	r := gin.Default() //router

	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.GET("/info", info)

	//routes start with get, and get upgraded to ws.
	r.GET("/subscribe", subscribe)
	r.GET("/publish", publish)

	return r
}

func info(c *gin.Context) {
	c.String(200, "Web API for wscalc.com\nWritten by Tyler Beverley. Copywrite (C) 2020, all rights reserverd")
}

func serveIndex(c *gin.Context) {
	c.File("./public/index.html")
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

	type response struct {
		NumCalculations int                        `json:"numCalculations"`
		Calculations    []calculations.Calculation `json:"calculations"`
	}

	var lastTime time.Time = time.Now()
	cs := calculations.GetInstance()

	//wait for updates, then send new data to user.
	for {

		//if out last time doesnt match the updated stack.
		//then update the user.
		if !(lastTime.Equal(cs.Peek().Timestamp)) {

			//fmt.Println("Updating the user!")
			lastTime = cs.Peek().Timestamp

			// update the client.
			calculations := cs.Peek10()

			r := response{
				NumCalculations: len(calculations),
				Calculations:    calculations,
			}

			err = ws.WriteJSON(r)
			if err != nil {
				log.Printf("Error writing client bound json: %s\n", err.Error())
				break
			}
		}

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

	type request struct {
		Equation string `json:"equation"`
		User     string `json:"user"`
	}

	cs := calculations.GetInstance()
	//loop listen / write.
	for {
		var r request
		err = ws.ReadJSON(&r)
		if err != nil {
			log.Printf("Error reading server bound json: %s\n", err.Error())
			break
		} else {
			//fmt.Printf("Got input from user! %#v\n", r)
			calc := calculations.NewCalculation(r.Equation)
			cs.Push(*calc)
		}
	}

}
