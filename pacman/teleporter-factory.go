package pacman

import "github.com/humblecandyman/pacman/utils"

type TeleporterFactory struct {
	TerminalA utils.Vector
	TerminalB utils.Vector

	Size utils.Vector
}

func (factory TeleporterFactory) Create() (*Teleporter, *Teleporter) {
	teleporterA := &Teleporter{
		position: factory.TerminalA,
		size:     factory.Size,
	}

	teleporterB := &Teleporter{
		position: factory.TerminalB,
		size:     factory.Size,
	}

	teleporterA.terminal = teleporterB
	teleporterB.terminal = teleporterA

	return teleporterA, teleporterB
}
