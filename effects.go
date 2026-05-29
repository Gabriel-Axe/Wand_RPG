package main

import (
	"fmt"
)

type Effect interface {
	Apply(target *Unit, attacker *Unit, attack *Attack) error
}


func (p PoisonEffect) Apply(target *Unit, attacker *Unit, attack *Attack) error {

	err := DeductMana(attacker, attack)
	if err != nil {
		return  err
	}


	target.Effects = append(target.Effects, StatusEffect{
		Type: "poison",
		Damage: p.DamagePerTurn,
		Duration: p.Turns,
	})

	return nil
}

func (p FreezeEffect) Apply(target *Unit, attacker *Unit, attack Attack) error {

	err := DeductMana(attacker, &attack)
	if err != nil {
		return  err
	}

	target.Effects = append(target.Effects, StatusEffect{
		Type: "freeze",
		Slowdown: p.Slowdown,
		Duration: p.Turns,
	})
	return  nil
}

func (p FireEffect) Apply(target *Unit, attacker *Unit, attack *Attack) error {

	err := DeductMana(attacker, attack)
	if err != nil {
		return  err
	}

	target.Effects = append(target.Effects, StatusEffect{
		Type: "fire",
		Damage: p.DamagePerTurn,
		Duration: p.Turns,
	})

	return nil
}

func DeductMana(atacker *Unit, attack *Attack) error {
	if atacker.ManaPool - attack.ManaUsage < 0 {
		return fmt.Errorf("Insuficient mana from %s to perform %s", atacker.Name, attack.Name)
	}

	atacker.ManaPool -=  attack.ManaUsage
	return nil
}
