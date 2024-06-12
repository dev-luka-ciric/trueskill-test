package main

import (
	"fmt"
	"slices"

	"github.com/mafredri/go-trueskill"
)

func calculateTeam(players []trueskill.Player) trueskill.Player {
	mean, stddev := 0.0, 0.0
	for _, p := range players {
		mean += p.Mean()
		stddev += p.StdDev()
	}

	return trueskill.NewPlayer(mean, stddev)
}

func main() {
	ts := trueskill.New(trueskill.DrawProbabilityZero())
	players1 := []trueskill.Player{
		ts.NewPlayer(),
		ts.NewPlayer(),
	}

	players2 := []trueskill.Player{
		ts.NewPlayer(),
		ts.NewPlayer(),
	}

	for i := 0; i < 20; i++ {
		slices.Reverse(players1)
		var prob1, prob2 float64
		players1, prob1 = ts.AdjustSkills(players1, false)
		players2, prob2 = ts.AdjustSkills(players2, false)

		fmt.Println("Iteration", i+1)
		fmt.Println("Probabilities")

		fmt.Println(players2)
		fmt.Println(prob1)
		fmt.Println(prob2)
		fmt.Println("Players 1")
		fmt.Println(players1)
		fmt.Println("Players 2")
		fmt.Println("------------------------------------------------")
	}

	fmt.Println("STARTING")

	team1 := calculateTeam(players1)
	team2 := calculateTeam(players2)

	teams := []trueskill.Player{team1, team2}
	teams, prob := ts.AdjustSkills(teams, false)
	fmt.Println(prob)
	fmt.Println(teams)

	// players1 := []trueskill.Player{
	// 	ts.NewPlayer(),
	// 	ts.NewPlayer(),
	// }

	// for i := 0; i < 1000000; i++ {
	// 	players1[1] = trueskill.NewPlayer(players1[0].Mean(), players1[0].StdDev())
	// 	// var prob float64

	// 	// players1, prob = ts.AdjustSkills(players1, false)

	// 	players1, _ = ts.AdjustSkills(players1, false)
	// 	// fmt.Println("Iteration", i+1)
	// 	// fmt.Println("Probabilities")
	// 	// fmt.Println(prob)
	// 	// fmt.Println("Players 1")
	// 	// fmt.Println(players1)
	// 	// fmt.Println("------------------------------------------------")
	// }

	// fmt.Println("Players 1")
	// fmt.Println(players1)

	// fmt.Println(prob)
	// for _, p := range players1 {
	// 	fmt.Println(p)
	// }

	// for _, p := range players2 {
	// 	fmt.Println(p)
	// }
}
