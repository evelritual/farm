package main

import (
	"fmt"
	"log"
	"time"

	"github.com/evelritual/goose"
	"github.com/evelritual/goose/input"

	"github.com/evelritual/farm/pkg/event"
	"github.com/evelritual/farm/pkg/unit"
)

type Game struct {
	eventManager *event.Manager
	keyboard     input.Keyboard

	units []unit.Unit
}

func (g *Game) Init() error {
	// Create Event Manager
	g.eventManager = event.NewManager()

	// Create Input
	g.keyboard = goose.NewKeyboard()

	// Create Player
	p, err := unit.NewPlayer()
	if err != nil {
		return fmt.Errorf("error creating player: %v", err)
	}
	g.units = append(g.units, p)
	return nil
}

func (g *Game) Close() error {
	for _, u := range g.units {
		err := u.Close()
		if err != nil {
			log.Printf("error closing unit: %v", err)
		}
	}
	return nil
}

func (g *Game) FixedUpdate(elapsedTime time.Duration) error {
	for _, u := range g.units {
		err := u.FixedUpdate(elapsedTime)
		if err != nil {
			return fmt.Errorf("error updating unit: %v", err)
		}
	}
	return nil
}

func (g *Game) Update() error {
	if g.keyboard.IsKeyPress(input.KeyArrowDown) {
		g.eventManager.Submit(event.EventDownPress)
	} else if g.keyboard.IsKeyRelease(input.KeyArrowDown) {
		g.eventManager.Submit(event.EventDownDepress)
	}
	if g.keyboard.IsKeyPress(input.KeyArrowLeft) {
		g.eventManager.Submit(event.EventLeftPress)
	} else if g.keyboard.IsKeyRelease(input.KeyArrowLeft) {
		g.eventManager.Submit(event.EventLeftDepress)
	}
	if g.keyboard.IsKeyPress(input.KeyArrowRight) {
		g.eventManager.Submit(event.EventRightPress)
	} else if g.keyboard.IsKeyRelease(input.KeyArrowRight) {
		g.eventManager.Submit(event.EventRightDepress)
	}
	if g.keyboard.IsKeyPress(input.KeyArrowUp) {
		g.eventManager.Submit(event.EventUpPress)
	} else if g.keyboard.IsKeyRelease(input.KeyArrowUp) {
		g.eventManager.Submit(event.EventUpDepress)
	}
	return nil
}

func (g *Game) Draw() error {
	for _, u := range g.units {
		err := u.Draw()
		if err != nil {
			return fmt.Errorf("error drawing unit: %v", err)
		}
	}
	return nil
}

func main() {
	err := goose.Run(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
