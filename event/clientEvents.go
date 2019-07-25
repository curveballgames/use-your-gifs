package event

import (
	"encoding/json"
)

type (
	// NewRoundEvent defines an event that informs the clients that a new round is starting.
	NewRoundEvent struct {
		Type       string `json:"type"`
		Controller string `json:"controller"`
	}
)

// CreateGameStartedEvent creates a game started event to send to clients.
func CreateGameStartedEvent() []byte {
	evt, _ := json.Marshal(Event{
		Type: "game_started",
	})

	return evt
}

// CreateNewRoundEvent creates an event to send to clients to start a new prompting round.
func CreateNewRoundEvent(controller string) []byte {
	evt, _ := json.Marshal(NewRoundEvent{
		Type:       "new_round",
		Controller: controller,
	})

	return evt
}
