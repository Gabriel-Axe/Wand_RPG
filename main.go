package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", pong)
	http.HandleFunc("/game/start", HandleGameStart)
	http.HandleFunc("/game/status", HandleGetGameStatus)
	http.HandleFunc("/game/turn/pass", HandlePassTurn)
	http.HandleFunc("/game/combat/attack", HandleAttackRequest)
	http.HandleFunc("/game/combat/defend", HandleDefenseRequest)
	http.HandleFunc("/game/status/attack", HandleAttackerStatusRequest)
	http.HandleFunc("/game/status/defend", HandleDefenderStatusRequest)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
