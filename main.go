package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	// "os"
	"strconv"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

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
	http.Handle("/ping", http.HandlerFunc(Pong))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Some error ocurred: ", nil)
	}
}

