package main

import "time"

type Game struct {
	ID              string
	LastInteraction time.Time
}

func (g *Game) UpdateInteraction() {
	g.LastInteraction = time.Now()
}
