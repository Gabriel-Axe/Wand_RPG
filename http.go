package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var currentGame *Game

func handle_get_game_status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentGame)
	// gs := get_game_status(currentGame)
	// json.NewEncoder(w).Encode(gs)
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
	json.NewEncoder(w).Encode(currentGame)
}

func handle_attack_request(w http.ResponseWriter, r *http.Request) {
	var req AttackRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Fprint(w, `{"error": "invalid JSON"}`)
		return
	}

	if req.Target_id == nil || req.Unit_id == nil {
		fmt.Fprintf(w, `{"error": "Either target or unit id is null"}`)
		return
	}

	if req.Attack_id == nil {
		fmt.Fprintf(w, `{"error": "Atack ID is null"}`)
		return
	}

	make_attack(*req.Unit_id, *req.Target_id, *req.Attack_id)

	json.NewEncoder(w).Encode(currentGame)
}

func handle_defender_status_request(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(currentGame.Defender)
}

func handle_attacker_status_request(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(currentGame.Attacker)
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
	if currentGame.Defender == nil {
		fmt.Fprint(w, `{"error": "defender is nil"}`)
		return
	}
	if currentGame.Attacker == nil {
		fmt.Fprint(w, `{"error": "attacker is nil"}`)
		return
	}

	if req.Unit_id == nil {
		fmt.Fprintf(w, `{"error": "Either target or unit id is null"}`)
	}

	toggle_defend(*currentGame.Defender, *req.Unit_id)

	json.NewEncoder(w).Encode(currentGame)
}

func handle_pass_turn(w http.ResponseWriter, req *http.Request) {
	next_turn()
	json.NewEncoder(w).Encode(currentGame)
}
