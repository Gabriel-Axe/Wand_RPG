package main

type Item struct {
	Name string
	Type string // wand, potion, weapon, poison...
}

type Wand struct {
	Item
	Spells []Effect
	ManaPool int
	RechargeRate int
}

func NewWand(name string, spells[]Effect, manaPool int, rechargeRate int) Wand {
	return Wand{
		Item: Item{
			Name: name,
			Type: "wand",
		},
		Spells: spells,
		RechargeRate: rechargeRate,
	}
}
