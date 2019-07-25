package handlers

import (
	"encoding/json"

	"github.com/curveballgames/use-your-gifs/client"
	"github.com/curveballgames/use-your-gifs/event"
	"github.com/curveballgames/use-your-gifs/server"
	"github.com/olahol/melody"
)

// HandleEvent is the entry point for any WebSocket event.
func HandleEvent(s *melody.Session, msg []byte) {
	var eventPayload map[string]interface{}

	if err := json.Unmarshal(msg, &eventPayload); err != nil {
		s.Write(event.CreateErrorEvent(err, "handle_event"))
		return
	}

	println(eventPayload)
	eventType := eventPayload["type"].(string)
	println("Event received: " + eventType)

	switch eventType {
	case "new_player":
		handleNewPlayer(eventPayload, s)
	case "start_game":
		handleStartGame(eventPayload, s)
	case "start_round":
		handleStartRound(eventPayload, s)
	}
}

// handleNewPlayer handles a client connecting.
func handleNewPlayer(eventPayload map[string]interface{}, clientSession *melody.Session) {
	roomCode := eventPayload["room_code"].(string)
	playerName := eventPayload["player_name"].(string)

	if err := server.RegisterPlayer(roomCode, playerName); err != nil {
		clientSession.Write(event.CreateErrorEvent(err, "register_player"))
		return
	}

	client.RegisterPlayer(roomCode, playerName, clientSession)
	clientSession.Write(event.CreatePlayerJoinedEvent(playerName))
}

// handleNewPlayer handles a client connecting.
func handleStartGame(eventPayload map[string]interface{}, clientSession *melody.Session) {
	if err := server.StartGame(eventPayload["room_code"].(string)); err != nil {
		clientSession.Write(event.CreateErrorEvent(err, "start_game"))
	}
}

func handleStartRound(eventPayload map[string]interface{}, serverSession *melody.Session) {
	client.StartNewRound(eventPayload["room_code"].(string), eventPayload["selected_player"].(string))
}
