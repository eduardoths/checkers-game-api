package structs

type Board [ROW_LENGTH * COLUMN_LENGTH]*Checker

const (
	ROW_LENGTH    = 8
	COLUMN_LENGTH = 8
)

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
		rowNumber := i / 8
		isEvenRow := rowNumber%2 == 0
		isEvenColumn := i%2 == 0
		if isEvenColumn != isEvenRow { // xor equivalent
			this[i] = &Checker{Owner: player, IsKing: false}
		}
	}
}
