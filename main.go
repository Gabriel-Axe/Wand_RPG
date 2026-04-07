package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
	g := setup_game()
	var player_input string

	for {
		for _, player := range g.players {
			get_player_team(player)
			action_menu_show()

			fmt.Scan(&player_input)
			if is_pressing_key("q", player_input) {
				os.Exit(0) 
			} else if is_pressing_key("l", player_input) { 
				list_opposit_player_team(player, g)
			} else if is_pressing_key("a", player_input) {
				attack_opposite_player_team(player, g)
			}
		}
		g.turn += 1
		fmt.Println("Current turn:", g.turn)
	}
}

func get_opposite_player(current_player player, g game) player {
	var p2 player
	if current_player.id == 1 {
		p2 = g.players[1]
	} else {
		p2 = g.players[0]
	}

	return p2
}

func list_opposit_player_team(current_player player, g game) {
	p2 := get_opposite_player(current_player, g)
	// for _, unit in range p2.unicurrent_player.units
	fmt.Println("# ------------------- #")
	get_player_team(p2)		
	fmt.Println("# ------------------- #")
}

func attack_opposite_player_team(current_player player, g game) {
	fmt.Println("Select a unit to attack: ")
	fmt.Println("1: ")
	list_opposit_player_team(current_player, g)
	other_player := get_opposite_player(current_player, g)

	cur_ply_unit := current_player.main_unit
	// for _, unit in range 
	other_player.main_unit.health -= cur_ply_unit.damage
	if other_player.main_unit.health >= 0 {
		os.Exit(0)
	}
}

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
	fmt.Println("Atacar o time inimigo: a")
}

func get_player_team(p player) {
	fmt.Println("Time do jogador", p.id, ":")
	list_unit_stats(p.main_unit)
}

func setup_game() game {
	u1 := unit{id: 1, name: "Goblin", health: 100, damage: 10}
	u2 := unit{id: 1, name: "Skeleton", health: 100, damage: 10}

	p1 := player{id: 1, name: "Alexander", main_unit: u1}
	p2 := player{id: 2, name: "Oliver", main_unit: u2}

	g := game{players: []player{p1, p2}, turn: 1}

	return g
}

func handle_http() {
	http.Handle("/ping", http.HandlerFunc(Pong))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Some error ocurred: ", nil)
	}
}

func Pong(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

type unit struct {
	id int
	name string
	health int
	damage int
}

type player struct {
	id int
	name string
	main_unit unit
}

type game struct {
	players []player // TODO: Make this hold exactly 2
	turn int
}
