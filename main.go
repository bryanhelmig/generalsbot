package main

import (
	"log"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

// ["chat_message","chat_custom_queue_henrysgame,henry123",{"text":"Anonymous joined the custom lobby."}]
// ["queue_update",{"playerIndex":0,"numPlayers":1,"numForce":0,"teams":[1],"usernames":[null]}]
// ["game_start",{"playerIndex":1,"replay_id":"SeB6AI1ce","chat_room":"game_1488051901476rUTJcHXmZ4FJWHswAABy","team_chat_room":"game_1488051901476rUTJcHXmZ4FJWHswAABy_team_2","usernames":["catcatcat","[Bot] Henry"],"teams":[2,2]}]
// ["game_update",{"scores":[{"total":1,"tiles":1,"i":0,"dead":false},{"total":1,"tiles":1,"i":1,"dead":false}],"turn":1,"stars":[19,0],"attackIndex":0,"generals":[261,310],"map_diff":[0,800,19,21,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-4,-4,-4,-3,-3,-3,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-4,-3,-4,-3,-3,-3,-3,-3,-3,-3,-4,-3,-3,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-4,-3,-3,-3,-3,-3,-3,-3,-4,-4,-3,-3,-3,-3,-3,-3,-3,-4,-4,-3,-3,-3,-3,-4,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-4,-4,-3,-4,-4,-4,-3,-3,-3,-3,-3,-3,-3,-3,-4,-3,-3,-3,-3,-3,-4,-3,-3,-4,-3,-4,-3,-4,-4,-3,-3,-3,-3,-3,-3,-3,-4,-4,-3,-3,-3,-4,-3,-4,-3,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-4,-4,-3,-3,-3,-3,-3,-4,-3,-3,-3,-3,-4,-4,-3,-3,-3,-4,-3,-4,-4,-4,-3,-3,-4,-3,-3,-3,-3,-3,-4,-3,-3,-4,-3,-3,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-4,-3,-3,-1,-1,-1,-3,-3,-4,-3,-4,-3,-3,-4,-3,-3,-3,-4,-3,-3,-3,-3,-2,0,-1,-3,-3,-3,-3,-3,-3,-4,-3,-4,-3,-3,-3,-3,-3,-4,-4,-2,-1,-1,-3,-4,-3,-3,-3,-4,-3,-3,-1,-1,-1,-3,-4,-4,-3,-3,-3,-3,-3,-4,-3,-3,-3,-3,-3,-4,-3,-1,1,-2,-4,-3,-4,-4,-3,-3,-3,-3,-3,-3,-4,-4,-3,-3,-3,-3,-1,-1,-1,-3,-4,-3,-3,-3,-4,-3,-3,-4,-3,-3,-3,-4,-3,-3,-4,-3,-3,-3,-3,-3,-3,-3,-3,-3,-3,-4,-3,-4,-3,-3,-3,-3,-4,-4,-3,-4,-3,-3,-3,-4,-4,-3,-3,-3,-3,-3,-3,-3,-3,-4,-3,-3,-4,-4,-4,-3,-3,-3,-3,-3,-4,-3,-3,-3,-4,-3,-3],"cities_diff":[0,0]}]

const userID = "ob4zsd672"
const userName = "[Bot] Henry"
const gameID = "henrysgame"
const tileEmpty = -1
const tileMountain = -2
const tileFog = -3
const tileFogObstacle = -4 // Cities and Mountains show up as Obstacles in the fog of war.

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

	c := connect()

	err = c.On("game_start", func(h *gosocketio.Channel, msg GameStart) {
		log.Println("On game_start: ", msg)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("game_update", func(h *gosocketio.Channel, msg GameUpdate) {
		log.Println("On game_update: ", msg)
	})
	if err != nil {
		log.Fatal(err)
	}

	// err = c.On("chat_message", func(h *gosocketio.Channel, args ChatMessage) {
	// 	log.Println("On chat_message: ", args)
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = c.On("queue_update", func(h *gosocketio.Channel, args QueueUpdate) {
	// 	log.Println("On queue_update: ", args)
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	time.Sleep(1 * time.Second)

	err = c.Emit("get_username", []string{userID})
	if err != nil {
		log.Fatal(err)
	}
	// err = c.Emit("set_username", []string{userID, userName})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err = c.Emit("join_private", []string{gameID, userID})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Emit("set_force_start", []string{gameID, "true"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Joined custom game at http://bot.generals.io/games/" + gameID)

	time.Sleep(30 * time.Second)

	c.Close()

	log.Println("Complete")
}
