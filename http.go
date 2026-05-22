package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var currentGame *game

func handle_get_game_status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	gs := get_game_status(currentGame)

	json.NewEncoder(w).Encode(gs)
}

func pong(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func handle_game_start(w http.ResponseWriter, req *http.Request) {

	if currentGame != nil {
		fmt.Fprintf(w, `{"error": "Game already begun"}`)
		return
	}

	currentGame = setup_game()
	gs := get_game_status(currentGame)
	json.NewEncoder(w).Encode(gs)
}

func handle_attack_request(w http.ResponseWriter, req *http.Request) {
	if currentGame.attacker_turn == true {
		make_attack(*currentGame.attacker, *currentGame.defender, 0, 0)
	} else {
		make_attack(*currentGame.defender, *currentGame.attacker, 0, 0)
	}

	gs := get_game_status(currentGame)
	json.NewEncoder(w).Encode(gs)
}

func handle_defender_status_request(w http.ResponseWriter, req *http.Request) {
	stats := see_defender_stats()
	json.NewEncoder(w).Encode(stats)
}

func handle_attacker_status_request(w http.ResponseWriter, req *http.Request) {
	stats := see_attacker_stats()
	json.NewEncoder(w).Encode(stats)
}

func handle_defense_request(w http.ResponseWriter, req *http.Request) {
	if currentGame.attacker_turn == true {
		make_defend(*currentGame.defender, 0)
	} else {
		make_defend(*currentGame.attacker, 0)
	}
}

func handle_pass_turn(w http.ResponseWriter, req *http.Request) {
	next_turn()
	gs := get_game_status(currentGame)
	json.NewEncoder(w).Encode(gs)
}
