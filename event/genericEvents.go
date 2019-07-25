// Package event defines common events that could be shared by any WebSocket listener.
package event

import "encoding/json"

type (
	// Event defines a generic event with just a type.
	Event struct {
		Type string `json:"type"`
	}

	// ErrorEvent defines an error from the server.
	ErrorEvent struct {
		Type         string `json:"type"`
		Subtype      string `json:"subType"`
		ErrorMessage string `json:"error_message"`
	}

	// PlayerJoinedEvent defines a request to the server to connect a player to the room.
	PlayerJoinedEvent struct {
		Type       string `json:"type"`
		PlayerName string `json:"player_name"`
	}
)

// CreateErrorEvent creates a WebSocket-ready error payload.
func CreateErrorEvent(e error, subtype string) []byte {
	evt, err := json.Marshal(ErrorEvent{
		Type:         "error",
		Subtype:      subtype,
		ErrorMessage: e.Error(),
	})

	if err != nil {
		println(err.Error())
	}

	return evt
}

// CreatePlayerJoinedEvent creates an event that tells the clients or server that a player has joined.
func CreatePlayerJoinedEvent(playerName string) []byte {
	evt, err := json.Marshal(PlayerJoinedEvent{
		Type:       "player_joined",
		PlayerName: playerName,
	})

	if err != nil {
		println(err.Error())
	}

	return evt
}

// CreateServerDisconnectedEvent creates an event to tell all that the server has disconnected.
func CreateServerDisconnectedEvent() []byte {
	evt, err := json.Marshal(Event{
		Type: "server_disconnected",
	})

	if err != nil {
		println(err.Error())
	}

	return evt
}
