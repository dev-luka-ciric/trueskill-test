package pkg

import "github.com/mafredri/go-trueskill"

type Player struct {
	SteamId string
	Skill   trueskill.Player
}

func NewDefaultPlayer(steamId string) Player {
	return Player{
		SteamId: steamId,
		Skill:   ts.NewPlayer(),
	}
}

func NewPlayer(steamId string, mean, stddev float64) Player {
	return Player{
		SteamId: steamId,
		Skill:   trueskill.NewPlayer(mean, stddev),
	}
}

func (p *Player) GetPlayerSkill() trueskill.Player {
	return p.Skill
}
