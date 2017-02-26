package main

// Move is for emit('attack')
type Move struct {
	start int
	end   int
	is50  bool
}

func (m *Move) isNull() bool {
	return m.start == m.end
}

// Strategy can rank a situation for their move, and if selected, do an Attack
type Strategy interface {
	wantsMove(Game) int // 0 doesn't want to move, 100 is very confident
	suggestMove(Game) Move
}

// RandomStrategy is not very smart lol.
type RandomStrategy struct{}

func (r *RandomStrategy) wantsMove(game Game) int {
	return 10
}

func (r *RandomStrategy) suggestMove(game Game) Move {
	return Move{0, 0, false}
}

func findStrategy(game Game) Move {
	return Move{0, 0, false}
}
