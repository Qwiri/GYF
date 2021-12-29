package handlers

import (
	"encoding/json"
	"github.com/Qwiri/GYF/backend/pkg/model"
	"github.com/Qwiri/GYF/backend/pkg/util"
	"github.com/gofiber/websocket/v2"
)

func strMarshal(a interface{}) string {
	if a == nil {
		return "nil"
	}
	data, err := json.Marshal(a)
	if err != nil {
		return "errored"
	}
	return string(data)
}

func handleWhoAmI(conn *websocket.Conn, game *model.Game, client *model.Client, _ string, message []string) error {
	util.Write(conn, "YOU "+strMarshal(client)+" GAME "+strMarshal(game)+" MSG "+strMarshal(message))
	return nil
}
