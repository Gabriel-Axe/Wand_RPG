package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// g := setup_game()
	// game_loop(g)
	// get_input()

	http.HandleFunc("/ping", pong)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func pong(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func handle_http() {
	http.Handle("/ping", http.HandlerFunc(pong))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Some error ocurred: ", nil)
	}
}

