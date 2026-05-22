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

func choose_team(player *player) {

	teamMaxSize := 3
	team := make([]*unit, 0, teamMaxSize)

	fmt.Printf("Select your team %s: \n", player.name)

	fmt.Println("Goblin: 1")
	fmt.Println("Elven: 2")
	fmt.Println("Werewolf: 3")

	for len(team) < teamMaxSize {
		fmt.Printf("Choose unit %d/%d: ", len(team)+1, teamMaxSize)
		choice := get_input()

		var newUnit *unit
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

	player.team = team
}

func make_attack(attacker player, receiver player, attacker_unit_id int, receiver_unit_id int) {
	a_unit := attacker.team[attacker_unit_id]
	r_unit := receiver.team[receiver_unit_id]

	r_unit.health -= a_unit.damage
}

func select_unit_from_team(t []*unit) *unit {

	for _, unit := range(t) {
		fmt.Printf("%d. %s: %d\n", unit.id, unit.name, unit.health)
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

func list_player_team(current_player *player, g game) {
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
