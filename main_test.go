package main

import (
	"encoding/json"
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

func getGameUpdate() GameUpdate {
	var err error

	var byt []byte
	byt, err = ioutil.ReadFile("json/game_update.json")
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
	gameUpdate := getGameUpdate()
	log.Println(gameUpdate)
}

func TestGameMap(t *testing.T) {
	game := Game{}

	gameStart := getGameStart()
	game.start(gameStart)

	gameUpdate := getGameUpdate()
	game.update(gameUpdate)

	log.Println("\n" + game.makeMap())
	// log.Println("\n" + game.makeCityMap())
}
