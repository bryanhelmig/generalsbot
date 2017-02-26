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

func done(c *gosocketio.Client) {
	c.Close()
}

func main() {
	var err error

	game := Game{
		Strategies: []Strategy{
			new(RandomStrategy),
		},
	}

	c := connect()

	c.On("game_start", func(h *gosocketio.Channel, gameStart GameStart) {
		game.start(gameStart)
	})

	c.On("game_update", func(h *gosocketio.Channel, gameUpdate GameUpdate) {
		game.update(gameUpdate)
		log.Println("\n" + game.makeMap())
		// game.printCityMap()
		move := game.suggestMove()
		log.Printf("%+v", move)
		if !move.isNull() {
			err = c.Emit("attack", []int{move.start, move.end})
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	c.On("game_lost", func(h *gosocketio.Channel) {
		done(c)
	})

	c.On("game_won", func(h *gosocketio.Channel) {
		done(c)
	})

	time.Sleep(1 * time.Second)

	c.Emit("get_username", []string{userID})
	c.Emit("join_private", []string{gameID, userID})
	c.Emit("set_force_start", []string{gameID, "true"})

	log.Println("Joined custom game at http://bot.generals.io/games/" + gameID)

	time.Sleep(60 * 60 * time.Second)

	done(c)

	log.Println("Complete")
}
