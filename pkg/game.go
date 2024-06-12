package pkg

import (
	"math"
	"slices"

	"github.com/mafredri/go-trueskill"
)

type Game struct {
	Team1 *Team
	Team2 *Team
}

func NewGame(team1, team2 *Team) Game {
	return Game{
		Team1: team1,
		Team2: team2,
	}
}

func (g *Game) MatchQuality() float64 {
	team1 := g.Team1.CalculateTeam()
	team2 := g.Team2.CalculateTeam()

	return ts.MatchQuality([]trueskill.Player{team1, team2})
}

func calculateNewSkillsForPlayer(player *Player, skill trueskill.Player) {
	d := math.Pow(player.Skill.Sigma(), 2) / math.Pow(skill.Sigma(), 2)
	mu := d * skill.Mu()
	sigma := d * skill.Sigma()
	player.Skill = trueskill.NewPlayer(mu, sigma)
}

func (g *Game) UpdateRatings(firstWon bool) {
	team1 := g.Team1.CalculateTeam()
	team2 := g.Team2.CalculateTeam()

	teams := []trueskill.Player{team1, team2}
	if !firstWon {
		slices.Reverse(teams)
	}

	skills, _ := ts.AdjustSkills(teams, false)

	if !firstWon {
		slices.Reverse(skills)
	}

	for _, p := range g.Team1.GetPlayers() {
		calculateNewSkillsForPlayer(p, skills[0])
	}

	for _, p := range g.Team2.GetPlayers() {
		calculateNewSkillsForPlayer(p, skills[1])
	}
}

func (g *Game) UpdateRatingsDraw() {
	team1 := g.Team1.CalculateTeam()
	team2 := g.Team2.CalculateTeam()

	teams := []trueskill.Player{team1, team2}
	skills, _ := ts.AdjustSkills(teams, true)

	for _, p := range g.Team1.GetPlayers() {
		calculateNewSkillsForPlayer(p, skills[0])
	}

	for _, p := range g.Team2.GetPlayers() {
		calculateNewSkillsForPlayer(p, skills[1])
	}
}
