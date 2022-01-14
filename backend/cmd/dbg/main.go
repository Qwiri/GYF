package main

import (
	"encoding/json"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/apex/log"
	ws "github.com/fasthttp/websocket"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	GameID            = "" // empty = create new game
	NumberConnections = len(Names)
	Names             = []string{
		"Michael",
		"Christopher",
		"Jessica",
		"Matthew",
		"Ashley",
		"Jennifer",
		"Joshua",
		"Amanda",
		"Daniel",
		"David",
		"James",
		"Robert",
		"John",
		"Joseph",
		"Andrew",
		"Ryan",
		"Brandon",
		"Jason",
		"Justin",
		"Sarah",
		"William",
		"Jonathan",
		"Stephanie",
		"Brian",
		"Nicole",
		"Nicholas",
		"Anthony",
		"Heather",
		"Eric",
		"Elizabeth",
		"Adam",
		"Megan",
		"Melissa",
		"Kevin",
		"Steven",
		"Thomas",
		"Timothy",
		"Christina",
		"Kyle",
		"Rachel",
		"Laura",
		"Lauren",
		"Amber",
		"Brittany",
		"Danielle",
		"Richard",
		"Kimberly",
		"Jeffrey",
		"Amy",
		"Crystal",
		"Michelle",
		"Tiffany",
		"Jeremy",
		"Benjamin",
		"Mark",
		"Emily",
		"Aaron",
		"Charles",
		"Rebecca",
		"Jacob",
		"Stephen",
		"Patrick",
		"Sean",
		"Erin",
		"Zachary",
		"Jamie",
		"Kelly",
		"Samantha",
		"Nathan",
		"Sara",
		"Dustin",
		"Paul",
		"Angela",
		"Tyler",
		"Scott",
		"Katherine",
		"Andrea",
		"Gregory",
		"Erica",
		"Mary",
		"Travis",
		"Lisa",
		"Kenneth",
		"Bryan",
		"Lindsey",
		"Kristen",
		"Jose",
		"Alexander",
		"Jesse",
		"Katie",
		"Lindsay",
		"Shannon",
		"Vanessa",
		"Courtney",
		"Christine",
		"Alicia",
		"Cody",
		"Allison",
		"Bradley",
	}
)

type DebugConnection struct {
	Socket *ws.Conn
	Name   string
}

func main() {
	// create new game and get game ID
	if GameID == "" {
		// create new game
		res, err := http.DefaultClient.Get("http://localhost:8080/game/create")
		if err != nil {
			panic(err)
		}

		// read game ID
		var data []byte
		if data, err = io.ReadAll(res.Body); err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// unmarshal game
		var game model.Game
		if err = json.Unmarshal(data, &game); err != nil {
			panic(err)
		}

		log.Infof("Created game: %s", game.ID)
		GameID = game.ID
	}

	log.Infof("Connecting to game: %s", GameID)

	// list of connection for further use
	var connections []*DebugConnection

	// connect players
	for i := 0; i < NumberConnections; i++ {
		time.Sleep(5 * time.Millisecond)

		name := Names[i%len(Names)]
		// make new connection
		conn, _, err := ws.DefaultDialer.Dial("ws://localhost:8080/game/socket/"+GameID, nil)
		if err != nil {
			log.Warnf("Client %s failed: %v", name, err)
			continue
		}

		log.Infof("Created connection %s", name)

		connections = append(connections, &DebugConnection{
			Socket: conn,
			Name:   name,
		})

		// debug incoming messages
		go func() {
			for {
				mt, md, err := conn.ReadMessage()
				if err != nil {
					log.Warnf("Client %s failed to receive message: %v", err)
					break
				}
				log.Debugf("[%s] :(%d): %s", mt, string(md))
			}
		}()

		if err = conn.WriteMessage(ws.TextMessage, []byte("JOIN "+name)); err != nil {
			log.Warnf("Client %s failed to send message: %v", err)
		}
	}

	go func() {
		time.Sleep(5 * time.Second)
		for _, c := range connections {
			time.Sleep(5 * time.Millisecond)
			if err := c.Socket.WriteMessage(ws.TextMessage, []byte("CHAT was geht ABAP?")); err != nil {
				log.Warnf("[%s] cannot write message: %v", err)
			}
		}
	}()

	log.Info("Clients connected. Press CTRL-C to disconnect all clients")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	log.Info("Disconnecting ...")
	for _, c := range connections {
		time.Sleep(5 * time.Millisecond)
		if err := c.Socket.Close(); err != nil {
			log.Warnf("[%s] cannot close: %v", c.Name, err)
		}
	}
}
