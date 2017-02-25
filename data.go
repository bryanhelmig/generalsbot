package main

import (
	"bytes"
	"log"
	"strconv"
)

// GameMap contains all the goodies
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

func (m *GameMap) makeMap() string {
	var buffer bytes.Buffer

	printables := map[int]string{
		tileEmpty:       " ",
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
		if army != 0 {
			cell = strconv.Itoa(army)
		} else if terr < 0 {
			cell = printables[terr]
		}
		buffer.WriteString(" " + cell + " ")
	}

	return buffer.String()
}

func (m *GameMap) printMap() {
	log.Println(m.makeMap())
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

// GameStart is from game_start.json
type GameStart struct {
	PlayerIndex  int      `json:"playerIndex"`
	ReplayID     string   `json:"replay_id"`
	ChatRoom     string   `json:"chat_room"`
	TeamChatRoom string   `json:"team_chat_room"`
	Usernames    []string `json:"usernames"`
	Teams        []int    `json:"teams"`
}

// GameUpdate is from game_update.json
type GameUpdate struct {
	Scores []struct {
		Total int  `json:"total"`
		Tiles int  `json:"tiles"`
		I     int  `json:"i"`
		Dead  bool `json:"dead"`
	} `json:"scores"`
	Turn        int   `json:"turn"`
	Stars       []int `json:"stars"`
	AttackIndex int   `json:"attackIndex"`
	Generals    []int `json:"generals"`
	MapDiff     []int `json:"map_diff"`
	CityDiff    []int `json:"city_diff"`
}
