package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func TestNothing(t *testing.T) {
	log.Println("test!")
	// main()
}

func getGameUpdate() GameUpdate {
	var err error

	var byt []byte
	byt, err = ioutil.ReadFile("json/game_update.json")
	if err != nil {
		panic(err)
	}

	var data GameUpdate
	err = json.Unmarshal(byt, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func TestGameUpdateInterface(t *testing.T) {
	data := getGameUpdate()
	log.Println(data)
}

func TestGameStartInterface(t *testing.T) {
	var err error

	var byt []byte
	byt, err = ioutil.ReadFile("json/game_start.json")
	if err != nil {
		panic(err)
	}

	var data GameStart
	err = json.Unmarshal(byt, &data)
	if err != nil {
		panic(err)
	}
	log.Println(data)
}

func TestGameMap(t *testing.T) {
	gameMap := GameMap{}
	data := getGameUpdate()
	gameMap.patch(data.MapDiff)
	log.Println("\n" + gameMap.makeMap())
}
