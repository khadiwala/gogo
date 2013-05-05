package game

import (
	"math/rand"
	"time"
)

type node struct {
	visits int
	wins   int
}

type tuple struct {
	row int
	col int
}

func evaluate(board *Board, turn string) map[byte]int {
	rand.Seed(time.Now().UTC().UnixNano())
	length := board.size * board.size
	moves := make([]tuple, 0, length)
	for col := 0; col < board.size; col++ {
		for row := 0; row < board.size; row++ {
			moves = append(moves, tuple{row, col})
		}
	}
	indicies := rand.Perm(length)
	for i := 0; i < length; i++ {
		m := moves[indicies[i]]
		board.Play(m.row, m.col, turn)
		turn = enemies[turn]
	}
	return board.Score()
}
