package RoomInfo

import (
	"github.com/BrainGame/PlayerInfo"
)
type Room struct {
	RoomID string
	Seats map[string] *PlayerInfo.Player // seat corresponding to player: seatID -> player
	Status string // idle, playing, 
}

// the json send from creating room
type RoomConfig struct {
	RoomID string `json:"roomID"`
	SeatsNum int  `json:"seatsNum"`
	RoomName string `json:"roomName"`
}