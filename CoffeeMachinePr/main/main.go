package main

import (
	"Practika/CoffeeMachinePr/CoffeeMachine"
)

func main() {
	cm := CoffeeMachine.NewMachine(5000, 3000, 3000, 5000, 250)
	cm.СoffeeType()
	CoffeeMachine.Mach(cm)
}
