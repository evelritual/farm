package unit

import (
	"fmt"
	"time"

	"github.com/evelritual/goose"
	"github.com/evelritual/goose/graphics"

	"github.com/evelritual/farm/pkg/event"
)

type Player struct {
	Name string

	textureAtlas graphics.TextureAtlas

	posX   int32
	posY   int32
	speed  int32
	speedX int32
	speedY int32

	events chan *event.Event
}

func NewPlayer() (*Player, error) {
	pta, err := goose.NewTextureAtlas("assets/img/player2.png", 32, 32)
	if err != nil {
		return nil, fmt.Errorf("error loading texture: %v", err)
	}

	// Subscribe to the Event Manager
	c, err := event.GlobalManager.Subscribe("player")
	if err != nil {
		return nil, fmt.Errorf("error subscribing to event manager: %v", err)
	}

	p := &Player{
		Name:         "Evelyn",
		textureAtlas: pta,
		events:       c,
		speed:        1,
	}
	go p.handleEvent()
	return p, nil
}

func (p *Player) Close() error {
	err := p.textureAtlas.Close()
	if err != nil {
		return fmt.Errorf("error closing texture atlas: %v", err)
	}
	return nil
}

func (p *Player) Draw() error {
	err := p.textureAtlas.Draw(0, p.posX, p.posY, 1.0, 1.0, 0.0)
	if err != nil {
		return fmt.Errorf("error drawing: %v", err)
	}

	return nil
}

func (p *Player) FixedUpdate(elapsedTime time.Duration) error {
	p.posX += p.speedX
	p.posY += p.speedY
	return nil
}

func (p *Player) handleEvent() {
	for e := range p.events {
		switch *e {
		case event.EventDownPress:
			p.speedY = p.speed
		case event.EventDownDepress:
			p.speedY = 0
		case event.EventLeftPress:
			p.speedX = -p.speed
		case event.EventLeftDepress:
			p.speedX = 0
		case event.EventRightPress:
			p.speedX = p.speed
		case event.EventRightDepress:
			p.speedX = 0
		case event.EventUpPress:
			p.speedY = -p.speed
		case event.EventUpDepress:
			p.speedY = 0
		}
	}
}
