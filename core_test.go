package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func getGameStart() GameStart {

	var err error

	var byt []byte
	byt, err = ioutil.ReadFile("json/game_start.json")
	if err != nil {
		panic(err)
	}

	var gameStart GameStart
	err = json.Unmarshal(byt, &gameStart)
	if err != nil {
		panic(err)
	}
	return gameStart
}

func TestGameStartInterface(t *testing.T) {
	gameStart := getGameStart()
	log.Println(gameStart)
}

func getGameUpdate(num int) GameUpdate {
	var err error

	var byt []byte
	byt, err = ioutil.ReadFile(fmt.Sprintf("json/game_update_%d.json", num))
	if err != nil {
		panic(err)
	}

	var gameUpdate GameUpdate
	err = json.Unmarshal(byt, &gameUpdate)
	if err != nil {
		panic(err)
	}
	return gameUpdate
}

func TestGameUpdateInterface(t *testing.T) {
	gameUpdate := getGameUpdate(1)
	log.Println(gameUpdate)
}

func TestGameMap(t *testing.T) {
	game := Game{
		Strategies: []Strategy{
			new(RandomStrategy),
		},
	}

	gameStart := getGameStart()
	game.start(gameStart)

	gameUpdate := getGameUpdate(1)
	game.update(gameUpdate)

	gameUpdate = getGameUpdate(2)
	game.update(gameUpdate)

	log.Println(game.matrix())

	log.Println("\n" + game.makeMap())

	log.Println("!!!!", game.suggestMove())
}
