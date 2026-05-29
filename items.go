package main

import "fmt"

func (i Item) GetName() string { return i.Name }
func (i Item) GetType() string { return i.Type }

func (w Wand) Use(target *Unit) error {

	mana_cost := 0
	for _, spell := range w.Spells {
		mana_costs += spell.ManaUsage
	}

	if mana_cost > w.ManaPool {
		return  fmt.Error("Wand does not have enough mana to send spells")
	}

	for _, spell := range w.Spells {
		CalculateDamage(target, spell.Damage)
	}

	return nil
}

func PassEffectInTarget

func (p Potion) Use(target *Unit) error {
	target.Effects = append(target.Effects, p.Effects...)
	return nil
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

func NewMistPotion() Potion {

	f := FreezeEffect{
		Slowdown: 3,
		Turns: 4,
	}
	i := FireEffect{
		DamagePerTurn: 5,
		Turns: 4,
	}

	return  Potion{
		Effects: []StatusEffect{
			f.ToStatusEffect(),
			i.ToStatusEffect(),
		},
	}
}
