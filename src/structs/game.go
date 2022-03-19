package structs

import "github.com/google/uuid"

type Game struct {
	ID              uuid.UUID
	Board           *Board
	PlayerOne       *Player
	PlayerTwo       *Player
	IsPlayerOneTurn bool
}

func (this *Game) CurrentPlayerID() uuid.UUID {
	if this.IsPlayerOneTurn {
		return this.PlayerOne.ID
	}
	return this.PlayerTwo.ID
}
