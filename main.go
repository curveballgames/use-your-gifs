// Package main is the entry point for the application
package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/curveballgames/use-your-gifs/router"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"google.golang.org/appengine"
)

const (
	port = ":57925"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	r.GET("/server", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/client", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(router.HandleConnection)
	m.HandleMessage(router.HandleMessage)
	m.HandleDisconnect(router.HandleDisconnection)

	r.Run(port)

	appengine.Main()
}