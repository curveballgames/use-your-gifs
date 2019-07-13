// Pakcage client defines logic for associated with the clients, or "players" of the game.
package client

import (
	"github.com/curveballgames/use-your-gifs/event"
	"github.com/olahol/melody"
)

type (
	// Player is a simple struct to identify a player/ client in a room.
	Player struct {
		Name          string
		ClientSession *melody.Session
	}

	// Room is a room full of players, keyed by name.
	Room struct {
		Players map[string]*Player
	}
)

const (
	maxPlayerCount = 6
)

var (
	roomToPlayersMap       = make(map[string]*Room)
	playerSessionToRoomMap = make(map[*melody.Session]string)
)

// CanConnectPlayer determines whether a player can be connected for the given room. The room must exist for players to be able to connect.
func CanConnectPlayer(roomCode string) bool {
	return roomToPlayersMap[roomCode] != nil && len(roomToPlayersMap[roomCode].Players) < maxPlayerCount
}

// DoesPlayerAlreadyExist determines whether a room already has a player with the given name
func DoesPlayerAlreadyExist(roomCode string, playerName string) bool {
	return roomToPlayersMap[roomCode].Players[playerName] != nil
}

// HandleRoomCreation handles the creation of a new room for a server.
func HandleRoomCreation(roomCode string) {
	roomToPlayersMap[roomCode] = &Room{
		Players: make(map[string]*Player),
	}
}

// RegisterPlayer registers a new player in a room.
func RegisterPlayer(roomCode string, playerName string, clientSession *melody.Session) {
	roomToPlayersMap[roomCode].addPlayer(playerName, clientSession)
	playerSessionToRoomMap[clientSession] = roomCode
}

// HandleDisconnection handles a client disconnecting.
func HandleDisconnection(s *melody.Session) {
	roomName := playerSessionToRoomMap[s]

	if roomName == "" {
		return
	}

	room := roomToPlayersMap[roomName]

	if room == nil {
		return
	}

	for playerName := range room.Players {
		if room.Players[playerName].ClientSession == s {
			delete(room.Players, playerName)
			break
		}
	}
}

// DisconnectAll disconnected all clients in a given room.
func DisconnectAll(roomCode string) {
	room := roomToPlayersMap[roomCode]

	if room == nil {
		return
	}

	for playerName := range room.Players {
		room.Players[playerName].ClientSession.CloseWithMsg(event.CreateServerDisconnectedEvent())
	}

	delete(roomToPlayersMap, roomCode)
}

func (r *Room) addPlayer(playerName string, clientSession *melody.Session) {
	r.Players[playerName] = &Player{
		Name:          playerName,
		ClientSession: clientSession,
	}
}
