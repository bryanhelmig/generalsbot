package main

import (
	"bytes"
	"log"
	"strconv"

	"github.com/mgutz/ansi"
)

// Game holds all the default state
type Game struct {
	PlayerIndex int
	Turn        int
	Cities      GameMap
	Map         GameMap
	Generals    []int

	Start   GameStart
	Updates []GameUpdate
}

func (g *Game) start(gameStart GameStart) {
	g.PlayerIndex = gameStart.PlayerIndex
	g.Start = gameStart
}

func (g *Game) update(gameUpdate GameUpdate) {
	g.Turn = gameUpdate.Turn
	g.Generals = gameUpdate.Generals
	g.Map.patch(gameUpdate.MapDiff)
	g.Cities.patch(gameUpdate.CityDiff)
	g.Updates = append(g.Updates, gameUpdate)
}

func (g *Game) maybeMove() Move {
	strategy := RandomStrategy{}
	if strategy.wantsMove(*g) != 0 {
		return strategy.suggestMove(*g)
	}
	return Move{0, 0, false}
}

func (g *Game) makeMap() string {
	return g.Map.makeMap(g.PlayerIndex)
}

func (g *Game) printMap() {
	g.Map.printMap(g.PlayerIndex)
}

func (g *Game) makeCityMap() string {
	return g.Cities.makeMap(-1)
}

func (g *Game) printCityMap() {
	g.Cities.printMap(-1)
}

// GameMap tracks a list of integers that represent the armies and terrain.
type GameMap struct {
	Raw []int
}

func (m *GameMap) width() int {
	return m.Raw[0]
}

func (m *GameMap) height() int {
	return m.Raw[1]
}

func (m *GameMap) size() int {
	return m.width() * m.height()
}

func (m *GameMap) armies() []int {
	return m.Raw[2 : m.size()+2]
}

func (m *GameMap) terrain() []int {
	return m.Raw[m.size()+2 : m.size()+2+m.size()]
}

// From https://github.com/vzhou842/generals.io-Node.js-Bot-example/blob/master/main.js#L69
func (m *GameMap) patch(diff []int) {
	var out []int
	i := 0

	for i < len(diff) {
		if diff[i] != 0 {
			out = append(out, m.Raw[len(out):len(out)+diff[i]]...)
		}
		i++
		if i < len(diff) && diff[i] != 0 {
			out = append(out, diff[i+1:i+1+diff[i]]...)
			i += diff[i]
		}
		i++
	}

	m.Raw = out
}
func (m *GameMap) makeMap(playerIndex int) string {
	var buffer bytes.Buffer

	printables := map[int]string{
		tileEmpty:       ".",
		tileMountain:    "^",
		tileFog:         "~",
		tileFogObstacle: "?",
	}

	armies := m.armies()
	terrain := m.terrain()

	for i := 0; i < m.size(); i++ {
		var cell string
		army := armies[i]
		terr := terrain[i]
		if i != 0 && i%m.width() == 0 {
			buffer.WriteString("\n")
		}
		color := "red"
		if terr == playerIndex {
			color = "green"
		}
		if playerIndex < 0 {
			color = "white"
		}
		if army != 0 {
			cell = ansi.ColorCode(color) + strconv.Itoa(army) + ansi.ColorCode("reset")
		} else if terr < 0 {
			cell = printables[terr]
		}
		buffer.WriteString(" " + cell + " ")
	}

	return buffer.String()
}

func (m *GameMap) printMap(playerIndex int) {
	log.Println(m.makeMap(playerIndex))
}
