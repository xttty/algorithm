package poker

import "testing"

func TestCalculate(t *testing.T) {
	// pokers := buildPoker()
	// c1 := pokers[0:5]
	// c2 := pokers[5:10]
	var c1 = []Card{
		{
			Num:  10,
			Suit: 1,
		},
		{
			Num:  2,
			Suit: 1,
		},
		{
			Num:  2,
			Suit: 4,
		},
		{
			Num:  7,
			Suit: 2,
		},
		{
			Num:  7,
			Suit: 3,
		},
	}
	var c2 = []Card{
		{
			Num:  2,
			Suit: 3,
		},
		{
			Num:  2,
			Suit: 2,
		},
		{
			Num:  7,
			Suit: 1,
		},
		{
			Num:  7,
			Suit: 4,
		},
		{
			Num:  10,
			Suit: 4,
		},
	}
	l1, _ := getLevel(c1)
	l2, _ := getLevel(c2)
	res, _ := compare(c1, c2)
	t.Log(c1, c2)
	t.Log(l1, l2)
	t.Log(res)
}
