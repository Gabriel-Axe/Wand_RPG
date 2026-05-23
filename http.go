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

func handle_attack_request(w http.ResponseWriter, r *http.Request) {
	var req AttackRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprint(w, `{"error": "invalid JSON"}`)
		return
	}

	if req.Target_id == nil || req.Unit_id == nil {
		fmt.Fprintf(w, `{"error": "Either target or unit id is null"}`)
	}

	make_attack(*req.Unit_id, *req.Target_id)

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

func handle_defense_request(w http.ResponseWriter, r *http.Request) {
	var req DefenseRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprint(w, `{"error": "invalid JSON"}`)
		return
	}

	if currentGame == nil {
		fmt.Fprint(w, `{"error": "no game started"}`)
		return
	}
	if currentGame.defender == nil {
		fmt.Fprint(w, `{"error": "defender is nil"}`)
		return
	}
	if currentGame.attacker == nil {
		fmt.Fprint(w, `{"error": "attacker is nil"}`)
		return
	}

	if req.Unit_id == nil {
		fmt.Fprintf(w, `{"error": "Either target or unit id is null"}`)
	}

	toggle_defend(*currentGame.defender, *req.Unit_id)

	gs := get_game_status(currentGame)
	json.NewEncoder(w).Encode(gs)
}

func handle_pass_turn(w http.ResponseWriter, req *http.Request) {
	next_turn()
	gs := get_game_status(currentGame)
	json.NewEncoder(w).Encode(gs)
}
