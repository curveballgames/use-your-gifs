// Package router defines methods for handling the routing of incoming WebSocket requests.
package router

import (
	"github.com/curveballgames/use-your-gifs/client"
	"github.com/curveballgames/use-your-gifs/handlers"
	"github.com/curveballgames/use-your-gifs/server"
	"github.com/olahol/melody"
)

// HandleConnection handles a new WebSocket connection from either a client or server.
func HandleConnection(s *melody.Session) {
	if s.Request.RequestURI == "/server" {
		server.HandleConnection(s)
	}
}

// HandleMessage handles any message from a client or server, routing it accordingly.
func HandleMessage(s *melody.Session, msg []byte) {
	handlers.HandleEvent(s, msg)
}

// HandleDisconnection handles a disconnection of a server or client.
func HandleDisconnection(s *melody.Session) {
	if s.Request.RequestURI == "/server" {
		server.HandleDisconnection(s)
	} else {
		client.HandleDisconnection(s)
	}
}
