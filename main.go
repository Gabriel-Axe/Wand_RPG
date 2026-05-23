package main

import (
	"fmt"
	"net/http"
)

func main() {

	// g := setup_game()
	// game_loop(g)
	// get_input()

	http.HandleFunc("/ping", pong)
	// http.HandleFunc("/game/start", handle_game_start)
	currentGame = setup_game() // WARN: Temporary, im just to lazy to keep doing this every test
	http.HandleFunc("/game/status", handle_get_game_status)
	http.HandleFunc("/game/turn/pass", handle_pass_turn)

	http.HandleFunc("/game/combat/attack", handle_attack_request)
	http.HandleFunc("/game/combat/defend", handle_defense_request)
	http.HandleFunc("/game/status/attack", handle_attacker_status_request)
	http.HandleFunc("/game/status/defend", handle_defender_status_request)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
