package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var currentGame *game

func handle_get_game_status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	game_status := gameStatusResponse(currentGame)

	json.NewEncoder(w).Encode(game_status)
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
	game_status := gameStatusResponse(currentGame)
	json.NewEncoder(w).Encode(game_status)
}

func handle_attack_request(w http.ResponseWriter, req *http.Request) {
	if currentGame.attacker_turn == true {
		make_attack(*currentGame.attacker, *currentGame.defender, 0, 0)
	} else {
		make_attack(*currentGame.defender, *currentGame.attacker, 0, 0)
	}

	game_status := gameStatusResponse(currentGame)
	json.NewEncoder(w).Encode(game_status)
}

func handle_defender_status_request(w http.ResponseWriter, req *http.Request) {
}

func handle_defense_request(w http.ResponseWriter, req *http.Request) {
}

func handle_pass_turn(w http.ResponseWriter, req *http.Request) {
	currentGame.turn++
	game_status := gameStatusResponse(currentGame)
	json.NewEncoder(w).Encode(game_status)
}
