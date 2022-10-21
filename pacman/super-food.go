package pacman

import (
	"github.com/humblecandyman/pacman/physics"
	"github.com/humblecandyman/pacman/utils"
)

type SuperFood struct {
	Food

	onEaten utils.Callback
}

func (food *SuperFood) OnCollision(collision physics.Collision) {
	// food.Food.OnCollision(collision)
	food.isDead = true
	food.onEaten()
}
