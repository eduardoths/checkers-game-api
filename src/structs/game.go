package structs

import "github.com/google/uuid"

type Game struct {
	ID              uuid.UUID
	Board           *Board
	PlayerOne       *Player
	PlayerTwo       *Player
	IsPlayerOneTurn bool
}
