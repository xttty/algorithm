package poker

import "testing"

func TestProbability(t *testing.T) {
	poker := buildPoker()
	cnt := 0
	f := func(cs Cards) {
		cnt++
	}
	iterate(poker[0:2], poker[2:5], 0, f)
	t.Log(cnt)
	// var gamers = []Gamer{
	// 	{
	// 		Name: "game1",
	// 		Cs:   make(Cards, gameCardsCnt),
	// 	},
	// 	{
	// 		Name: "game2",
	// 		Cs:   make(Cards, gameCardsCnt),
	// 	},
	// 	{
	// 		Name: "game3",
	// 		Cs:   make(Cards, gameCardsCnt),
	// 	},
	// }
	// copy(gamers[0].Cs, poker[0:2])
	// copy(gamers[1].Cs, poker[2:4])
	// copy(gamers[2].Cs, poker[4:6])
	// CalculateWinRate(gamers, poker[6:9], poker[9:15])
	// t.Log(poker[6:9])
	// for _, gamer := range gamers {
	// 	t.Log(gamer)
	// }
}
