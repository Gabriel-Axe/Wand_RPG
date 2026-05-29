package main

import (
	"fmt"
)


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

func (w *Unit) DeductMana(attack *Attack) error {
	if w.ManaPool - attack.ManaUsage < 0 {
		return fmt.Errorf("Insuficient mana from %s to perform %s", atacker.Name, attack.Name)
	}

	w.ManaPool -=  attack.ManaUsage
	return nil
}

func (w *Wand) DeductMana(attack *Attack) error {
	if w.ManaPool - attack.ManaUsage < 0 {
		return fmt.Errorf("Insuficient mana from %s to perform %s", atacker.Name, attack.Name)
	}

	w.ManaPool -=  attack.ManaUsage
	return nil
}

func (f FreezeEffect) ToStatusEffect() StatusEffect {
	return StatusEffect{
		Type: "freeze",
		Slowdown: f.Slowdown,
		Duration: f.Turns,
	}
}

func (p PoisonEffect) ToStatusEffect() StatusEffect {
	return StatusEffect{
		Type: "poison",
		Damage: p.DamagePerTurn,
		Duration: p.Turns,
	}
}
func (f FireEffect) ToStatusEffect() StatusEffect {
	return StatusEffect{
		Type: "fire",
		Slowdown: f.DamagePerTurn,
		Duration: f.Turns,
	}
}
