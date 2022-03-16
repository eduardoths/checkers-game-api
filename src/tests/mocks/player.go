package mocks

import (
	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/google/uuid"
)

var (
	PlayerID = uuid.MustParse("333edea0-d4cc-4343-a4a9-19025eddaa29")
)

func FakePlayer() structs.Player {
	return structs.Player{
		ID: PlayerID,
	}
}
