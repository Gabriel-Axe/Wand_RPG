package main

import (
	"log"
	"flag"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
	// handle_http()
}

func handle_http() {
	http.Handle("/ping", http.HandlerFunc(Pong))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Some error ocurred: ", nil)
	}
}

func Pong(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

type unit struct {
	name string
	health int
	damage int
}

type player struct {
	name string
	main_unit unit
}

type game struct {
	players []player // TODO: Make this hold exactly 2
	turn int
}
