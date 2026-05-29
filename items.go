package main

func (i Item) GetName() string { return i.Name }
func (i Item) GetType() string { return i.Type }

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
