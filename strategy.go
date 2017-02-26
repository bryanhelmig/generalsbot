package main

import (
	"math/rand"
)

// Move is for emit('attack')
type Move struct {
	start int
	end   int
}

func (m *Move) isNull() bool {
	return m.start == m.end
}

func nullMove() Move {
	return Move{0, 0}
}

// Strategy can rank a situation for their move, and if selected, do an Attack
type Strategy interface {
	// 0 doesn't want to guide strategy
	// 100 is very confident about having a good strategy
	wantsMove(*Game) int
	suggestMove(*Game) Move
}

// RandomStrategy is not very smart lol.
type RandomStrategy struct {
	PossibleMoves []Move
}

func (r *RandomStrategy) wantsMove(game *Game) int {
	return 1
}

func (r *RandomStrategy) maybeAddPossibleMove(fromCell, toCell Cell) {
	if toCell.canMoveTo() {
		r.PossibleMoves = append(r.PossibleMoves, Move{fromCell.Index, toCell.Index})
	}
}

func (r *RandomStrategy) suggestMove(game *Game) Move {
	startingCells := []Cell{}
	matrix := game.matrix()

	for _, columns := range matrix {
		for _, cell := range columns {
			if cell.canMoveFrom() {
				startingCells = append(startingCells, cell)
			}
		}
	}

	r.PossibleMoves = []Move{}

	for i := 0; i < len(startingCells); i++ {
		startingCell := startingCells[i]
		r.maybeAddPossibleMove(startingCell, matrix[startingCell.Row-1][startingCell.Col]) // up
		r.maybeAddPossibleMove(startingCell, matrix[startingCell.Row][startingCell.Col+1]) // right
		r.maybeAddPossibleMove(startingCell, matrix[startingCell.Row+1][startingCell.Col]) // down
		r.maybeAddPossibleMove(startingCell, matrix[startingCell.Row][startingCell.Col-1]) // left
	}

	if len(r.PossibleMoves) != 0 {
		return r.PossibleMoves[rand.Intn(len(r.PossibleMoves))]
	}

	return nullMove()
}
