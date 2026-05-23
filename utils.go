package main

import (
	// "encoding/json"
	"fmt"
	"os"
	"strings"

	// "strconv"
	"bufio"
)

func list_unit_stats(u Unit) {
	fmt.Println("\tname:",  u.Name)
	fmt.Println("\thealth:", u.Health)
	fmt.Println("\tdamage:", u.Damage)
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

func setup_game() *Game {
	// p1 := &player{id: 1, name: "Alexander"}
	// p2 := &player{id: 2, name: "Oliver"}

	// WARN: Making the mother of all temporary solutions Jack, cant fret over every egg

	var team_1 [3] *Unit
	team_1[0] = make_goblin()
	team_1[1] = make_goblin()
	team_1[2] = make_goblin()

	team_1[0].ID = 1
	team_1[1].ID = 2
	team_1[2].ID = 3

	var team_2 [3] *Unit
	team_2[0] = make_elven()
	team_2[1] = make_elven()
	team_2[2] = make_elven()

	team_2[0].ID = 1
	team_2[1].ID = 2
	team_2[2].ID = 3

	p1 := &player{ID: 1, Name: "Alexander", Team: team_1[:]}
	p2 := &player{ID: 2, Name: "Oliver", Team: team_2[:]}

	g := Game{Attacker: p1, Defender: p2, Turn: 1}

	return &g
}

func get_game_status(g *Game) map[string]interface{} {
	if g == nil {
		return map[string]interface{}{"error": "No game started"}
	}

	// players := make([]map[string]interface{}, len(g.players))
	players := make([]map[string]interface{}, 2)
	players[0] = map[string]interface{}{
		"id": g.Attacker.ID,
	}
	players[1] = map[string]interface{}{
		"id": g.Defender.ID,
	}

	return map[string]interface{}{
		"defender": players,
		"turn": g.Turn,
	}
}

func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func get_player_team(p *player) {
	fmt.Println("Time do jogador", p.ID, ":")
	for _, unit := range (p.Team) {
		list_unit_stats(*unit)
	}
}

func make_goblin() *Unit {
	return &Unit{Name: "Goblin", Health: 70, Damage: 10}
}

func make_werewolf() *Unit {
	return &Unit{Name: "Werewolf", Health: 130, Damage: 30}
}

func make_elven() *Unit {
	return &Unit{Name: "Elven", Health: 100, Damage: 20}
}
