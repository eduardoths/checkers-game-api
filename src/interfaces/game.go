package interfaces

import "github.com/eduardoths/checkers-game/src/structs"

type GameService interface {
	NewGame(PlayerOne, PlayerTwo *structs.Player) (*structs.Game, error)
}
