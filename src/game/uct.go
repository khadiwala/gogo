package game

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	visits int
	wins   int
	c      []*node
}

type vertex struct {
	row int
	col int
}

func allMoves(size int) []vertex {
	length := size * size
	moves := make([]vertex, 0, length)
	for col := 0; col < size; col++ {
		for row := 0; row < size; row++ {
			moves = append(moves, vertex{row, col})
		}
	}
	return moves
}

func uct(board *Board, turn string) {
	for _, m := range allMoves(board.size) {
		c := board.Copy()
		c.Play(m.row, m.col, turn)
		fmt.Println(m, evaluate(&c, enemies[turn]))
	}
}

func seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func random_evaluate(board *Board, turn string) map[byte]int {
	c := board.Copy()
	board = &c
	for i := 0; i < 10*board.size; i++ {
		row := rand.Intn(board.size)
		col := rand.Intn(board.size)
		board.Play(row, col, turn)
		turn = enemies[turn]
	}
	return board.Score()
}

func evaluate(board *Board, turn string) map[byte]int {
	moves := allMoves(board.size)
	length := board.size * board.size
	indicies := rand.Perm(length)
	for i := 0; i < length; i++ {
		m := moves[indicies[i]]
		board.Play(m.row, m.col, turn)
		turn = enemies[turn]
	}
	return board.Score()
}
