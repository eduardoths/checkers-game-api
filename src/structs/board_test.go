package structs_test

import (
	"testing"

	"github.com/eduardoths/checkers-game/src/structs"
	"github.com/eduardoths/checkers-game/src/tests/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	playerOneID = uuid.MustParse("f01e82e4-2d03-4c06-a737-1f8d572e5bc2")
	playerTwoID = uuid.MustParse("ddc96f35-138b-4012-959b-636c19c828bc")
)

func TestNewBoard(t *testing.T) {
	playerOne := mocks.FakePlayer()
	playerTwo := mocks.FakePlayer()
	playerOne.ID = playerOneID
	playerTwo.ID = playerTwoID
	p1 := &playerOne
	p2 := &playerTwo
	newChecker := func(p *structs.Player) *structs.Checker { return &structs.Checker{Owner: p} }
	board := structs.NewBoard(p1, p2)
	want := &structs.Board{
		nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1),
		newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil,
		nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1), nil, newChecker(p1),
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
		newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil,
		nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2),
		newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil, newChecker(p2), nil,
	}
	t.Run("It should fill player one correctly", func(t *testing.T) {
		assert.Equal(t, want[:23], board[:23])
	})

	t.Run("It should fill middle rows with nil", func(t *testing.T) {
		assert.Equal(t, want[23:40], board[23:40])
	})

	t.Run("It should player two rows correctly", func(t *testing.T) {
		assert.Equal(t, want[40:], board[40:])
	})
}
