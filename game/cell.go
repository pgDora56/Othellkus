package game

type Cell struct {
	Placed bool
	data   []CellData // TeamID -> CellData
}

type CellData struct {
	Score    int
	PlayerID int // -1: おいてない, 0: 誰のでもない, 1~3: おいた
}

func createCell() Cell {
	cds := make([]CellData, 0, 4)
	for j := 0; j < 4; j++ {
		cds = append(cds, CellData{
			Score:    0,
			PlayerID: -1,
		})
	}
	return Cell{
		Placed: false,
		data:   cds,
	}
}

func (c *Cell) Set(teamID, playerID, score int) error {
	if teamID < 0 || teamID > 4 {
		return ErrInvalidTeamID
	} else if playerID < -1 || playerID > 4 {
		return ErrInvalidPlayerID
	} else if score <= 0 {
		return ErrInvalidScore
	}

	c.data[teamID] = CellData{
		Score:    score,
		PlayerID: playerID,
	}
	return nil
}

func (c *Cell) Put(teamID, playerID, score int) error {
	if teamID < 0 || teamID > 4 {
		return ErrInvalidTeamID
	} else if playerID < 0 || playerID > 4 {
		return ErrInvalidPlayerID
	} else if score <= 0 {
		return ErrInvalidScore
	} else if c.data[teamID].PlayerID != -1 {
		return ErrAlreadyPlaced
	}

	c.data[teamID].Score = score
	c.data[teamID].PlayerID = playerID

	return nil
}

func (c *Cell) GetCellData(teamID int) CellData {
	return c.data[teamID]
}

func (c *Cell) Reset(teamID int) {
	c.data[teamID] = CellData{
		Score:    0,
		PlayerID: -1,
	}
}
