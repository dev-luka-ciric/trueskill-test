package main

import (
	"fmt"
	"trueskill-test/pkg"
)

func main() {
	p1 := pkg.NewDefaultPlayer("1")
	p2 := pkg.NewDefaultPlayer("2")
	p3 := pkg.NewDefaultPlayer("3")
	p4 := pkg.NewDefaultPlayer("4")

	team1 := pkg.NewTeam(&p1)
	team2 := pkg.NewTeam(&p2)
	team3 := pkg.NewTeam(&p3)
	team4 := pkg.NewTeam(&p4)

	game1 := pkg.NewGame(&team1, &team2)
	game2 := pkg.NewGame(&team3, &team4)

	fmt.Println("Game 1 Match Quality: ", game1.MatchQuality())
	fmt.Println("Game 2 Match Quality: ", game2.MatchQuality())

	game1.UpdateRatings(true)
	game2.UpdateRatingsDraw()

	fmt.Println("Player 1 Skill: ", p1.GetPlayerSkill())
	fmt.Println("Player 2 Skill: ", p2.GetPlayerSkill())
	fmt.Println("Player 3 Skill: ", p3.GetPlayerSkill())
	fmt.Println("Player 4 Skill: ", p4.GetPlayerSkill())
}
