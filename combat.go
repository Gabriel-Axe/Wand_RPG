package main

import (
	"fmt"
	"strconv"
)

// func attack_opposite_player_team(current_player *player, g game) {
//
// 	other_player := get_opposite_player(current_player, g)
//
// 	player_team := current_player.team
// 	other_team := other_player.team
// 	var attacker *unit
// 	var attacked *unit
//
// 	fmt.Println("Select your attacker: ")
// 	attacker = select_unit_from_team(player_team)
//
// 	fmt.Println("Select a unit to attack: ")
// 	attacked = select_unit_from_team(other_team)
//
// 	attacked.health -= attacker.damage
// 	if attacked.health <= 0 {
// 		fmt.Printf("A unidade %s de %s morreu!", attacked.name, attacked.name)
// 		other_player.team[attacked.id] = nil
// 	}
// }

const ATTACKER_DEFENDING_DAMAGE_MULTIPLIER int = 70
const DEFENDER_DEFENDING_DAMAGE_MULTIPLIER int = 85

func choose_team(player *player) {

	teamMaxSize := 3
	team := make([]*Unit, 0, teamMaxSize)

	fmt.Printf("Select your team %s: \n", player.Name)

	fmt.Println("Goblin: 1")
	fmt.Println("Elven: 2")
	fmt.Println("Werewolf: 3")

	for len(team) < teamMaxSize {
		fmt.Printf("Choose unit %d/%d: ", len(team)+1, teamMaxSize)
		choice := get_input()

		var newUnit *Unit
		switch choice {
		case "1":
			newUnit = make_goblin()
			break
		case "2":
			newUnit = make_elven()
			break
		case "3":
			newUnit = make_werewolf()
			break
		default:
			fmt.Println("Unknow input: %s", choice)
			continue
		}
		team = append(team, newUnit)
		fmt.Printf("Added! Team now has %d/%d units\n", len(team), teamMaxSize)
	}

	player.Team = team
}

func see_defender_stats() []map[string]interface{} {
	defender_team := currentGame.Defender.Team
	stats := make([]map[string]interface{}, len(defender_team))

	for i, u := range defender_team {
		stats[i] = map[string]interface{}{
			"id": u.ID,
			"name": u.Name,
			"damage": u.Damage,
			"health": u.Health,
			"is_defending": u.IsDefending,
		}
	}

	return stats
}

func see_attacker_stats() []map[string]interface{} {
	attacker_team := currentGame.Attacker.Team
	stats := make([]map[string]interface{}, len(attacker_team))

	for i, u := range attacker_team {
		stats[i] = map[string]interface{}{
			"id": u.ID,
			"name": u.Name,
			"damage": u.Damage,
			"health": u.Health,
			"is_defending": u.IsDefending,
		}
	}

	return stats
}

func next_turn() {
	currentGame.Turn++
	holder := currentGame.Attacker 
	currentGame.Attacker = currentGame.Defender
	currentGame.Defender = holder
}

func toggle_defend(defender player, unit_id int) {
	unit := defender.Team[unit_id]
	if unit.IsDefending {
		unit.IsDefending = false
	} else {
		unit.IsDefending = true
	}
	next_turn()
}

func make_attack(attacker_unit_id int, defender_unit_id int) {
	attacker := currentGame.Attacker
	defender := currentGame.Defender

	a_unit := attacker.Team[attacker_unit_id]
	d_unit := defender.Team[defender_unit_id]

	d_is_defending := d_unit.IsDefending
	a_is_defending := a_unit.IsDefending

	final_damage := a_unit.Damage

	if a_is_defending == true {
		final_damage = (final_damage * ATTACKER_DEFENDING_DAMAGE_MULTIPLIER) / 100
	}
	if d_is_defending == true {
		final_damage = (final_damage * DEFENDER_DEFENDING_DAMAGE_MULTIPLIER) / 100
	}

	// WARN: Remove this, substitute for a proper logger
	fmt.Printf("Dealing %d damage on IsDefending unit", final_damage)
	d_unit.Health -= final_damage

	next_turn()
}

func select_unit_from_team(t []*Unit) *Unit {

	for _, unit := range(t) {
		fmt.Printf("%d. %s: %d\n", unit.ID, unit.Name, unit.Health)
	}
	fmt.Println()

	var id int

	for {
		choice := get_input()
		id, err := strconv.Atoi(choice)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Number chosen: %d\n", id)

		// if id < 0 || id < len(t) || {
		// 	fmt.Printf("Please choose a humber between 1 and %d\n", len(t))
		// 	continue
		// }

		// break
	}

	chosen_unit := t[id]
	return  chosen_unit
}

// func list_opposit_player_team(current_player *player, g game) {
// 	p2 := get_opposite_player(current_player, g)
// 	// for _, unit in range p2.unicurrent_player.units
// 	fmt.Println("# ------------------- #")
// 	get_player_team(p2)		
// 	fmt.Println("# ------------------- #")
// }

func list_player_team(current_player *player, g Game) {
	fmt.Println("# ------------------- #")
	get_player_team(current_player)		
	fmt.Println("# ------------------- #")
}

// func get_opposite_player(current_player *player, g game) *player {
// 	var p2 *player
// 	if current_player.id == 1 {
// 		p2 = g.players[1]
// 	} else {
// 		p2 = g.players[0]
// 	}
//
// 	return p2
// }
