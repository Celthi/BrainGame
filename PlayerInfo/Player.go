package PlayerInfo

import (
	"github.com/BrainGame/Games"
)
// The player
type Player struct {
	Role Games.Role
	Life Games.Life
	InRoomID string
	Name string
}
