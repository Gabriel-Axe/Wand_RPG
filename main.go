package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
	g := setup_game()

	fmt.Println(g.turn)
	fmt.Println(g.players[0].name)
	fmt.Println(g.players[0].main_unit.name)
}

func setup_game() game {
	u := unit{name: "Goblin", health: 100, damage: 10}
	p := player{name: "Alexander", main_unit: u}
	g := game{players: []player{p}, turn: 0}

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
	name string
	main_unit unit
}

type game struct {
	players []player // TODO: Make this hold exactly 2
	turn int
}
