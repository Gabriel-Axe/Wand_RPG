package main

type Item struct {
	Name string
	Type string // wand, potion, weapon, poison...
}

type ItemHolder interface {
    GetName() string
    GetType() string
}

func (i Item) GetName() string { return i.Name }
func (i Item) GetType() string { return i.Type }

type Wand struct {
	Item
	Spells []Attack
	ManaPool int
	RechargeRate int
}

func NewWand(name string, spells[]Attack, manaPool int, rechargeRate int) Wand {
	return Wand{
		Item: Item{
			Name: name,
			Type: "wand",
		},
		Spells: spells,
		RechargeRate: rechargeRate,
	}
}
