package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func list_unit_stats(u Unit) {
	fmt.Println("\tname:",  u.Name)
	fmt.Println("\thealth:", u.Health)
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

func quick_game_setup() *Game {

	team_1 := make_goblin_team()

	team_1[0].ID = 1
	team_1[1].ID = 2
	team_1[2].ID = 3

	team_2 := make_elven_team()

	team_2[0].ID = 1
	team_2[1].ID = 2
	team_2[2].ID = 3

	p1 := &player{ID: 1, Name: "Alexander", Team: team_1[:]}
	p2 := &player{ID: 2, Name: "Oliver", Team: team_2[:]}

	g := Game{Attacker: p1, Defender: p2, Turn: 1}

	return &g
}

func make_goblin_team() [3]*Unit {
	var team [3] *Unit
	team[0] = make_goblin()
	team[1] = make_goblin()
	team[2] = make_goblin()

	return team
}
func make_elven_team() [3]*Unit {
	var team [3] *Unit
	team[0] = make_goblin()
	team[1] = make_goblin()
	team[2] = make_goblin()

	return team
}
func make_werewolf_team() [3]*Unit {
	var team [3] *Unit
	team[0] = make_werewolf()
	team[1] = make_werewolf()
	team[2] = make_werewolf()

	return team
}

func get_game_status(g *Game) map[string]interface{} {
	if g == nil {
		return map[string]interface{}{"error": "No game started"}
	}

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
	return &Unit{
		Name: "Goblin",
		Health: 70,
		Attacks: []Attack{
			AttackHeadbutt,
		},
	}
}

func make_werewolf() *Unit {
	return &Unit{
		Name: "Werewolf",
		Health: 130,
		Attacks: []Attack{
			AttackSlash,
			AttackHeadbutt,
		},
	}
}

func make_elven() *Unit {
	return &Unit{
		Name: "Elven",
		Health: 100,
		Attacks: []Attack{
			AttackSlash,
			AttackRoundkick,
		},
	}
}
