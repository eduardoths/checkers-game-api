package interfaces

import "github.com/eduardoths/checkers-game/src/structs"

type GameService interface {
	NewGame(playerOne, playerTwo *structs.Player) (*structs.Game, error)
}
