package main

import (
	"fmt"
	"log"
)

const ATTACKER_DEFENDING_DAMAGE_MULTIPLIER int = 70
const DEFENDER_DEFENDING_DAMAGE_MULTIPLIER int = 85

func GetDefenderStats() []map[string]interface{} {
	defender_team := currentGame.Defender.Team
	stats := make([]map[string]interface{}, len(defender_team))

	for i, u := range defender_team {
		stats[i] = map[string]interface{}{
			"id": u.ID,
			"name": u.Name,
			"health": u.Health,
			"is_defending": u.IsDefending,
		}
	}

	return stats
}

func GetAttackerStats() []map[string]interface{} {
	attacker_team := currentGame.Attacker.Team
	stats := make([]map[string]interface{}, len(attacker_team))

	for i, u := range attacker_team {
		stats[i] = map[string]interface{}{
			"id": u.ID,
			"name": u.Name,
			"health": u.Health,
			"is_defending": u.IsDefending,
		}
	}

	return stats
}

func CalculateDamage(target *Unit, ammount int) {
	target.Health -= ammount
}

func NextTurn(g *Game) {

	for _, u := range g.Attacker.Team {
		u.ProcessEffects()
	}
	for _, u := range g.Defender.Team {
		u.ProcessEffects()
	}

	log.Println("Advancing turn")
	g.Turn++
	
	log.Println("Swapping attacker and defender")
	holder := g.Attacker 
	g.Attacker = g.Defender
	g.Defender = holder
}

func ToggleDefend(g *Game, unit_id int) {
	unit := g.Defender.Team[unit_id]
	log.Printf("Togling defense for %s\n", unit.Name)

	if unit.IsDefending {
		unit.IsDefending = false
	} else {
		unit.IsDefending = true
	}
}

func MakeAttack(g *Game, attacker_unit_id int, defender_unit_id int, attack_type int) error {
	attacker := g.Attacker
	defender := g.Defender

	a_unit := attacker.Team[attacker_unit_id]
	d_unit := defender.Team[defender_unit_id]
	log.Printf("%s is attacking %s\n", a_unit.Name, d_unit.Name)

	d_is_defending := d_unit.IsDefending
	a_is_defending := a_unit.IsDefending

	var atack Attack
	if attack_type < len(a_unit.Attacks) {
		atack = a_unit.Attacks[attack_type]
	} else {
		return fmt.Errorf("The attack_type with id %d is higher then the number of attacks (%d)", attack_type, len(a_unit.Attacks))
	}

	final_damage := atack.Damage

	log.Println("Calculating defense deductions")
	if a_is_defending == true {
		final_damage = (final_damage * ATTACKER_DEFENDING_DAMAGE_MULTIPLIER) / 100
	}

	if d_is_defending == true {
		final_damage = (final_damage * DEFENDER_DEFENDING_DAMAGE_MULTIPLIER) / 100
	}

	// fmt.Printf("Dealing %d damage on IsDefending unit", final_damage)
	log.Println("Calculating damage")
	CalculateDamage(d_unit, final_damage)

	return nil
}

func (u *Unit) ProcessEffects() {
	var remaining []StatusEffect
	for _, e := range u.Effects {
		u.Health -= e.Damage
		e.Duration--
		if e.Duration > 0 {
			remaining = append(remaining, e)
		}
	}

	u.Effects = remaining
}
