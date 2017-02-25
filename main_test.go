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

func TestGameUpdateInterface(t *testing.T) {
	var err error

	var byt []byte
	byt, err = ioutil.ReadFile("json/game_update.json")
	if err != nil {
		panic(err)
	}

	var dat GameUpdate
	err = json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
	log.Println(dat)
}

func TestGameStartInterface(t *testing.T) {
	var err error

	var byt []byte
	byt, err = ioutil.ReadFile("json/game_start.json")
	if err != nil {
		panic(err)
	}

	var dat GameStart
	err = json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
	log.Println(dat)
}
