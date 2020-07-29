package main

import (
	//"log"
	//"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //router
	r.GET("/", info)
}

func info(c *gin.Context) {
	c.String(200, "Web API for wscalc.com\nWritten by Tyler Beverley. Copywrite (C) 2020, all rights reserverd")
}
