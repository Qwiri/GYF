package model

import (
	"github.com/gofiber/websocket/v2"
	"time"
)

type Client struct {
	Name            string
	Connection      *websocket.Conn
	Leader          bool
	LastInteraction time.Time
}

type Topic struct {
	Description string
	Submissions map[*Client]*Submission
}

type Submission struct {
	Creator *Client
	URL     string
}

func NewClient(conn *websocket.Conn, username string) (c *Client) {
	return &Client{
		Name:            username,
		Connection:      conn,
		Leader:          false,
		LastInteraction: time.Now(),
	}
}

func NewTopic(description string) *Topic {
	return &Topic{
		Description: description,
		Submissions: make(map[*Client]*Submission),
	}
}

func NewSubmission(creator *Client, url string) *Submission {
	return &Submission{creator, url}
}
