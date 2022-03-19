package mocks

import (
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

var (
	PlayerOneID = uuid.MustParse("333edea0-d4cc-4343-a4a9-19025eddaa29")
	PlayerTwoID = uuid.MustParse("020fe7c0-1c97-48d1-b70b-2c830d5af25a")
)

func fakePlayer() structs.Player {
	return structs.Player{}
}

func FakePlayerOne() structs.Player {
	player := fakePlayer()
	player.ID = PlayerOneID
	return player
}

func FakePlayerTwo() structs.Player {
	player := FakePlayerOne()
	player.ID = PlayerTwoID
	return player
}
