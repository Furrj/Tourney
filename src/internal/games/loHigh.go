package games

import (
	"math/rand"
)

type LoHigh struct{}

func (LoHigh) RunGame(iterations uint) uint8 {
	p1Score := 0
	p2Score := 0
	draws := 0

	for i := uint(1); i <= iterations; i++ {
		result := gameRound()

		if result == 1 {
			p1Score++
		} else if result == 2 {
			p2Score++
		} else {
			draws++
		}
	}

	if p1Score > p2Score {
		return 1
	} else if p1Score < p2Score {
		return 2
	} else {
		return 3
	}
}

func gameRound() uint8 {
	p1Hand := rand.Intn(101)
	p2Hand := rand.Intn(101)

	if p1Hand > p2Hand {
		return 1
	} else if p1Hand < p2Hand {
		return 2
	} else {
		return 3
	}
}
