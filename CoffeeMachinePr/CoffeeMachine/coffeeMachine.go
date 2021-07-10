package CoffeeMachine

import "Practika/CoffeeMachinePr/CoffeeModel"

type Machine struct {
	Money  int
	Coffee int
	Milk   int
	Water  int
	Cup    int

	CoffeeTypes map[string]CoffeeModel.Coffee
}

func NewMachine(money, coffee, milk, water, cup int) *Machine {
	return &Machine{
		Money:  money,
		Coffee: coffee,
		Milk:   milk,
		Water:  water,
		Cup:    cup,
	}
}

func (mach *Machine) СoffeeType() {
	var CoffeeTypes = make(map[string]CoffeeModel.Coffee)
	CoffeeTypes["Эспрессо"] = CoffeeModel.Coffee{
		Name:   "Эспрессо",
		Coffee: 30,
		Water:  50,
		Milk:   0,
		Money:  50,
		Cup:    1,
	}
	CoffeeTypes["Каппучино"] = CoffeeModel.Coffee{
		Name:   "Каппучино",
		Coffee: 30,
		Water:  50,
		Milk:   50,
		Money:  100,
		Cup:    1,
	}
	CoffeeTypes["Бамовский зерновой"] = CoffeeModel.Coffee{
		Name:   "Бамовский зерновой",
		Coffee: 50,
		Water:  50,
		Milk:   50,
		Money:  75,
		Cup:    1,
	}
	CoffeeTypes["Ристретто"] = CoffeeModel.Coffee{
		Name:   "Ристретто",
		Coffee: 30,
		Water:  30,
		Milk:   0,
		Money:  50,
		Cup:    1,
	}
	CoffeeTypes["Американо"] = CoffeeModel.Coffee{
		Name:   "Американо",
		Coffee: 30,
		Water:  100,
		Milk:   0,
		Money:  60,
		Cup:    1,
	}
	mach.CoffeeTypes = CoffeeTypes
}
