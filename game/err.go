package game

import "errors"

var (
	ErrInvalidTeamID   = errors.New("invalid team id")
	ErrInvalidPlayerID = errors.New("invalid player id")
	ErrInvalidScore    = errors.New("invalid score")
	ErrAlreadyPlaced   = errors.New("already placed cell")
	ErrCantPlace       = errors.New("can't place here")
	ErrNoMark          = errors.New("this player has no piece")
	ErrNotEnoughScore  = errors.New("this player doesn't enough score")
	ErrOutOfRange      = errors.New("cell is out of range")
)
