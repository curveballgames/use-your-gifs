// Package server defines logic associated with the game server (the "host" of the game).
package server

import (
	"errors"

	"github.com/curveballgames/use-your-gifs/client"
	"github.com/curveballgames/use-your-gifs/event"
	"github.com/curveballgames/use-your-gifs/util"
	"github.com/olahol/melody"
)

const (
	minPlayers = 3
)

var (
	connections = make(map[string]*melody.Session)
)

// HandleConnection handles a new server connection, generating a new room.
func HandleConnection(s *melody.Session) {
	roomCode := util.GenerateRoomCode()
	connections[roomCode] = s

	client.HandleRoomCreation(roomCode)

	s.Write(event.CreateRoomCreatedEvent(roomCode))
}

// HandleDisconnection handles the disconnection of a server instance.
func HandleDisconnection(s *melody.Session) {
	for key, val := range connections {
		if val == s {
			delete(connections, key)
			client.DisconnectAll(key)
			break
		}
	}
}

// RegisterPlayer registers a new player in a given room. If the room does not exist, a player is not created.
func RegisterPlayer(roomCode string, playerName string) error {
	if connections[roomCode] == nil {
		return errors.New("No room with given room code")
	}

	if !client.CanConnectPlayer(roomCode) {
		return errors.New("Room full")
	}

	if client.DoesPlayerAlreadyExist(roomCode, playerName) {
		return errors.New("Player already exists with given name in room")
	}

	connections[roomCode].Write(event.CreatePlayerJoinedEvent(playerName))

	return nil
}

// StartGame requests that the server starts the game.
func StartGame(roomCode string) error {
	if connections[roomCode] == nil {
		return errors.New("No room with given room code")
	}

	if client.GetNumPlayers(roomCode) < minPlayers {
		return errors.New("Not enough players have joined the game")
	}

	connections[roomCode].Write(event.CreateStartGameEvent())
	client.StartGameForAll(roomCode)

	return nil
}
