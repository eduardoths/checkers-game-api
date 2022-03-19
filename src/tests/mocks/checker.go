package mocks

import "github.com/eduardoths/checkers-game/src/structs"

func FakeChecker() structs.Checker {
	player := FakePlayer()
	return structs.Checker{
		Owner:  &player,
		IsKing: false,
	}
}

func FakeCheckerFromPlayer(player *structs.Player) structs.Checker {
	checker := FakeChecker()
	checker.Owner = player
	return checker
}
