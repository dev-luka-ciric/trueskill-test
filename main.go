package main

import (
	"fmt"

	"github.com/mafredri/go-trueskill"
)

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

	players1, _ = ts.AdjustSkills(players1, false)
	players2, prob := ts.AdjustSkills(players2, false)

	players, prob := ts.AdjustSkills([]trueskill.Player{
		players2[1],
		players1[0],
	}, false)

	for _, p := range players {
		fmt.Println(p)
	}

	fmt.Println(prob)
	for _, p := range players1 {
		fmt.Println(p)
	}

	for _, p := range players2 {
		fmt.Println(p)
	}
}
