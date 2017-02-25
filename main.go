package main

import (
	"log"
	"time"

	"github.com/bryanhelmig/golang-socketio"
	"github.com/bryanhelmig/golang-socketio/transport"
)

func connect() *gosocketio.Client {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("botws.generals.io", 80, false),
		transport.GetDefaultWebsocketTransport())

	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Println("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func main() {
	var err error

	var playerIndex int
	gameMap := GameMap{}

	c := connect()

	err = c.On("game_start", func(h *gosocketio.Channel, data GameStart) {
		playerIndex = data.PlayerIndex
		log.Println("On game_start: ", data)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("game_update", func(h *gosocketio.Channel, data GameUpdate) {
		gameMap.patch(data.MapDiff)
		log.Println("Update taken", gameMap.width(), gameMap.height())
		log.Println(gameMap.makeMap())
	})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	err = c.Emit("get_username", []string{userID})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Emit("join_private", []string{gameID, userID})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Emit("set_force_start", []string{gameID, "true"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Joined custom game at http://bot.generals.io/games/" + gameID)

	time.Sleep(5 * 60 * time.Second)

	c.Close()

	log.Println("Complete")
}
