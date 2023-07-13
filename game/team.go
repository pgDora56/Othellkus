package game

type Team struct {
	Color   string
	Members []Member
}

type Member struct {
	Name        string
	Letter      string
	RemainMark  int
	RemainScore int
}

func createTeam() Team {
	mems := make([]Member, 0, 3)
	for j := 0; j < 3; j++ {
		mems = append(mems, Member{
			Name:        "",
			Letter:      "",
			RemainMark:  7,
			RemainScore: 100,
		})
	}
	return Team{
		Color:   "",
		Members: mems,
	}
}

func (t Team) UseScore(memberID, score int) {
	t.Members[memberID].RemainMark -= 1
	t.Members[memberID].RemainScore -= score
}
