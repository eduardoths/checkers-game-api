package domain

import "errors"

// Movement errors
var (
	ErrInvalidMovement             = errors.New("error: invalid checker movement")
	ErrNoCheckerAtSelectedPosition = errors.New("error: no checker at selected position")
	ErrNotPlayersTurn              = errors.New("error: it's no the player's turn")
	ErrInvalidFieldMovementsArray  = errors.New("invalid_field:movements array is invalid")
)
