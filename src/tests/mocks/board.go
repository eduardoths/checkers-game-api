package mocks

import "github.com/eduardoths/checkers-game/src/structs"

func FakeBoard() structs.Board {
	playerOne := FakePlayerOne()
	playerTwo := FakePlayerTwo()
	p1 := &playerOne
	p2 := &playerTwo
	newChecker := func(p *structs.Player) *structs.Checker {
		checker := FakeCheckerFromPlayer(p)
		return &checker
	}
	return structs.Board{
		nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1),
		newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil,
		nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1),
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
		newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil,
		nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2),
		newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil,
	}
}
