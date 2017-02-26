package main

import (
	"log"
	"testing"
)

func TestRandomStrategy(t *testing.T) {
	game := Game{}

	gameStart := getGameStart()
	game.start(gameStart)

	gameUpdate := getGameUpdate(1)
	game.update(gameUpdate)

	gameUpdate = getGameUpdate(2)
	game.update(gameUpdate)

	strategy := RandomStrategy{}
	for i := 0; i < 10; i++ {
		log.Println(strategy.wantsMove(&game))
		log.Println(strategy.suggestMove(&game))
	}
}
