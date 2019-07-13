// Package event defines events that can be sent to or received from the server.
package event

import "encoding/json"

type (
	// RoomCreatedEvent defines a message to the server that a room has been created.
	RoomCreatedEvent struct {
		Type     string `json:"type"`
		RoomCode string `json:"room_code"`
	}
)

// CreateRoomCreatedEvent creates an event to tell the server a room has been successfully created.
func CreateRoomCreatedEvent(roomCode string) []byte {
	evt, err := json.Marshal(RoomCreatedEvent{
		Type:     "room_created",
		RoomCode: roomCode,
	})

	if err != nil {
		println(err.Error())
	}

	return evt
}

// CreateStartGameEvent creates a start game event, telling the server to being the game with the current players.
func CreateStartGameEvent() []byte {
	evt, err := json.Marshal(Event{
		Type: "start_game",
	})

	if err != nil {
		println(err.Error())
	}

	return evt
}
