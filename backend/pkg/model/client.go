package model

import (
	"github.com/gofiber/websocket/v2"
)

type (
	ClientMap   map[string]*Client
	ClientArray []*Client
)

type Client struct {
	Name       string
	Connection *websocket.Conn
	Leader     bool
}

func NewClient(conn *websocket.Conn, username string) (c *Client) {
	return &Client{
		Name:       username,
		Connection: conn,
		Leader:     false,
	}
}

/// Util

func (M ClientMap) Delete(client *Client) {
	for k, v := range M {
		if v == client {
			delete(M, k)
		}
	}
}

func (M ClientMap) Array() (res ClientArray) {
	res = make(ClientArray, len(M))
	i := 0
	for _, v := range M {
		res[i] = v
		i += 1
	}
	return
}

func (A ClientArray) Contains(client *Client) bool {
	for _, c := range A {
		if c == client {
			return true
		}
	}
	return false
}

func (A ClientArray) Names() (res []string) {
	res = make([]string, len(A))
	for i, v := range A {
		res[i] = v.Name
	}
	return
}
