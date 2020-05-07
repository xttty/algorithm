package poker

type cardsOperationFunc func(cs Cards)

// CalculateWinRate 计算各个玩家的胜率
func CalculateWinRate(gamers []Gamer, commonPool, restPool Cards) error {
	if len(commonPool)+len(restPool) < 5 {
		return ErrPoolNotEnough
	}
	if len(commonPool) > 5 {
		return ErrCommonPool
	}
	for i := range gamers {
		if len(gamers[i].Cs) != gameCardsCnt {
			return ErrGamer
		}
		gamers[i].winCnt = 0
	}
	compareCnt := 0
	calFunc := func(cs Cards) {
		compareCnt++
		maxCs, temp := make(Cards, gameCardsCnt+commonPoolCnt), make(Cards, gameCardsCnt+commonPoolCnt)
		copy(temp[0:gameCardsCnt], gamers[0].Cs)
		copy(temp[gameCardsCnt:], cs)
		copy(maxCs, temp)
		maxFiveCs, _ := getMaxCards(maxCs)
		for i := 1; i < len(gamers); i++ {
			copy(temp, gamers[i].Cs)
			maxFiveCsTemp, _ := getMaxCards(temp)
			res, err := compare(maxFiveCs, maxFiveCsTemp)
			if err == nil && res == fail {
				copy(maxCs, temp)
				copy(maxFiveCs, maxFiveCsTemp)
			}
		}
		for idx := range gamers {
			copy(temp, gamers[idx].Cs)
			maxFiveCsTemp, _ := getMaxCards(temp)
			res, err := compare(maxFiveCs, maxFiveCsTemp)
			if err == nil && res == draw {
				gamers[idx].winCnt++
			}
		}
	}
	if len(commonPool) < commonPoolCnt {
		restLen := commonPoolCnt - len(commonPool)
		commonPool = append(commonPool, restPool[0:restLen]...)
		restPool = restPool[restLen:]
		calFunc(commonPool)
		iterate(commonPool, restPool, commonPoolCnt-restLen, calFunc)
	} else {
		calFunc(commonPool)
	}
	for i := range gamers {
		gamers[i].WinRate = float64(gamers[i].winCnt) / float64(compareCnt)
	}

	return nil
}

func getMaxCards(pool Cards) (Cards, error) {
	if len(pool) < 5 {
		return pool, ErrPoolNotEnough
	}

	maxRes := make(Cards, 5)
	temp := make(Cards, 5)
	copy(maxRes, pool[0:5])
	copy(temp, maxRes)
	iterate(temp, pool[5:], 0, func(cs Cards) {
		res, err := compare(maxRes, cs)
		if err == nil && res != win {
			copy(maxRes, cs)
		}
	})

	return maxRes, nil
}

func iterate(curCards, pool Cards, pos int, op cardsOperationFunc) {
	if len(pool) < len(curCards) {
		return
	}
	if pos >= len(curCards) {
		op(curCards)
	}

}
