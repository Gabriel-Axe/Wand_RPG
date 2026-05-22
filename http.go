package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var currentGame *game

func handle_get_game_status(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if currentGame == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "no game started"}`)
		return
	}

	json.NewEncoder(w).Encode(currentGame)
}

func pong(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func handle_game_start(w http.ResponseWriter, req *http.Request) {
	currentGame = setup_game()
	fmt.Fprintf(w, "uuuh: %v", currentGame.players)
}
