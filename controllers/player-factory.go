package controllers

import "github.com/hajimehoshi/ebiten/v2"

type PlayerFactory struct {
	GoUpKey      ebiten.Key
	GoRightKey   ebiten.Key
	GoDownKey    ebiten.Key
	GoLeftKey    ebiten.Key
	Controllable Controllable
}

func (factory PlayerFactory) Create() *Player {
	return &Player{
		goUpKey:      factory.GoUpKey,
		goRightKey:   factory.GoRightKey,
		goDownKey:    factory.GoDownKey,
		goLeftKey:    factory.GoLeftKey,
		controllable: factory.Controllable,
	}
}
