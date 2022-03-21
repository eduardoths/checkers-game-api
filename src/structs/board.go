package structs

import "github.com/eduardoths/checkers-game/src/domain"

type Board [ROW_LENGTH * COLUMN_LENGTH]*Checker

const (
	ROW_LENGTH    = 8
	COLUMN_LENGTH = 8
	BOARD_INIT    = 0
	BOARD_END     = 63
)

var validMovements = []int{-18, -14, -9, -7, 7, 9, 14, 18}

func NewBoard(playerOne, playerTwo *Player) *Board {
	board := new(Board)
	board.fillPlayerOneCheckers(playerOne)
	board.fillPlayerTwoCheckers(playerTwo)
	return board
}

func (this *Board) fillPlayerOneCheckers(playerOne *Player) {
	startPos := 0
	endPos := ROW_LENGTH * 3
	this.fillWithPlayer(startPos, startPos+endPos, playerOne)
}

func (this *Board) fillPlayerTwoCheckers(playerTwo *Player) {
	startPos := 40
	endPos := ROW_LENGTH * 3
	this.fillWithPlayer(startPos, startPos+endPos, playerTwo)
}

func (this *Board) fillWithPlayer(start, end int, player *Player) {
	for i := start; i < end; i++ {
		isValidColumn := this.isColumnValid(i)
		if isValidColumn {
			this[i] = &Checker{Owner: player, IsKing: false}
		}
	}
}

func (this *Board) isColumnValid(pos int) bool {
	rowNumber := pos / 8
	isEvenRow := rowNumber%2 == 0
	isEvenColumn := pos%2 == 0
	if isEvenColumn != isEvenRow { // xor equivalent
		return true
	}
	return false
}

func (this *Board) GetCheckerFromPos(pos int) *Checker {
	if this.positionIsOnBoard(pos) {
		return this[pos]
	}
	return nil
}

func (this *Board) IsValidPosition(pos int) bool {
	validations := []func(int) bool{
		this.positionIsOnBoard,
		this.isColumnValid,
		this.positionHasNoChecker,
	}
	for _, validation := range validations {
		isValid := validation(pos)
		if !isValid {
			return false
		}
	}
	return true
}

func (this *Board) positionHasNoChecker(pos int) bool {
	checker := this.GetCheckerFromPos(pos)
	return checker == nil
}

func (this *Board) positionIsOnBoard(pos int) bool {
	if pos >= BOARD_INIT && pos <= BOARD_END {
		return true
	}
	return false
}

func (this *Board) Move(from int, mobeBy int) error {
	return domain.ErrInvalidMovement
}
