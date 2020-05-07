package poker

import (
	"errors"
	"math/rand"
	"sort"
	"time"
)

const (
	win  = iota // 赢
	fail        // 输
	draw        // 平局
)

const (
	spade   = iota // 黑桃
	heart          // 红桃
	club           // 梅花
	diamond        // 方片
)

const (
	two = 2 + iota
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queue
	king
	ace
)

const (
	hightCard     = iota // 单张
	onePair              // 一对
	twoPairs             // 两对
	threeOfKind          // 三条
	straight             // 顺子
	flush                // 同花
	fullHouse            // 葫芦
	fourOfKind           // 四条
	straightFlush        // 同花顺
)

const commonPoolCnt = 5
const gameCardsCnt = 2

// ErrPoolNotEnough 牌数不够五个
var ErrPoolNotEnough = errors.New("cards pool is not enough")

// ErrCommonPool 公共牌池错误
var ErrCommonPool = errors.New("common pool occurred error")

// ErrGamer 玩家错误
var ErrGamer = errors.New("Gamer occurred error")

// ErrCardsCnt 最终选出的卡牌数不等于五
var ErrCardsCnt = errors.New("count of cards is not equal five")

// Gamer 玩家
type Gamer struct {
	Name    string
	Cs      Cards
	WinRate float64
	winCnt  int
}

// Card 扑克结构
type Card struct {
	Num  int
	Suit int //花色
}

// Cards 排序使用
type Cards []Card

type level struct {
	attr int
	cs   Cards
}

func (cs *Cards) Len() int {
	return len(*cs)
}
func (cs *Cards) Swap(i, j int) {
	(*cs)[i], (*cs)[j] = (*cs)[j], (*cs)[i]
}
func (cs *Cards) Less(i, j int) bool {
	return (*cs)[i].Num < (*cs)[j].Num
}

func compare(a, b Cards) (int, error) {
	if len(a) != commonPoolCnt || len(b) != commonPoolCnt {
		return draw, ErrCardsCnt
	}
	tempA, tempB := make(Cards, len(a)), make(Cards, len(b))
	copy(tempA, a)
	copy(tempB, b)
	la, err := getLevel(tempA)
	if err != nil {
		return draw, err
	}
	lb, err := getLevel(tempB)
	if err != nil {
		return draw, err
	}
	if la.attr == lb.attr {
		for i := commonPoolCnt - 1; i >= 0; i-- {
			if la.cs[i].Num == lb.cs[i].Num {
				continue
			} else if la.cs[i].Num > lb.cs[i].Num {
				return win, nil
			} else {
				return fail, nil
			}
		}
	} else if la.attr > lb.attr {
		return win, nil
	} else {
		return fail, nil
	}
	return draw, nil
}

func getLevel(cs Cards) (level, error) {
	if len(cs) != commonPoolCnt {
		return level{}, ErrCardsCnt
	}
	sort.Sort(&cs)
	l := level{}
	isStraight := true
	isFlush := true
	for i := 0; i < commonPoolCnt-1; i++ {
		if cs[i].Num != cs[i+1].Num-1 {
			isStraight = false
		}
		if cs[i].Suit != cs[i+1].Suit {
			isFlush = false
		}
	}
	l.cs = cs
	if isStraight && isFlush {
		l.attr = straightFlush
	} else {
		if cs[0].Num == cs[1].Num && cs[1].Num == cs[2].Num && cs[2].Num == cs[3].Num {
			l.attr = fourOfKind
			l.cs[0], l.cs[4] = l.cs[4], l.cs[0]
		} else if cs[1].Num == cs[2].Num && cs[2].Num == cs[3].Num && cs[3].Num == cs[4].Num {
			l.attr = fourOfKind
		} else if cs[0].Num == cs[1].Num && cs[2].Num == cs[3].Num && cs[3].Num == cs[4].Num {
			l.attr = fullHouse
			temp := make(Cards, commonPoolCnt)
			copy(temp[0:2], cs[0:2])
			copy(temp[2:], cs[2:])
			l.cs = temp
		} else if cs[0].Num == cs[1].Num && cs[1].Num == cs[2].Num && cs[3].Num == cs[4].Num {
			l.attr = fullHouse
			temp := make(Cards, commonPoolCnt)
			copy(temp[0:2], cs[3:])
			copy(temp[2:], cs[0:3])
			l.cs = temp
		} else if isFlush {
			l.attr = flush
		} else if isStraight {
			l.attr = straight
		} else if cs[0].Num == cs[1].Num && cs[1].Num == cs[2].Num {
			l.attr = threeOfKind
			temp := make(Cards, commonPoolCnt)
			copy(temp[2:], cs[0:3])
			temp[1] = cs[4]
			temp[0] = cs[3]
			l.cs = temp
		} else if cs[1].Num == cs[2].Num && cs[2].Num == cs[3].Num {
			l.attr = threeOfKind
			temp := make(Cards, commonPoolCnt)
			copy(temp[2:], cs[1:4])
			temp[1] = cs[4]
			temp[0] = cs[0]
			l.cs = temp
		} else if cs[2].Num == cs[3].Num && cs[3].Num == cs[4].Num {
			l.attr = threeOfKind
		} else {
			cardsMap := make(map[int][]Card)
			parsCnt := 0
			for _, c := range cs {
				if temp, ok := cardsMap[c.Num]; ok {
					cardsMap[c.Num] = append(temp, c)
					parsCnt++
				} else {
					cardsMap[c.Num] = append([]Card{}, c)
				}
			}
			if parsCnt == 2 {
				l.attr = twoPairs
			} else if parsCnt == 1 {
				l.attr = onePair
			} else {
				l.attr = hightCard
			}
			// 对子放前面
			idx := 4
			if parsCnt > 0 {
				for i := ace; i >= two; i-- {
					if temp, ok := cardsMap[i]; ok && len(temp) > 1 {
						for _, c := range temp {
							l.cs[idx] = c
							idx--
						}
					}
				}
			}
			for i := ace; i >= two; i-- {
				if temp, ok := cardsMap[i]; ok && len(temp) == 1 {
					l.cs[idx] = temp[0]
					idx--
				}
			}
		}
	}

	return l, nil
}

func buildPoker() Cards {
	all := make(Cards, 52)
	idx := 0
	for i := two; i <= ace; i++ {
		for j := spade; j <= diamond; j++ {
			all[idx].Num = i
			all[idx].Suit = j
			idx++
		}
	}
	// shuffle
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, j := range r.Perm(len(all)) {
		all[i], all[j] = all[j], all[i]
	}
	return all
}
