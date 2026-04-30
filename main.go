package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	// "os"
	"strconv"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {

	g := setup_game()
	game_loop(g)
	// get_input()
}

func game_loop(g game) {
	// var player_input string
	//
	// for {
	// 	for _, player := range g.players {
	// 		for {
	// 			action_menu_show()
	//
	// 			fmt.Scan(&player_input)
	// 			if is_pressing_key("q", player_input) {
	// 				os.Exit(0) 
	// 			} else if is_pressing_key("l", player_input) { 
	// 				list_opposit_player_team(player, g)
	// 			} else if is_pressing_key("m", player_input) { 
	// 				list_player_team(player, g)
	// 			} else if is_pressing_key("a", player_input) {
	// 				attack_opposite_player_team(player, g)
	// 				break
	// 			} else {
	// 				fmt.Printf("Unknow key: %s", player_input)
	// 			}
	//
	// 			continue
	// 		}
	// 	}
	// 	g.turn += 1
	// 	fmt.Println("Current turn:", g.turn)
	// }
}

func get_opposite_player(current_player player, g game) *player {
	var p2 *player
	if current_player.id == 1 {
		p2 = g.players[1]
	} else {
		p2 = g.players[0]
	}

	return p2
}

func list_opposit_player_team(current_player *player, g *game) {
	p2 := get_opposite_player(*current_player, *g)
	// for _, unit in range p2.unicurrent_player.units
	fmt.Println("# ------------------- #")
	get_player_team(p2)		
	fmt.Println("# ------------------- #")
}

func list_player_team(current_player *player, g *game) {
	fmt.Println("# ------------------- #")
	get_player_team(current_player)		
	fmt.Println("# ------------------- #")
}

func attack_opposite_player_team(current_player *player, g *game) {
	var player_input string

	other_player := get_opposite_player(*current_player, *g)
	other_team := other_player.team

	var other_unit *unit
	var player_unit *unit

	fmt.Println("Select a unit to attack: ")
	for _, unit := range(other_team) {
		fmt.Printf("%s: %d\n", unit.name, unit.health)
	}

	// chose := 0
	// var choices [3]int

	for {
		fmt.Scan(&player_input)
		input, err := strconv.Atoi(player_input)
		if err != nil {
			panic(err)
		}

		if input < len(other_team) || input < 0 {
			fmt.Printf("Please choose a humber between 1 and %d\n", len(other_team))
			continue
		}

		other_unit = other_player.team[input]
		break
	}

	fmt.Println("Select your attacker: ")
	for _, unit := range(current_player.team) {
		fmt.Printf("%s: %d\n", unit.name, unit.health)
	}

	for {
		fmt.Scan(&player_input)
		input, err := strconv.Atoi(player_input)
		if err != nil {
			panic(err)
		}

		if input < len(current_player.team) || input < 0 {
			fmt.Printf("Please choose a humber between 1 and %d\n", len(current_player.team))
			continue
		}

		player_unit = current_player.team[input]
		break
	}

	// other_player.main_unit.health -= cur_ply_unit.damage
	other_unit.health -= player_unit.damage
	if other_unit.health <= 0 {
		fmt.Printf("A unidade %s de %s morreu!", other_unit.name, other_player.name)
		other_player.team[other_unit.id] = nil
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
	fmt.Println("Listar time: m")
	fmt.Println("Atacar o time inimigo: a")
}

func get_player_team(p *player) {
	fmt.Println("Time do jogador", p.id, ":")
	for _, unit := range (p.team) {
		list_unit_stats(*unit)
	}
}

func setup_game() game {
	p1 := &player{id: 1, name: "Alexander"}
	p2 := &player{id: 2, name: "Oliver"}

	choose_team(p1)
	choose_team(p2)

	g := game{players: []*player{p1, p2}, turn: 1}

	return g
}

func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

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

func make_goblin() *unit {
	return &unit{name: "Goblin", health: 70, damage: 10}
}

func make_werewolf() *unit {
	return &unit{name: "Werewolf", health: 130, damage: 30}
}

func make_elven() *unit {
	return &unit{name: "Elven", health: 100, damage: 20}
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
	team []*unit
}

type game struct {
	players []*player // TODO: Make this hold exactly 2
	turn int
}
