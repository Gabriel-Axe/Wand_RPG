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
			if player_input == "q" {
				os.Exit(0)			
			} else if player_input == "l" {
				fmt.Println("oi")
			}
		}

		return
	}
}

func action_menu_show() {
	fmt.Println("Ação: ")
	fmt.Println("Sair: q")
	fmt.Println("Dar oi: i")
}

func get_player_team(p player) {
	fmt.Println("Time do jogador", p.id, ":")
	fmt.Println("\tname:",  p.main_unit.name)
	fmt.Println("\thealth:", p.main_unit.health)
	fmt.Println("\tdamage:", p.main_unit.damage)
}

func setup_game() game {
	u1 := unit{name: "Goblin", health: 100, damage: 10}
	u2 := unit{name: "Skeleton", health: 100, damage: 10}

	p1 := player{id: 1, name: "Alexander", main_unit: u1}
	p2 := player{id: 2, name: "Oliver", main_unit: u2}

	g := game{players: []player{p1, p2}, turn: 0}

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
