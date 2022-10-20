package controllers

type Controllable interface {
	Do(command string)
}
