package mocks

import "github.com/eduardoths/checkers-game/src/structs"

func fakeChecker() structs.Checker {
	player := FakePlayer()
	return structs.Checker{
		Owner:  &player,
		IsKing: false,
	}
}

func FakeCheckerFromPlayer(player *structs.Player) structs.Checker {
	checker := fakeChecker()
	checker.Owner = player
	return checker
}
