package model

import "github.com/gofiber/websocket/v2"

func (g *Game) Broadcast(response *Response) {
	for _, client := range g.Clients {
		_ = response.Respond(client.Connection)
	}
}

func (g *Game) BroadcastExcept(conn *websocket.Conn, response *Response) {
	for _, client := range g.Clients {
		if client.Connection != conn {
			_ = response.Respond(client.Connection)
		}
	}
}
