package game

import (
	"errors"
	"math"
)

type Game struct {
	Board []Cell
	Teams []Team
}

func Create() Game {
	g := Game{}
	g.Board = make([]Cell, 0, 9*9)
	for i := 0; i < 9*9; i++ {
		g.Board = append(g.Board, createCell())
	}

	g.Teams = make([]Team, 0, 4)
	for i := 0; i < 4; i++ {
		g.Teams = append(g.Teams, createTeam())
	}

	return g
}

func (g Game) GetCell(x, y int) (*Cell, error) {
	if x < 0 || x > 8 || y < 0 || y > 8 {
		return nil, ErrOutOfRange
	}
	return &g.Board[x*9+y], nil
}

var (
	DIRECTIONS = [][]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
)

func (g Game) Set(x, y, teamID, playerID, score int) error {
	cell, err := g.GetCell(x, y)
	if err != nil {
		return err
	}
	return cell.Set(teamID, playerID, score)
}

func (g Game) Put(x, y, teamID, playerID, score int) (int, error) {
	if g.Teams[teamID].Members[playerID].RemainMark <= 0 {
		return 0, ErrNoMark
	}
	if g.Teams[teamID].Members[playerID].RemainScore < score {
		return 0, ErrNotEnoughScore
	}

	cell, err := g.GetCell(x, y)
	if err != nil {
		return 0, err
	}
	err = cell.Put(teamID, playerID, score)
	if err != nil {
		return 0, err
	}

	getScore := float64(0)
	for _, dir := range DIRECTIONS {
		xt := x
		yt := y
		cnt := 0
		for {
			nextX := xt + dir[0]
			nextY := yt + dir[1]

			checkCell, err := g.GetCell(nextX, nextY)
			if errors.Is(err, ErrOutOfRange) {
				break
			} else if err != nil {
				return 0, err
			}

			if !checkCell.Placed {
				// unplaced
				break
			}
			celldata := checkCell.GetCellData(teamID)
			if celldata.PlayerID != -1 {
				// Found team piece
				if cnt == 0 {
					break
				}
				point := float64(celldata.Score*score) * math.Pow(1.2, float64(cnt))
				if celldata.PlayerID != 0 && playerID != celldata.PlayerID {
					point *= 1.5
				}
				getScore += point
				break
			}

			xt += dir[0]
			yt += dir[1]
			cnt += 1
		}
	}

	if getScore == 0 {
		cell.Reset(teamID)
		return 0, ErrCantPlace
	}
	cell.Placed = true
	g.Teams[teamID].UseScore(playerID, score)

	return int(math.Ceil(getScore)), nil
}

func (g Game) PutAlmighty(x, y int) error {
	cell, err := g.GetCell(x, y)
	if err != nil {
		return err
	}
	cell.Placed = true
	return nil
}
