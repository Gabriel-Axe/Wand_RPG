package main

import (
	// "encoding/json"
	"fmt"
	"os"
	"strings"

	// "strconv"
	"bufio"
)

func list_unit_stats(u unit) {
	fmt.Println("\tname:",  u.name)
	fmt.Println("\thealth:", u.health)
	fmt.Println("\tdamage:", u.damage)
}

func is_pressing_key(expected string, player_input string) bool {
	if player_input == expected { return true }
	return  false
}

func action_menu_show() {
	fmt.Println("Ação: ")
	fmt.Println("Sair: q")
	fmt.Println("Listar time inimigo: l")
	fmt.Println("Listar time: m")
	fmt.Println("Atacar o time inimigo: a")
}

func setup_game() *game {
	// p1 := &player{id: 1, name: "Alexander"}
	// p2 := &player{id: 2, name: "Oliver"}

	// WARN: Making the mother of all temporary solutions Jack, cant fret over every egg

	var team_1 [3] *unit
	team_1[0] = make_goblin()
	team_1[1] = make_elven()
	team_1[2] = make_werewolf()

	team_1[0].id = 1
	team_1[1].id = 2
	team_1[2].id = 3

	var team_2 [3] *unit
	team_2[0] = make_goblin()
	team_2[1] = make_elven()
	team_2[2] = make_werewolf()

	team_2[0].id = 1
	team_2[1].id = 2
	team_2[2].id = 3

	p1 := &player{id: 1, name: "Alexander", team: team_1[:], is_attacker: true}
	p2 := &player{id: 2, name: "Oliver", team: team_2[:], is_attacker: false}

	g := game{attacker: p1, defender: p2, turn: 1}

	return &g
}

func get_game_status(g *game) map[string]interface{} {
	if g == nil {
		return map[string]interface{}{"error": "No game started"}
	}

	// players := make([]map[string]interface{}, len(g.players))
	players := make([]map[string]interface{}, 2)
	players[0] = map[string]interface{}{
		"id": g.attacker.id,
	}
	players[1] = map[string]interface{}{
		"id": g.defender.id,
	}

	return map[string]interface{}{
		"players": players,
		"turn": g.turn,
	}
}

func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func get_player_team(p *player) {
	fmt.Println("Time do jogador", p.id, ":")
	for _, unit := range (p.team) {
		list_unit_stats(*unit)
	}
}
