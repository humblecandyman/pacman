package controllers

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	controllable Controllable
	goUpKey      ebiten.Key
	goRightKey   ebiten.Key
	goDownKey    ebiten.Key
	goLeftKey    ebiten.Key
}

func (player Player) Update() {
	if ebiten.IsKeyPressed(player.goUpKey) {
		player.controllable.Do("go-up")
	}
	if ebiten.IsKeyPressed(player.goRightKey) {
		player.controllable.Do("go-right")
	}
	if ebiten.IsKeyPressed(player.goDownKey) {
		player.controllable.Do("go-down")
	}
	if ebiten.IsKeyPressed(player.goLeftKey) {
		player.controllable.Do("go-left")
	}
}
