package pkg

import "github.com/mafredri/go-trueskill"

type Team struct {
	Players []*Player
}

func NewDefaultTeam() Team {
	return Team{
		Players: []*Player{},
	}
}

func NewTeam(players ...*Player) Team {
	return Team{
		Players: players,
	}
}

func (t *Team) AddPlayer(p *Player) {
	t.Players = append(t.Players, p)
}

func (t *Team) GetPlayers() []*Player {
	return t.Players
}

func (t *Team) CalculateTeam() trueskill.Player {
	mean, stddev := 0.0, 0.0
	for _, p := range t.Players {
		mean += p.Skill.Mean()
		stddev += p.Skill.StdDev()
	}

	return trueskill.NewPlayer(mean, stddev)
}
