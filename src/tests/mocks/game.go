package mocks

import (
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

var (
	GameID = uuid.MustParse("21fce374-36ae-40a1-9dec-7bc75c1e0e4c")
)

func FakeGame() structs.Game {
	playerOne := FakePlayer()
	playerTwo := FakePlayerTwo()
	board := structs.NewBoard(&playerOne, &playerTwo)
	return structs.Game{
		ID:              GameID,
		PlayerOne:       &playerOne,
		PlayerTwo:       &playerTwo,
		Board:           board,
		IsPlayerOneTurn: true,
	}
}
