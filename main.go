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
	http.HandleFunc("/game/start", handle_game_start)
	http.HandleFunc("/game/status", handle_get_game_status)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
